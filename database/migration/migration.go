package migration

import (
	"simpel-gateway/database"
	"simpel-gateway/internal/model"
)

var tables = []interface{}{
	&model.Server{},
	&model.ApkSetting{},
}

func Migrate() {
	conn := database.GetConnection() // Get db connection
	conn.AutoMigrate(tables...)      // migrate the tables
}
