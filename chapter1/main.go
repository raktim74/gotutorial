package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type person struct {
	Name string
	Age  int
	sal  int
	job  string
}

func main() {
	fmt.Println("Baba1")
	a := 5
	var b string = "30"
	var bInt int
	var err error
	bInt, err = strconv.Atoi(b)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Sum: ", sum(a, bInt))
	fmt.Println("Concatenation: ", concat_values(strconv.Itoa(a), strconv.Itoa(bInt)))
	fmt.Println("Floating Values String: ", float_values(strconv.FormatFloat(float64(a), 'f', 1, 64), strconv.FormatFloat(float64(bInt), 'f', 1, 64)))
	af := float64(a)
	bf := float64(bInt)
	result := float_values_actual(af, bf)
	fmt.Println("Float: ", strconv.FormatFloat(result, 'f', 1, 64))

	//array
	arr := [5]any{1, 2, 3, 4, "bba"}
	arr1 := []any{1, 2, 3, 4, "bba"}
	fmt.Println("Array: ", array_data(arr[:], 4))
	fmt.Println("Array: ", array_data(arr1, 3))

	fmt.Println("Array1: ", array_data_reflect(arr[:], 4))
	fmt.Println("Array1: ", array_data_reflect(arr1, 3))

	p := person{Name: "Raktim", Age: 30, sal: 200, job: "Engineer"}
	fmt.Println(p.Greet())

	return_data := Hello(p)
	retrunedPerson, ok := return_data.(person)
	if ok {
		fmt.Println("data ", retrunedPerson.sal)
	} else {
		fmt.Println("Not a person type")
	}

	//date format
	dateStr := "2025-01-28 16:30:45"
	layout := "2006-01-02 15:04:05"
	// parsedDate, err := time.Parse(layout, dateStr)
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date", err)
		return
	}
	fmt.Println("Date: ", parsedDate)

	todayDate := time.Now()
	formatted := todayDate.Format("02/01/2006 15:04:05")
	fmt.Println("Date Fomatted: ", formatted)
	nextdate := todayDate.Add(24 * time.Hour)
	// nextdate := todayDate.AddDate(0, 1, 0).Format("02/01/2006 15:04:05")
	fmt.Println("Next Date: ", nextdate)
	prevdate := todayDate.AddDate(-1, 0, 0).Format("02/01/2006 15:04:05")
	fmt.Println("Prev Date: ", prevdate)

	timezoneEx, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	currentTime := time.Now().In(timezoneEx)
	fmt.Println("Tiemzone Ex ", currentTime)

}

func (p person) Greet() string {
	return "Hello: " + p.Name
}

func Hello(obj any) any {
	return obj
}

func sum(x, y int) any {
	return x + y
}

func concat_values(x, y string) string {
	return x + y
}
func float_values(x, y string) string {
	return x + y
}

func float_values_actual(x, y float64) float64 {
	return x + y
}

func array_data[T any](arr []T, indx int) T {
	return arr[indx]
}

// using reflect
func array_data_reflect(arr any, indx int) any {
	val := reflect.ValueOf(arr)
	// check if arr is slice or array
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		panic("Arr must be slice or Array")
	}
	// bounds check
	if indx < 0 || indx >= val.Len() {
		panic("index out of range")
	}
	// get element at indx and interface
	return val.Index(indx).Interface()
}
