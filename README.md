# collection

Coming soon ...

## Support function

* [KeyBy](#KeyBy)
* [KeyBySlice](#KeyBySlice)
* [PluckUint64](#PluckUint64)
* [PluckString](#PluckString)

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

### KeyBySlice

```
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
		Name: "Bill",
		Id:   1,
	},
	{
		Name: "Bill",
		Id:   3,
	},
}

result := KeyBySlice(people, "Id").(map[int][]*Person)
result[1] // [0xc000008120 0xc000008138]
result[1][0] // &{Alice 1}
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

## Respect

https://laravel.com/docs/10.x/collections
https://github.com/elliotchance/pie (Most of the code is copied from here.)

## License
MIT