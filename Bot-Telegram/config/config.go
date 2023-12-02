package config

import (
	"flag"
	"log"
)

type Config struct {
	TgBotToken string
	//MongoConnectionString string
	//SQLiteConnectionString string
}

func MustLoad() Config {
	TgBotToken := flag.String(
		"tg-bot-token",
		"6778276160:AAFtp8OvIMWam-3Sqym3WTG99lmCU6F-GeA",
		"token for access to telegram bot",
	)

	//mongoConnectionString := flag.String(
	//	"mongo-connection-string",
	//	"",
	//	"connection string for MongoDB",
	//)

	//sqliteConnectionString := flag.String(
	//	"sqlite-connection-string",
	//	"",
	//	"connection string for SQLite",
	//)

	flag.Parse()

	if *TgBotToken == "" {
		log.Fatal("token is not specified")
	}

	//if *mongoConnectionString == "" {
	//	log.Fatal("mongo connection string is not specified")
	//}

	//if *sqliteConnectionString == "" {
	//	log.Fatal("sqlite connection string is not specified")
	//}

	return Config{
		TgBotToken: *TgBotToken,
		//MongoConnectionString: *mongoConnectionString,
		//SQLiteConnectionString: *sqliteConnectionString,
	}
}
