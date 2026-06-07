package handlers

import (
	"HOTA/internal/repositories"
	"errors"
	"fmt"
)

func UpdateUser() error {
	var id int
	fmt.Println("Введите ID")
	fmt.Scan(&id)

	user := repositories.SersheID(id)
	if user == nil {
		return fmt.Errorf("Такого пользователя с id: %d не существует", id)
	}

	var nik string
	fmt.Printf("Введите имя [%s]\n", user.Nickname)
	fmt.Scan(&nik)

	var rol string
	fmt.Printf("Введите новую роль [%s]\n", user.Role)
	fmt.Scan(&rol)

	var gh string
	fmt.Printf("Введите GitHub [%s]\n", user.GitHub)
	fmt.Scan(&gh)

	var tg string
	fmt.Printf("Введите Telegram [%s]\n", user.Telegram)
	fmt.Scan(&tg)

	var stats string
	fmt.Printf("Введите статус [%s]\n", user.Status)
	fmt.Scan(&stats)

	user.Nickname = nik
	user.Role = rol
	user.GitHub = gh
	user.Telegram = tg
	user.Status = stats

	ok := repositories.UpdateUser(*user)

	if ok == false {
		return errors.New("Данного пользователя не удалось обновить")
	}

	return nil
}
