package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"

	"github.com/Ostrovski/sxtant/sxapi"
	"github.com/Ostrovski/sxtant/web/handlers"
)

func main() {
	port := flag.Int("port", 3000, "TCP port")
	flag.Parse()

	session, err := mgo.Dial("sxtant-mg.local")
	if err != nil {
		panic(fmt.Sprintf("MongoDB connection failed ", err))
	}
	defer session.Close()

	sxApiClient := sxapi.NewClient("https://api.stackexchange.com/2.2", session.Clone())
	sitesFetcher := sxapi.NewSitesFetcher(sxApiClient)

	router := gin.Default()
	// if dev {
	router.StaticFS("/static", http.Dir("web/static"))
	// }
	router.LoadHTMLGlob("web/templates/**/*.tmpl")

	siteHandler := handlers.SiteHandler(sitesFetcher)
	router.GET("/", siteHandler)
	router.GET("/sites/:site", siteHandler)

	router.GET("/sock", handlers.SockHandler())

	if err := router.Run(fmt.Sprintf(":%v", *port)); err != nil {
		panic(fmt.Sprintf("Server start failed", err))
	}
}
