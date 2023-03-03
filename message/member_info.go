package message

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

func MemberInfoMessage(name string, phone string, region string, birthday string, cartype string, url string) *linebot.FlexMessage {
	messageTemplate := []byte(fmt.Sprintf(`
	{
		"type": "bubble",
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "spacing": "md",
		  "contents": [
			{
			  "type": "text",
			  "text": "會員資料",
			  "weight": "bold",
			  "size": "xl",
			  "gravity": "center",
			  "wrap": true,
			  "contents": []
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "spacing": "sm",
			  "margin": "lg",
			  "contents": [
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "姓名",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "電話",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
					  "contents": []
					}
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "地區(縣市)",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
				  	  "contents": []
				    }
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "生日(月/日)",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "color": "#666666",
				      "flex": 4,
					  "wrap": true,
				  	  "contents": []
				    }
				  ]
				},
				{
				  "type": "box",
				  "layout": "baseline",
				  "spacing": "sm",
				  "contents": [
					{
					  "type": "text",
					  "text": "車型/等級",
					  "size": "sm",
					  "color": "#AAAAAA",
					  "flex": 2,
					  "contents": []
					},
					{
					  "type": "text",
					  "text": "%s",
					  "size": "sm",
					  "color": "#666666",
					  "flex": 4,
					  "wrap": true,
					  "contents": []
					}
				  ]
				}
			  ]
			},
			{
			  "type": "box",
			  "layout": "vertical",
			  "margin": "xxl",
			  "contents": [
				{
				  "type": "spacer"
				},
				{
				  "type": "image",
				  "url": "%s",
				  "size": "xl",
				  "aspectMode": "cover"
				},
				{
				  "type": "text",
				  "text": "此 QRCode 用於累積會員點數使用",
				  "size": "xs",
				  "color": "#AAAAAA",
				  "align": "center",
				  "margin": "xxl",
				  "wrap": true,
				  "contents": []
				}
			  ]
			}
		  ]
		}
	  }`, name, phone, region, birthday, cartype, url))

	container, err := linebot.UnmarshalFlexMessageJSON(messageTemplate)
	if err != nil {
		return nil
	}
	message := linebot.NewFlexMessage("member_info", container)

	return message
}
