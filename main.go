package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/TheGrizzlyDev/egraphql/parser"
	"os"
	"fmt"
)

type TreeShapeListener struct {
	*parser.BaseEGraphQLListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewEGraphQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewEGraphQLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.TypeDefinition()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}