package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MutationOperation struct {
	Sequence         string `json:"sequence"`
	SequenceDataType string `json:"sequence_data_type"`
	Operation        string `json:"operation"`
	OperationKey     int    `json:"operation_key"`
}

type MutationResult struct {
	Id             uuid.UUID `json:"id,omitempty"`
	Result         string    `json:"result,omitempty"`
	ResultDataType string    `json:"result_data_type,omitempty"`
}

func Mutate(c *gin.Context) {

}

func Find(c *gin.Context) {

}
