package interfaces

import "gorm.io/gorm"

type DB interface {
	Connection() *gorm.DB
}
