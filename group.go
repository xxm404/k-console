package main

import (
    "net/http"
    "github.com/IBM/sarama"
    "github.com/gin-gonic/gin"
)

func ListGroups(admin sarama.ClusterAdmin) func (c *gin.Context) {
    return func(c *gin.Context) {
        groups, err := admin.ListConsumerGroups()
        if err != nil {
            c.String(http.StatusBadRequest, err.Error())
        }
        c.JSON(http.StatusOK, groups)
    }
}

