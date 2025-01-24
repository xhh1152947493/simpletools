package utils

import "time"

func Timestamp2DateInLoc(timestamp int64) string {
	return time.Unix(timestamp, 0).In(time.Local).Format(time.DateTime)
}
