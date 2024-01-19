package main

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

func TestRemoveBlocks(t *testing.T) {
	tests := []struct {
		name     string
		body     *hclwrite.Body
		expected bool
	}{
		{
			name: "Remove single block",
			body: func() *hclwrite.Body {
				file := hclwrite.NewEmptyFile()
				body := file.Body()
				body.AppendNewBlock("moved", nil)
				return body
			}(),
			expected: true,
		},
		{
			name: "Remove multiple blocks",
			body: func() *hclwrite.Body {
				file := hclwrite.NewEmptyFile()
				body := file.Body()
				body.AppendNewBlock("moved", nil)
				body.AppendNewBlock("import", nil)
				body.AppendNewBlock("removed", nil)
				return body
			}(),
			expected: true,
		},
		{
			name: "No blocks to remove",
			body: func() *hclwrite.Body {
				file := hclwrite.NewEmptyFile()
				body := file.Body()
				body.AppendNewBlock("other", nil)
				return body
			}(),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updated := RemoveBlocks(test.body)
			if updated != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, updated)
			}
		})
	}
}
