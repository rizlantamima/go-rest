package main

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/rizlantamima/go-rest/config"
	"gopkg.in/yaml.v2"
)

func main() {
	data, err := ioutil.ReadFile("server.yml")
	if err != nil {
		log.Error(err)
		return
	}

	var config config.Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Error(err)
		return
	}

	e := echo.New()
	e.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello, World!")
	})
	var serverAddress string = config.Server.Host + ":" + config.Server.Port
	e.Logger.Fatal(e.Start(serverAddress))
}
