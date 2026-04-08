package tree_sitter_typhoon_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_typhoon "github.com/koneko096/tree-sitter-typhoon/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_typhoon.Language())
	if language == nil {
		t.Errorf("Error loading Typhoon grammar")
	}
}
