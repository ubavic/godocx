package ctypes

import (
	"encoding/xml"
	"reflect"
	"strings"
	"testing"

	"github.com/ubavic/godocx/wml/stypes"
)

func TestTableRowHeight_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    TableRowHeight
		expected string
	}{
		{
			name:     "Test with Val only",
			input:    TableRowHeight{Val: new(500)},
			expected: `<w:val w:val="500"></w:val>`,
		},
		{
			name:     "Test with HRule only",
			input:    TableRowHeight{HRule: new(stypes.HeightRuleAtLeast)},
			expected: `<w:val w:hRule="atLeast"></w:val>`,
		},
		{
			name:     "Test with Val and HRule",
			input:    TableRowHeight{Val: new(1000), HRule: new(stypes.HeightRuleExact)},
			expected: `<w:val w:val="1000" w:hRule="exact"></w:val>`,
		},
		{
			name:     "Test with nil values",
			input:    TableRowHeight{},
			expected: `<w:val></w:val>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)
			start := xml.StartElement{Name: xml.Name{Local: "w:val"}}

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

func TestTableRowHeight_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name       string
		inputXML   string
		expected   TableRowHeight
		expectFail bool // Whether unmarshalling is expected to fail
	}{
		{
			name:     "Test with Val attribute",
			inputXML: `<w:val w:val="750"></w:val>`,
			expected: TableRowHeight{Val: new(750)},
		},
		{
			name:     "Test with HRule attribute",
			inputXML: `<w:val w:hRule="auto"></w:val>`,
			expected: TableRowHeight{HRule: new(stypes.HeightRuleAuto)},
		},
		{
			name:     "Test with Val and HRule attributes",
			inputXML: `<w:val w:val="500" w:hRule="atLeast"></w:val>`,
			expected: TableRowHeight{Val: new(500), HRule: new(stypes.HeightRuleAtLeast)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result TableRowHeight
			err := xml.Unmarshal([]byte(tt.inputXML), &result)

			if err != nil {
				t.Fatalf("Error unmarshaling XML: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected TableRowHeight %+v but got %+v", tt.expected, result)
			}
		})
	}
}
