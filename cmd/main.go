package main

import (
	"HOTA/internal/models"
	"HOTA/internal/repositories"
	"fmt"

	"strings"
)

func main() {

	for {
		fmt.Println()
		models.PrintMeny()
        fmt.Println()
		var action int
		fmt.Scan(&action)
		fmt.Println()

		switch action {
		case 1: // Добавить пользователя
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

			fmt.Println("Пользователь создан")
			fmt.Println()
		case 2: // Поиск по нику
			var nickname string
			fmt.Println("Введите Ник: ")
			fmt.Scan(&nickname)

			user := repositories.SersheNickname(nickname)
			if user == nil {
				fmt.Println("Такого пользователя не существует")
				continue
			}
			models.PrintUser(*user)

		case 3: // Поиск по ID

			var id int
			fmt.Println("Ведите ID")
			fmt.Scan(&id)

			user := repositories.SersheID(id)

			if user == nil {
				fmt.Printf("Такого пользователя с id: %d не существует", id)
				continue
			}
			models.PrintUser(*user)

		case 4: // Поиск по стеку
			var stack string
			fmt.Println("Введите стек: ")
			fmt.Scan(&stack)

			users := repositories.FindByStack(stack)

			if len(users) == 0 {
				fmt.Println("Ничего не найдено")
			}
			for _, user := range users {
				fmt.Printf("ID: %d,\nНик: %s,\nРоль: %s,\nСтек: %s\n", user.ID, user.Nickname, user.Role, strings.Join(user.Staсk, ","))

			}

		case 5: // Вывести всех пользователей
			users := repositories.GetAllUsers()
			for _, user := range users {
				// strings.Join(user.Staсk, ",") //Преобразовываем из слайса в строку
				fmt.Printf("ID: %d,\nНик: %s,\nРоль: %s,\nСтек: %s, \nГитхаб: %s, \nTG: %s, \nСтатус: %s", user.ID, user.Nickname, user.Role, strings.Join(user.Staсk, ","), user.GitHub, user.Telegram, user.Status)
			}

		case 6: // Обновление данных

			var id int
			fmt.Println("Введите ID")
			fmt.Scan(&id)

			user := repositories.SersheID(id)
			if user == nil {
				fmt.Printf("Такого пользователя с id: %d не существует", id)
				continue
			}
			// РЕДАКТИРОВАНИЯ ПОЛЬЗОВАТЕЛЯ

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
			fmt.Printf("Введите Telegram [%s]\n", user.Status)
			fmt.Scan(&stats)

			user.Nickname = nik
			user.Role = rol
			user.GitHub = gh
			user.Telegram = tg
			user.Status = stats

			ok := repositories.UpdateUser(*user)

			if ok == false {
				fmt.Println("Данного пользователя не удалось обновить")
			} else {
				fmt.Println("Пользовтель успешно обновлен")
			}

		case 7: // Удаление пользователя
			var id int
			fmt.Println("Введите ID")
			fmt.Scan(&id)

			ok := repositories.DeleteUser(id)

			if ok == false {
				fmt.Printf("Такого пользователя с id: %d не существует", id)
			} else {
				fmt.Println("Пользователь успешно удален")
			}
		case 8: //статистика
			count := repositories.CountUser()
			fmt.Printf("Количество зарегистрнированных пользователей: %d\n", count)

			stastRoll := repositories.CountByRole()
			fmt.Println("Статистика по ролям:")

			for role, count := range stastRoll {
				fmt.Printf("%s: %d\n", role, count)
			}

		default:
			fmt.Printf("Данного выбора: %d не существует", action)
		}

	}

}
