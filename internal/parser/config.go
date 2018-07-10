package parser

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grogdb/grogdb/internal/parser/config"
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

func ParseSchema(raw []byte) (*Config, error) {
	input := antlr.NewInputStream(string(raw))
	lexer := config.NewConfigLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	p := config.NewConfigParser(tokenStream)

	errorListener := NewParseErrorListener()
	p.AddErrorListener(errorListener)
	p.BuildParseTrees = true

	v := schemaVisitor{}
	_ = v.Visit(p.Document())

	// listener := &schemaVisitor{}
	// antlr.VisNewParseTreeWalker().BaseParseTreeVisitor{}ParseTreeWalkerDefault.Walk(listener, p.Schema())

	if errorListener.Err() != nil {
		return nil, errorListener.Err()
	}

	return nil, nil
}

// type Config struct {
// }
//
// type schemaVisitor struct {
// 	*config.BaseConfigVisitor
// }
//
// func (l *schemaVisitor) Visit(tree antlr.ParseTree) interface{} {
// 	return tree.Accept(l)
// }
//
// // func (l *schemaVisitor) VisitChildren(node antlr.RuleNode) interface{}     {
// // 	return node.Accept(l)
// // }
// //
// // func (l *schemaVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
// // 	return node.Accept(l)
// // }
// //
// // func (l *schemaVisitor) VisitErrorNode(node antlr.ErrorNode) interface{}   {
// // 	return node.Accept(l)
// // }
//
// // Visit a parse tree produced by ConfigParser#document.
// func (l *schemaVisitor) VisitDocument(ctx *config.DocumentContext) interface{} {
// 	fmt.Println(".VisitDocument")
// 	l.VisitSchemaBlock(ctx.SchemaBlock())
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#schemaBlock.
// func (l *schemaVisitor) VisitSchemaBlock(ctx *config.SchemaBlockContext) interface{} {
// 	fmt.Println(".VisitSchemaBlock")
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#schemaBody.
// func (l *schemaVisitor) VisitSchemaBody(ctx *config.SchemaBodyContext) interface{} {
// 	fmt.Println(".VisitSchemaBody")
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#schema.
// func (l *schemaVisitor) VisitSchemaTypes(ctx *config.SchemaTypesContext) interface{} {
// 	fmt.Println(".VisitSchemaTypes")
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#typeDecl.
// func (l *schemaVisitor) VisitTypeDecl(ctx *config.TypeDeclContext) interface{} {
// 	fmt.Println(".VisitTypeDecl")
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#typeName.
// func (l *schemaVisitor) VisitTypeName(ctx *config.TypeNameContext) interface{} {
// 	fmt.Println(".VisitTypeName")
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#typeBody.
// func (l *schemaVisitor) VisitTypeBody(ctx *config.TypeBodyContext) interface{} {
// 	fmt.Println(".VisitTypeBody")
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#field.
// func (l *schemaVisitor) VisitField(ctx *config.FieldContext) interface{} {
// 	fmt.Printf(".VisitField [name=%s, type=%s]\n", ctx.FieldName().GetText(), ctx.FieldType().GetText())
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#fieldName.
// func (l *schemaVisitor) VisitFieldName(ctx *config.FieldNameContext) interface{} {
// 	fmt.Println(".VisitFieldName")
// 	return nil
// }
//
// // Visit a parse tree produced by ConfigParser#fieldType.
// func (l *schemaVisitor) VisitFieldType(ctx *config.FieldTypeContext) interface{} {
// 	fmt.Println(".VisitFieldType")
// 	return nil
// }

//
// type Schema struct {
// 	Types []Type
// }
//
// type Type struct {
// 	Name   string
// 	Fields []Field
// }
//
// type Field struct {
// 	Name string
// 	Type string
// }
//
// type schemaListener struct {
// 	schema   *Schema
// 	tmpTypes []Type
// 	tmpType  *Type
// }
//
// // EnterTypeName is called when entering the typeName production.
// func (l *schemaListener) EnterTypeName(ctx *config.TypeNameContext) {
// 	l.tmpType = &Type{
// 		Name: ctx.TypeIdent().GetText(),
// 	}
// }
//
// // ExitSchema is called when exiting the schema production.
// func (l *schemaListener) ExitSchema(ctx *config.SchemaContext) {
// 	l.schema = &Schema{
// 		Types: l.tmpTypes,
// 	}
// 	l.tmpTypes = nil
// }
//
// // ExitTypeDecl is called when exiting the typeDecl production.
// func (l *schemaListener) ExitTypeDecl(ctx *config.TypeDeclContext) {
// 	l.tmpTypes = append(l.tmpTypes, *l.tmpType)
// 	l.tmpType = nil
// }
//
// // ExitField is called when exiting the field production.
// func (l *schemaListener) ExitField(ctx *config.FieldContext) {
// 	field := &Field{
// 		Name: ctx.FieldName().GetText(),
// 		Type: ctx.FieldType().GetText(),
// 	}
// 	l.tmpType.Fields = append(l.tmpType.Fields, *field)
// }
//
// func (l *schemaListener) VisitErrorNode(node antlr.ErrorNode) {
// 	// NOP
// }
//
// // EnterSchema is called when entering the schema production.
// func (l *schemaListener) EnterSchema(ctx *config.SchemaContext) {
// 	// NOP
// }
//
// // EnterTypeDecl is called when entering the typeDecl production.
// func (l *schemaListener) EnterTypeDecl(ctx *config.TypeDeclContext) {
// 	// NOP
// }
//
// // EnterTypeBody is called when entering the typeBody production.
// func (l *schemaListener) EnterTypeBody(ctx *config.TypeBodyContext) {
// 	// NOP
// }
//
// // EnterField is called when entering the field production.
// func (l *schemaListener) EnterField(ctx *config.FieldContext) {
// 	// NOP
// }
//
// // EnterFieldName is called when entering the fieldName production.
// func (l *schemaListener) EnterFieldName(ctx *config.FieldNameContext) {
// 	// NOP
// }
//
// // EnterFieldType is called when entering the fieldType production.
// func (l *schemaListener) EnterFieldType(ctx *config.FieldTypeContext) {
// 	// NOP
// }
//
// // ExitTypeName is called when exiting the typeName production.
// func (l *schemaListener) ExitTypeName(ctx *config.TypeNameContext) {
// 	// NOP
// }
//
// // ExitTypeBody is called when exiting the typeBody production.
// func (l *schemaListener) ExitTypeBody(ctx *config.TypeBodyContext) {
// 	// NOP
// }
//
// // ExitFieldName is called when exiting the fieldName production.
// func (l *schemaListener) ExitFieldName(ctx *config.FieldNameContext) {
// 	// NOP
// }
//
// // ExitFieldType is called when exiting the fieldType production.
// func (l *schemaListener) ExitFieldType(ctx *config.FieldTypeContext) {
// 	// NOP
// }
//
// func (l *schemaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	// NOP
// }
//
// func (l *schemaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
// 	// NOP
// }
//
// func (l *schemaListener) VisitTerminal(node antlr.TerminalNode) {
// 	// NOP
// }
