package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type User struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

var jsonStr = `{
		"name": "liu",
		"age": 30
	}`

func TestJson(t *testing.T) {
	u := new(User)
	err := json.Unmarshal([]byte(jsonStr), u)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*u)

	if v, err := json.Marshal(u); err == nil {
		fmt.Println(v)
	} else {
		t.Error(err)
	}
}
