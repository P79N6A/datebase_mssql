package tools

import (
	"fmt"
	"time"
)

func GetDate() string {
	now := time.Now()
	day, _ := time.ParseDuration("24h")
	datet := now.Add(day)
	return fmt.Sprint(datet.Format("2006-01-02"))
}
