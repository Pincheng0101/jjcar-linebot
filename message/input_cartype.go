package message

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func InputCarTypeMessage() *linebot.FlexMessage {
	title := "請選擇車型/等級"
	carTypes := []string{"手排(二人座)", "手排(五人座)", "自排(五人座)", "自排(五人座)TSS"}

	headerText := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   title,
		Weight: linebot.FlexTextWeightTypeBold,
		Size:   linebot.FlexTextSizeTypeLg,
		Align:  linebot.FlexComponentAlignTypeCenter,
	}

	header := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Contents: []linebot.FlexComponent{
			headerText,
		},
	}

	buttonGroup := &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeVertical,
		Flex:   linebot.IntPtr(1),
		Contents: func() []linebot.FlexComponent {
			buttons := make([]linebot.FlexComponent, 0)
			for _, carType := range carTypes {
				button := &linebot.ButtonComponent{
					Type:   linebot.FlexComponentTypeButton,
					Action: linebot.NewMessageAction(carType, carType),
					Flex:   linebot.IntPtr(1),
				}
				buttons = append(buttons, button)
			}
			return buttons
		}(),
	}

	footer := &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeHorizontal,
		Contents: []linebot.FlexComponent{buttonGroup},
	}

	container := &linebot.BubbleContainer{
		Type:      linebot.FlexContainerTypeBubble,
		Direction: linebot.FlexBubbleDirectionTypeLTR,
		Header:    header,
		Footer:    footer,
	}

	message := linebot.NewFlexMessage("input_cartype", container)

	return message
}
