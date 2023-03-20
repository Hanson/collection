# collection

Coming soon..

## Support function

* KeyBy

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