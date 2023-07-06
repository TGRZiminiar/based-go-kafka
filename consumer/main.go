package main

import (
	"consumer/controllers"
	"consumer/repositories"
	"consumer/services"
	"context"
	"events"
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	db := initDatabase()

	consumer, err := sarama.NewConsumerGroup(viper.GetStringSlice("kafka.servers"), viper.GetString("kafka.group"), nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)
	userEventHandler := services.NewUserEventHandler(userController)
	userConsumerHandler := services.NewConsumerHandler(userEventHandler)

	fmt.Println("Account Consumer started")
	for {
		consumer.Consume(context.Background(), events.Topics, userConsumerHandler)
	}

}

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

// consumer, err := sarama.NewConsumerGroup([]string{"127.0.0.1:9092"}, "accountConsumer", nil)
// kafka-consumer-groups.sh --list --bootstrap-server localhost:9092
