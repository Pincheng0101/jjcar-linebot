package message

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func InputCarTypeMessage() *linebot.FlexMessage {
	messageTemplate := []byte(`
	{
		"type": "bubble",
		"direction": "ltr",
		"header": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "請選擇車型/等級",
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
			  "type": "box",
			  "layout": "vertical",
			  "contents": [
				{
				  "type": "box",
				  "layout": "vertical",
				  "flex": 1,
				  "contents": [
					{
					  "type": "button",
					  "action": {
						"type": "message",
						"label": "手排(二人座)",
						"text": "手排(二人座)"
					  }
					},
					{
					  "type": "button",
					  "action": {
						"type": "message",
						"label": "手排(五人座)",
						"text": "手排(五人座)"
					  }
					},
					{
					  "type": "button",
					  "action": {
						"type": "message",
						"label": "自排(五人座)",
						"text": "自排(五人座)"
					  }
					},
					{
					  "type": "button",
					  "action": {
						"type": "message",
						"label": "自排(五人座)TSS",
						"text": "自排(五人座)TSS"
					  }
					}
				  ]
				}
			  ]
			}
		  ]
		}
	  }`)

	container, err := linebot.UnmarshalFlexMessageJSON(messageTemplate)
	if err != nil {
		return nil
	}
	message := linebot.NewFlexMessage("input_cartype", container)

	return message
}
