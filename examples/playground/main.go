package main

import (
	"log"

	gest "github.com/martinpetrdev/Gest"
	"github.com/martinpetrdev/Gest/examples/playground/greeting"
)

var AppModule = gest.Module{
	Name:    "app",
	Imports: []*gest.Module{&greeting.GreetingModule},
}

func main() {
	app, err := gest.Create(AppModule)
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
