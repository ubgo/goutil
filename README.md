# Goutil  [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/knesklab/util/blob/master/LICENSE)

A small sets of golang indpendent packages for daily usage in any project.

## Installation
```
go get github.com/ubgo/goutil
```

### FuncName : string
FuncName will return the current function's name.
It can be used for a better log debug system.(I'm NOT sure.)
```go
goutil.FuncName() // output: myfunc
```

### FloatPrecision(num float64, precision int) float64
Limit the decimal value e.g. 2.34343434 will become 2.34 if precision is set to 2
```go
goutil.FloatPrecision(2.34343434, 2) // output: 2.34
```

### FullName(firstName string, lastName string) string 
Concat FirstName and LastName
```go
goutil.FullName("Lucian", "Khanakia") // output: Lucian Khanakia
```

### Env(key, defaultValue string) string {
Env get key environment variable if exist otherwise return defalutValue
```go
goutil.Env("MODE", "local") // output: local if set MODE=prod is not set in terminal
```

### RandString(n int) string 
It generates random string Define n as number to limit the length of the random string.
```go
goutil.RandString(10) // output: Xydere12sw
```

### SnakeCase(str string) string 
Convert any string to snakecase.
```go
goutil.SnakeCase("Lucian Khanakia") // output: luci_khanakia
goutil.SnakeCase("kebab-case") // output: kebab_case
goutil.SnakeCase("Id0Value") // output: id0_value
goutil.SnakeCase("ID0Value") // output: id0_value
goutil.SnakeCase("HTTPRequest") // output: http_request

```

### ToJSON(val interface{}) (string, error)
Convert any struct to JSON string
```go
type Student struct {
  Name   string `json:"name"`
  RoleNo int    `json:"roleNo"`
}

student := Student{
  Name:   "Lucian",
  RoleNo: 7,
}

fmt.Println(goutil.ToJSON(student))
// out: {"name":"Lucian","roleNo":7}

fmt.Println(goutil.ToJSONIndent(student))
// out:
{
  "name": "Lucian",
  "roleNo": 7
}

goutil.PrintToJSON(student)
// out:
{
  "name": "Lucian",
  "roleNo": 7
}
```

### Use - Accepts any value as argument
Use Make the unused value used so golang will not give error while compiling
```go
goutil.Use("Luci")
goutil.Use([]string{"test"})
```

### StringIndex(slice []string, val string) (int, bool)
Check if value exist on a given string slice then return the index
```go
result, ok := StringIndex([]string{"luci", "aman", "khanakia"}, "aman")
// output: 1, true
```


### UintIndex(slice []uint, val uint) (int, bool)
Check if value exist on a given uint slice then return the index
```go
result, ok := UintIndex([]uint{1, 2, 3}, 2)
// output: 1, true
```

### IntIndex(slice []int, val int) (int, bool)
Check if value exist on a given int slice then return the index
```go
result, ok := IntIndex([]int{1, 2, 3}, 2)
// output: 1, true
```

### CleanString(s string) string {
Clean the string by removing all the special chars and other weird chars and returns alphanumeric only
```go
result := CleanString("##khanakia")
// output: khanakia
```

### HashString(val string) string
Convert string to sha256
```go
result := HashString("luci")
// output: 2fdcbc8615c275ffbe49106cf85fbab1566b92559a251a5535a217f211dfa3f2
```

## Contribute

If you would like to contribute to the project, please fork it and send us a pull request.  Please add tests
for any new features or bug fixes.

## Stay in touch

* Author - [Aman Khanakia](https://twitter.com/mrkhanakia)
* Website - [https://khanakia.com](https://khanakia.com/)

## License

goutil is [MIT licensed](LICENSE).
