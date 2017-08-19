package controllers

import (
	"github.com/revel/revel"
	"helloWorld/app/models"
)

type App struct {
	GormController
}

func (c App) Index() revel.Result {
    user := models.User{Name: "Jinzhup"}
    c.Tx.NewRecord(user)
    c.Tx.Create(&user)
    return c.RenderJSON(user)
}
