package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type MutationOperation struct {
	Sequence     string `json:"sequence,required"`
	Operation    string `json:"operation,required"`
	OperationKey int    `json:"operation_key,omitempty"`
}

type MutationResult struct {
	Id     uuid.UUID `json:"id"`
	Result string    `json:"result"`
}

var (
	requestValidator = validator.New()
	operationsThatRequireKey = map[string]bool{
		"encode_shift_left":  true,
		"decode_shift_left":  true,
		"encode_shift_right": true,
		"decode_shift_right": true,
	}
)

func validateMutationOperationProperties(operation MutationOperation) bool {
	basicCheck := operation.Sequence != "" && operation.Operation != ""
	if basicCheck {
		if _, ok := operationsThatRequireKey[operation.Operation]; ok {
			return operation.OperationKey != 0
		} else {
			return true
		}
	} else {
		return false
	}
}

func Mutate(c *gin.Context) {
	var requestBody MutationOperation
	// check that our request data has been properly bound to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid data contained in request body.",
			},
		)
		return
	}

	validationErr := requestValidator.Struct(requestBody)
	if validationErr != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid data contained in request body.",
			},
		)
		return
	}

	// validate the bound properties
	if !validateMutationOperationProperties(requestBody) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid data contained in request body.",
			},
		)
		return
	}

	// perform the operation
	switch requestBody.Operation {
	case "count":
		c.JSON(
			http.StatusOK,
			MutationResult{
				Id:     uuid.New(),
				Result: strconv.Itoa(Count(requestBody.Sequence)),
			},
		)
		return
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
		return
	default:
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": fmt.Sprintf("Invalid operation type specified. Operation %s is not included in this API.", requestBody.Operation),
			},
		)
		return
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
