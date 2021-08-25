package model

import "gorm.io/gorm"

// Entity is the base model entity of your application. Feel free to extend/define more entities
// TODO: rename and have fun
type Entity struct {
	gorm.Model
}
