package ast

import (
	"testing"
)

func TestNodeCreation(t *testing.T) {
	// Example test for creating a node
	node := &Node{
		Type:  "ExampleNode",
		Value: "TestValue",
	}

	if node.Type != "ExampleNode" {
		t.Errorf("Expected node type 'ExampleNode', got '%s'", node.Type)
	}
	if node.Value != "TestValue" {
		t.Errorf("Expected node value 'TestValue', got '%s'", node.Value)
	}
}

func TestASTTraversal(t *testing.T) {
	// Example test for traversing the AST
	root := &Node{
		Type: "Root",
		Children: []*Node{
			&Node{Type: "Child1"},
			&Node{Type: "Child2"},
		},
	}

	if len(root.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(root.Children))
	}
}