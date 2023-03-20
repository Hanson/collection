package collection

import (
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
