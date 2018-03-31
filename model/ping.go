package model

import (
	"template/engine"

	"github.com/jinzhu/gorm"
)

type Ping struct {
	gorm.Model
	Log string
}

func init() {
	engine.GetORM().AutoMigrate(&Ping{})
}
