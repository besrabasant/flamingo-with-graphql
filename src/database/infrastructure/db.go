package infrastructure

import "gorm.io/gorm"

type DBService struct {
	db *gorm.DB
}

func (dbService *DBService) SetConnection(dbConnection *gorm.DB) *DBService {
	dbService.db = dbConnection
	return dbService
}

func (dbService *DBService) Connection() *gorm.DB {
	return dbService.db
}
