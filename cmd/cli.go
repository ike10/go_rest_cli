package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	fmt.Println("Cli is running")
	app := cli.NewApp()
	app.Name = "Go REST APi consumer"
	app.Usage = "Query your rest api with http Methods"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "From",
			Value: "Api",
			Usage: "The Endpoint you are calling",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "get",
			Usage: "Sends a GET http request to an API endpoint",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				response, err := http.Get(c.String("From"))
				if err != nil {
					return err
				}
				responseData, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(responseData))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
