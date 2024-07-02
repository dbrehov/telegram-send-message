package main

import (
    "log"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    botToken := "7175657723:AAGScSFnJ_AOrCrnFqXSwMKFv9wvpQUz8iI" // Ваш токен бота
    chatID := int64(5114383390) // Замените на ваш chat ID
    messageText := "Hello, this is a test message from my Go application!"

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }

    msg := tgbotapi.NewMessage(chatID, messageText)
    _, err = bot.Send(msg)
    if err != nil {
        log.Panic(err)
    }

    log.Println("Message sent successfully")
}

