package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("hello world")
	printHello("kara", 5)
}

func printHello(name string, timezone int) {
	// fmt.Println("hello" + name)

	// var timezoneStr string
	// if timezone == 5 {
	// 	timezoneStr = "EST"
	// }
	tz, err := getTimeZone(timezone)
	if err != nil {
		// msg := fmt.Sprintf("error getting time zone: %s", err)
		// fmt.Println(msg)
		fmt.Printf("error getting time zone: %s", err.Error())
		return
	}

	greeting := fmt.Sprintf("hello %s, my time zone is %s", name, tz)

	fmt.Println(greeting)
}

// error zero value = nil
func getTimeZone(tz int) (string, error) {
	// if tz == 5 {
	// 	return "EST", nil
	// }

	allTimeZones := getTimeZones()
	resp, ok := allTimeZones[tz]
	if !ok {
		return "", errors.New("invalid tz")
	}

	return resp, nil
}

func getTimeZones() map[int]string {
	m := make(map[int]string) // declare + initialize map

	// var m map[int]string // this is a zero value map, is not initialized

	m[5] = "EST"

	return m
}

/* built in types
strings = ""
int = 0
boolean = false
struct{} = nil

[]interface{} = nil -> []string, []int, []bool,
map[interface]interface = nil // map[string]int, map[string]string
*/

/* built in functions
make() - declare + initialize slices and maps
defer() -
go() -

*/

/*
%d
%.2f
*/
