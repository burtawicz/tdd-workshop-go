package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type MutationOperation struct {
	Sequence     string `json:"sequence"`
	Operation    string `json:"operation,required"`
	OperationKey int    `json:"operation_key,omitempty"`
}

type MutationResult struct {
	Id     uuid.UUID `json:"id"`
	Result string    `json:"result"`
}

func Mutate(c *gin.Context) {
	var requestBody MutationOperation
	// check that our request data has been properly bound to the struct
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid data contained in request body.",
			},
		)
		return
	}

	switch requestBody.Operation {
	case "count":
		c.JSON(
			http.StatusOK,
			MutationResult{
				Id:     uuid.New(),
				Result: strconv.Itoa(Count(requestBody.Sequence)),
			},
		)
	case "reverse":
		fallthrough
	case "separate":
		fallthrough
	case "encode_base64":
		fallthrough
	case "decode_base64":
		fallthrough
	case "encode_rot13":
		fallthrough
	case "decode_rot13":
		fallthrough
	case "encode_shift_left":
		fallthrough
	case "decode_shift_left":
		fallthrough
	case "encode_shift_right":
		fallthrough
	case "decode_shift_right":
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Not implemented.",
			},
		)
	default:
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf("Invalid operation type specified. Operation %s is not included in this API.", requestBody.Operation),
			})
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "We're currently working on our implementation. Stay tuned.",
		},
	)
}

func Find(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "We're currently working on our implementation. Stay tuned.",
		},
	)
}
