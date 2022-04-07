/*
 * MassBank Minio API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"github.com/uly55e5/mbServices/src/redisstore"
	"log"
	"net/http"
	"os"

	openapi "github.com/uly55e5/mbServices/src/openapi"
)

func main() {
	file, err := os.OpenFile("/var/log/mb-minio.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)
	log.Printf("Server started")

	err = redisstore.Connect()
	if err != nil {
		panic(err)
	}

	DefaultApiService := openapi.NewDefaultApiService()
	DefaultApiController := openapi.NewDefaultApiController(DefaultApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
