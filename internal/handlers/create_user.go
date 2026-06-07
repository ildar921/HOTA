package handlers

import (
	"HOTA/internal/models"
	"HOTA/internal/repositories"
	"fmt"
	"strings"
)

func CreateUser() {
	user := models.User{}

	fmt.Print("Ник: ")
	fmt.Scan(&user.Nickname)

	fmt.Print("Роль: ")
	fmt.Scan(&user.Role)

	var stack string
	fmt.Print("Стек (через запятую): ")
	fmt.Scan(&stack)
	user.Staсk = strings.Split(stack, ",")

	fmt.Print("GitHub: ")
	fmt.Scan(&user.GitHub)

	fmt.Print("Telegram: ")
	fmt.Scan(&user.Telegram)

	fmt.Print("Статус: ")
	fmt.Scan(&user.Status)
	repositories.AppendUser(user)
}
