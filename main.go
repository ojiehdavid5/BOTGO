package main

import (
 "log"
 "github.com/chuks/BOTGO/client"
 "github.com/chuks/BOTGO/config"
 "github.com/chuks/BOTGO/handlers"

 "github.com/gofiber/fiber/v2"
 "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
 app := fiber.New()
 app.Use(cors.New(cors.Config{
  AllowOrigins: "*",
  AllowHeaders: "*",
 }))
 bot, _ := clients.Init()
 handler.Init(bot)
 port, err := config.LoadConfig("PORT")
 if err != nil {
  log.Fatalf("failed to load port: %v", err)
 }
 app.Listen(":" + port)
}