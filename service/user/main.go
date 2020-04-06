package main

import (
	"github.com/micro/go-micro"
	"gitlab.com/otis-team/backend/db/client"
	pb "gitlab.com/otis-team/backend/service/user/proto/user"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
	)
	service.Init()

	dynamoClient := client.DynamoClient{}
	dynamoClient.Init()

	handler := &Handler{dynamoClient}

	pb.RegisterUserServiceHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}