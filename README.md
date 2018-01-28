# Transformer

Parses XML and JSON payloads using "reverse templates" and applies a Golang template in it.

This is useful for API Gateways and quick prototypes.

## Examples

Remember that you can convert from JSON and XML to any textual format you want.

### JSON to XML

```go
import "github.com/shferreira/transformer"

data := `{ "response": { "id": 12, "data": { "value": "x" }}}`
from := `{ "response": { "id": "{{id}}", "data": { "value": "{{value}}" }}}`
to := `<response id="{{.id}}"><value>{{.value}}</value></response>`

transformed, err := transformer.Transform(data, from, to, transformer.JSONParser)
```

```xml
<response id="12"><value>x</value></response>
```

### XML to JSON

```go
import "github.com/shferreira/transformer"

data := `<response id="12"><value>x</value></response>`
from := `<response id="{{id}}"><value>{{value}}</value></response>`
to := `{ "response": { "id": {{.id}}, "data": { "value": "{{.value}}" }}}`

transformed, err := transformer.Transform(data, from, to, transformer.XMLParser)
```

```js
{ "response": { "id": 12, "data": { "value": "x" }}}
```
