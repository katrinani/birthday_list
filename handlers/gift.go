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

// CreateGift godoc
// @Summary Создать новый подарок
// @Description Добавляет новый подарок в список желаний
// @Tags gifts
// @Accept  json
// @Produce json
// @Param   gift body models.Gift true "Данные подарка"
// @Success 201 {object} models.Gift
// @Failure 400 {object} map[string]string
// @Router /gifts [post]
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

// ReserveGift godoc
// @Summary Зарезервировать подарок
// @Description Позволяет пользователю зарезервировать подарок
// @Tags gifts
// @Accept  json
// @Produce json
// @Param   id path int true "ID подарка"
// @Param   user_id formData int true "ID пользователя, который резервирует подарок"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /gifts/{id}/reserve [put]
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
