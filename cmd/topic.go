package cmd

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TopicDetail struct {
	// NumPartitions contains the number of partitions to create in the topic, or
	// -1 if we are either specifying a manual partition assignment or using the
	// default partitions.
	NumPartitions int32
	// ReplicationFactor contains the number of replicas to create for each
	// partition in the topic, or -1 if we are either specifying a manual
	// partition assignment or using the default replication factor.
	ReplicationFactor int16
	// ReplicaAssignment contains the manual partition assignment, or the empty
	// array if we are using automatic assignment.
	ReplicaAssignment map[int32][]int32
	// ConfigEntries contains the custom topic configurations to set.
	ConfigEntries map[string]*string
}

type PubBody struct {
	Topic   string `json:"topic" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// ListTopics godoc
// @Summary List Kafka topics
// @Description get list of all Kafka topics from the cluster
// @Tags Kafka,Topic
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]TopicDetail "List of Kafka topics"
// @Failure 400 {string} string "Bad Request when listing topics fails"
// @Router /topics [get]
func ListTopic(admin sarama.ClusterAdmin) func(c *gin.Context) {
	return func(c *gin.Context) {
		topics, err := admin.ListTopics()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, topics)
	}
}

// PubToTopic publishes message to a specified Kafka topic.
// @Summary Publish message to Kafka topic
// @Description Publishes a message to the specified Kafka topic using Sarama SyncProducer.
// @Tags Kafka,Topic
// @Accept  json
// @Produce  json
// @Param   pubBody  body      PubBody  true  "Publish Body"
// @Success 200 {string} string "Message is stored in topic(partition/offset)"
// @Failure 400 {object} map[string]interface{} "error": "Error description"
// @Router /publish [post]
func PubToTopic(producer sarama.SyncProducer) func(c *gin.Context) {
	return func(c *gin.Context) {
		var pubBody PubBody

		if err := c.ShouldBindJSON(&pubBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msg := &sarama.ProducerMessage{
			Topic: pubBody.Topic,
			Value: sarama.StringEncoder(pubBody.Content),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Message is stored in topic(%s)/partition(%d)/offset(%d)", msg.Topic, partition, offset)})
	}
}
