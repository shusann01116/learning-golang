package morestructs

type Person struct {
	// <field> <type> `<key:"meta">`
	Name string `json:"name" validate:"required" gorm:"column:name"`
}
