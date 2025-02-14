package endpoints

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllPresent(c echo.Context) error {
	// получаем имя чей список хотим посмотреть
	birthdayBoy := c.QueryParam("birthday_boy")
	// узнаем кто смотрит
	username := c.QueryParam("username")

	if username == birthdayBoy {
		// отправляем сокращенный список без Booked и Giver
		if _, exists := presents[birthdayBoy]; !exists {
			return c.String(http.StatusOK, "Нашли, но сокращ")
		} else {
			return c.String(http.StatusNotFound, "Не найдено ни одного подарка")
		}
	} else {
		// отправляем полный тк смотрит не именинник
		return c.String(http.StatusOK, "Нашли полный")
	}
}
