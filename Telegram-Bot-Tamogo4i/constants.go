package main

const (
	EMODJI_COIN         = "\U0001FA99" // (coin)
	EMODJI_SMILE        = "\U0001F642"
	EMODJI_SUNGLASSES   = "\U0001F60E"
	EMODJI_WOW          = "\U0001F604"
	EMODJI_DONT_KNOW    = "\U0001F937"
	EMODJI_SAD          = "\U0001F63F"
	EMODJI_BICEPS       = "\U0001F4AA"
	EMODJI_BUTTON_START = "\U000025B6"
	EMODJI_BUTTON_STOP  = "\U000025C0"

	// Блок текста для кнопок
	BUTTON_TEXT_PRINT_INTRO     = EMODJI_BUTTON_START + "Посмотреть вступление" + EMODJI_BUTTON_STOP
	BUTTON_TEXT_SKIP_INTRO      = EMODJI_BUTTON_START + "Пропустить вступление" + EMODJI_BUTTON_STOP
	BUTTON_TEXT_BALANCE         = EMODJI_BUTTON_START + "Текущий баланс" + EMODJI_BUTTON_STOP
	BUTTON_TEXT_USEFUL_ACTIVIES = EMODJI_BUTTON_START + "Полезные действия" + EMODJI_BUTTON_STOP
	BUTTON_TEXT_REWARDS         = EMODJI_BUTTON_START + "Награды" + EMODJI_BUTTON_STOP
	BUTTON_TEXT_PRINT_MENU      = EMODJI_BUTTON_START + "ОСНОВНОЕ МЕНЮ" + EMODJI_BUTTON_STOP

	// Блок кода для кнопак
	BUTON_CODE_PRINT_INTRO       = "print_intro"
	BUTON_CODE_SKIP_INTRO        = "skip_intro"
	BUTON_CODE_BALANCE           = "show_balance"
	BUTON_CODE_USEFUL_ACTIVITIES = "show_useful_activities"
	BUTON_CODE_REWARDS           = "show_rewards"
	BUTON_CODE_PRINT_MENU        = "print_menu"

	TOKEN_NAME_IN_OS             = "tamogo4i_bot"
	UPDATE_CONFIG_TIMEOUT        = 60
	MAX_USER_COINS        uint16 = 500
)
