package collection

import (
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
