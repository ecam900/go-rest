package main

import (
	"fmt"
	"net/http"

	"github.com/ecam900/go-rest/internal/comment"
	"github.com/ecam900/go-rest/internal/database"
	transportHTTP "github.com/ecam900/go-rest/internal/transport/http"
	"github.com/joho/godotenv"
)

// App struct which contains things like pointers
// to database connections
type App struct{}

// Run - Sets up  our application
func (app *App) Run() error {
	fmt.Println("Setting Up...")
	godotenv.Load(".env")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	// handler := transportHTTP.NewHandler(commentService)
	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Prod Ready Rest API Boilerplate")
	// Load environment variables with godotenv
	godotenv.Load(".env")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")

	}
}
