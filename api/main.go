package main

import (
	"api/framework/config"
	"api/framework/registry"
	"api/framework/router"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

func main() {
	app := config.App{}
	app.Initialize()

	r := registry.NewRegistry(app.DB)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
