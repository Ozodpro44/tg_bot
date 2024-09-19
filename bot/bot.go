package main

import (
    "log"
    "os"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "google.golang.org/grpc"
    pb "path/to/your/protobuf"
    "context"
)

func main() {
    // Load bot token from environment variable
    botToken := os.Getenv("BOT_TOKEN")

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Panic(err)
    }
    bot.Debug = true

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message != nil { // If a message is received
            if update.Message.IsCommand() {
                switch update.Message.Command() {
                case "order":
                    handleOrder(bot, update.Message)
                }
            }
        }
    }
}

func handleOrder(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
    msg := tgbotapi.NewMessage(message.Chat.ID, "Please provide details of the order.")

    // Here, you'll typically collect further order details (e.g., asking for location, date, etc.)
    // Assuming the message text is the order description.

    order := pb.OrderRequest{
        UserId:    message.From.ID,
        ChatId:    message.Chat.ID,
        OrderText: message.Text,
    }

    // Connect to Order Service via gRPC
    conn, err := grpc.Dial("order-service:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect to Order Service: %v", err)
    }
    defer conn.Close()

    client := pb.NewOrderServiceClient(conn)

    // Call the CreateOrder function via gRPC
    _, err = client.CreateOrder(context.Background(), &order)
    if err != nil {
        msg.Text = "Error processing your order, please try again."
    } else {
        msg.Text = "Your order has been placed successfully!"
    }

    bot.Send(msg)
}
