package main

import (
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"strconv"
)

type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

func CurrentHeight(c *gin.Context) {
	height, err := RawDB.CurrentHeight()
	response := &Response{}
	if err != nil {
		response.Code = 50000
	} else {
		response.Code = 20000
		response.Result = height
	}
	c.JSON(http.StatusOK, response)
}

func GetConvertItemAll(c *gin.Context) {
	items, err := RawDB.GetConvertItemAll()
	response := &Response{}
	if err != nil {
		response.Code = 50000
	} else {
		response.Code = 20000
		response.Result = items
	}
	c.JSON(http.StatusOK, response)
}

func GetConvertItemByMid(c *gin.Context) {
	mid, _ := strconv.ParseInt(c.Query("mid"), 10, 64)
	items, err := RawDB.GetConvertItem(big.NewInt(mid))
	response := &Response{}
	if err != nil {
		response.Code = 50000
	} else {
		response.Code = 20000
		response.Result = items
	}
	c.JSON(http.StatusOK, response)
}
