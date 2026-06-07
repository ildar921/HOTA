package models

// Стуктура нашего пользователя
type User struct {
	ID       int
	Nickname string
	Role     string
	Staсk    []string
	GitHub   string //ссылка гитхаб
	Telegram string //ссылка тг
	Status   string //в сети или нет
}
