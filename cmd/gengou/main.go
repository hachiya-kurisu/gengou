package main

import (
	"blekksprut.net/gengou"
	"flag"
	"fmt"
	"golang.org/x/text/width"
	"os"
	"time"
)

func main() {
	widen := flag.Bool("w", false, "full-width")
	version := flag.Bool("v", false, "version")
	layout := flag.String("f", "2006.01.02", "time format")

	flag.Parse()

	if *version {
		fmt.Printf("%s %s\n", os.Args[0], gengou.Version)
		os.Exit(0)
	}

	if flag.NArg() < 1 {
		now := time.Now()
		eraYear := gengou.EraYear(now)
		if *widen {
			fmt.Println(width.Widen.String(eraYear))
		} else {
			fmt.Println(eraYear)
		}
	} else {
		for _, arg := range flag.Args() {
			date, err := time.Parse(*layout, arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "failed to parse date: %s\n", arg)
				continue
			}
			eraYear := gengou.EraYear(date)
			if *widen {
				fmt.Println(width.Widen.String(eraYear))
			} else {
				fmt.Println(eraYear)
			}
		}
	}

}
