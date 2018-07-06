package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grogdb/grogdb/internal/parser/schema"
	"strconv"
)

func NewParseErrorListener() *ParseErrorListener {
	return &ParseErrorListener{}
}

type ParseErrorListener struct {
	antlr.DefaultErrorListener
	errors []error
}

func (l *ParseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	err := fmt.Errorf("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
	l.errors = append(l.errors, err)
}

func (l ParseErrorListener) Err() error {
	if len(l.errors) > 0 {
		return l.errors[0]
	}
	return nil
}

func ParseSchema(raw []byte) (*Schema, error) {
	input := antlr.NewInputStream(string(raw))
	lexer := schema.NewSchemaLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	p := schema.NewSchemaParser(tokenStream)
	errorListener := NewParseErrorListener()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true

	listener := &schemaListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Schema())

	if errorListener.Err() != nil {
		return nil, errorListener.Err()
	}

	return listener.schema, nil
}

type Schema struct {
	Types []Type
}

type Type struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
}

type schemaListener struct {
	schema   *Schema
	tmpTypes []Type
	tmpType  *Type
}

// EnterTypeName is called when entering the typeName production.
func (l *schemaListener) EnterTypeName(ctx *schema.TypeNameContext) {
	l.tmpType = &Type{
		Name: ctx.TypeIdent().GetText(),
	}
}

// ExitSchema is called when exiting the schema production.
func (l *schemaListener) ExitSchema(ctx *schema.SchemaContext) {
	l.schema = &Schema{
		Types: l.tmpTypes,
	}
	l.tmpTypes = nil
}

// ExitTypeDecl is called when exiting the typeDecl production.
func (l *schemaListener) ExitTypeDecl(ctx *schema.TypeDeclContext) {
	l.tmpTypes = append(l.tmpTypes, *l.tmpType)
	l.tmpType = nil
}

// ExitField is called when exiting the field production.
func (l *schemaListener) ExitField(ctx *schema.FieldContext) {
	field := &Field{
		Name: ctx.FieldName().GetText(),
		Type: ctx.FieldType().GetText(),
	}
	l.tmpType.Fields = append(l.tmpType.Fields, *field)
}

func (l *schemaListener) VisitErrorNode(node antlr.ErrorNode) {
	// NOP
}

// EnterSchema is called when entering the schema production.
func (l *schemaListener) EnterSchema(ctx *schema.SchemaContext) {
	// NOP
}

// EnterTypeDecl is called when entering the typeDecl production.
func (l *schemaListener) EnterTypeDecl(ctx *schema.TypeDeclContext) {
	// NOP
}

// EnterTypeBody is called when entering the typeBody production.
func (l *schemaListener) EnterTypeBody(ctx *schema.TypeBodyContext) {
	// NOP
}

// EnterField is called when entering the field production.
func (l *schemaListener) EnterField(ctx *schema.FieldContext) {
	// NOP
}

// EnterFieldName is called when entering the fieldName production.
func (l *schemaListener) EnterFieldName(ctx *schema.FieldNameContext) {
	// NOP
}

// EnterFieldType is called when entering the fieldType production.
func (l *schemaListener) EnterFieldType(ctx *schema.FieldTypeContext) {
	// NOP
}

// ExitTypeName is called when exiting the typeName production.
func (l *schemaListener) ExitTypeName(ctx *schema.TypeNameContext) {
	// NOP
}

// ExitTypeBody is called when exiting the typeBody production.
func (l *schemaListener) ExitTypeBody(ctx *schema.TypeBodyContext) {
	// NOP
}

// ExitFieldName is called when exiting the fieldName production.
func (l *schemaListener) ExitFieldName(ctx *schema.FieldNameContext) {
	// NOP
}

// ExitFieldType is called when exiting the fieldType production.
func (l *schemaListener) ExitFieldType(ctx *schema.FieldTypeContext) {
	// NOP
}

func (l *schemaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// NOP
}

func (l *schemaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// NOP
}

func (l *schemaListener) VisitTerminal(node antlr.TerminalNode) {
	// NOP
}
