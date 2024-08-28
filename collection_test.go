package collection

import (
	"fmt"
	"testing"
)

func TestKeyBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	people := []*Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}
	result := KeyBy(people, "Name").(map[string]*Person)
	if result["Alice"].Age != 25 {
		t.Errorf("Expected Alice to be 25 years old, but got %d", result["Alice"].Age)
	}
	if result["Bob"].Age != 30 {
		t.Errorf("Expected Bob to be 30 years old, but got %d", result["Bob"].Age)
	}
	if result["Charlie"].Age != 35 {
		t.Errorf("Expected Charlie to be 35 years old, but got %d", result["Charlie"].Age)
	}
}

func TestKeyBySlice(t *testing.T) {
	type Person struct {
		Name string
		Id   int
	}
	people := []*Person{
		{
			Name: "Alice",
			Id:   1,
		},
		{
			Name: "Alice345",
			Id:   1,
		},
		{
			Name: "Bill",
			Id:   3,
		},
	}
	m := KeyBySlice(people, "Id").(map[int][]*Person)
	fmt.Println(m[1][0])

	for i, v := range m {

		fmt.Println(i, v)
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
