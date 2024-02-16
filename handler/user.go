package handler

import (
	"github.com/labstack/echo/v4"
	// import my user package
	"github.com/anomaly44/go-tmpl/model"
	"github.com/anomaly44/go-tmpl/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "Rob@fixtrack.be",
	}
	return render(c, user.Show(u))
}
