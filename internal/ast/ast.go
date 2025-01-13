package ast

// NodeType represents the type of a node in the AST.
type NodeType int

const (
    ProgramNode NodeType = iota
    VarDeclarationNode
    IfStatementNode
    ForLoopNode
    PrintStatementNode
    BinaryExpressionNode
    LiteralNode
)

// Node is the interface that all AST nodes implement.
type Node interface {
    GetNodeType() NodeType
}

// Program represents the root of the AST.
type Program struct {
    Statements []Node
}

func (p *Program) GetNodeType() NodeType {
    return ProgramNode
}

// VarDeclaration represents a variable declaration in the AST.
type VarDeclaration struct {
    Name  string
    Value Node
}

func (v *VarDeclaration) GetNodeType() NodeType {
    return VarDeclarationNode
}

// IfStatement represents an if statement in the AST.
type IfStatement struct {
    Condition Node
    ThenBlock []Node
    ElseBlock []Node
}

func (i *IfStatement) GetNodeType() NodeType {
    return IfStatementNode
}

// ForLoop represents a for loop in the AST.
type ForLoop struct {
    Init Node
    Condition Node
    Increment Node
    Body []Node
}

func (f *ForLoop) GetNodeType() NodeType {
    return ForLoopNode
}

// PrintStatement represents a print statement in the AST.
type PrintStatement struct {
    Value Node
}

func (p *PrintStatement) GetNodeType() NodeType {
    return PrintStatementNode
}

// BinaryExpression represents a binary expression in the AST.
type BinaryExpression struct {
    Left  Node
    Operator string
    Right Node
}

func (b *BinaryExpression) GetNodeType() NodeType {
    return BinaryExpressionNode
}

// Literal represents a literal value in the AST.
type Literal struct {
    Value interface{}
}

func (l *Literal) GetNodeType() NodeType {
    return LiteralNode
}