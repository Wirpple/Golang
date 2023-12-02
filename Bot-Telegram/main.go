package main

import (
	"log"
	tgClient "project_test/clients/telegram"
	"project_test/config"
	"project_test/consumer/event-consumer"
	"project_test/events/telegram"
	"project_test/storage/files"
	//"read-adviser-bot/storage/mongo"
)

const (
	tgBotHost         = "api.telegram.org"
	storagePath       = "files_storage"
	batchSize         = 100
	sqliteStoragePath = "data/sqlite/storage.db"
)

func main() {
	cfg := config.MustLoad()

	storage := files.New(storagePath)
	//storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	//storage, err := sqlite.New(sqliteStoragePath)
	//if err != nil {
	//	log.Fatal("can`t connect to storage: ", err)
	//}
	//
	//if err := storage.Init(context.TODO()); err != nil {
	//	log.Fatal("can`t init storage: ", err)
	//}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
