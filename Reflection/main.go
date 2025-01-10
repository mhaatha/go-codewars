package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// Kind
	// Digunakan untuk melihat tipe dasar sebuah value
	// Bedanya dengan TypeOf adalah, TypeOf melihat tipe secara spesifik
	// Contoh:
	// var x MyInt
	// t := reflect.TypeOf(x)
	// fmt.Println(t)         // Output: main.MyInt (type spesifik)
	// fmt.Println(t.Kind())  // Output: int (kategori dasar)
	for _, v := range []any{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}

	// TypeOf
	// As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))
}
