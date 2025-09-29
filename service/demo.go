package service

import (
	"templates/database"
	"templates/database/mongo"
)

type DemoService interface{}

type demoServiceIml struct {
	repo database.DemoRepository
}

func NewDemoService() DemoService {
	return &demoServiceIml{
		repo: mongo.NewDemoRepository(),
	}
}
