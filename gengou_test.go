package gengou_test

import (
	"fmt"
	"testing"
	"time"
	"blekksprut.net/gengou"
)

var layout = "2006.01.02 MST"

func TestEraYear(t *testing.T) {
	date, _ := time.Parse(layout, "2019.04.30 JST")
	year := gengou.EraYear(date)
	if year != "平成31年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestReiwa(t *testing.T) {
	date, _ := time.Parse(layout, "2019.05.01 JST")
	year := gengou.EraYear(date)
	if year != "令和元年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestReiwa2(t *testing.T) {
	date, _ := time.Parse(layout, "2020.01.01 JST")
	year := gengou.EraYear(date)
	if year != "令和2年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestReiwaLastDayOfYear(t *testing.T) {
	date, _ := time.Parse(layout, "2019.12.31 JST")
	year := gengou.EraYear(date)
	if year != "令和元年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestEraDay(t *testing.T) {
	date, _ := time.Parse(layout, "2019.12.31 JST")
	year := gengou.EraDate(date)
	if year != "令和元年12月31日" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestEraNotFound(t *testing.T) {
	date, _ := time.Parse(layout, "0644.12.31 JST")
	year := gengou.EraYear(date)
	if year != "644年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func ExampleFind() {
	now := time.Now()
	era, err := gengou.Find(now)
	if err != nil {
		panic(err)
	}
	fmt.Println(era.Name)
	// Output: 令和
}

func ExampleEraDate() {
	date, _ := time.Parse(layout, "1991.07.29 JST")
	fmt.Println(gengou.EraDate(date))
	// Output: 平成3年7月29日
}
