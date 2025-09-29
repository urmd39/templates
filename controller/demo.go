package controller

import "templates/service"

type DemoController interface{}

type demoControllerIml struct {
	service service.DemoService
}

func NewDemoController() DemoController {
	return &demoControllerIml{
		service: service.NewDemoService(),
	}
}
