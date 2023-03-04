package message

import (
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
)

func BalanceMessage(balance int) *linebot.FlexMessage {
	container := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "您的點數餘額為：" + strconv.Itoa(balance) + " 點",
				},
			},
		},
	}

	balanceMessage := linebot.NewFlexMessage("點數查詢", container)

	return balanceMessage
}
