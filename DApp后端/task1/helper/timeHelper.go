package helper

import "time"

func FormatTimestamp(ts int64) string {
	return time.Unix(ts, 0).UTC().Format("2006-01-02 15:04:05")
}

func FormatTimestampCN(ts int64) string {
	return time.Unix(ts, 0).In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")
}
