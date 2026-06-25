package godocx

import (
	_ "embed"
	"os"
	"path/filepath"

	"github.com/gomutex/godocx/docx"
	"github.com/gomutex/godocx/packager"
)

//go:embed templates/default.docx
var defaultDocx []byte

// Creates a new document from the default template.
func NewDocument() (*docx.RootDoc, error) {
	return packager.Unpack(&defaultDocx)
}

// Creates a new document from a valid in-memory representation.
func ParseDocument(doc []byte) (*docx.RootDoc, error) {
	return packager.Unpack(&doc)
}

// Opens a document from the given file name.
//
// Deprecated: Read file by yourself and use `ParseDocument`
func OpenDocument(fileName string) (*docx.RootDoc, error) {
	docxContent, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}
	return packager.Unpack(&docxContent)
}
