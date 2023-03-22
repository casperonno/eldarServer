package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := fiber.New()
	app.Get("/greet", func(ctx *fiber.Ctx) error {

		userToGreet := ctx.Query("name", "incognito user")
		greeting := fmt.Sprintf("Hello %s", userToGreet)
		err := ctx.SendString(greeting)
		if err != nil {
			errMsg := fmt.Sprintf("failed to send status code respond. error message !: \n %s", err.Error())
			fmt.Println(errMsg)
			return fmt.Errorf(errMsg)
		}
		return nil
	})

	runService(app)

}

func runService(app *fiber.App) {

	// Create a channel to receive signals
	sigCh := make(chan os.Signal, 1)

	// Notify the channel when the program receives an interrupt or termination signal
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go listenService(app, cancel)
	fmt.Println("server is listening...")

	for {
		select {
		case <-ctx.Done():
			//release resources if there any
			shutDownService(app)
			return
		case <-sigCh:
			fmt.Print("\n app got system signal ! \n")
			shutDownService(app)
			return
		}

	}
}

func shutDownService(app *fiber.App) {
	fmt.Println("server stopped.. Shutting down gracefully !!!")
	err := app.Shutdown()
	if err != nil {
		fmt.Println("service shut down with error.", "err:", err.Error())
	}
}

func listenService(app *fiber.App, cancel context.CancelFunc) {

	err := app.Listen(":4500")
	if err != nil {
		fmt.Println("error initializing server -> ", err.Error())
		cancel()
	}

}
