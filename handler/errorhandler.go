package handler

import (
	"errors"
	"fmt"
	"github.com/feryadialoi/go-error-handling-practice/errorutil"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

func HandleIfGetAccountError(ctx *gin.Context, err error) bool {
	if err != nil {
		if errors.Is(err, errorutil.EntityNotFoundError) {
			ctx.JSON(http.StatusNotFound, ResponseError(err.Error()))
			return true
		}
		ctx.JSON(http.StatusInternalServerError, ResponseError(err.Error()))
		return true
	}
	return false
}

func HandleIfTopUpError(ctx *gin.Context, err error) bool {
	if err != nil {
		if errors.Is(err, errorutil.EntityNotFoundError) {
			ctx.JSON(http.StatusNotFound, ResponseError(err.Error()))
			return true
		}
		ctx.JSON(http.StatusInternalServerError, ResponseError(err.Error()))
		return true
	}
	return false
}

func HandleIfTransferError(ctx *gin.Context, err error) bool {
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, errorutil.EntityNotFoundError) {
			ctx.JSON(http.StatusNotFound, ResponseError(err.Error()))
			return true
		}
		if errors.Is(err, errorutil.BalanceInsufficientError) {
			ctx.JSON(http.StatusBadRequest, ResponseError(err.Error()))
			return true
		}
		ctx.JSON(http.StatusInternalServerError, ResponseError(err.Error()))
		return true
	}
	return false
}

func HandleIfShouldBindJSONError(ctx *gin.Context, err error, request interface{}) bool {
	if err != nil {
		if validateErr, ok := err.(validator.ValidationErrors); ok {
			errorMap := ValidationErrorMessage(validateErr, request)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMap})
			return true
		}
		if ginErr, ok := err.(gin.Error); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": ginErr.Error()})
			return true
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}

func ValidationErrorMessage(validateErr validator.ValidationErrors, request interface{}) map[string]string {
	var errorMap = map[string]string{}
	for _, fieldError := range validateErr {
		structField, ok := reflect.TypeOf(request).Elem().FieldByName(fieldError.StructField())
		if !ok {
			errorMap[fieldError.Field()] = fieldError.Error()
			continue
		}
		fieldName := structField.Tag.Get("json")
		if fieldName == "" {
			errorMap[fieldError.Field()] = fieldError.Error()
			continue
		}
		errorMap[fieldName] = fieldError.Error()
	}
	return errorMap
}

func ResponseError(error string) map[string]interface{} {
	return map[string]interface{}{
		"error": error,
	}
}
