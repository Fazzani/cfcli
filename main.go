package main

import (
	json "encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/fazzani/cfcli/domain"
	svc "github.com/fazzani/cfcli/service"
	cli "github.com/urfave/cli/v2"

	gorequest "github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
)

func main() {
	cwd, err := os.Getwd()
	configFilePath := path.Join(cwd, "config.yaml")

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
				EnvVars: []string{"CONFIG_PATH"},
			},
		}, Action: func(c *cli.Context) error {
			fmt.Println("serve:", c.Bool("serve"))
			fmt.Println("option:", c.Bool("option"))
			fmt.Println("message:", c.String("message"))
			configFilePath = c.String("config")
			return nil
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	zoneID := "4e7521218955d660919420f8c0e16cd6"

	config, err := svc.Default(configFilePath)
	fmt.Printf("config =====> %s\n", config)
	//todo: overide it from CLI

	request := gorequest.New()
	resp, body, errs := request.Get(config.API.BaseURL+"zones/"+zoneID).
		Set("Content-Type", "application/json").
		Set("account_id", config.Authentication.AccountID).
		Set("x-auth-email", config.Authentication.Email).
		Set("x-auth-key", config.Authentication.Key).
		End()

	if errs != nil {
		panic(errs)
	}

	results := domain.Zone{}
	err = json.Unmarshal([]byte(body), &results)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	log.WithFields(log.Fields{
		"baseUrl": config.API.BaseURL,
		"resp":    resp,
		"body":    body,
	}).Info("Cloudflare Request")
}
