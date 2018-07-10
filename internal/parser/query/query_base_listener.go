// Code generated from internal/parser/query/Query.g4 by ANTLR 4.7.1. DO NOT EDIT.

package query // Query
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQueryListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryListener struct{}

var _ QueryListener = &BaseQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSchema is called when production schema is entered.
func (s *BaseQueryListener) EnterSchema(ctx *SchemaContext) {}

// ExitSchema is called when production schema is exited.
func (s *BaseQueryListener) ExitSchema(ctx *SchemaContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseQueryListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseQueryListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseQueryListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseQueryListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeBody is called when production typeBody is entered.
func (s *BaseQueryListener) EnterTypeBody(ctx *TypeBodyContext) {}

// ExitTypeBody is called when production typeBody is exited.
func (s *BaseQueryListener) ExitTypeBody(ctx *TypeBodyContext) {}

// EnterField is called when production field is entered.
func (s *BaseQueryListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseQueryListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseQueryListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseQueryListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseQueryListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseQueryListener) ExitFieldType(ctx *FieldTypeContext) {}
