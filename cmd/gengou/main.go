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
	w := flag.Bool("w", false, "full-width")
	d := flag.Bool("d", false, "show full date")
	v := flag.Bool("v", false, "version")
	f := flag.String("f", "2006.01.02", "time format")

	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], gengou.Version)
		os.Exit(0)
	}

	fun := gengou.EraYear
	if *d {
		fun = gengou.EraDate
	}

	dates := flag.Args()
	if len(dates) < 1 {
		now := time.Now()
		dates = append(dates, now.Format(*f))
	}

	for _, raw := range dates {
		date, err := time.Parse(*f, raw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse date: %s\n", raw)
			continue
		}
		if *w {
			fmt.Println(width.Widen.String(fun(date)))
		} else {
			fmt.Println(fun(date))
		}
	}
}
