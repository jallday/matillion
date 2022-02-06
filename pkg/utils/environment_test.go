package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvString(t *testing.T) {
	expected := "DEFAULT"
	v := GetEnvString("TEST_ENV", "DEFAULT")
	assert.EqualValues(t, expected, v, "expected it to return the default value")

	os.Setenv("TEST_ENV", "NOT_DEFAULT")
	defer os.Setenv("TEST_ENV", "")
	v = GetEnvString("TEST_ENV", "DEFAULT")
	assert.NotEqual(t, expected, v, "expected it to not return the default value")
}

func TestGetboolString(t *testing.T) {
	expected := true
	v := GetEnvBool("TEST_ENV", expected)
	assert.EqualValues(t, expected, v, "expected it to return the default value")

	os.Setenv("TEST_ENV", "false")
	defer os.Setenv("TEST_ENV", "")
	v = GetEnvBool("TEST_ENV", expected)
	assert.NotEqual(t, expected, v, "expected it to not return the default value")
}

func TestGetIntString(t *testing.T) {
	expected := 100
	v := GetEnvInt("TEST_ENV", expected)
	assert.EqualValues(t, expected, v, "expected it to return the default value")

	os.Setenv("TEST_ENV", "10000")
	defer os.Setenv("TEST_ENV", "")
	v = GetEnvInt("TEST_ENV", expected)
	assert.NotEqual(t, expected, v, "expected it to not return the default value")
}

func TestGetEnvStringArray(t *testing.T) {
	expected := []string{"one", "two", "three"}
	v := GetEnvStringArray("TEST_ENV", expected)
	assert.EqualValues(t, len(expected), len(v), "expected it to return the default value")

	os.Setenv("TEST_ENV", "10000")
	defer os.Setenv("TEST_ENV", "")
	v = GetEnvStringArray("TEST_ENV", expected)
	assert.NotEqual(t, len(expected), len(v), "expected it to not return the default value")
}
