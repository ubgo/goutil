package goutil

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
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

type Name struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Get FirstName and LastName from the FullName
func ParseName(fullName string) Name {
	nameParts := strings.Split(strings.TrimSpace(fullName), " ")

	var name Name
	if len(nameParts) > 0 {
		name.FirstName = nameParts[0]
	}

	if len(nameParts) > 1 {
		name.LastName = nameParts[1]
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

func RandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())

	chars := []rune("123456789")
	// length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "ExcbsVQs"
	return str
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

func RandMD5String() string {
	secret := uuid.New().String()
	key := []byte(secret)
	hash := md5.Sum(key)
	return hex.EncodeToString(hash[:])
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

func StringIndexWithLowerCase(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == strings.ToLower(val) {
			return i, true
		}
	}
	return -1, false
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

func IntIndex(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func UUIDIndex(slice []uuid.UUID, val uuid.UUID) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func StringArrayToUUIDArray(val []string) []uuid.UUID {
	var uids []uuid.UUID
	for _, v := range val {
		uids = append(uids, uuid.MustParse(v))
	}
	return uids
}

// Remove all special characters e.g. Aman C。Salcedo will become Aman C Salcedo
func CleanString(s string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		// log.Fatal(err)
		return ""
	}
	processedString := reg.ReplaceAllString(s, "")
	return processedString
}

func HashString(val string) string {
	key := []byte(val)
	h := sha256.New()
	h.Write(key)
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

// Strip all email from strings e.g. My email is abc@demo.com. will become My email is.
func StripEmails(s string) string {
	const regex = `\S*@\S*\s?`
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, "")
}

func GetTypeName(myvar interface{}) string {
	valueOf := reflect.ValueOf(myvar)

	if valueOf.Type().Kind() == reflect.Ptr {
		return (reflect.Indirect(valueOf).Type().Name())
	} else {
		return (valueOf.Type().Name())
	}
}

// Get Nested TypeName as String func to make the Preload Typesafe instead of hardcoded
// fmt.Println(util.GormTypeName(uxmpim_type.Item{}, uxmpim_type.ItemStatus{}))
// output: Item.ItemStatus
func GormTypeName(values ...interface{}) string {
	var typeNames []string
	for _, val := range values {
		typeNames = append(typeNames, GetTypeName(val))
	}

	if len(typeNames) == 0 {
		return ""
	}

	return strings.Join(typeNames, ".")
}

func TypeToString(data interface{}) string {
	mdata, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(mdata)
}

func IsErrNil(err error) bool {
	if err == nil {
		return true
	}

	return false
}

func IsErrNotNil(err error) bool {
	if err != nil {
		return true
	}

	return false
}
