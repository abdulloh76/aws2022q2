package main

import (
	"awsshop/config"
	"awsshop/internal/entity"
	"awsshop/internal/productDto"
	"awsshop/internal/repository"
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	config.LoadConfig(".")
	envs := config.GetEnvConfigs()
	db, err := config.GetConnection(envs)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	db.AutoMigrate(&entity.Stock{})
	db.AutoMigrate(&entity.Product{})

	var dto productDto.ProductBody
	rBody := []byte(request.Body)
	parseErr := json.Unmarshal(rBody, &dto)
	if parseErr != nil {
		return Response{StatusCode: 400}, err
	}

	product := productDto.RequestDTOtoEntity(&dto)
	createErr := repository.InsertProduct(db, &product)
	if createErr != nil {
		return Response{StatusCode: 500}, err
	}

	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"message": "successfully created",
	})
	if err != nil {
		return Response{StatusCode: 500}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
