package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/microcosm-cc/bluemonday"
)

var api *anaconda.TwitterApi
var p *bluemonday.Policy

func main() {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))

	api = anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	p = bluemonday.StrictPolicy()

	ln, err := net.Listen("tcp", ":79")

	if err != nil {
		fmt.Printf("Could not bind to port")
		os.Exit(1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error occurred: %s", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	userName, _, _ := reader.ReadLine()

	if len(userName) > 0 {
		fmt.Printf("[%s] %s %s\n",
			time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
			conn.RemoteAddr().String(),
			string(userName),
		)

		user, err := api.GetUsersShow(string(userName), nil)

		if err != nil {
			conn.Write([]byte(err.Error()))
			return
		}

		fmt.Fprintf(conn, "Login: %-33s Name: %-34s\n", user.ScreenName, user.Name)
		fmt.Fprintf(conn, "Location: %-30s URL: %-35s\n", user.Location, user.URL)
		fmt.Fprintf(conn, "Description: %s\n", user.Description)
		fmt.Fprintf(conn, "Last tweet %s from %s\n", user.Status.CreatedAt, p.Sanitize(user.Status.Source))
		fmt.Fprintf(conn, "Plan: %s\n", user.Status.Text)
	}
}
