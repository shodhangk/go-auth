package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {

	//Load environmenatal variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	un := os.Getenv("databaseUser")
	pwd := os.Getenv("databasePassword")
	dbName := os.Getenv("databaseName")
	dbHost := os.Getenv("databaseHost")
	dbPort := os.Getenv("databasePort")

	//Define DB connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", un, pwd, dbHost, dbPort, dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	//connect to DB URI
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger})
	fmt.Println()
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}

	fmt.Println("Init Database")

	fmt.Println(DB)

	// Migrate the schema
	DB.AutoMigrate(
		&User{},
		&Item{},
		&Cart{},
		&Order{},
		&OrderItem{},
		&CartItem{})
	DB.SetupJoinTable(&Cart{}, "Items", &CartItem{})
	DB.SetupJoinTable(&Order{}, "Items", &OrderItem{})
	// DB.Create(&Item{Name: "Tomatio", Price: 22, Quantity: 5})
	// DB.Create(&Item{Name: "Ladies Finger", Price: 100, Quantity: 5})
	// DB.Create(&Item{Name: "Carrot", Price: 50, Quantity: 5})
	// DB.Create(&Item{Name: "Cabbage", Price: 50, Quantity: 2})

}
