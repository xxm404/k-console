package cmd

import (
    "net/http"
    "github.com/IBM/sarama"
    "github.com/gin-gonic/gin"
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

// ListTopics godoc
// @Summary List Kafka topics
// @Description get list of all Kafka topics from the cluster
// @Tags Kafka,Topic
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]TopicDetail "List of Kafka topics"
// @Failure 400 {string} string "Bad Request when listing topics fails"
// @Router /topics [get]
func ListTopic(admin sarama.ClusterAdmin) func (c *gin.Context) {
    return func(c *gin.Context) {
        topics, err := admin.ListTopics()
        if err != nil {
            c.String(http.StatusBadRequest, err.Error())
        }
        c.JSON(http.StatusOK, topics)
    }
}

