# Transformer

Parses XML and JSON payloads using "reverse templates" and applies a Golang template in it.

This is useful for API Gateways and quick prototypes.

## Installation

```bash
go get github.com/shferreira/transformer
```

## Usage

First of all, import the library:

```
import "github.com/shferreira/transformer"
```

### From JSON to any format

The `JSONParser` parses JSON objects. It doesn't understand arrays, though.

```go
data := `{ "id": 12, "value": "x" }`
from := `{ "id": "{{id}}", "value": "{{value}}" }`
to := `<data id="{{.id}}">{{.value}}</data>`

transformed, err := transformer.Transform(data, from, to, transformer.JSONParser)
```

```xml
<data id="12">x</data>
```

### From XML to any format

Likewise, the `XMLParser` understand XML values and attributes. It doesn't understand namespaces.

```go
data := `<data id="12">x</response>`
from := `<data id="{{id}}">{{value}}</response>`
to := `{ "id": {{.id}}, "value": "{{.value}}" }`

transformed, err := transformer.Transform(data, from, to, transformer.XMLParser)
```

```js
{ "id": 12, "value": "x" }
```

# License

```
MIT License

Copyright (c) 2018 Silvio Henrique Ferreira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
