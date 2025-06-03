package services

import (
 "fmt"

 "github.com/chuks/BOTGO/keyboards"
 "github.com/chuks/BOTGO/repositories"
 "github.com/chuks/BOTGO/database"

 tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
 	"gorm.io/gorm"

)
var DB *gorm.DB = database.Init()
// Services for handling tasks in the Telegram bot
// This file contains functions to start the bot, set tasks, delete tasks, and show all tasks.
// It uses the tgbotapi package for Telegram Bot API interactions and repositories for database operations.
// Start initializes the bot and sends a welcome message to the user.
// It also sets up the command keyboard for user interaction.
// SetTask prompts the user to write a todo item.
// SetTaskCallback handles the callback from the user after they set a task.
// DeleteTask prompts the user to select a todo item to delete.
// DeleteTaskCallback handles the callback from the user after they select a task to delete.
// ShowAllTasks retrieves and displays all tasks for the user.
// It uses the repositories package to interact with the database for task management.
// It is important to handle errors properly and send appropriate messages to the user in case of any issues.
// It is also important to ensure that the bot is running and connected to the Telegram API.
// It is recommended to use proper logging instead of panic for error handling in production code.
// It is also recommended to use context for better control over the bot's lifecycle and to avoid blocking operations.
// Start initializes the bot and sends a welcome message to the user.
// It also sets up the command keyboard for user interaction.

// It is the entry point for the bot and should be called when the bot starts.
// It uses the tgbotapi package for Telegram Bot API interactions and keyboards for command keyboard setup.

// It is important to handle errors properly and send appropriate messages to the user in case of any issues.
// It is also important to ensure that the bot is running and connected to the Telegram API.
// It is recommended to use proper logging instead of panic for error handling in production code.
// It is also recommended to use context for better control over the bot's lifecycle and to avoid blocking operations.		



func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
 text := "Hi, here you can create todos for your todolist."
 msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
 msg.ReplyMarkup = keyboards.CmdKeyboard()
 if _, err := bot.Send(msg); err != nil {
  panic(err)
 }
}

func SetTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
 text := "Please, write todo."
 msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
 if _, err := bot.Send(msg); err != nil {
  panic(err)
 }
}

func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
 text := "Todo successfully created"

 err := repositories.SetTask(update)
 if err != nil {
  text = "Couldnt set task"
 }

 msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
 if _, err := bot.Send(msg); err != nil {
  panic(err)
 }
}

func DeleteTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
 data, _ := repositories.GetAllTasks(update.Message.Chat.ID)
 var btns []tgbotapi.InlineKeyboardButton
 for i := 0; i < len(data); i++ {
  btn := tgbotapi.NewInlineKeyboardButtonData(data[i].Task, "delete_task="+data[i].ID.String())
  btns = append(btns, btn)
 }

 var rows [][]tgbotapi.InlineKeyboardButton
 for i := 0; i < len(btns); i += 2 {
  if i < len(btns) && i+1 < len(btns) {
   row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
   rows = append(rows, row)
  } else if i < len(btns) {
   row := tgbotapi.NewInlineKeyboardRow(btns[i])
   rows = append(rows, row)
  }
 }
 fmt.Println(len(rows))
 var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
 //keyboard.InlineKeyboard = rows

 text := "Please, select todo you want to delete"
 msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
 msg.ReplyMarkup = keyboard
 if _, err := bot.Send(msg); err != nil {
  panic(err)
 }
}

func DeleteTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, taskId string) {
 text := "Task successfully deleted"

 err := repositories.DeleteTask(taskId)
 if err != nil {
  text = "Couldnt delete task"
 }

 msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
 if _, err := bot.Send(msg); err != nil {
  panic(err)
 }
}

func ShowAllTasks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
 text := "Tasks: \n"

 tasks, err := repositories.GetAllTasks(update.Message.Chat.ID)
 if err != nil {
  text = "Couldnt get tasks"
 }

 for i := 0; i < len(tasks); i++ {
  text += tasks[i].Task + " \n"
 }

 msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
 if _, err := bot.Send(msg); err != nil {
  panic(err)
 }
}