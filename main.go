package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"nms/src/backend/env"
	"nms/src/backend/model"
	"nms/src/backend/util/connection"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

//? Embed static files to binary
//go:embed src/frontend/web/*
var frontend embed.FS

func main() {
	e := echo.New()

	/* -------------------------------- FRONT END ------------------------------- */
	contentHandler := echo.WrapHandler(http.FileServer(http.FS(frontend)))
	contentRewrite := middleware.Rewrite(map[string]string{"/*": "/src/frontend/web/$1"})
	e.GET("/*", contentHandler, contentRewrite)

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
