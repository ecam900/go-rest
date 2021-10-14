package main

import "fmt"

// App struct which contains things like pointers
// to database connections
type App struct {}

// Run - Sets up  our application 
func (app *App) Run() error {
	fmt.Println("Setting Up...")
	return nil
}

func main() {
	fmt.Println("Prod Ready Rest API Boilerplate")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")

	}
}