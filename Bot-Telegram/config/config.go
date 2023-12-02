package config

import (
	"flag"
	"log"
)

type Config struct {
	TgBotToken            string
	MongoConnectionString string
}

func MustLoad() Config {
	tgBotTokenToken := flag.String(
		"tg-bot-token",
		"1",
		"token for access to telegram bot",
	)

	mongoConnectionString := flag.String(
		"mongo-connection-string",
		"1",
		"connection string for MongoDB",
	)

	flag.Parse()

	if *tgBotTokenToken == "" {
		log.Fatal("token is not specified")
	}
	if *mongoConnectionString == "" {
		log.Fatal("mongo connection string is not specified")
	}

	return Config{
		TgBotToken:            *tgBotTokenToken,
		MongoConnectionString: *mongoConnectionString,
	}
}
