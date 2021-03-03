package controller

import (
	"github.com/MathisBurger/nginx-automation/config"
	"github.com/MathisBurger/nginx-automation/utils"
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

// controller to configure wordpress configuration
func ConfigureWordpressController(c *fiber.Ctx) error {

	// checking cors permission
	if !utils.CheckCORS(c.IP()) {
		return c.JSON(configureWordpressResponse{
			"Your origin is not allowed",
			200,
			"ok",
			"None",
		})
	}

	domain := c.Query("domain")

	// defined path
	cfgPath := "/etc/nginx/rproxy/http/enabled" + domain + ".conf"

	// creates file if not exists
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

	// read sample template config
	data, _ := ioutil.ReadFile("/root/installation-service/sample/wordpress.conf")

	// modify config
	modified := []byte(strings.ReplaceAll(strings.ReplaceAll(string(data), "{{DOMAIN}}", domain), "{{UPSTREAM}}", cfg.WebserverAddress))

	// write modified config
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
