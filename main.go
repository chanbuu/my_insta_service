package main

import "github.com/my_service/route"

func main() {

	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":1323"))

}
