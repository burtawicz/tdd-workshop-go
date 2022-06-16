package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MutationOperation struct {
	Sequence     string `json:"sequence"`
	Operation    string `json:"operation"`
	OperationKey int    `json:"operation_key"`
}

type MutationResult struct {
	Id     uuid.UUID `json:"id"`
	Result string    `json:"result"`
}

func Mutate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "We're currently working on our implementation. Stay tuned.",
	})
}

func Find(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "We're currently working on our implementation. Stay tuned.",
	})
}
