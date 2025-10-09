package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	templates := template.Must(template.ParseGlob("views/*.html"))
	return &Template{
		templates: templates,
	}
}

type work struct {
	Name string
}

func newWork(name string) work {
	return work{
		Name: name,
	}
}

type works = []work

type Data struct {
	Works works
}

func newData() Data {
	return Data{
		Works: []work{},
	}
}

type FormData struct {
	Value map[string]string
	Error map[string]string
}

func newFormData() FormData {
	return FormData{
		Value: map[string]string{},
		Error: map[string]string{},
	}
}

type Page struct {
	Data Data
	Form FormData
}

func newPage() Page {
	return Page{
		Data: newData(),
		Form: newFormData(),
	}
}

func hasWork(works works, name string) bool {
	for _, work := range works {
		if work.Name == name {
			return true
		}
	}
	return false
}

func main() {

	e := echo.New()
	e.Renderer = newTemplate()

	page := newPage()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.POST("/submitForm", func(c echo.Context) error {
		workName := c.FormValue("work")

		if hasWork(page.Data.Works, workName) {
			formData := newFormData()
			formData.Value["work"] = workName
			formData.Error["work"] = "This work already exists."

			return c.Render(422, "form", formData)
		}

		if workName == "" {
			formData := newFormData()
			formData.Error["work"] = "Work cannot be empty."

			return c.Render(422, "form", formData)
		}

		toDoWork := newWork(workName)

		page.Data.Works = append(page.Data.Works, toDoWork)

		c.Render(200, "form", newFormData())
		return c.Render(200, "oob-work", toDoWork)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
