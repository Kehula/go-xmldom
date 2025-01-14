package xmldom

import (
	"bytes"
)

const (
	DEFAULT_XML_HEADER = `<?xml version="1.0" encoding="UTF-8"?>`
)

func NewDocument(name string) *Document {
	d := &Document{
		ProcInst: DEFAULT_XML_HEADER,
	}
	d.Root = &Node{
		Document: d,
		Name:     name,
	}
	return d
}

type Document struct {
	ProcInst      string
	Directives    []string
	NamespaceList []*Namespace
	Root          *Node
}

func (d *Document) getNamespaceByURI(uri string) *Namespace {
	if uri == "" {
		return nil
	}
	if d.NamespaceList != nil {
		for _, ns := range d.NamespaceList {
			if ns.URI == uri {
				return ns
			}
		}
	}

	// create a new namespace
	ns := &Namespace{
		Name: uri,
		URI:  uri,
	}
	d.NamespaceList = append(d.NamespaceList, ns)

	return ns
}

func (d *Document) XML() string {
	buf := new(bytes.Buffer)
	buf.WriteString(d.ProcInst)
	for _, directive := range d.Directives {
		buf.WriteString(directive)
	}
	buf.WriteString(d.Root.XML())
	return buf.String()
}

func (d *Document) XMLPretty() string {
	buf := new(bytes.Buffer)
	if len(d.ProcInst) > 0 {
		buf.WriteString(d.ProcInst)
		buf.WriteByte('\n')
	}
	for _, directive := range d.Directives {
		buf.WriteString(directive)
		buf.WriteByte('\n')
	}
	buf.WriteString(d.Root.XMLPretty())
	return buf.String()
}

func (d *Document) XMLPrettyEx(indent string) string {
	buf := new(bytes.Buffer)
	if len(d.ProcInst) > 0 {
		buf.WriteString(d.ProcInst)
		buf.WriteByte('\n')
	}
	for _, directive := range d.Directives {
		buf.WriteString(directive)
		buf.WriteByte('\n')
	}
	buf.WriteString(d.Root.XMLPrettyEx(indent))
	return buf.String()
}
