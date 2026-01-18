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

func TestIsFileContentEmpty(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected bool
	}{
		{
			name:     "Empty content",
			content:  []byte{},
			expected: true,
		},
		{
			name:     "Only newline",
			content:  []byte("\n"),
			expected: true,
		},
		{
			name:     "Only whitespace",
			content:  []byte("  \n\t\n  "),
			expected: true,
		},
		{
			name:     "Content with block",
			content:  []byte("resource \"aws_instance\" \"example\" {}\n"),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			empty := isFileContentEmpty(test.content)
			if empty != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, empty)
			}
		})
	}
}
