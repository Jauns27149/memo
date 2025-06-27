package util

import (
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
