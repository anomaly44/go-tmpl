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

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Errors: map[string]string{},
		Values: map[string]string{},
	}
}

type Recipe struct {
	Name string
	Id   string
}

type Data struct {
	Recipes []Recipe
}

func NewData() *Data {
	return &Data{
		Recipes: []Recipe{
			{
				Name: "Mayo",
				Id:   "5555",
			},
			{
				Name: "Ketchup",
				Id:   "6666",
			},
		},
	}
}

type PageData struct {
	Data Data
	Form FormData
}

func NewRecipe(name string, id string) Recipe {
	return Recipe{
		Name: name,
		Id:   id,
	}
}

func NewPageData() PageData {
	return PageData{
		Data: *NewData(),
		Form: NewFormData(),
	}
}
