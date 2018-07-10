// Code generated from internal/parser/config/Config.g4 by ANTLR 4.7.1. DO NOT EDIT.

package config // Config
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 81, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 9,
	13, 3, 2, 3, 2, 7, 2, 29, 10, 2, 12, 2, 14, 2, 32, 11, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 42, 10, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 55, 10, 7, 12, 7, 14,
	7, 58, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 6, 10,
	68, 10, 10, 13, 10, 14, 10, 69, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 2, 3, 3, 2, 8, 16, 2, 73, 2, 30, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2,
	6, 38, 3, 2, 2, 2, 8, 45, 3, 2, 2, 2, 10, 48, 3, 2, 2, 2, 12, 52, 3, 2,
	2, 2, 14, 59, 3, 2, 2, 2, 16, 63, 3, 2, 2, 2, 18, 65, 3, 2, 2, 2, 20, 73,
	3, 2, 2, 2, 22, 76, 3, 2, 2, 2, 24, 78, 3, 2, 2, 2, 26, 29, 5, 4, 3, 2,
	27, 29, 5, 8, 5, 2, 28, 26, 3, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3,
	2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32,
	30, 3, 2, 2, 2, 33, 34, 7, 2, 2, 3, 34, 3, 3, 2, 2, 2, 35, 36, 7, 3, 2,
	2, 36, 37, 5, 6, 4, 2, 37, 5, 3, 2, 2, 2, 38, 39, 7, 4, 2, 2, 39, 41, 3,
	2, 2, 2, 40, 42, 11, 2, 2, 2, 41, 40, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42,
	43, 3, 2, 2, 2, 43, 44, 7, 5, 2, 2, 44, 7, 3, 2, 2, 2, 45, 46, 7, 6, 2,
	2, 46, 47, 5, 10, 6, 2, 47, 9, 3, 2, 2, 2, 48, 49, 7, 4, 2, 2, 49, 50,
	5, 12, 7, 2, 50, 51, 7, 5, 2, 2, 51, 11, 3, 2, 2, 2, 52, 56, 5, 14, 8,
	2, 53, 55, 5, 14, 8, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54,
	3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 13, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2,
	59, 60, 7, 7, 2, 2, 60, 61, 5, 16, 9, 2, 61, 62, 5, 18, 10, 2, 62, 15,
	3, 2, 2, 2, 63, 64, 7, 16, 2, 2, 64, 17, 3, 2, 2, 2, 65, 67, 7, 4, 2, 2,
	66, 68, 5, 20, 11, 2, 67, 66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 67, 3,
	2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 72, 7, 5, 2, 2, 72,
	19, 3, 2, 2, 2, 73, 74, 5, 22, 12, 2, 74, 75, 5, 24, 13, 2, 75, 21, 3,
	2, 2, 2, 76, 77, 7, 17, 2, 2, 77, 23, 3, 2, 2, 2, 78, 79, 9, 2, 2, 2, 79,
	25, 3, 2, 2, 2, 7, 28, 30, 41, 56, 69,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'options'", "'{'", "'}'", "'schema'", "'type'", "'ID'", "'int'", "'float'",
	"'string'", "'bool'", "'datetime'", "'geopoint'", "", "", "", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "TypeIdentArray", "TypeIdent",
	"FieldIdent", "LBRACK", "RBRACK", "WS", "NEWLINE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"document", "optionsBlock", "optionsBody", "schemaBlock", "schemaBody",
	"schemaTypes", "typeDecl", "typeName", "typeBody", "field", "fieldName",
	"fieldType",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ConfigParser struct {
	*antlr.BaseParser
}

func NewConfigParser(input antlr.TokenStream) *ConfigParser {
	this := new(ConfigParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Config.g4"

	return this
}

// ConfigParser tokens.
const (
	ConfigParserEOF            = antlr.TokenEOF
	ConfigParserT__0           = 1
	ConfigParserT__1           = 2
	ConfigParserT__2           = 3
	ConfigParserT__3           = 4
	ConfigParserT__4           = 5
	ConfigParserT__5           = 6
	ConfigParserT__6           = 7
	ConfigParserT__7           = 8
	ConfigParserT__8           = 9
	ConfigParserT__9           = 10
	ConfigParserT__10          = 11
	ConfigParserT__11          = 12
	ConfigParserTypeIdentArray = 13
	ConfigParserTypeIdent      = 14
	ConfigParserFieldIdent     = 15
	ConfigParserLBRACK         = 16
	ConfigParserRBRACK         = 17
	ConfigParserWS             = 18
	ConfigParserNEWLINE        = 19
	ConfigParserCOMMENT        = 20
	ConfigParserLINE_COMMENT   = 21
)

// ConfigParser rules.
const (
	ConfigParserRULE_document     = 0
	ConfigParserRULE_optionsBlock = 1
	ConfigParserRULE_optionsBody  = 2
	ConfigParserRULE_schemaBlock  = 3
	ConfigParserRULE_schemaBody   = 4
	ConfigParserRULE_schemaTypes  = 5
	ConfigParserRULE_typeDecl     = 6
	ConfigParserRULE_typeName     = 7
	ConfigParserRULE_typeBody     = 8
	ConfigParserRULE_field        = 9
	ConfigParserRULE_fieldName    = 10
	ConfigParserRULE_fieldType    = 11
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfigParserEOF, 0)
}

func (s *DocumentContext) AllOptionsBlock() []IOptionsBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOptionsBlockContext)(nil)).Elem())
	var tst = make([]IOptionsBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOptionsBlockContext)
		}
	}

	return tst
}

func (s *DocumentContext) OptionsBlock(i int) IOptionsBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionsBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOptionsBlockContext)
}

func (s *DocumentContext) AllSchemaBlock() []ISchemaBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISchemaBlockContext)(nil)).Elem())
	var tst = make([]ISchemaBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISchemaBlockContext)
		}
	}

	return tst
}

func (s *DocumentContext) SchemaBlock(i int) ISchemaBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISchemaBlockContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *ConfigParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfigParserRULE_document)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConfigParserT__0 || _la == ConfigParserT__3 {
		p.SetState(26)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ConfigParserT__0:
			{
				p.SetState(24)
				p.OptionsBlock()
			}

		case ConfigParserT__3:
			{
				p.SetState(25)
				p.SchemaBlock()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		p.Match(ConfigParserEOF)
	}

	return localctx
}

// IOptionsBlockContext is an interface to support dynamic dispatch.
type IOptionsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionsBlockContext differentiates from other interfaces.
	IsOptionsBlockContext()
}

type OptionsBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionsBlockContext() *OptionsBlockContext {
	var p = new(OptionsBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_optionsBlock
	return p
}

func (*OptionsBlockContext) IsOptionsBlockContext() {}

func NewOptionsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionsBlockContext {
	var p = new(OptionsBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_optionsBlock

	return p
}

func (s *OptionsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionsBlockContext) OptionsBody() IOptionsBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionsBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionsBodyContext)
}

func (s *OptionsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionsBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterOptionsBlock(s)
	}
}

func (s *OptionsBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitOptionsBlock(s)
	}
}

func (p *ConfigParser) OptionsBlock() (localctx IOptionsBlockContext) {
	localctx = NewOptionsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfigParserRULE_optionsBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(ConfigParserT__0)
	}
	{
		p.SetState(34)
		p.OptionsBody()
	}

	return localctx
}

// IOptionsBodyContext is an interface to support dynamic dispatch.
type IOptionsBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionsBodyContext differentiates from other interfaces.
	IsOptionsBodyContext()
}

type OptionsBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionsBodyContext() *OptionsBodyContext {
	var p = new(OptionsBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_optionsBody
	return p
}

func (*OptionsBodyContext) IsOptionsBodyContext() {}

func NewOptionsBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionsBodyContext {
	var p = new(OptionsBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_optionsBody

	return p
}

func (s *OptionsBodyContext) GetParser() antlr.Parser { return s.parser }
func (s *OptionsBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionsBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionsBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterOptionsBody(s)
	}
}

func (s *OptionsBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitOptionsBody(s)
	}
}

func (p *ConfigParser) OptionsBody() (localctx IOptionsBodyContext) {
	localctx = NewOptionsBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConfigParserRULE_optionsBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(ConfigParserT__1)
	}

	p.SetState(39)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		p.SetState(38)
		p.MatchWildcard()

	}
	{
		p.SetState(41)
		p.Match(ConfigParserT__2)
	}

	return localctx
}

// ISchemaBlockContext is an interface to support dynamic dispatch.
type ISchemaBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaBlockContext differentiates from other interfaces.
	IsSchemaBlockContext()
}

type SchemaBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaBlockContext() *SchemaBlockContext {
	var p = new(SchemaBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_schemaBlock
	return p
}

func (*SchemaBlockContext) IsSchemaBlockContext() {}

func NewSchemaBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaBlockContext {
	var p = new(SchemaBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_schemaBlock

	return p
}

func (s *SchemaBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaBlockContext) SchemaBody() ISchemaBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaBodyContext)
}

func (s *SchemaBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterSchemaBlock(s)
	}
}

func (s *SchemaBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitSchemaBlock(s)
	}
}

func (p *ConfigParser) SchemaBlock() (localctx ISchemaBlockContext) {
	localctx = NewSchemaBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConfigParserRULE_schemaBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(ConfigParserT__3)
	}
	{
		p.SetState(44)
		p.SchemaBody()
	}

	return localctx
}

// ISchemaBodyContext is an interface to support dynamic dispatch.
type ISchemaBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaBodyContext differentiates from other interfaces.
	IsSchemaBodyContext()
}

type SchemaBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaBodyContext() *SchemaBodyContext {
	var p = new(SchemaBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_schemaBody
	return p
}

func (*SchemaBodyContext) IsSchemaBodyContext() {}

func NewSchemaBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaBodyContext {
	var p = new(SchemaBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_schemaBody

	return p
}

func (s *SchemaBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaBodyContext) SchemaTypes() ISchemaTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaTypesContext)
}

func (s *SchemaBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterSchemaBody(s)
	}
}

func (s *SchemaBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitSchemaBody(s)
	}
}

func (p *ConfigParser) SchemaBody() (localctx ISchemaBodyContext) {
	localctx = NewSchemaBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConfigParserRULE_schemaBody)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(ConfigParserT__1)
	}
	{
		p.SetState(47)
		p.SchemaTypes()
	}
	{
		p.SetState(48)
		p.Match(ConfigParserT__2)
	}

	return localctx
}

// ISchemaTypesContext is an interface to support dynamic dispatch.
type ISchemaTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaTypesContext differentiates from other interfaces.
	IsSchemaTypesContext()
}

type SchemaTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaTypesContext() *SchemaTypesContext {
	var p = new(SchemaTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_schemaTypes
	return p
}

func (*SchemaTypesContext) IsSchemaTypesContext() {}

func NewSchemaTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaTypesContext {
	var p = new(SchemaTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_schemaTypes

	return p
}

func (s *SchemaTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaTypesContext) AllTypeDecl() []ITypeDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem())
	var tst = make([]ITypeDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDeclContext)
		}
	}

	return tst
}

func (s *SchemaTypesContext) TypeDecl(i int) ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *SchemaTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterSchemaTypes(s)
	}
}

func (s *SchemaTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitSchemaTypes(s)
	}
}

func (p *ConfigParser) SchemaTypes() (localctx ISchemaTypesContext) {
	localctx = NewSchemaTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConfigParserRULE_schemaTypes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.TypeDecl()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConfigParserT__4 {
		{
			p.SetState(51)
			p.TypeDecl()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TypeName() ITypeNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeNameContext)
}

func (s *TypeDeclContext) TypeBody() ITypeBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeBodyContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *ConfigParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConfigParserRULE_typeDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(ConfigParserT__4)
	}
	{
		p.SetState(58)
		p.TypeName()
	}
	{
		p.SetState(59)
		p.TypeBody()
	}

	return localctx
}

// ITypeNameContext is an interface to support dynamic dispatch.
type ITypeNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeNameContext differentiates from other interfaces.
	IsTypeNameContext()
}

type TypeNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeNameContext() *TypeNameContext {
	var p = new(TypeNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_typeName
	return p
}

func (*TypeNameContext) IsTypeNameContext() {}

func NewTypeNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeNameContext {
	var p = new(TypeNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_typeName

	return p
}

func (s *TypeNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeNameContext) TypeIdent() antlr.TerminalNode {
	return s.GetToken(ConfigParserTypeIdent, 0)
}

func (s *TypeNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterTypeName(s)
	}
}

func (s *TypeNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitTypeName(s)
	}
}

func (p *ConfigParser) TypeName() (localctx ITypeNameContext) {
	localctx = NewTypeNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ConfigParserRULE_typeName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(ConfigParserTypeIdent)
	}

	return localctx
}

// ITypeBodyContext is an interface to support dynamic dispatch.
type ITypeBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeBodyContext differentiates from other interfaces.
	IsTypeBodyContext()
}

type TypeBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeBodyContext() *TypeBodyContext {
	var p = new(TypeBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_typeBody
	return p
}

func (*TypeBodyContext) IsTypeBodyContext() {}

func NewTypeBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeBodyContext {
	var p = new(TypeBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_typeBody

	return p
}

func (s *TypeBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeBodyContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *TypeBodyContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *TypeBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterTypeBody(s)
	}
}

func (s *TypeBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitTypeBody(s)
	}
}

func (p *ConfigParser) TypeBody() (localctx ITypeBodyContext) {
	localctx = NewTypeBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ConfigParserRULE_typeBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(ConfigParserT__1)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConfigParserFieldIdent {
		{
			p.SetState(64)
			p.Field()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
		p.Match(ConfigParserT__2)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldName() IFieldNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ConfigParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ConfigParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.FieldName()
	}
	{
		p.SetState(72)
		p.FieldType()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) FieldIdent() antlr.TerminalNode {
	return s.GetToken(ConfigParserFieldIdent, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *ConfigParser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ConfigParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(ConfigParserFieldIdent)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfigParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfigParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) TypeIdent() antlr.TerminalNode {
	return s.GetToken(ConfigParserTypeIdent, 0)
}

func (s *FieldTypeContext) TypeIdentArray() antlr.TerminalNode {
	return s.GetToken(ConfigParserTypeIdentArray, 0)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfigListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *ConfigParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ConfigParserRULE_fieldType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ConfigParserT__5)|(1<<ConfigParserT__6)|(1<<ConfigParserT__7)|(1<<ConfigParserT__8)|(1<<ConfigParserT__9)|(1<<ConfigParserT__10)|(1<<ConfigParserT__11)|(1<<ConfigParserTypeIdentArray)|(1<<ConfigParserTypeIdent))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
