package main

import (
  "log"

  "github.com/go-telegram-bot-api/telegram-bot-api"

   "io/ioutil"

   "fmt"

   "os/exec"
)
//Keyboard in the bot
var stringKeyboard = tgbotapi.NewReplyKeyboard(
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("Fail2ban"),
    tgbotapi.NewKeyboardButton("SSH"),
  ),
  tgbotapi.NewKeyboardButtonRow(
    tgbotapi.NewKeyboardButton("Netstat"),
  ),
)

func main() {
   exec.Command("~/ServerMonitor/.server.sh") //Run shell script
//Telegram bot
//Token and message update function
  bot, err := tgbotapi.NewBotAPI("BOT_TOKEN")
  if err != nil {
    log.Panic(err)
  }

  bot.Debug = true

  log.Printf("Authorized on account %s", bot.Self.UserName)

  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil {
      continue
    }
    msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

//Commands and information output by clicking on the buttons.
    switch update.Message.Text {
    case "open":
      msg.ReplyMarkup = stringKeyboard
    case "close":
      msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
    case "Fail2ban":
      fail2ban, err := ioutil.ReadFile("fail.txt")
       if err != nil{
        fmt.Println(err)
                    }
      stats := (string(fail2ban))
      msg.Text = stats
    case "SSH":
       last, err := ioutil.ReadFile("ssh.txt")
        if err != nil{
         fmt.Println(err)
                    }
      statssh := (string(last))
      msg.Text = statssh

    case "Netstat":
    netstat, err := ioutil.ReadFile("net.txt")
       if err != nil{
        fmt.Println(err)
                    }
      statnet := (string(netstat))
      msg.Text = statnet
    }
//Message sending function (telegram)
    bot.Send(msg)
  }
}
