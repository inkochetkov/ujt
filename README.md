# ujt

Unmarshaled json to need struct

valid and check in slice byte

## Method

#### UnmarshaledJSONToMap

UnmarshaledJSONToMap(b []byte) (map[string]any, error)

```
checks the structure for validity and converts it to a map
```

#### UnmarshaledJSONToArr

UnmarshaledJSONToArr(b []byte) ([]any, error)

```
checks the structure for validity and converts it to a slice
```

#### UnmarshaledJSONToCustom

UnmarshaledJSONToCustom(b []byte, customStruct any) (any, error)

```
checks the structure for validity and converts the structure passed as the second argument
```
