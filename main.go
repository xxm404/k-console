package main

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	sf "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"github.com/xxm404/k-console/cmd"
	docs "github.com/xxm404/k-console/docs"
	"log"
	"net/http"
	"os"
)

func main() {
	kafkaString := os.Getenv("KAFKA_STRING")
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{kafkaString}, config)
	if err != nil {
		log.Fatal("Failed to start Sarama producer ", err)
	}

	admin, err := sarama.NewClusterAdmin([]string{kafkaString}, config)
	if err != nil {
		log.Fatal("Failed to connect to Kafka admin ", err)
	}
	client, err := sarama.NewClient([]string{kafkaString}, config)
	if err != nil {
		log.Fatal("Failed to connect to Kafka client ", err)
	}
	defer admin.Close()
	defer client.Close()
	defer producer.Close()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")

	v1.GET("/topics", cmd.ListTopic(admin))
	v1.GET("/groups", cmd.ListGroups(admin))
	v1.GET("/brokers", cmd.ListBrokers(client))
	v1.POST("/publish", cmd.PubToTopic(producer))

	r.GET("/swagger/*any", gs.WrapHandler(sf.Handler))
	r.Run(":7777")
}
