# collection

Coming soon..

## Support function

* [KeyBy](#KeyBy)
* [PluckUint64](#PluckUint64)

## Usage

### KeyBy

```
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

result["Alice"].Age // 25
```

### PluckUint64

```
type testStruct struct {
	ID   uint64
	Name string
}

tests := []*testStruct{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
}

result := PluckUint64(tests, "ID")
// result: []uint64{1, 2, 3}
```