package db

import (
	"github.com/NatSilprasert/ticket_booking_app/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Event{},
		&models.Ticket{},
	)
}