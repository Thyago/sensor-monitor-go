package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type APIError struct {
	Error string `json:"error"`
}

type ResponseData struct {
	Data interface{} `json:"data"`
}

type ResponseDataArray struct {
	Data []interface{} `json:"data"`
}

func unmarshalBody(c *gin.Context, v interface{}) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{"Failed"})
		log.Println(err)
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIError{"Failed"})
		log.Println(err)
		return err
	}
	return nil
}

func setResponseError(c *gin.Context, code int, message string, err error) {
	c.JSON(code, APIError{message})
	log.Println(err)
}

func getUrlParamUINT64(c *gin.Context, param string) (uint64, error) {
	sID, err := strconv.ParseUint(c.Param(param), 10, 64)
	if err != nil {
		setResponseError(c, http.StatusBadRequest, fmt.Sprintf("%v is required", param), err)
		return 0, err
	}
	return sID, nil
}
