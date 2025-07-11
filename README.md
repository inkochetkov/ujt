# ujt

Unmarshaled json to need struct

valid and check in slice byte

## Method

#### UnmarshaledJSONToMap(b []byte) (map[string]any, error)

checks the structure for validity and converts it to a map

```
dataMap,err:=ujt.UnmarshaledJSONToMap([]byte{})
if err!=nil{
    ...
}
```

#### UnmarshaledJSONToArr(b []byte) ([]any, error)

checks the structure for validity and converts it to a slice

```
dataArr,er:=ujt.UnmarshaledJSONToMap([]byte{})
if err!=nil{
    ...
}
```

#### UnmarshaledJSONToCustom(b []byte, customStruct any) (any, error)

checks the structure for validity and converts the structure passed as the second argument

```
type youStruct struct{
    ...
}

dataYouStruct,err:=ujt.UnmarshaledJSONToMap([]byte{}, youStruct)
if err!=nil{
    ...
}
```
