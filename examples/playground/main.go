package main

import (
	"log"

	gest "github.com/martinpetrdev/Gest"
)

var AppModule = gest.Module{
	Name: "app",
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
