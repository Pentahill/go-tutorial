package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a ptr to struct type")
	}

	if (reflect.TypeOf(st)).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a ptr to struct type")
	}

	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}

		fmt.Println("field type: ", field.Type)

		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func TestFillField(t *testing.T) {
	user := new(User)

	t.Log("name: ", user.Name, " age: ", user.Age)
	fillBySettings(user, map[string]interface{}{"Name": "liu", "Age": 20})
	t.Log("name: ", user.Name)
}

func TestTypeOf(t *testing.T) {
	num := 1.0
	reflect.ValueOf(&num).Elem().SetFloat(2.0)
}
