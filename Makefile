ROCKSDB_DIR :=  /opt/homebrew/opt/rocksdb
SNAPPY_DIR  :=  /opt/homebrew/opt/snappy
LZ4_DIR     :=  /opt/homebrew/opt/lz4
ZSTD_DIR    :=  /opt/homebrew/opt/zstd
LIBS        :=  -lrocksdb
GO         := go
CC         := gcc
CXX        := g++

CGO_FLAGS  := CGO_ENABLED=1 CC=$(CC) CXX=$(CXX)
INCLUDE    := -I$(ROCKSDB_DIR)/include
LIBS       := -L$(ROCKSDB_DIR)/lib -L$(SNAPPY_DIR)/lib -L$(LZ4_DIR)/lib -L$(ZSTD_DIR)/lib $(LIBS)
BINARY     := rocksql


# Default target
all: build

# Build target
build:
	$(CGO_FLAGS) \
	CGO_CFLAGS="$(INCLUDE)" \
	CGO_LDFLAGS="$(LIBS)" \
	$(GO) build -o $(BINARY)


# Clean target
clean:
	rm -f $(BINARY)

run:
	make build
	./$(BINARY)


test:
	$(CGO_FLAGS) \
	CGO_CFLAGS="$(INCLUDE)" \
	CGO_LDFLAGS="$(LIBS)" \
	$(GO) test ./...

grammar:
	antlr -Dlanguage=Go -package grammars grammars/RocksQL.g4

fmt:
	go fmt ./...

clean-grammar:
	rm grammars/RocksQL.tokens
	rm grammars/RocksQLLexer.tokens
	rm grammars/rocksql_lexer.go
	rm grammars/rocksql_parser.go
	rm grammars/RocksQL.interp
	rm grammars/RocksQLLexer.interp
	rm grammars/rocksql_base_listener.go
	rm grammars/rocksql_listener.go


.PHONY: all build clean grammar clean-grammar
