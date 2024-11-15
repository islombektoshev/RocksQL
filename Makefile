grammar:
	antlr -Dlanguage=Go -package grammars grammars/RocksQL.g4


clean-grammar:
	rm grammars/RocksQL.tokens
	rm grammars/RocksQLLexer.tokens
	rm grammars/rocksql_lexer.go
	rm grammars/rocksql_parser.go
	rm grammars/RocksQL.interp
	rm grammars/RocksQLLexer.interp
	rm grammars/rocksql_base_listener.go
	rm grammars/rocksql_listener.go
