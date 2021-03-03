package controller

import (
	"github.com/MathisBurger/nginx-automation/config"
	"github.com/MathisBurger/nginx-automation/utils"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
	"strings"
)

type configureAngularResponse struct {
	Message    string `json:"message"`
	HttpStatus int    `json:"http_status"`
	Status     string `json:"status"`
	Error      string `json:"error"`
}

func ConfigureAngularController(c *fiber.Ctx) error {
	if !utils.CheckCORS(c.IP()) {
		return c.JSON(configureAngularResponse{
			"Your origin is not allowed",
			200,
			"ok",
			"None",
		})
	}
	domain := c.Query("domain")
	cfgPath := "/etc/nginx/rproxy/http/enabled/" + domain + ".conf"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		_, err := os.Create(cfgPath)
		if err != nil {
			return c.JSON(configureAngularResponse{
				"Error while installing angular app",
				200,
				"ok",
				err.Error(),
			})
		}
	}
	cfg, _ := config.ParseConfig()
	data, _ := ioutil.ReadFile("/root/installation-service/sample/angular.conf")
	modified := []byte(strings.ReplaceAll(strings.ReplaceAll(string(data), "{{DOMAIN}}", domain), "{{UPSTREAM}}", cfg.WebserverAddress))
	err := ioutil.WriteFile(cfgPath, modified, 0644)
	if err != nil {
		return c.JSON(configureAngularResponse{
			"Error while installing angular app",
			200,
			"ok",
			err.Error(),
		})
	}

	return c.JSON(configureAngularResponse{
		"Successfully configured angular app",
		200,
		"ok",
		"None",
	})
}
