package cmd

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListGroups godoc
// @Summary List Kafka consumer groups
// @Description Get a list of all consumer groups from the Kafka cluster.
// @Tags Kafka,ConsumerGroup
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "OK: Map of consumer group names to their respective protocol types."
// @Failure 400 {string} string "Bad Request: Error message if unable to list groups."
// @Router /groups [get]
func ListGroups(admin sarama.ClusterAdmin) func(c *gin.Context) {
	return func(c *gin.Context) {
		groups, err := admin.ListConsumerGroups()
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusOK, groups)
	}
}
