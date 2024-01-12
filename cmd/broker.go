package cmd

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strconv"
)

type BrokerDetail struct {
	ID   int32  `json:"id"`
	Host string `json:"host"`
	Port int    `json:"port"`
	Rack string `json:"rack,omitempty"`
}

// ListBrokers godoc
// @Summary List brokers
// @Description get list of Kafka brokers from the sarama client
// @Tags Kafka,Broker
// @Accept  json
// @Produce  json
// @Success 200 {array} BrokerDetail "List of broker details"
// @Failure 400 {object} string "Bad request when the host and port cannot be parsed"
// @Router /brokers [get]
func ListBrokers(client sarama.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		brokers := client.Brokers()
		var brokerList []BrokerDetail
		for _, b := range brokers {
			addr := b.Addr()
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				host = addr
				port = "0"
			}
			intPort, err := strconv.Atoi(port)
			id := b.ID()
			rack := b.Rack()
			brokerList = append(brokerList, BrokerDetail{
				ID:   id,
				Host: host,
				Port: intPort,
				Rack: rack,
			})
		}
		c.JSON(http.StatusOK, brokerList)
	}
}
