package handler

import (
	"fmt"

	"github.com/anomaly44/go-tmpl/model"
	"github.com/anomaly44/go-tmpl/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "Rob@fixtrack.be",
	}
	return render(c, user.Show(u))
}

func (h UserHandler) HandleLoginShow(c echo.Context) error {
	return render(c, user.Login(model.NewLoginFormData()))
}

func (h UserHandler) HandleLoginSubmit(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	formData := model.LoginFormData{
		Values: map[string]string{"email": email, "password": password},
	}
	// print formdata
	fmt.Printf("formdata is %v \n", formData)
	if email == "rob@fixtrack.be" {
		// redirect to user page
		return c.Redirect(302, "/user")
	}
	fmt.Printf("email is %s and password is %s \n", email, password)
	return render(c, user.Login(formData))
}
