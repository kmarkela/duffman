## Examples of possible json formats:

1. 
```json
[{"fruit":"apple"},{"fruit":"banana"},{"fruit":"cherry"}]
```

2. 
```json
{"fruits": ["apple", "banana", "cherry"]}
```

3. 
```json
{
  "person": {
    "firstName": "John",
    "lastName": "Doe",
    "age": 30
  }
}
```

4. 
```json
{
  "company": {
    "name": "Tech Corp",
    "address": {
      "street": "123 Tech Road",
      "city": "Innovate City",
      "postalCode": "12345"
    },
    "employees": [
      {
        "firstName": "Alice",
        "lastName": "Johnson",
        "age": 28,
        "role": "Engineer"
      },
      {
        "firstName": "Bob",
        "lastName": "Brown",
        "age": 35,
        "role": "Manager"
      }
    ]
  }
}
```


Usage: 
```golang
func main() {
	jsonStr := `{
  "company": {
    "name": "Tech Corp",
    "address": {
      "street": "123 Tech Road",
      "city": "Innovate City",
      "postalCode": "12345"
    },
    "employees": [
      {
        "firstName": "Alice",
        "lastName": "Johnson",
        "age": 28,
        "role": "Engineer"
      },
      {
        "firstName": "Bob",
        "lastName": "Brown",
        "age": 35,
        "role": "Manager"
      }
    ]
  }
}`
	flatMap, err := Unmarshal(jsonStr)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	fmt.Println("Flattened Map:", flatMap)

	// Marshal flattened map back to JSON
	jsonData, err := Marshal(flatMap)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}
	fmt.Println("JSON Data:", string(jsonData))
}
```