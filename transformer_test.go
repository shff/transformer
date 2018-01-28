package transformer

import (
	"testing"
)

func TestShouldTransformJSON(t *testing.T) {
	data := `{ "response": { "id": 12, "data": { "value": "x" }}}`
	from := `{ "response": { "id": "{{id}}", "data": { "value": "{{value}}" }}}`
	to := `<response id="{{.id}}"><value>{{.value}}</value></response>`

	transformed, err := Transform(data, from, to, JSONParser)
	if err != nil {
		panic(err)
	}
	if transformed != `<response id="12"><value>x</value></response>` {
		t.Fail()
	}
}

func TestShouldTransformXML(t *testing.T) {
	data := `<response id="12"><value>x</value></response>`
	from := `<response id="{{id}}"><value>{{value}}</value></response>`
	to := `{ "response": { "id": {{.id}}, "data": { "value": "{{.value}}" }}}`

	transformed, err := Transform(data, from, to, XMLParser)
	if err != nil {
		panic(err)
	}
	if transformed != `{ "response": { "id": 12, "data": { "value": "x" }}}` {
		print(transformed)
		t.Fail()
	}
}
