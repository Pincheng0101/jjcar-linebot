package message

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func BaseMessage(text string) *linebot.FlexMessage {
	container := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: text,
				},
			},
		},
	}

	message := linebot.NewFlexMessage("base", container)

	return message
}
