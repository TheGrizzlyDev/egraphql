generate:
	antlr4 -Dlanguage=Go EGraphQL.g4

setup:
	./scripts/setup.sh