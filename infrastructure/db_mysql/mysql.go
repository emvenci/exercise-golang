package mysql

import (
	"exercise-backend/config"
	"log"
	"sync"

	"bitbucket.org/liamstask/goose/lib/goose"
	"github.com/jinzhu/gorm"
)

var lock = &sync.Mutex{}

var dbConn *gorm.DB

// GetMysqlConn returns mysql db connection.
// Ensures that only one connection exists (singleton pattern).
func GetMysqlConn() (db *gorm.DB, err error) {
	if dbConn == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbConn == nil {
			return startDB()
		}
	}
	return dbConn, nil
}

// startDB creates DB connection and runs migrations
func startDB() (db *gorm.DB, err error) {
	// Initialize db connection
	db, err = gorm.Open("mysql", config.Conf.DBPath)
	if err != nil {
		log.Printf("database opening error: %v\n", err)
		return
	}

	db.DB().SetMaxOpenConns(1)
	log.Println("database connection established...")

	// Create migrations config
	migrateConf := &goose.DBConf{
		MigrationsDir: config.Conf.MigrationsPath,
		Env:           "development",
		Driver: goose.DBDriver{
			Name:    "mysql",
			OpenStr: config.Conf.DBPath,
			Import:  "github.com/go-sql-driver/mysql",
			Dialect: &goose.MySqlDialect{},
		},
	}

	// Get the latest migration
	var latest int64
	latest, err = goose.GetMostRecentDBVersion(migrateConf.MigrationsDir)
	if err != nil {
		log.Printf("getting latest migration error: %v\n", err)
		return
	}

	// Run migration
	err = goose.RunMigrationsOnDb(migrateConf, migrateConf.MigrationsDir, latest, db.DB())
	if err != nil {
		log.Printf("running migrations error: %v\n", err)
		return
	}

	// Insert db connection into var and return
	dbConn = db
	return
}
