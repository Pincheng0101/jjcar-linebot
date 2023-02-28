package message

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func messageTemplate() []byte {
	return []byte(`
	{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "尚未註冊會員，是否註冊？",
			  "weight": "bold",
			  "size": "lg",
			  "align": "center",
			  "contents": []
			}
		  ]
		},
		"footer": {
		  "type": "box",
		  "layout": "horizontal",
		  "contents": [
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "是",
				"text": "是"
			  }
			},
			{
			  "type": "button",
			  "action": {
				"type": "message",
				"label": "否",
				"text": "否"
			  }
			}
		  ]
		}
	  }
	`)

}

func RegistryMessage() *linebot.FlexMessage {
	container, err := linebot.UnmarshalFlexMessageJSON(messageTemplate())
	if err != nil {
		return nil
	}
	message := linebot.NewFlexMessage("registry", container)

	return message
}
