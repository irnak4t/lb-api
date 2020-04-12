package mysql

import (
	"github.com/irnak4t/lb-api/db/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Open() (*gorm.DB, error) {
	cfg := config.Get()
	args := cfg.Db.User + ":" + cfg.Db.Password + "@/" + cfg.Db.Database + "?parseTime=true"

	return gorm.Open("mysql", args)
}
