package main

import (
    "net/http"
    "github.com/IBM/sarama"
    "github.com/gin-gonic/gin"
)

func ListTopic(admin sarama.ClusterAdmin) func (c *gin.Context) {
    return func(c *gin.Context) {
        topics, err := admin.ListTopics()
        if err != nil {
            c.String(http.StatusBadRequest, err.Error())
        }
        c.JSON(http.StatusOK, topics)
    }
}

func ListConsumerGroups(admin sarama.ClusterAdmin) func (c *gin.Context) {
    return func(c *gin.Context) {
        consumerGroups, err := admin.ListConsumerGroups()
        if err != nil {
            c.String(http.StatusBadRequest, err.Error())
        }
        c.JSON(http.StatusOK, consumerGroups)
    }
}

