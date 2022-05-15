package migrations

import (
	"github.com/joshua-ather/sv_be/app/models"
)

func (m Migrate) CreatePost() {
	m.db.Debug().AutoMigrate(&models.Post{}) //database migration
}
