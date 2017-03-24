package internal

import (
	"github.com/jinzhu/gorm"
)

type Context struct {
	DB        *gorm.DB
	RawURL     string
	PrivateKey string
}
