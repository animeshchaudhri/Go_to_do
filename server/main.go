package main

import (
	"fmt"
	"log"

	"github.com/animesh/go-to-do/bootstrap"
	"github.com/animesh/go-to-do/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "localhost"), env.GetEnv("APP_PORT", "4000"))))

}
