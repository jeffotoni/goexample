package main

import (
    "log"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
    bot, err := tgbotapi.NewBotAPI("TOKEN_DO_SEU_BOT")
    if err != nil {
        log.Fatalf("Erro ao iniciar o bot: %v", err)
    }

    bot.Debug = true

    log.Printf("Autorizado como @%s", bot.Self.UserName)

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates, err := bot.GetUpdatesChan(u)
    if err != nil {
        log.Fatalf("Erro ao obter as atualizações do bot: %v", err)
    }

    for update := range updates {
        if update.Message == nil {
            continue
        }

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        reply := "Olá, sou um bot criado pelo BotFather! O que posso fazer por você?"
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
        msg.ReplyToMessageID = update.Message.MessageID

        _, err := bot.Send(msg)
        if err != nil {
            log.Printf("Erro ao enviar a mensagem: %v", err)
        }
    }
}
