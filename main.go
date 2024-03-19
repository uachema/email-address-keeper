package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type templates struct {
	templates *template.Template
}

func (t *templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplates() *templates {
	return &templates{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
}

type contact struct {
	ID    uint
	Name  string
	Email string
}

type contacts struct {
	contacts []contact
}

func (c *contacts) createNewContact(name, email string) *contact {
	contact := &contact{
		ID:    uint(len(c.contacts) + 1),
		Name:  name,
		Email: email,
	}
	c.contacts = append(c.contacts, *contact)
	return contact
}

func newContacts() *contacts {
	contacts := &contacts{}
	contacts.createNewContact("Umair Ali", "umair@gmail.com")
	contacts.createNewContact("Azlan Ali", "azlan@gmail.com")
	return contacts
}

func main() {
	contacts := newContacts()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplates()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", contacts.contacts)
	})
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
