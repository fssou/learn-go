package container

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Container interface {
	Close() error
	Init() error
}
type MainContainer struct {
	containers []Container
}

func NewMainContainer() *MainContainer {
	containers := []Container{
		newHttpContainer(),
	}
	return &MainContainer{
		containers,
	}
}

func (c *MainContainer) Init() error {

	for _, container := range c.containers {
		if err := container.Init(); err != nil {
			typeOfContainer := reflect.TypeOf(container)
			if typeOfContainer.Kind() == reflect.Ptr {
				typeOfContainer = typeOfContainer.Elem()
			}
			log.Printf("failed to initializr container [%s]: %s\n", typeOfContainer.Name(), err.Error())
		}
	}
	return nil
}

func (c *MainContainer) Close() error {
	errs := errors.New("failed to close main container")
	for _, container := range c.containers {
		if err := container.Close(); err != nil {
			errs = errors.New(fmt.Sprintf("%s: %s", errs.Error(), err.Error()))
			typeOfContainer := reflect.TypeOf(container)
			if typeOfContainer.Kind() == reflect.Ptr {
				typeOfContainer = typeOfContainer.Elem()
			}
			log.Printf("failed to close container [%s]: %s\n", typeOfContainer.Name(), err.Error())
		}
	}
	return errs
}
