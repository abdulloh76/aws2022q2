package main

import (
	"awsshop/config"
	"awsshop/internal/entity"
	"awsshop/internal/repository"
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (Response, error) {
	config.LoadConfig(".")
	envs := config.GetEnvConfigs()
	db, err := config.GetConnection(envs)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	db.AutoMigrate(&entity.Stock{})
	db.AutoMigrate(&entity.Product{})
	products, _ := repository.FindProducts(db)

	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"data": products,
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
