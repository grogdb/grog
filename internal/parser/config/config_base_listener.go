// Code generated from internal/parser/config/Config.g4 by ANTLR 4.7.1. DO NOT EDIT.

package config // Config
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConfigListener is a complete listener for a parse tree produced by ConfigParser.
type BaseConfigListener struct{}

var _ ConfigListener = &BaseConfigListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConfigListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConfigListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConfigListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConfigListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseConfigListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseConfigListener) ExitDocument(ctx *DocumentContext) {}

// EnterOptionsBlock is called when production optionsBlock is entered.
func (s *BaseConfigListener) EnterOptionsBlock(ctx *OptionsBlockContext) {}

// ExitOptionsBlock is called when production optionsBlock is exited.
func (s *BaseConfigListener) ExitOptionsBlock(ctx *OptionsBlockContext) {}

// EnterOptionsBody is called when production optionsBody is entered.
func (s *BaseConfigListener) EnterOptionsBody(ctx *OptionsBodyContext) {}

// ExitOptionsBody is called when production optionsBody is exited.
func (s *BaseConfigListener) ExitOptionsBody(ctx *OptionsBodyContext) {}

// EnterSchemaBlock is called when production schemaBlock is entered.
func (s *BaseConfigListener) EnterSchemaBlock(ctx *SchemaBlockContext) {}

// ExitSchemaBlock is called when production schemaBlock is exited.
func (s *BaseConfigListener) ExitSchemaBlock(ctx *SchemaBlockContext) {}

// EnterSchemaBody is called when production schemaBody is entered.
func (s *BaseConfigListener) EnterSchemaBody(ctx *SchemaBodyContext) {}

// ExitSchemaBody is called when production schemaBody is exited.
func (s *BaseConfigListener) ExitSchemaBody(ctx *SchemaBodyContext) {}

// EnterSchemaTypes is called when production schemaTypes is entered.
func (s *BaseConfigListener) EnterSchemaTypes(ctx *SchemaTypesContext) {}

// ExitSchemaTypes is called when production schemaTypes is exited.
func (s *BaseConfigListener) ExitSchemaTypes(ctx *SchemaTypesContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseConfigListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseConfigListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseConfigListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseConfigListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterTypeBody is called when production typeBody is entered.
func (s *BaseConfigListener) EnterTypeBody(ctx *TypeBodyContext) {}

// ExitTypeBody is called when production typeBody is exited.
func (s *BaseConfigListener) ExitTypeBody(ctx *TypeBodyContext) {}

// EnterField is called when production field is entered.
func (s *BaseConfigListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseConfigListener) ExitField(ctx *FieldContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseConfigListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseConfigListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseConfigListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseConfigListener) ExitFieldType(ctx *FieldTypeContext) {}
