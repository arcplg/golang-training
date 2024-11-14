package storage

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ngocthanh06/chatapp/internal/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DatabaseService interface {
	AutoMigration()
}

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Load env is error: %s", err)
	}

	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port = os.Getenv("DB_PORT")
	host = os.Getenv("DB_HOST")
	schema = os.Getenv("DB_SCHEMA")

}

type service struct {
	db *gorm.DB
}

func setDB(db *gorm.DB) *service {
	return &service{
		db: db,
	}
}

func GetDB() *gorm.DB {
	return dbInstance.db
}

var (
	database   string
	password   string
	username   string
	port       string
	host       string
	schema     string
	dbInstance *service
)

func (s service) AutoMigration() {
	if err := s.db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.Message{},
		&models.RoomMember{},
	); err != nil {
		log.Fatalf("Migrate error: %v", err)
	}
}

func New() DatabaseService {
	if dbInstance != nil {
		return dbInstance
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln("Connection database not working!")
	}

	log.Println("Connection database is successfully!")

	dbInstance = setDB(db)

	// run auto migration
	dbInstance.AutoMigration()

	return dbInstance
}
