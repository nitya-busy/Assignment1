package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ivalue(value interface{}, indent string) {
	val := reflect.ValueOf(value)
	typ := reflect.TypeOf(value)

	switch val.Kind() {

	case reflect.Map:
		fmt.Println(indent, "Type:", typ, "| Value: MAP")
		for _, key := range val.MapKeys() {
			fmt.Println(indent+"Key:", key.Interface())
			ivalue(val.MapIndex(key).Interface(), indent+"  ")
		}

	case reflect.Slice:
		fmt.Println(indent, "Type:", typ, "| Value:SLICE")
		for i := 0; i < val.Len(); i++ {
			fmt.Println(indent+"Index:", i)
			ivalue(val.Index(i).Interface(), indent+" ")
		}

	default:
		fmt.Println(indent, "Type:", typ, "| Value:", value)
	}
}

func main() {

	input := `{
		"name" : "Tolexo Online Pvt. Ltd",
		"age_in_years" : 8.5,
		"origin" : "Noida",
		"head_office" : "Noida, Uttar Pradesh",
		"address" : [
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			},
			{
				"street" : "91 Springboard",
				"landmark" : "Axis Bank",
				"city" : "Noida",
				"pincode" : 201301,
				"state" : "Uttar Pradesh"
			}
		],
		"sponsers" : {
			"name" : "One"
		},
		"revenue" : "19.8 million$",
		"no_of_employee" : 630,
		"str_text": ["one","two"],
		"int_text" : [1,3,4]
	}`

	var data map[string]interface{}
	json.Unmarshal([]byte(input), &data)

	ivalue(data, "")
}
