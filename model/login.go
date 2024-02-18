package model

type LoginFormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewLoginFormData() LoginFormData {
	return LoginFormData{
		Errors: map[string]string{},
		Values: map[string]string{},
	}
}
