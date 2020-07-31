package main

import (
	"api/framework/registry"
	"api/framework/router"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

func main() {
	a := App{}
	a.Initialize()

	r := registry.NewRegistry(a.DB)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
