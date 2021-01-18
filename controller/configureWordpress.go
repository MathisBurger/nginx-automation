package controller

import (
	"github.com/MathisBurger/nginx-automation/config"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"os"
	"strings"
)

type configureWordpressResponse struct {
	Message    string `json:"message"`
	HttpStatus int    `json:"http_status"`
	Status     string `json:"status"`
	Error      string `json:"error"`
}

func ConfigureWordpressController(c *fiber.Ctx) error {
	domain := c.Query("domain")
	cfgPath := "./" + domain + ".conf"
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		_, err := os.Create(cfgPath)
		if err != nil {
			return c.JSON(configureWordpressResponse{
				"Error while installing wordpress",
				200,
				"ok",
				err.Error(),
			})
		}
	}
	cfg, _ := config.ParseConfig()
	data, _ := ioutil.ReadFile("./sample/wordpress.conf")
	modified := []byte(strings.ReplaceAll(strings.ReplaceAll(string(data), "{{DOMAIN}}", domain), "{{UPSTREAM}}", cfg.WebserverAddress))
	err := ioutil.WriteFile(cfgPath, modified, 0644)
	if err != nil {
		return c.JSON(configureWordpressResponse{
			"Error while installing wordpress",
			200,
			"ok",
			err.Error(),
		})
	}

	return c.JSON(configureWordpressResponse{
		"Successfully configured Wordpress",
		200,
		"ok",
		"None",
	})
}
