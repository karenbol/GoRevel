package models


import (
    "github.com/revel/revel"
)

type Person struct {
    Id              int  
    Name,Lastname   string  
}

func (b *Person) Validate(v *revel.Validation) {

    v.Check(b.Name,
        revel.ValidRequired(),
        revel.ValidMaxSize(25))

    v.Check(b.Lastname,
        revel.ValidRequired(),
        revel.ValidMaxSize(25))   
}