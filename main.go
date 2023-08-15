package main

import (
    "os"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/IBM/sarama"
)

func main() {
    kafkaString := os.Getenv("KAFKA_STRING")
    config := sarama.NewConfig()
    admin, err := sarama.NewClusterAdmin([]string{kafkaString}, config)
    if err != nil {
        log.Fatal("Failed to connect to Kafka ", err)
    }
    defer admin.Close()

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
    r.GET("/topics", ListTopic(admin))
    r.GET("/groups", ListGroups(admin))

    r.Run(":7777")
}


