
package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	bot, err := linebot.New(
		"a73b62d06a29b77d3b57b3d3b0aa0e7b",
		"VhC7qpsC9Op/QN1MDc61EGAN5Jqiq2fl5RlyzGZjVJr0CnZE7gs2G52HOt9pWPEzFYvY74eRqzC939lWERLSxYZk1uaFMSQpy0v92hjZfVvyFoOX9VzMSAULznGrP5sa5wE+viP8gkG2d939jxiV3QdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text=="a"{
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("สวัสดีค่ะ")).Do(); err != nil {
								log.Print(err)
						}
					}else{
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://i.ytimg.com/vi/n7gcats5uCQ/maxresdefault.jpg", "https://i.ytimg.com/vi/n7gcats5uCQ/maxresdefault.jpg")).Do(); err != nil {
							log.Print(err)
					}
					}
					// if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					// 	log.Print(err)
					// }
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":5670", nil); err != nil {
		log.Fatal(err)
	}
}
