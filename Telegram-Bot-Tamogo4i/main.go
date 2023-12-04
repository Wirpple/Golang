package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
	"time"
)

var gBot *tgbotapi.BotAPI
var gToken string
var gChatId int64

var gUsersInChat Users

var gUsefulActivities = Activities{
	// Саморазвитие
	{"yoga", "Йога (15 мин)", 1},
	{"meditation", "Медитация (15 мин)", 1},
	{"language", "Изучение иностранного языка (15 мин)", 1},
	{"swimming", "Плаванье (15 мин)", 1},
	{"walk", "Прогулка (15 мин)", 1},
	{"chores", "Уборка", 1},

	// Работа
	{"work_learning", "Изучение материалов по работе (15 мин)", 1},
	{"portfolio_work", "Работа над проектом для портфолио (15 мин)", 1},
	{"resume_edit", "Редактирование резюме (15 мин)", 1},

	// Творчество
	{"creative", "Творческое созданёие (15 мин)", 1},
	{"reading", "Чтение худ. литературы (15 мин)", 1},
}

var gRewards = Activities{
	// Просмотр
	{"watch_series", "Просмотр сериала (1 серия)", 10},
	{"watch_movies", "Просмотр фильмов (1 фильм)", 30},
	{"social_nets", "Просмотр соц. сетей (30 мин)", 10},

	// Еда
	{"eat_sweets", "300 ккал вкусняшек", 60},
}

type User struct {
	id    int64
	name  string
	coins uint16
}

type Users []*User

type Activity struct {
	code, name string
	coins      uint16
}

type Activities []*Activity

func init() {
	// Uncomment adn update token value to set environment variable for telegram Bot Token given by BotFather.
	// Delete this line after setting the env var. Keep the token out of the public domain!
	//_ = os.Setenv(TOKEN_NAME_IN_OS, "INSERT_YOUR_TOKEN")
	_ = os.Setenv(TOKEN_NAME_IN_OS, "6677016908:AAEIrjtNEjEDO7EQBAWl6dv4WGsav-X2O3w")

	gToken = os.Getenv(TOKEN_NAME_IN_OS)
	if gToken == "" {
		panic(fmt.Errorf("failed to load enviroment variable %s", TOKEN_NAME_IN_OS))
	}

	var err error
	if gBot, err = tgbotapi.NewBotAPI(gToken); err != nil {
		log.Panic(err)
	}

	gBot.Debug = true
}

func isStartMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func isCallbackQuery(update *tgbotapi.Update) bool {
	return update.CallbackQuery != nil && update.CallbackQuery.Data != ""
}

func delay(second uint8) {
	time.Sleep(time.Second * time.Duration(second))
}

func sendStringMessage(msg string) {
	gBot.Send(tgbotapi.NewMessage(gChatId, msg))
}

func printSystemMessageWithDelay(delayInSecond uint8, message string) {
	gBot.Send(tgbotapi.NewMessage(gChatId, message))
	delay(delayInSecond)
}

func printIntro(update *tgbotapi.Update) {
	printSystemMessageWithDelay(2, `Привет! `+EMODJI_SUNGLASSES)
	printSystemMessageWithDelay(7, `Есть множество полезных действий, совершая которые на регулярной основе мы улучашем качество своей жизни. Но часто гараздо веселее, проще или вкуснее сделать что-то вредное. Не так ли?`)
	printSystemMessageWithDelay(7, `С большей вероятностью мы предпочтем залипнуть в YouTube Shorts вместо урока английского, купим M&M's вместо овощей, полежим на кровати вместо йоги`)
	printSystemMessageWithDelay(1, EMODJI_SAD)
	printSystemMessageWithDelay(10, `Каждый играл хоть в одну игру, где нужно прокачивать персонажа, делая его сильнее, умнее или красивее. Делать приятно, потому что каждое действие приносит результат. В реальной же жизни только систематические действия через время начинают быть заметны. Давай это изменим?`)
	printSystemMessageWithDelay(1, EMODJI_SMILE)
	printSystemMessageWithDelay(14, `Перед тобой две таблицы: "Полезные действия" и "Вознаграждения". В первой таблице перечислены несложные короткие активности, за выполнение каждой из которых ты получишь указанное количество монет. Во второй же таблице ты увидишь перечень активностей, сделать которые ты можешь только после того, как оплатишь их заработанными на предыдущем шаге монетами.`)
	printSystemMessageWithDelay(1, EMODJI_COIN)
	printSystemMessageWithDelay(10, `Например, ты пол часа занимаешьсяч йогой, за что получаешь 2 монеты. После этого у тебя 2 часа изучания программирования за что ты получаешь 8 монет. Теперь ты можешь посмотреть 1 серию "Интернов" и выйти в ноль. Все просто!`)
	printSystemMessageWithDelay(6, `Отмечай совершенные полезные активности, чтобы не потерять монеты. И не забывай "купить" вознаграждение, перед тем как его совершить.`)
}

func getKeyboardRow(buttonText, buttonCode string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(buttonText, buttonCode))
}

func askToPrintInfo() {
	msg := tgbotapi.NewMessage(gChatId, "Во вступительных сообщениях ты можешь найти смысл данного бота, и правила игры. Что думаешь?")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		getKeyboardRow(BUTTON_TEXT_PRINT_INTRO, BUTON_CODE_PRINT_INTRO),
		getKeyboardRow(BUTTON_TEXT_SKIP_INTRO, BUTON_CODE_SKIP_INTRO),
	)

	gBot.Send(msg)
}

func showMenu() {
	msg := tgbotapi.NewMessage(gChatId, "Выбери один из вариантов: ")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		getKeyboardRow(BUTTON_TEXT_BALANCE, BUTON_CODE_BALANCE),
		getKeyboardRow(BUTTON_TEXT_USEFUL_ACTIVIES, BUTON_CODE_USEFUL_ACTIVITIES),
		getKeyboardRow(BUTTON_TEXT_REWARDS, BUTON_CODE_REWARDS),
	)

	gBot.Send(msg)
}

func showBalance(user *User) {
	msg := fmt.Sprintf("%s, твой кошелек пока пуст %s \nЗатрекай полезное действие, чтобы получить монеты", user.name, EMODJI_DONT_KNOW)
	if coins := user.coins; coins > 0 {
		msg = fmt.Sprintf("%s, у тебя %d %s", user.name, coins, EMODJI_COIN)
	}

	gBot.Send(tgbotapi.NewMessage(gChatId, msg))

	showMenu()
}

func callbackQueryIsMissing(update *tgbotapi.Update) bool {
	return update.CallbackQuery == nil || update.CallbackQuery.From == nil
}

func getUserFromUpdate(update *tgbotapi.Update) (user *User, found bool) {
	if callbackQueryIsMissing(update) {
		return
	}

	userID := update.CallbackQuery.From.ID
	for _, usersInChat := range gUsersInChat {
		if userID == usersInChat.id {
			return usersInChat, true
		}
	}

	return
}

func storageUserFromUpdate(update *tgbotapi.Update) (user *User, found bool) {
	if callbackQueryIsMissing(update) {
		return
	}

	from := update.CallbackQuery.From
	user = &User{id: from.ID, name: strings.TrimSpace(from.FirstName + " " + from.LastName), coins: 0}
	gUsersInChat = append(gUsersInChat, user)

	return user, true
}

func showActivities(activities Activities, message string, isUseful bool) {
	activitiesButtonsRows := make([]([]tgbotapi.InlineKeyboardButton), 0, len(gUsefulActivities)+1)
	for _, activity := range activities {
		activityDescription := ""
		if isUseful {
			activityDescription = fmt.Sprintf("+ %d %s: %s", activity.coins, EMODJI_COIN, activity.name)
		} else {
			activityDescription = fmt.Sprintf("- %d %s: %s", activity.coins, EMODJI_COIN, activity.name)
		}
		activitiesButtonsRows = append(activitiesButtonsRows, getKeyboardRow(activityDescription, activity.code))
	}

	activitiesButtonsRows = append(activitiesButtonsRows, getKeyboardRow(BUTTON_TEXT_PRINT_MENU, BUTON_CODE_PRINT_MENU))

	msg := tgbotapi.NewMessage(gChatId, message)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(activitiesButtonsRows...)
	gBot.Send(msg)
}

func showUsefulActivities() {
	showActivities(gUsefulActivities, "Трекай полезное действие или возращайся в главное меню: ", true)
}

func showRewards() {
	showActivities(gRewards, "Трекай полезное действие или возращайся в главное меню: ", false)
}

func findActivities(activities Activities, choiceCode string) (activity *Activity, found bool) {
	for _, activity := range activities {
		if choiceCode == activity.code {
			return activity, true
		}
	}

	return
}

func processUsefulActivity(activities *Activity, user *User) {
	errorMsg := ""
	if activities.coins == 0 {
		errorMsg = fmt.Sprintf("У активности %s не указана стоимость", activities.name)
	} else if user.coins+activities.coins > MAX_USER_COINS {
		errorMsg = fmt.Sprintf("У тебя не может быть больше %d %s ", MAX_USER_COINS, EMODJI_COIN)
	}

	resultMessage := ""
	if errorMsg != "" {
		resultMessage = fmt.Sprintf("%s прости, но %s %s Твой баланс остался без изменений", user.name, errorMsg, EMODJI_SAD)
	} else {
		user.coins += activities.coins
		resultMessage = fmt.Sprintf("%s действие %s выполнено! %d %s поступило к тебе на счет. Так деражать! %s%s Теперь у тебя %d %s",
			user.name, activities.name, activities.coins, EMODJI_COIN, EMODJI_BICEPS, EMODJI_SUNGLASSES, user.coins, EMODJI_COIN)
	}

	sendStringMessage(resultMessage)
}

func processReward(activity *Activity, user *User) {
	errorMsg := ""
	if activity.coins == 0 {
		errorMsg = fmt.Sprintf(`у вознаграждения "%s" не указана стоимость`, activity.name)
	} else if user.coins < activity.coins {
		errorMsg = fmt.Sprintf(`у тебя сейчас %d %s. Ты не можешь позволить себе "%s" за %d %s`, user.coins, EMODJI_COIN, activity.name, activity.coins, EMODJI_COIN)
	}

	resultMessage := ""
	if errorMsg != "" {
		resultMessage = fmt.Sprintf("%s прости, но %s %s Твой баланс остался без изменений, вознаграждение не доступно %s", user.name, errorMsg, EMODJI_SAD)
	} else {
		user.coins -= activity.coins
		resultMessage = fmt.Sprintf(`%s, вознаграждение "%s" оплачено, приступай! %d %s было снято с твоего счета. Теперь у тебя %d %s`, user.name, activity.name, activity.coins, EMODJI_COIN, user.coins, EMODJI_COIN)
	}

	sendStringMessage(resultMessage)
}

func updateProcessing(update *tgbotapi.Update) {
	user, found := getUserFromUpdate(update)
	if !found {
		if user, found = storageUserFromUpdate(update); !found {
			gBot.Send(tgbotapi.NewMessage(gChatId, "Не получается индентифицировать пользователя"))
		}
	}

	choiceCode := update.CallbackQuery.Data
	log.Printf("[%T] %s", time.Now(), choiceCode)

	switch choiceCode {
	case BUTON_CODE_BALANCE:
		showBalance(user)
	case BUTON_CODE_USEFUL_ACTIVITIES:
		showUsefulActivities()
	case BUTON_CODE_REWARDS:
		showRewards()
	case BUTON_CODE_PRINT_INTRO:
		printIntro(update)
		showMenu()
	case BUTON_CODE_SKIP_INTRO:
		showMenu()
	case BUTON_CODE_PRINT_MENU:
		showMenu()
	default:
		if usefulActivities, found := findActivities(gUsefulActivities, choiceCode); found {
			processUsefulActivity(usefulActivities, user)

			delay(2)
			showUsefulActivities()
			return
		}

		if reward, found := findActivities(gRewards, choiceCode); found {
			processReward(reward, user)

			delay(2)
			showRewards()
			return
		}

		log.Printf(`[%T] !!!!! ERROR: Unknown code %s`, time.Now(), choiceCode)
		msg := fmt.Sprintf(`%s, прости, я не знаю код '%s' %s Пожалуйста сообщи моему создателю об ошибке`, user.name, choiceCode, EMODJI_SAD)
		sendStringMessage(msg)
	}
}

func main() {
	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UPDATE_CONFIG_TIMEOUT

	for update := range gBot.GetUpdatesChan(updateConfig) {
		if isCallbackQuery(&update) {
			updateProcessing(&update)
		} else if isStartMessage(&update) { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			gChatId = update.Message.Chat.ID
			askToPrintInfo()
		}
	}
}
