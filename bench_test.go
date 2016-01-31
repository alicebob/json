package json

import (
	"encoding/json"
	"testing"
)

var stringV = `"foo"`

func BenchmarkStringUs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			tv string
		)
		if err := Decode(stringV, &tv); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStringThem(b *testing.B) {
	v := []byte(stringV)
	for i := 0; i < b.N; i++ {
		var (
			tv string
		)
		if err := json.Unmarshal(v, &tv); err != nil {
			b.Fatal(err)
		}
	}
}

var intV = `12345`

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			tv int
		)
		if err := Decode(intV, &tv); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIntThem(b *testing.B) {
	v := []byte(intV)
	for i := 0; i < b.N; i++ {
		var (
			tv int
		)
		if err := json.Unmarshal(v, &tv); err != nil {
			b.Fatal(err)
		}
	}
}

type str struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
	Baz string `json:"baz"`
	N   int    `json:"n"`
}

var structV = `{"foo":"foov","bar":"barv","n":101}`

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			tv str
		)
		if err := Decode(structV, &tv); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStructThem(b *testing.B) {
	v := []byte(structV)
	for i := 0; i < b.N; i++ {
		var (
			tv str
		)
		if err := json.Unmarshal(v, &tv); err != nil {
			b.Fatal(err)
		}
	}
}
