package gengou

import (
	"testing"
	"time"
)

var layout = "2006.01.02 MST"

func TestEraYear(t *testing.T) {
	date, _ := time.Parse(layout, "2019.04.30 JST")
	year := EraYear(date)
	if year != "平成31年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestReiwa(t *testing.T) {
	date, _ := time.Parse(layout, "2019.05.01 JST")
	year := EraYear(date)
	if year != "令和元年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestReiwa2(t *testing.T) {
	date, _ := time.Parse(layout, "2020.01.01 JST")
	year := EraYear(date)
	if year != "令和2年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestReiwaLastDayOfYear(t *testing.T) {
	date, _ := time.Parse(layout, "2019.12.31 JST")
	year := EraYear(date)
	if year != "令和元年" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestEraDay(t *testing.T) {
	date, _ := time.Parse(layout, "2019.12.31 JST")
	year := EraDate(date)
	if year != "令和元年12月31日" {
		t.Errorf("formatting failed: %s", year)
	}
}

func TestEraNotFound(t *testing.T) {
	date, _ := time.Parse(layout, "0644.12.31 JST")
	year := EraYear(date)
	if year != "644年" {
		t.Errorf("formatting failed: %s", year)
	}
}
