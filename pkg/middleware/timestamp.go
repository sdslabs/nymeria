package middleware

import (
	"time"
)

func CurrentTimeStamp() string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	t := time.Now().In(loc)
	return t.Format("2006-01-02 15:04:05")
}
