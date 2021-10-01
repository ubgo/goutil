package goutil

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// FuncName will return the current function's name.
// It can be used for a better log debug system.(I'm NOT sure.)
func FuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

// Limit the decimal value e.g. 2.34343434 will become 2.34 if precision is set to 2
func FloatPrecision(num float64, precision int) float64 {
	p := math.Pow10(precision)
	value := float64(int(num*p)) / p
	return value
}

func FullName(firstName string, lastName string) string {
	name := firstName
	if len(lastName) > 0 {
		name = name + " " + lastName
	}
	return name
}

// Env get key environment variable if exist otherwise return defalutValue
func Env(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// Define n as number to limit the length of the random string
func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Convert any string to snakecase StudentID will become student_id
func SnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("([A-Z])([A-Z][a-z])")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	output := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	output = matchAllCap.ReplaceAllString(output, "${1}_${2}")
	output = strings.ReplaceAll(output, "-", "_")
	return strings.ToLower(output)
}

// Convert any struct to JSON String
func ToJSON(val interface{}) (string, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Convert any struct to JSON String with Pretty Print
func ToJSONIndent(val interface{}) (string, error) {
	b, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func PrintToJSON(val interface{}) {
	b, err := ToJSONIndent(val)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(b)
}

// Use Make the unused value used so golang will not give error while compiling
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func StringIndex(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func UintIndex(slice []uint, val uint) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func IntIndex(slice []uint, val uint) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
