package handlers

import (
	"baseToDo/db"
	"baseToDo/models"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// CreateGift создает новый подарок.
func CreateGift(c echo.Context) error {
	gift := new(models.Gift)
	if err := c.Bind(gift); err != nil {
		return err
	}

	// подключение к базе
	base, _ := db.Connect()
	defer func(base *sql.DB) {
		err := base.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(base)

	_, err := base.Exec(
		"INSERT INTO gifts (user_id, name, description, photo_url) VALUES ($1, $2, $3, $4)",
		gift.UserID, gift.Name,
		gift.Description, gift.PhotoURL,
	)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, gift)
}

// ReserveGift помечает конкретный подарок занятым определенным пользователем.
func ReserveGift(c echo.Context) error {
	giftID, _ := strconv.Atoi(c.Param("id"))
	userID, _ := strconv.Atoi(c.FormValue("user_id"))

	base, _ := db.Connect()
	defer func(base *sql.DB) {
		err := base.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(base)

	_, err := base.Exec("UPDATE gifts SET reserved_by = $1 WHERE id = $2", userID, giftID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Gift reserved"})
}
