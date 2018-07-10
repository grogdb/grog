package schema

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/grogdb/grogdb/internal/parser/config"
)

type Schema struct {
}

func Parse(raw []byte) (*Schema, error) {
	input := antlr.NewInputStream(string(raw))
	lexer := config.NewConfigLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	p := config.NewConfigParser(tokenStream)

	// errorListener := NewParseErrorListener()
	// p.AddErrorListener(errorListener)
	// p.BuildParseTrees = true
	//
	// v := schemaVisitor{}
	// _ = v.Visit(p.Document())
	//
	// // listener := &schemaVisitor{}
	// // antlr.VisNewParseTreeWalker().BaseParseTreeVisitor{}ParseTreeWalkerDefault.Walk(listener, p.Schema())
	//
	// if errorListener.Err() != nil {
	// 	return nil, errorListener.Err()
	// }

	return nil, nil
}
