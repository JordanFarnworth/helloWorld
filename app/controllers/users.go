package controllers

import (
	"github.com/revel/revel"
	"helloWorld/app/models"
)

type Users struct {
	*revel.Controller
}

var users = []models.User{
	models.User{1, "Jordan Farnworth", "jordanfarnworth", "jordanfarn23@gmail.com", "fuckyou"},
	models.User{2, "Chelsea Farnworth", "chel_sea", "cmf22@gmail.com", "jordanishot"},
	models.User{3, "American Dream", "Lager", "email@email", "password"},
}

func (c Users) List() revel.Result {
	return c.RenderJSON(users)
}

func (c Users) Show(userID int) revel.Result {
	var selectedUser models.User

	for _, user := range users {
		if user.ID == userID {
			selectedUser = user
		}
	}

	if selectedUser.ID == 0 {
		return c.NotFound("Could not find user")
	}

	return c.RenderJSON(selectedUser)
}
