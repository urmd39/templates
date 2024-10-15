package infrastructure

import (
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"

	"log"
)

var db *gorm.DB

func init() {
	dbOpen, err := OpenConnection()
	if err != nil {
		log.Printf("Not connect to database\n")
		log.Panic(err)
	}
	db = dbOpen
}

// OpenConnection open one session
func OpenConnection() (*gorm.DB, error) {
	connectSQL := "host=" + DBPostgresHost + " port= " + DBPostgresPort + " user=" + DBPostgresUsername + " dbname= " + DBPostgresName + " password = " + DBPostgresPassword + " sslmode=" + DBPostgresSSLMode
	db, err := gorm.Open(postgres.Open(connectSQL), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		ErrLog.Println("connection error, connection string: ", connectSQL)
		ErrLog.Println(err)
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		ErrLog.Println(err)
		return nil, err
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(20)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return db, nil
}

// GetDB get data session
func GetDB() *gorm.DB {
	return db
}

// InitDatabase init tables in database
func InitDatabase() error {
	err := db.AutoMigrate()

	return err
}
