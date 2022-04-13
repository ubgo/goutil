package goutil

import (
	"os"
	"testing"
)

func Test_FloatPrecision(t *testing.T) {
	result := FloatPrecision(2.34343434, 2)
	if result != 2.34 {
		t.Errorf("Returned = %f; want 2.34", result)
	}
}

func Test_FullName(t *testing.T) {
	val := FullName("Lucian", "Khanakia")
	if val != "Lucian Khanakia" {
		t.Errorf("Returned = %s; want Lucian Khanakia", val)
	}
}

// Run this in terminal - `mode="prod" go test`
func Test_Env1(t *testing.T) {
	os.Setenv("ENV_VAR", "prod")
	defer os.Unsetenv("ENV_VAR")
	result := Env("ENV_VAR", "local")

	if result != "prod" {
		t.Errorf("Returned = %s; want prod", result)
	}
}

func Test_Env2(t *testing.T) {
	result := Env("ENV_VAR1", "local")

	if result != "local" {
		t.Errorf("Returned = %s; want local", result)
	}
}

func Test_RandString(t *testing.T) {
	val := RandString(10)
	if len(val) != 10 {
		t.Errorf("Run(10) = %d; want 10", len(val))
	}
}

func Test_SnakeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
		{"snake_case", "snake_case"},
		{"Pascal_Snake", "pascal_snake"},
		{"SCREAMING_SNAKE", "screaming_snake"},
		{"kebab-case", "kebab_case"},
		{"Pascal-Kebab", "pascal_kebab"},
		{"SCREAMING-KEBAB", "screaming_kebab"},
		{"A", "a"},
		{"AA", "aa"},
		{"AAA", "aaa"},
		{"AAAA", "aaaa"},
		{"AaAa", "aa_aa"},
		{"HTTPRequest", "http_request"},
		{"BatteryLifeValue", "battery_life_value"},
		{"Id0Value", "id0_value"},
		{"ID0Value", "id0_value"},
	}

	for _, test := range tests {
		result := SnakeCase(test.input)
		if result != test.expected {
			t.Errorf("Returned = %s; want %s", result, test.expected)
		}
		// assert.Equal(t, test.expected, result)
	}
}

func Test_ToJSON(t *testing.T) {
	type Student struct {
		Name   string `json:"name"`
		RoleNo int    `json:"roleNo"`
	}
	student := Student{
		Name:   "Lucian",
		RoleNo: 7,
	}

	_, err := ToJSON(student)
	if err != nil {
		t.Error(err)
	}
}

func Test_StringIndex(t *testing.T) {
	result, _ := StringIndex([]string{"luci", "aman", "khanakia"}, "aman")
	if result != 1 {
		t.Errorf("Returned = %d; want %s", result, "1")
	}

	result, _ = StringIndex([]string{"luci", "aman", "khanakia"}, "na")
	if result != -1 {
		t.Errorf("Returned = %d; want %s", result, "-1")
	}
}

func Test_UintIndex(t *testing.T) {
	result, _ := UintIndex([]uint{1, 2, 3}, 2)
	if result != 1 {
		t.Errorf("Returned = %d; want %s", result, "1")
	}

	result, _ = UintIndex([]uint{1, 2, 3}, 12)
	if result != -1 {
		t.Errorf("Returned = %d; want %s", result, "-1")
	}
}

func Test_IntIndex(t *testing.T) {
	result, _ := IntIndex([]int{1, 2, 3}, 2)
	if result != 1 {
		t.Errorf("Returned = %d; want %s", result, "1")
	}

	result, _ = IntIndex([]int{1, 2, 3}, 12)
	if result != -1 {
		t.Errorf("Returned = %d; want %s", result, "-1")
	}
}

func Test_CleanString(t *testing.T) {
	result := CleanString("Aman C。Salcedo")
	if result != "Aman CSalcedo" {
		t.Errorf("Returned = %s; want %s", result, "Aman CSalcedo")
	}

	result = CleanString("##Khanakia")
	if result != "Khanakia" {
		t.Errorf("Returned = %s; want %s", result, "Khanakia")
	}
}

func Test_HashString(t *testing.T) {
	result := HashString("luci")
	want := "2fdcbc8615c275ffbe49106cf85fbab1566b92559a251a5535a217f211dfa3f2"
	if result != want {
		t.Errorf("Returned = %s; want %s", result, want)
	}
}
