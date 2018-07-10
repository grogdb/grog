// Code generated from internal/parser/config/Config.g4 by ANTLR 4.7.1. DO NOT EDIT.

package config // Config
import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConfigListener is a complete listener for a parse tree produced by ConfigParser.
type ConfigListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterOptionsBlock is called when entering the optionsBlock production.
	EnterOptionsBlock(c *OptionsBlockContext)

	// EnterOptionsBody is called when entering the optionsBody production.
	EnterOptionsBody(c *OptionsBodyContext)

	// EnterSchemaBlock is called when entering the schemaBlock production.
	EnterSchemaBlock(c *SchemaBlockContext)

	// EnterSchemaBody is called when entering the schemaBody production.
	EnterSchemaBody(c *SchemaBodyContext)

	// EnterSchemaTypes is called when entering the schemaTypes production.
	EnterSchemaTypes(c *SchemaTypesContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterTypeBody is called when entering the typeBody production.
	EnterTypeBody(c *TypeBodyContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitOptionsBlock is called when exiting the optionsBlock production.
	ExitOptionsBlock(c *OptionsBlockContext)

	// ExitOptionsBody is called when exiting the optionsBody production.
	ExitOptionsBody(c *OptionsBodyContext)

	// ExitSchemaBlock is called when exiting the schemaBlock production.
	ExitSchemaBlock(c *SchemaBlockContext)

	// ExitSchemaBody is called when exiting the schemaBody production.
	ExitSchemaBody(c *SchemaBodyContext)

	// ExitSchemaTypes is called when exiting the schemaTypes production.
	ExitSchemaTypes(c *SchemaTypesContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitTypeBody is called when exiting the typeBody production.
	ExitTypeBody(c *TypeBodyContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)
}
