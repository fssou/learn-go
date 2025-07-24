package bootstrap

import (
	"errors"
	"github.com/fssou/learn-go/internal/bootstrap/container"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func StartApplication() error {
	signalSystemChannel := make(chan os.Signal, 1)
	signal.Notify(signalSystemChannel, syscall.SIGINT, syscall.SIGTERM)

	mainContainer := container.NewMainContainer()

	if err := mainContainer.Init(); err != nil {
		return errors.New("failed to initializr main container: " + err.Error())
	}

	<-signalSystemChannel
	log.Println("Received shutdown signal, shutting down gracefully...")
	if err := mainContainer.Close(); err != nil {
		log.Println("Error shutting down application:", err)
		return err
	}
	return nil
}
