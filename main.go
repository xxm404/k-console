package main

import (
    "os"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/IBM/sarama"
    gs "github.com/swaggo/gin-swagger"
    sf "github.com/swaggo/files"
    "github.com/xxm404/k-console/cmd"
    docs "github.com/xxm404/k-console/docs"
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

    docs.SwaggerInfo.BasePath = "/api/v1"
    v1 := r.Group("/api/v1")

    v1.GET("/topics", cmd.ListTopic(admin))
    v1.GET("/groups", cmd.ListGroups(admin))

    r.GET("/swagger/*any", gs.WrapHandler(sf.Handler))
    r.Run(":7777")
}


