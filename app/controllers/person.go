package controllers

import (
   
    "github.com/revel/revel"
    "encoding/json"
    "expo/app/models"
)

type Person struct {
    GorpController
}

func (c Person) parseBidItem() (models.Person, error) {
    biditem := models.Person{}
    err := json.NewDecoder(c.Request.Body).Decode(&biditem)
    return biditem, err
}

func (c Person) Add() revel.Result {
    if biditem, err := c.parseBidItem(); err != nil {
        return c.RenderText("Unable to parse the BidItem from JSON.")
    } else {
        // Validate the model
        biditem.Validate(c.Validation)
        if c.Validation.HasErrors() {
            // Do something better here!
            return c.RenderText("You have error in your BidItem.")
        } else {
            if err := c.Txn.Insert(&biditem); err != nil {
                return c.RenderText(
                    "Error inserting record into database!")
            } else {
                return c.RenderJson(biditem)
            }
        }
    }
}

func (c Person) Get(id int) revel.Result {
    biditem := new(models.Person)
    err := c.Txn.SelectOne(biditem, 
        `SELECT * FROM personas WHERE id = ?`, id)
    if err != nil {
        return c.RenderText("Error.  Item probably doesn't exist.")
    }
    return c.RenderJson(biditem)
}

func (c Person) List() revel.Result {
    lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
    limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
    biditems, err := c.Txn.Select(models.Person{}, 
        `SELECT * FROM BidItem WHERE Id > ? LIMIT ?`, lastId, limit)
    if err != nil {
        return c.RenderText(
            "Error trying to get records from DB.")
    }
    return c.RenderJson(biditems)
}

func (c Person) Update(id int) revel.Result {
    biditem, err := c.parseBidItem()
    if err != nil {
        return c.RenderText("Unable to parse the BidItem from JSON.")
    }
    // Ensure the Id is set.
    biditem.Id = id
    success, err := c.Txn.Update(&biditem)
    if err != nil || success == 0 {
        return c.RenderText("Unable to update bid item.")
    }
    return c.RenderText("Updated %v", id)
}

func (c Person) Delete(id int) revel.Result {
    success, err := c.Txn.Delete(&models.Person{Id: id})
    if err != nil || success == 0 {
        return c.RenderText("Failed to remove BidItem")
    }
    return c.RenderText("Deleted %v", id)
}
/*package controllers

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
}*/

