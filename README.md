# Godocx

Godocx is a library written in pure Go providing a set of functions that allow you to write to and read from Docx file. This library needs Go version 1.24 or later

This library is a fork of unmaintened ![gomutex/godocx](https://github.com/gomutex/godocx)

## Usage

Here's a simple example of how you can use Godocx to create and modify DOCX documents:

## Installation

```bash
go get github.com/ubavic/godocx
```

### Examples

Explore additional examples and use cases over at GitHub repository dedicated to showcasing the capabilities of Golang Docx:
https://github.com/ubavic/godocx-examples


```go
package main

import (
	"log"

	"github.com/ubavic/godocx"
)

func main() {
	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	document.AddHeading("Document Title", 0)

	// Add a new paragraph to the document
	p := document.AddParagraph("A plain paragraph having some ")
	p.AddText("bold").Bold(true)
	p.AddText(" and some ")
	p.AddText("italic.").Italic(true)

	document.AddHeading("Heading, level 1", 1)
	document.AddParagraph("Intense quote").Style("Intense Quote")
	document.AddParagraph("first item in unordered list").Style("List Bullet")
	document.AddParagraph("first item in ordered list").Style("List Number")

	records := []struct{ Qty, ID, Desc string }{{"5", "A001", "Laptop"}, {"10", "B202", "Smartphone"}, {"2", "E505", "Smartwatch"}}

	table := document.AddTable()
	table.Style("LightList-Accent4")
	hdrRow := table.AddRow()
	hdrRow.AddCell().AddParagraph("Qty")
	hdrRow.AddCell().AddParagraph("ID")
	hdrRow.AddCell().AddParagraph("Description")

	for _, record := range records {
		row := table.AddRow()
		row.AddCell().AddParagraph(record.Qty)
		row.AddCell().AddParagraph(record.ID)
		row.AddCell().AddParagraph(record.Desc)
	}

	err = document.SaveTo("demo.docx")
	if err != nil {
		log.Fatal(err)
	}
}
```

## Demo Output

This is screenshot of demo document generated from the godocx library. 

![Screenshot of the demo output](https://github.com/ubavic/godocx-examples/raw/main/demo.png)

## Licenses

The Godocx library is licensed under the [MIT License](https://opensource.org/licenses/MIT).
