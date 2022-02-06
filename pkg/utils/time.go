package utils

import "time"

func IntTime() int {
	return int(time.Now().Unix())
}
