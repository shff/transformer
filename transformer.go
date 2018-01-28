package transformer

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"strconv"
	"strings"
)

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func Transform(content, sample, tmpl string, parser func(string) (map[string]string, error)) (string, error) {
	contentTree, err := parser(content)
	if err != nil {
		return "", err
	}

	sampleTree, err := parser(sample)
	if err != nil {
		return "", err
	}

	values := make(map[string]string)
	for k, v := range sampleTree {
		if strings.HasPrefix(v, "{{") {
			values[trimString(v)] = trimString(contentTree[k])
		}
	}

	t, err := template.New("transformer").Parse(tmpl)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(nil)
	err = t.Execute(buf, values)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func JSONParser(content string) (map[string]string, error) {
	var document map[string]interface{}
	err := json.Unmarshal([]byte(content), &document)
	if err != nil {
		return nil, err
	}

	children := make(map[string]string)
	walkJSON(document, children, "")

	return children, nil
}

func walkJSON(node map[string]interface{}, values map[string]string, path string) {
	for k, v := range node {
		newPath := path + "." + k
		switch v2 := v.(type) {
		case map[string]interface{}:
			walkJSON(v2, values, newPath)
		case string:
			values[newPath] = v2
		case float64:
			values[newPath] = strconv.FormatFloat(v2, 'f', -1, 64)
		}
	}
}

func XMLParser(content string) (map[string]string, error) {
	var document Node
	err := xml.Unmarshal([]byte(content), &document)
	if err != nil {
		return nil, err
	}

	children := make(map[string]string)
	walkXML(document, children, "")

	return children, nil
}

func walkXML(node Node, children map[string]string, path string) {
	children[path] = string(node.Content)
	for _, a := range node.Attrs {
		children[path+"._"+a.Name.Local] = a.Value
	}
	for _, n := range node.Nodes {
		walkXML(n, children, path+"."+n.XMLName.Local)
	}
}

func trimString(s string) string {
	return strings.TrimSpace(strings.Replace(strings.Replace(s, "}}", "", -1), "{{", "", -1))
}
