package middleware

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"go-admin/common/config"
	"go-admin/common/global"
	"go-admin/common/middleware"
	"go-admin/tools"
)

var WithContextDb = middleware.WithContextDb

func getGormFromDb(driver string, db *sql.DB, config *gorm.Config) (*gorm.DB, error) {
	switch driver {
	case "mysql":
		return gorm.Open(mysql.New(mysql.Config{Conn: db}), config)
	case "postgres":
		return gorm.Open(postgres.New(postgres.Config{Conn: db}), config)
	default:
		return nil, errors.New("not support this db driver")
	}
}

func GetGormFromConfig(cfg config.Conf) map[string]*gorm.DB {
	gormDB := make(map[string]*gorm.DB)
	if cfg.GetSaas() {
		var err error
		for k, v := range cfg.GetDbs() {
			gormDB[k], err = getGormFromDb(v.Driver, v.DB, &gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			})
			if err != nil {
				global.Logger.Fatal(tools.Red(k+" connect error :"), err)
			}
		}
		return gormDB
	}
	c := cfg.GetDb()
	db, err := getGormFromDb(c.Driver, c.DB, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		global.Logger.Fatal(tools.Red(c.Driver+" connect error :"), err)
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	gormDB["*"] = db
	return gormDB
}
