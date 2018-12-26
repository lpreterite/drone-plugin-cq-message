package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "cq-message"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "cqHost",
			Usage:  "cqhttp host",
			EnvVar: "PLUGIN_CQ_HOST",
		},
		cli.StringFlag{
			Name:   "cqAction",
			Usage:  "cqhttp action",
			EnvVar: "PLUGIN_CQ_ACTION",
		},
		cli.StringFlag{
			Name:   "cqToken",
			Usage:  "cqhttp header auth token",
			EnvVar: "PLUGIN_CQ_TOKEN",
		},
		cli.GenericFlag{
			Name:   "cqQuery",
			Usage:  "action params",
			EnvVar: "PLUGIN_CQ_QUERY",
			Value:  &StringMapFlag{},
		},
	}
	app.Action = func(c *cli.Context) error {

		_url, _ := url.Parse(c.String("cqHost"))
		_url.Path = c.String("cqAction")
		_query := _url.Query()
		_params := c.Generic("cqQuery").(*StringMapFlag).Get()
		for key := range _params {
			_query.Add(key, _params[key])
		}
		_url.RawQuery = _query.Encode()

		reqest, err := http.NewRequest("GET", _url.String(), nil)

		reqest.Header.Add("cache-control", "no-cache")
		reqest.Header.Add("Authorization", "Token "+c.String("cqToken"))

		if err != nil {
			return err
		}

		client := &http.Client{}
		response, err := client.Do(reqest)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		fmt.Println("Send Message in ", _url.String())
		return err
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
