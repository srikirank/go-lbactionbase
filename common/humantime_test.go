package common

import (
	"regexp"
	"testing"
	"time"
)

func TestUnixMicroTimeToHuman_Now_ReturnsToday(t *testing.T) {
	now := time.Now().UnixMicro()
	ht := UnixMicroTimeToHuman(now)
	if UnixMicroTimeToHuman(now) != "Today" {
		t.Fatalf("Expected Unix Now to translate to %s, but found %s", "Today", ht)
	}
}

func TestUnixMicroTimeToHuman_2DaysAgo_ReturnsLastWeek(t *testing.T) {
	twoDaysAgo := time.Now().Add(-time.Hour * 1 * 24).UnixMicro()
	ht := UnixMicroTimeToHuman(twoDaysAgo)
	if ht != "Last Week" {
		t.Fatalf("Expected Unix Now to translate to %s, but found %s", "Last Week", ht)
	}
}

func TestUnixMicroTimeToHuman_8DaysAgo_ReturnsLastWeek(t *testing.T) {
	twoDaysAgo := time.Now().Add(-time.Hour * 8 * 24).UnixMicro()
	ht := UnixMicroTimeToHuman(twoDaysAgo)
	if ht != "Last Month" {
		t.Fatalf("Expected Unix Now to translate to %s, but found %s", "Last Week", ht)
	}
}

func TestUnixMicroTimeToHuman_31DaysAgo_ReturnsLastWeek(t *testing.T) {
	twoDaysAgo := time.Now().Add(-time.Hour * 31 * 24).UnixMicro()
	ht := UnixMicroTimeToHuman(twoDaysAgo)
	p := "[a-zA-Z]+ [\\d]+ [\\d]+"
	m, _ := regexp.Match(p, []byte(ht))
	if !m {
		t.Fatalf("Expected Unix Now to match the pattern %s, but found %s", p, ht)
	}
}
