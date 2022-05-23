package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	pb "stocks/grpc_utils/proto/PortfolioStorage"
)

const (
	StartPrefix               = "/start"
	UpdatePortfolioPrefix  = "/update_portfolio"
	GetPortfolioInfoPrefix = "/get_portfolio"
)

func runTelegramBot() {
	token := os.Getenv("telegram_token")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)
	client := pb.NewPortfolioStorageClient("...")
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == StartPrefix {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я stocks-stocks investment telegram_bot, я могу показать ваши инвестиции")
			bot.Send(msg)
		} else if strings.HasPrefix(update.Message.Text, UpdatePortfolioPrefix) {
			// /update_portfolio APPL, GOOGL, AMZN

			var tickers []string
			if len(update.Message.Text) == len(UpdatePortfolioPrefix) {
				portfolioItems := make([]*pb.PortfolioItem, 0)
				_, err = client.UpdatePortfolio(context.Background(), &pb.UpdatePortfolioRequest{
					UserId:         update.Message.Chat.ID,
					PortfolioItems: portfolioItems,
				})
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что-то сломалось в ходе вашего запроса.")
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше портфолио сейчас пусто.")
					bot.Send(msg)
				}
				continue
			}
			var numbers []int
			// if len(update.Message.Text) == len(UPDATE_PORTFOLIO_PREFIX) {
			portfolioNumbers := make([]*pb.PortfolioItem, 0)
				_, err = client.UpdatePortfolio(context.Background(), &pb.UpdatePortfolioRequest{
					UserId:         update.Message.Chat.ID,
					PortfolioItems: portfolioNumber,
					})
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Что-то сломалось в ходе вашего запроса.")
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше портфолио сейчас пусто.")
					bot.Send(msg)
				}
				continue
			if len(tickers) == 0 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше портфолио сейчас пусто")
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваше портфолио обновилось")
				bot.Send(msg)
			}
		} else if update.Message.Text == GetPortfolioInfoPrefix {
			// /get_portfolio
			resp, err := client.GetPortfolio(context.Background(), &pb.GetPortfolioRequest{UserId: update.Message.Chat.ID})
			if err != nil {
				log.Panic( "ошибка при получении запроса портфолио", err)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неверная команда.")
			bot.Send(msg)
		}
	}
}

func parseTickers(update tgbotapi.Update) []string {
	var tickers = strings.Split(update.Message.Text[len(UpdatePortfolioPrefix):], ",")
	for i, ticker := range tickers {
		tickers[i] = strings.TrimSpace(ticker)
	}
	return tickers
}

type Error struct {
	ErrorCode   int
	Description string
}

func (e *Error) Error() string {
	return fmt.Sprintf("telegram: %d %s", e.ErrorCode, e.Description)
}

NewResponse := make([] *pb.UpdatePortfolioRequest, 0)
_, err = client.UpdatePortfolio(context.Background(), &pb.UpdatePortfolioRequest{
					UserId:         update.Message.Chat.ID,
					PortfolioItems: portfolioItems,
				})


func main() {
		err := errors.New("Ошибка! Поменяйте команду")
		fmt.Println("", err)
	runTelegramBot()
}
