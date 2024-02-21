package handler

import (
	"net/http"

	"github.com/anomaly44/go-tmpl/model"
	"github.com/anomaly44/go-tmpl/view/recipe"
	"github.com/labstack/echo/v4"
)

var page = model.NewPageData()

type RecipeHandler struct{}

func recipeExists(recipes []model.Recipe, id string) bool {
	for _, c := range recipes {
		if c.Id == id {
			return true
		}
	}
	return false
}

func (h RecipeHandler) HandleRecipeShow(c echo.Context) error {
	return render(c, recipe.Show(page))
}

func (h RecipeHandler) HandleRecipeSubmit(c echo.Context) error {
	name := c.FormValue("name")
	id := c.FormValue("id")

	if recipeExists(page.Data.Recipes, id) {
		formData := model.FormData{
			Errors: map[string]string{
				"id": "ID already exists",
			},
			Values: map[string]string{
				"name": name,
				"id":   id,
			},
		}
		// Set the status code to 422 before writing the body
		// c.Response().Status = http.StatusUnprocessableEntity
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
		c.Response().WriteHeader(http.StatusUnprocessableEntity)
		return recipe.RecipeForm(formData).Render(c.Request().Context(), c.Response().Writer)
	}

	rec := model.NewRecipe(name, id)
	page.Data.Recipes = append(page.Data.Recipes, rec)

	formData := model.NewFormData()
	err := recipe.RecipeForm(formData).Render(c.Request().Context(), c.Response().Writer)
	// err := c.Render(200, "contact-form", formData)

	if err != nil {
		return err
	}

	return recipe.OobRecipe(rec).Render(c.Request().Context(), c.Response().Writer)
}
