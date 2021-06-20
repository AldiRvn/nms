package main

import (
	"fmt"
	"log"
	"nms/src/backend/env"
	"nms/src/backend/model"
	"nms/src/backend/util/connection"

	"github.com/labstack/echo"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	e := echo.New()

	/* -------------------------------- FRONT END ------------------------------- */
	e.Static("/", "src/frontend/web")

	/* -------------------------------- BACK END -------------------------------- */
	svc := e.Group("/svc")

	database := svc.Group("/database")
	mongodb := new(connection.MongoDB)
	database.GET("/find", func(c echo.Context) error {
		return model.BuildResponse(c, mongodb.GetDatabaseNames(), nil)
	})
	database.GET("/collection/find", func(c echo.Context) error {
		return model.BuildResponse(c, mongodb.GetDatabaseNameWithCollectionName(), nil)
	})

	log.Fatalln(e.Start(
		fmt.Sprintf(":%d", env.APP_PORT),
	))
}
