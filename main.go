package main

import (
	"ShiftAlt/marketstack/api"
	"ShiftAlt/marketstack/sheets"
	"reflect"
	"strconv"
)

func main() {
	MarketstackResponse := api.GetMarketStackData()
	i := 2
	headerAdded := false
	for _, elem := range MarketstackResponse.Data.Eod {
		myval := []interface{}{}
		heading := []interface{}{}

		v := reflect.ValueOf(elem)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {
			// fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
			myval = append(myval, v.Field(i).Interface())
			heading = append(heading, typeOfS.Field(i).Name)
		}
		if !headerAdded {
			writeRange := "A1:M1"
			sheets.Write(writeRange, heading)
			headerAdded = true
		}
		writeRange := "A" + strconv.Itoa(i) + ":M" + strconv.Itoa(i)
		sheets.Write(writeRange, myval)
		i++
		// sheets.Read("A1:A1")
	}
}
