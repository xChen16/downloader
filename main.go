package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/downloader/config"
	"github.com/downloader/extract"
	"github.com/downloader/tools"
)

func init() {
	flag.BoolVar(&config.Debug, "d", false, "Debug mode")
	flag.BoolVar(&config.Version, "v", false, "Show version")
	flag.BoolVar(&config.InfoOnly, "i", false, "Info only")
	flag.StringVar(&config.Cookie, "c", "", "Cookie")

}

func main() {
	flag.Parse()
	args := flag.Args()
	if config.Version {
		fmt.Printf(
			"go downloader version : %s.\n", config.VERSION,
		)
		return
	}
	if len(args) < 1 {
		fmt.Println("command error")
		return
	}
	videoURL := args[0]
	u, err := url.ParseRequestURI(videoURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	domain := tools.Domain(u.Host)
	switch domain {
	case "bilibili":
		extract.Bilibili(videoURL)
	default:
		fmt.Println("unsupported URL")
	}

}
