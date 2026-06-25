package ctypes

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ubavic/godocx/wml/stypes"
)

func TestSpacing_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		input    Spacing
		expected string
	}{
		{
			name: "All fields set",
			input: Spacing{
				Before:            new(uint64(120)),
				After:             new(uint64(240)),
				BeforeLines:       new(10),
				BeforeAutospacing: new(stypes.OnOffOn),
				AfterAutospacing:  new(stypes.OnOffOff),
				Line:              new(360),
				LineRule:          new(stypes.LineSpacingRuleExact),
			},
			expected: `<w:spacing w:before="120" w:after="240" w:beforeLines="10" w:beforeAutospacing="on" w:afterAutospacing="off" w:line="360" w:lineRule="exact"></w:spacing>`,
		},
		{
			name: "Some fields set",
			input: Spacing{
				Before:   new(uint64(120)),
				After:    new(uint64(240)),
				Line:     new(360),
				LineRule: new(stypes.LineSpacingRuleAuto),
			},
			expected: `<w:spacing w:before="120" w:after="240" w:line="360" w:lineRule="auto"></w:spacing>`,
		},
		{
			name:     "No fields set",
			input:    Spacing{},
			expected: `<w:spacing></w:spacing>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result strings.Builder
			encoder := xml.NewEncoder(&result)

			start := xml.StartElement{Name: xml.Name{Local: "w:spacing"}}
			if err := tt.input.MarshalXML(encoder, start); err != nil {
				t.Fatalf("Error marshaling XML: %v", err)
			}

			// Finalize encoding
			encoder.Flush()

			got := strings.TrimSpace(result.String())
			if got != tt.expected {
				t.Errorf("Expected XML:\n%s\nGot:\n%s", tt.expected, got)
			}
		})
	}
}
