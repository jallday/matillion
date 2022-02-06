package main

import (
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/joshuaAllday/matillion/pkg/app"
	"gitlab.com/joshuaAllday/matillion/pkg/config"
	"gitlab.com/joshuaAllday/matillion/pkg/handlers"
	"gitlab.com/joshuaAllday/matillion/pkg/server"
)

func main() {
	s, err := server.New(config.LoadConfig())
	if err != nil {
		panic(err)
	}
	if err = s.Start(); err != nil {
		panic(err)
	}
	defer s.Stop()
	handlers.InitHandlers(s.Router(), app.New(s))

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

}
