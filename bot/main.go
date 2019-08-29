package main

import (
	"fmt"
	"log"
	"net/http"

	"unicode/utf8"

	"github.com/globalsign/mgo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Record is struct for inserting records into db
type Record struct {
	Name    string
	Phone   string
	Message string
	Error   string
}

const (
	databaseURL = "mongo:27017"

	databaseName = "info"

	collectionName = "records"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("some key")
	if err != nil {
		log.Panic(err)
	}

	db, err := mgo.Dial(databaseURL)
	if err != nil {
		fmt.Printf("Error is: %s", err)
	}
	defer db.Close()
	db.SetMode(mgo.Monotonic, true)

	r := db.DB(databaseName).C(collectionName)

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.Text == "/start" || update.Message.Text == "/help" {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "How to use\n380xxxxxxxxx message\n")
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)

		} else {

			var phone, message, dbMessage string

			flag := false
			for _, r := range update.Message.Text {
				if flag == false && string(r) != " " {
					phone += string(r)
				} else if string(r) == " " && flag == false {
					flag = true
				} else if flag == true {
					if string(r) == " " {
						message += "%20"
					} else {
						message += string(r)
					}
					dbMessage += string(r)
				}
			}

			record := new(Record)
			record.Name = update.Message.From.UserName
			record.Phone = phone
			record.Message = dbMessage

			if utf8.RuneCountInString(phone) != 12 {

				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Phone number is invalid")
				msg.ReplyToMessageID = update.Message.MessageID

				record.Error = "Phone number is invalid"

				err = r.Insert(record)
				if err != nil {
					fmt.Printf("Error is: %s", err)
				}

				bot.Send(msg)

			} else if utf8.RuneCountInString(message) == 0 {

				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please enter message")
				msg.ReplyToMessageID = update.Message.MessageID

				record.Error = "Empty message"

				err = r.Insert(record)
				if err != nil {
					fmt.Printf("Error is: %s", err)
				}

				bot.Send(msg)

			} else {
				query := fmt.Sprintf("http://world.msg91.com/api/sendhttp.php?authkey=SOME_KEY_AGAIN&mobiles=%s&message=%s&sender=%s&route=4&country=SOME_COUNTRY_NAME&flash=1&unicode=1", phone, message, update.Message.From.UserName)

				resp, err := http.Post(query, "", nil)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer resp.Body.Close()

				err = r.Insert(record)
				if err != nil {
					fmt.Printf("Error is: %s", err)
				}

				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "OK")
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
