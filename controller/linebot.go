package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pincheng0101/go-linebot-server-template/command"
	"github.com/pincheng0101/go-linebot-server-template/config"
	"github.com/pincheng0101/go-linebot-server-template/firestore"
	msg "github.com/pincheng0101/go-linebot-server-template/message"
	"github.com/pincheng0101/go-linebot-server-template/state"
	"github.com/pincheng0101/go-linebot-server-template/storage"
)

type Firebase struct {
	ServiceAccountFile string
	StorageBucket      string
}

type LineBotController struct {
	Bot        *linebot.Client
	UserStates state.UserStates
	Storage    *storage.Storage
	Firestore  *firestore.Firestore
	Firebase   Firebase
}

func NewLineBotController(channelSecret, channelAccessToken string) (*LineBotController, error) {
	cfg, _ := config.LoadConfig()

	bot, err := linebot.New(channelSecret, channelAccessToken)
	if err != nil {
		panic(err)
	}
	storage, err := storage.NewStorage()
	if err != nil {
		panic(err)
	}
	firestore, err := firestore.NewFirestore()
	if err != nil {
		panic(err)
	}
	return &LineBotController{
		Bot:        bot,
		UserStates: state.NewUserStates(),
		Storage:    storage,
		Firestore:  firestore,
		Firebase: Firebase{
			ServiceAccountFile: cfg.Firebase.ServiceAccountFile,
			StorageBucket:      cfg.Firebase.StorageBucket,
		},
	}, nil
}

type Quota struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

func (ctr *LineBotController) Quota(c *gin.Context) {
	quota, err := ctr.Bot.GetMessageQuota().Do()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"quota": Quota{
			Type:  quota.Type,
			Value: quota.Value,
		},
	})
}

type rewardPointsRequest struct {
	UserID string `json:"user_id"`
	Points int    `json:"points"`
}

func (ctr *LineBotController) RewardPoints(c *gin.Context) {
	var request rewardPointsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctr.Bot.PushMessage(request.UserID, msg.RewardPointsMessage(request.Points)).Do()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Status(http.StatusOK)
}

func (ctr *LineBotController) Callback(c *gin.Context) {
	events, err := ctr.Bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctr.handleEvents(events)
	c.Status(http.StatusOK)
}

func (ctr *LineBotController) handleEvents(events []*linebot.Event) {
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				ctr.handleTextMessageEvent(event, message.Text)
			}
		}
	}
}

func (ctr *LineBotController) handleTextMessageEvent(event *linebot.Event, message string) {
	userID := event.Source.UserID
	text := message
	qrcodeImageUrl := generateQRCodeImageUrl(ctr.Firebase.StorageBucket, userID)

	var userState *state.UserState
	if !ctr.UserStates.UserExists(userID) {
		user, _ := ctr.Firestore.GetUser(userID)
		userInfo := state.UserInfo{
			Name:     user.Name,
			Phone:    user.Phone,
			Region:   user.Region,
			Birthday: user.Birthday,
			CarType:  user.CarType,
			Point:    uint(user.Points),
		}
		if user != nil {
			userState = ctr.UserStates.CreateUserStateByUser(userID, userInfo)
		}
	} else {
		userState = ctr.UserStates.CreateUserStateIfNotExist(userID)
	}

	if text == "會員資料" {
		if userState.IsRegistered {
			if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.MemberInfoMessage(
				userState.UserInfo.Name,
				userState.UserInfo.Phone,
				userState.UserInfo.Region,
				userState.UserInfo.Birthday,
				userState.UserInfo.CarType,
				qrcodeImageUrl)).Do(); err != nil {
				return
			}
		}
		userState.UpdateBeforeAsk(userID, command.RegistryAskCommand)
		if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.RegistryMessage()).Do(); err != nil {
			return
		}
	}

	switch userState.BeforeAskCommand {
	case command.RegistryAskCommand:
		switch text {
		case "是":
			userState.UpdateBeforeAsk(userID, command.InputNameAskCommand)
			if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.BaseMessage("請輸入姓名")).Do(); err != nil {
				return
			}
		case "否":
			userState.ResetUserState()
			if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.BaseMessage("已取消註冊會員，如需註冊請重新註冊流程")).Do(); err != nil {
				return
			}
		default:
			if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.BaseMessage("目前處於註冊會員狀態，請輸入是否註冊")).Do(); err != nil {
				return
			}
		}
	case command.InputNameAskCommand:
		userState.UpdateName(text)
		userState.UpdateBeforeAsk(userID, command.InputPhoneAskCommand)
		if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.BaseMessage("請輸入電話")).Do(); err != nil {
			return
		}
	case command.InputPhoneAskCommand:
		userState.UpdatePhone(text)
		userState.UpdateBeforeAsk(userID, command.InputRegionAskCommand)
		if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.BaseMessage("請輸入地區(縣市)")).Do(); err != nil {
			return
		}
	case command.InputRegionAskCommand:
		userState.UpdateRegion(text)
		userState.UpdateBeforeAsk(userID, command.InputBirthdayAskCommand)
		if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.BaseMessage("請輸入生日(月/日)")).Do(); err != nil {
			return
		}
	case command.InputBirthdayAskCommand:
		userState.UpdateBirthday(text)
		userState.UpdateBeforeAsk(userID, command.InputCarTypeAskCommand)
		if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.InputCarTypeMessage()).Do(); err != nil {
			return
		}
	case command.InputCarTypeAskCommand:
		userState.UpdateCarType(text)
		userState.Registered()
		userState.ResetUserState()

		if err := ctr.Storage.UploadQrcode(userID); err != nil {
			fmt.Println(err)
			// Todo: replay create user failed message
			return
		} else {
			user := firestore.User{
				UserID:   userID,
				Name:     userState.UserInfo.Name,
				Phone:    userState.UserInfo.Phone,
				Region:   userState.UserInfo.Region,
				Birthday: userState.UserInfo.Birthday,
				CarType:  userState.UserInfo.CarType,
				Points:   0,
			}
			if err := ctr.Firestore.AddUser(user); err != nil {
				fmt.Println(err)
				if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.BaseMessage("新增使用者失敗")).Do(); err != nil {
					return
				}
			}
			if _, err := ctr.Bot.ReplyMessage(event.ReplyToken, msg.MemberInfoMessage(
				userState.UserInfo.Name,
				userState.UserInfo.Phone,
				userState.UserInfo.Region,
				userState.UserInfo.Birthday,
				userState.UserInfo.CarType,
				qrcodeImageUrl)).Do(); err != nil {
				return
			}
		}
	}
}

func generateQRCodeImageUrl(bucket string, userID string) string {
	return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s.png?alt=media", bucket, userID)
}
