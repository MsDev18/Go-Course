package main

import (
	"fmt"
	"io/fs"
	"testing"
)

var result string

func BenchmarkPrint(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fErr := &fs.PathError{
			Op: "open",
			Path: "./storage/date.txt",
			Err: fmt.Errorf("file not found ..."),
		}
		res := String(fErr)
		result = res
	}
}

func BenchmarkPrintTwo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fErr := &fs.PathError{
			Op: "open",
			Path: "./storage/date.txt",
			Err: fmt.Errorf("file not found ..."),
		}
		res := StringTwo(fErr)
		result = res
	}
}

var jsonRes []byte
func BenchmarkJson(b *testing.B) {
	data := SimpleData{
		ID: 10,
		Name: "MsDev18",
		Email: "MsDev18",
	}
	for i := 0 ; i < b.N ; i ++ {
		res, _ := Json(data)
		jsonRes = res
	}
}

func BenchmarkJsonTwo(b *testing.B) {
	data := SimpleDataTwo{
		ID: 10,
		Name: "MsDev18",
		Email: "MsDev18",
	}
	for i := 0 ; i < b.N ; i ++ {
		res, _ := JsonTwo(data)
		jsonRes = res
	}
}