package main

import (
	"testing"
)

func TestGetCertainScoreReports(t *testing.T) {
	reports := getCertainScoreReports(schoolReports, certainScore)
	if len(reports) != 3 {
		t.Error("test failed")
	}
	for _, report := range reports {
		if report.score <= 59 {
			t.Error("test failed")
		} else if report.score > 100 {
			t.Error("test failed")
		} else if report.score < 0 {
			t.Error("test failed")
		}
	}
}

func TestGetCertainClassReports(t *testing.T) {
	reports := getCertainClassReports(schoolReports, certainClass)
	if len(reports) != 2 {
		t.Error("test failed")
	}
	for _, report := range reports {
		if report.class != "two" {
			t.Error("test failed")
		}
	}
}

func TestGetCertainNameReports(t *testing.T) {
	reports := getCertainNameReports(schoolReports, certainNamelen)
	if len(reports) != 1 {
		t.Error("test failed")
	}
	for _, report := range reports {
		if len(report.name) <= 10 {
			t.Error("test failed")
		}
	}
}
