package main

import (
	"fmt"
	"producer/controllers"
	"producer/middleware"
	repositories "producer/repositories"
	"producer/services"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	dns := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		viper.GetString("db.driver"),
		viper.GetString("db.UserName"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	db := initDatabase()

	userRepo := repositories.NewUserRepository(db)
	eventProducer := services.NewEventProducer(producer)
	userService := services.NewUserService(eventProducer, userRepo)
	userController := controllers.NewUserController(userService)

	app := fiber.New()

	app.Post("/register", userController.CreateUser)
	app.Post("/login", userController.Login)
	app.Post("/current-user", middleware.ParseUser, userController.CurrentUser)

	app.Listen(":5000")

}
