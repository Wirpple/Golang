package telegram

const msgHelp = `Привет! 👾

Я могу сохранять и поддерживать ваши страницы. Также я могу предложить вам их почитать.

Чтобы сохранить страницу, просто пришлите мне ссылку на нее.

Чтобы получить случайную страницу из вашего списка, отправьте мне команду /rnd.
Внимание! После этого страница будет удалена из вашего списка!`

const msgHello = "\n\n" + msgHelp

const (
	msgUnknownCommand = "Неизвестная команда 🤔"
	msgNoSavedPages   = "У вас нет сохраненных страниц 🙊"
	msgSaved          = "Сохранил! 👌"
	msgAlreadyExists  = "У вас уже есть эта страница в списке 🤗"
)
