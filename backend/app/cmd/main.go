package main

import (
	router "app/routers"
)

func main() {
	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.HTML(http.StatusOK, "<h1>Successful</h1>")
	// })
	// e.Logger.Fatal(e.Start(":8080"))

	router := router.NewRouter()

	router.Logger.Fatal(router.Start(":8080"))
}
