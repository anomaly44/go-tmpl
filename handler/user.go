package handler

import (
	"fmt"
	"strings"

	"github.com/anomaly44/go-tmpl/model"
	"github.com/anomaly44/go-tmpl/view/user"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	auth_sessions_key string = "authenticate-sessions"
	auth_key          string = "authenticated"
	user_id_key       string = "user_id"
	username_key      string = "username"
	tzone_key         string = "time_zone"
)

type UserHandler struct{}

// Function to check if a name exists in a slice.
func userExists(user string, slice []string) bool {
	u := strings.ToLower(user)
	for _, v := range slice {
		if v == u {
			return true
		}
	}
	return false
}

var users = []string{"rob@fixtrack.be", "test@user.be"}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: c.Get(username_key).(string),
	}
	// print username userid en tzone from context
	fmt.Printf(
		"username is %s, userid is %d and tzone is %s \n",
		c.Get(username_key),
		c.Get(user_id_key),
		c.Get(tzone_key),
	)
	return render(c, user.Show(u))
}

func (h UserHandler) HandleLoginShow(c echo.Context) error {
	return render(c, user.Login(model.NewLoginFormData()))
}

func (h UserHandler) HandleLoginSubmit(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	// obtaining the time zone from the POST request of the login form
	tzone := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tzone = c.Request().Header["X-Timezone"][0]
	}

	formData := model.LoginFormData{
		Values: map[string]string{"email": email, "password": password},
	}
	// print formdata
	fmt.Printf("formdata is %v \n", formData)
	if userExists(email, users) {
		// Get Session and setting Cookies
		sess, _ := session.Get(auth_sessions_key, c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600, // in seconds
			HttpOnly: true,
		}
		// print timezone
		fmt.Printf("timezone is %s \n", tzone)
		// Set user as authenticated, their username,
		// their ID and the client's time zone
		sess.Values = map[interface{}]interface{}{
			auth_key: true,
			// user_id_key:  user.ID,
			user_id_key:  777,
			username_key: email,
			tzone_key:    tzone,
		}
		sess.Save(c.Request(), c.Response())

		// redirect to user page
		return c.Redirect(302, "/users")
	}
	fmt.Printf("email is %s and password is %s \n", email, password)
	return render(c, user.Login(formData))
}

func (h UserHandler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(auth_sessions_key, c)
		// if sess.Values[auth_key] == nil {
		// 	return c.Redirect(302, "/")
		// }
		if auth, ok := sess.Values[auth_key].(bool); !ok || !auth {
			fmt.Println(ok, auth)

			return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Please provide valid credentials")
		}

		if userId, ok := sess.Values[user_id_key].(int); ok && userId != 0 {
			c.Set(user_id_key, userId) // set the user_id in the context
		}

		if username, ok := sess.Values[username_key].(string); ok && len(username) != 0 {
			c.Set(username_key, username) // set the username in the context
		}

		if tzone, ok := sess.Values[tzone_key].(string); ok && len(tzone) != 0 {
			c.Set(tzone_key, tzone) // set the client's time zone in the context
		}

		return next(c)
	}
}
