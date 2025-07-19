package util

import (
	"errors"
	"log"
	"regexp"
	"strings"
	"time"
)

func findTime(s string) time.Time {
	reg := regexp.MustCompile(`[^^ ][0-9-]+ [0-9:]+[^$ ]`)
	createTime, err := time.Parse(time.DateTime, reg.FindString(s))
	if err != nil {
		log.Fatal(err)
	}
	return createTime
}

func HadTime(s string) bool {
	reg := regexp.MustCompile(`[0-9-]+ [0-9:]+$`)
	return reg.MatchString(s)
}

func GetItemNoTime(item string) string {
	reg := regexp.MustCompile(`^[^ ]*( |$)`)
	return strings.TrimSpace(reg.FindString(item))
}

func RestaurantHadTime(s string) bool {
	reg := regexp.MustCompile(`[0-9-] [0-9:]+$`)
	return reg.MatchString(s)
}

func ClearTime(s string) string {
	reg := regexp.MustCompile(`^[^ ]+`)
	return reg.FindString(strings.TrimSpace(s))
}

func GetTime(v string) (time.Time, error) {
	reg := regexp.MustCompile(`[0-9-]+ [0-9:]+$`)
	if !reg.MatchString(v) {
		return time.Time{}, errors.New("this text not found time")
	}
	parse, _ := time.Parse(time.DateTime, reg.FindString(v))
	return parse, nil
}
