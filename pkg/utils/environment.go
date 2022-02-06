package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetEnvString(key string, dft string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return dft
}

func GetEnvBool(key string, dft bool) bool {
	if ok, err := strconv.ParseBool(os.Getenv(key)); err == nil {
		return ok
	}
	return dft
}

func GetEnvInt(key string, dft int) int {
	if i, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return i
	}
	return dft
}

func GetEnvStringArray(key string, dft []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return dft
	}

	return strings.Split(v, ",")
}
