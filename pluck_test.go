package collection

import (
	"fmt"
	"github.com/elliotchance/pie/v2"
	"reflect"
	"testing"
)

func TestPluckUint64(t *testing.T) {
	type testStruct struct {
		ID   uint64
		Name string
	}
	testCases := []struct {
		name       string
		list       interface{}
		fieldName  string
		expected   []uint64
		shouldFail bool
	}{
		{
			name: "valid list",
			list: []testStruct{
				{ID: 1, Name: "John"},
				{ID: 2, Name: "Doe"},
			},
			fieldName:  "ID",
			expected:   []uint64{1, 2},
			shouldFail: false,
		},
		{
			name:       "invalid list",
			list:       "invalid",
			fieldName:  "ID",
			expected:   nil,
			shouldFail: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tc.shouldFail {
					t.Errorf("unexpected panic: %v", r)
				}
			}()

			result := PluckUint64(tc.list, tc.fieldName)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func TestPluckAndDiff(t *testing.T) {
	type person struct {
		ID   uint64
		Name string
	}

	people := []person{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Doe"},
		{ID: 3, Name: "Jane"},
		{ID: 4, Name: "Doe"},
	}

	fmt.Println(PluckUint64(people, "ID"))
	fmt.Println(PluckUint64(people, "ID").Diff([]uint64{3, 4, 5}))

	add, del := pie.OfNumeric(PluckUint64(people, "ID")).Diff([]uint64{3, 4, 5})
	fmt.Printf("%+v, %+v", add, del)
}

func TestPluckUint32(t *testing.T) {
	type person struct {
		ID   uint32
		Name string
	}

	people := []person{
		{ID: 1, Name: "John"},
		{ID: 2, Name: "Doe"},
		{ID: 3, Name: "Jane"},
		{ID: 4, Name: "Doe"},
	}

	fmt.Println(PluckUint32(people, "ID"))
}
