package message

import (
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func RewardPointsMessage(points int) *linebot.FlexMessage {
	container := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "收到點數：" + strconv.Itoa(points) + " 點",
				},
			},
		},
	}

	balanceMessage := linebot.NewFlexMessage("獎勵點數", container)

	return balanceMessage
}
