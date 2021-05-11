package main

import (
	"fmt"

	"github.com/shayd3/jwn-it/data"
	"github.com/shayd3/jwn-it/routes"
)

func main() {
	fmt.Println("🗄️ Connecting to database...")
	data.ConnectDatabase("jwnit.db", 0600)
	fmt.Println("🗄️ Database connection successful!")
	// Don't close the database until the API shuts down
	// defer will defer the execution of a function until the surrounding function returns
	defer data.DB.Close()
	fmt.Println("⚙️ Setting up Gin routes...")
	routes.SetupRouter()
	fmt.Println("⚙️ Routes successfully setup!")
	routes.SetupStaticContent()
	routes.Run()
	fmt.Println("🚀 Up and running!")
}


