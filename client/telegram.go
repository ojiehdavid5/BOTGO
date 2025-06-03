package clients

import (
    "github.com/chuks/BOTGO/config"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "bytes"
    "fmt"
    "net/http"
)
func deleteWebhook(botToken string) error {
    url := fmt.Sprintf("https://api.telegram.org/bot%s/deleteWebhook", botToken)
    resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("{}")))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to delete webhook, status: %s", resp.Status)
    }
    return nil
}

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
    deleteWebhook(token) // Delete any existing webhook to ensure the bot works with long polling

    bot.Debug = true
    return bot, nil // Return the initialized bot instance
}