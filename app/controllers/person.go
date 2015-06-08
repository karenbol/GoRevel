package controllers

import "github.com/revel/revel"

type Person struct {
	*revel.Controller
}

func (c Person) Store(name string, lastname string) revel.Result {

	c.Validation.Required(name).Message("Your name is required!")
    c.Validation.Required(lastname).Message("Your lastname is required!")
    
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(Person.Create)
    }
    return c.Render()
}

//ingresar datos de una persona
func (c Person) Create() revel.Result {
	return c.Render()
}

//ingresar datos de una persona
func (c Person) Show() revel.Result {
    return c.Render()
}

