package main

import (
	"github.com/shayd3/jwn-it/data"
	"github.com/shayd3/jwn-it/routes"
)

func main() {
	data.ConnectDatabase("jwnit.db", 0600)
	// Don't close the database until the API shuts down
	// defer will defer the execution of a function until the surrounding function returns
	defer data.DB.Close()
	routes.SetupRouter()
	routes.Run()
}


