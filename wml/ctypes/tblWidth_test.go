package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/ubavic/godocx/wml/stypes"
)

func TestTableWidth_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    TableWidth
		expected string
	}{
		{
			name:     "Test with Width only",
			input:    TableWidth{Width: new(500)},
			expected: `<w:tblW w:w="500"></w:tblW>`,
		},
		{
			name:     "Test with Type only",
			input:    TableWidth{WidthType: new(stypes.TableWidthDxa)},
			expected: `<w:tblW w:type="dxa"></w:tblW>`,
		},
		{
			name:     "Test with Width and Type",
			input:    TableWidth{Width: new(1000), WidthType: new(stypes.TableWidthAuto)},
			expected: `<w:tblW w:w="1000" w:type="auto"></w:tblW>`,
		},
		{
			name:     "Test with nil values",
			input:    TableWidth{},
			expected: `<w:tblW></w:tblW>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:tblW"}}

			err := tt.input.MarshalXML(encoder, start)
			if err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if err = encoder.Flush(); err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			if result.String() != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, result.String())
			}
		})
	}
}

func TestTableWidth_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   TableWidth
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Width attribute",
			inputXML: `<w:tblW w:w="750"></w:tblW>`,
			expected: TableWidth{Width: new(750)},
		},
		{
			name:     "Test with Type attribute",
			inputXML: `<w:tblW w:type="dxa"></w:tblW>`,
			expected: TableWidth{WidthType: new(stypes.TableWidthDxa)},
		},
		{
			name:     "Test with Width and Type attributes",
			inputXML: `<w:tblW w:w="500" w:type="pct"></w:tblW>`,
			expected: TableWidth{Width: new(500), WidthType: new(stypes.TableWidthPct)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TableWidth
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected TableWidth %+v but got %+v", tt.expected, result)
			}
		})
	}
}
