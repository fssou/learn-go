package container

import (
	"errors"
	"github.com/fssou/learn-go/internal/adapter/in/http"
	"github.com/fssou/learn-go/internal/adapter/in/http/handler"
	"log"
)


type httpContainer struct {
	server *http.Adapter
}

func newHttpContainer() *httpContainer {
	container := &httpContainer{}
	return container
}

func (container *httpContainer) Init() error {

	route := http.NewRouter(
		[]http.Handler{
			handler.NewHealth(),
			handler.NewHello(),
			handler.NewLake(),
		},
	)
	routerRegistry := route.Register()
	config := &http.AdapterConfig{
		Host: "0.0.0.0",
		Port: "3333",
	}
	server := http.NewAdapter(routerRegistry, config)
	container.server = server
	go func() {
		if err := server.Listen(); err != nil {
			log.Println("failed to init http container: " + err.Error())
		}
	}()
	return nil
}

func (container *httpContainer) Close() error {
	if err := container.server.Close(); err != nil {
		return errors.New("failed to close container: " + err.Error())
	}
	return nil
}
