package clients

import (
    "github.com/chuks/BOTGO/config"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Init initializes the Telegram bot and returns the BotAPI instance or an error.
func Init() (*tgbotapi.BotAPI, error) {
    token, err := config.LoadConfig("TELEGRAM_APITOKEN")
    if err != nil {
        return nil, err // Return nil and the error
    }

    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        return nil, err // Return nil and the error
    }

    bot.Debug = true
    return bot, nil // Return the initialized bot instance
}