package main

import (
	"context"
	"flag"
	"log"
	tgClient "project_test/clients/telegram"
	"project_test/consumer/event-consumer"
	"project_test/events/telegram"
	"project_test/storage/sqlite"
)

const (
	tgBotHost = "api.telegram.org"
	// storagePath = "files_storage"
	storageSQLitePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	//cfg := config.MustLoad()
	//storage := files.New(storagePath)
	//storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)
	storage, err := sqlite.New(storageSQLitePath)
	if err != nil {
		log.Fatalf("can`t connect to storage: ", err)
	}

	if err := storage.Init(context.TODO()); err != nil {
		log.Fatalf("can`t init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken() /*cfg.TgBotToken*/),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"t.me/numOne_read_adviser_bot",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
