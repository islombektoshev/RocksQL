// Code generated from grammars/RocksQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammars

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type RocksQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var RocksQLLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rocksqllexerLexerInit() {
	staticData := &RocksQLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'PUT'", "'GET'", "'DEL'", "'ADMIN'", "'COMPACT'", "'STATS'", "'CONFIG'",
		"'SET'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "IDENTIFIER",
		"WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "STRING",
		"NUMBER", "IDENTIFIER", "WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 119, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 75, 8, 8, 10, 8, 12, 8, 78, 9, 8,
		1, 8, 1, 8, 1, 9, 4, 9, 83, 8, 9, 11, 9, 12, 9, 84, 1, 9, 1, 9, 4, 9, 89,
		8, 9, 11, 9, 12, 9, 90, 3, 9, 93, 8, 9, 1, 10, 1, 10, 5, 10, 97, 8, 10,
		10, 10, 12, 10, 100, 9, 10, 1, 11, 4, 11, 103, 8, 11, 11, 11, 12, 11, 104,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 113, 8, 12, 10, 12, 12,
		12, 116, 9, 12, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 1, 0, 6, 2, 0,
		34, 34, 92, 92, 1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10,
		13, 13, 126, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27, 1, 0, 0, 0, 3, 31, 1, 0, 0,
		0, 5, 35, 1, 0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 45, 1, 0, 0, 0, 11, 53, 1,
		0, 0, 0, 13, 59, 1, 0, 0, 0, 15, 66, 1, 0, 0, 0, 17, 70, 1, 0, 0, 0, 19,
		82, 1, 0, 0, 0, 21, 94, 1, 0, 0, 0, 23, 102, 1, 0, 0, 0, 25, 108, 1, 0,
		0, 0, 27, 28, 5, 80, 0, 0, 28, 29, 5, 85, 0, 0, 29, 30, 5, 84, 0, 0, 30,
		2, 1, 0, 0, 0, 31, 32, 5, 71, 0, 0, 32, 33, 5, 69, 0, 0, 33, 34, 5, 84,
		0, 0, 34, 4, 1, 0, 0, 0, 35, 36, 5, 68, 0, 0, 36, 37, 5, 69, 0, 0, 37,
		38, 5, 76, 0, 0, 38, 6, 1, 0, 0, 0, 39, 40, 5, 65, 0, 0, 40, 41, 5, 68,
		0, 0, 41, 42, 5, 77, 0, 0, 42, 43, 5, 73, 0, 0, 43, 44, 5, 78, 0, 0, 44,
		8, 1, 0, 0, 0, 45, 46, 5, 67, 0, 0, 46, 47, 5, 79, 0, 0, 47, 48, 5, 77,
		0, 0, 48, 49, 5, 80, 0, 0, 49, 50, 5, 65, 0, 0, 50, 51, 5, 67, 0, 0, 51,
		52, 5, 84, 0, 0, 52, 10, 1, 0, 0, 0, 53, 54, 5, 83, 0, 0, 54, 55, 5, 84,
		0, 0, 55, 56, 5, 65, 0, 0, 56, 57, 5, 84, 0, 0, 57, 58, 5, 83, 0, 0, 58,
		12, 1, 0, 0, 0, 59, 60, 5, 67, 0, 0, 60, 61, 5, 79, 0, 0, 61, 62, 5, 78,
		0, 0, 62, 63, 5, 70, 0, 0, 63, 64, 5, 73, 0, 0, 64, 65, 5, 71, 0, 0, 65,
		14, 1, 0, 0, 0, 66, 67, 5, 83, 0, 0, 67, 68, 5, 69, 0, 0, 68, 69, 5, 84,
		0, 0, 69, 16, 1, 0, 0, 0, 70, 76, 5, 34, 0, 0, 71, 75, 8, 0, 0, 0, 72,
		73, 5, 92, 0, 0, 73, 75, 9, 0, 0, 0, 74, 71, 1, 0, 0, 0, 74, 72, 1, 0,
		0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 79,
		1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 5, 34, 0, 0, 80, 18, 1, 0, 0, 0,
		81, 83, 7, 1, 0, 0, 82, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 92, 1, 0, 0, 0, 86, 88, 5, 46, 0, 0, 87,
		89, 7, 1, 0, 0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0,
		0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 86, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 20, 1, 0, 0, 0, 94, 98, 7, 2, 0, 0, 95, 97, 7, 3, 0, 0,
		96, 95, 1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 22, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 101, 103, 7, 4, 0, 0,
		102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 104,
		105, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 6, 11, 0, 0, 107, 24,
		1, 0, 0, 0, 108, 109, 5, 47, 0, 0, 109, 110, 5, 47, 0, 0, 110, 114, 1,
		0, 0, 0, 111, 113, 8, 5, 0, 0, 112, 111, 1, 0, 0, 0, 113, 116, 1, 0, 0,
		0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117, 1, 0, 0, 0, 116,
		114, 1, 0, 0, 0, 117, 118, 6, 12, 0, 0, 118, 26, 1, 0, 0, 0, 9, 0, 74,
		76, 84, 90, 92, 98, 104, 114, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RocksQLLexerInit initializes any static state used to implement RocksQLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewRocksQLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func RocksQLLexerInit() {
	staticData := &RocksQLLexerLexerStaticData
	staticData.once.Do(rocksqllexerLexerInit)
}

// NewRocksQLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewRocksQLLexer(input antlr.CharStream) *RocksQLLexer {
	RocksQLLexerInit()
	l := new(RocksQLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &RocksQLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "RocksQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// RocksQLLexer tokens.
const (
	RocksQLLexerT__0       = 1
	RocksQLLexerT__1       = 2
	RocksQLLexerT__2       = 3
	RocksQLLexerT__3       = 4
	RocksQLLexerT__4       = 5
	RocksQLLexerT__5       = 6
	RocksQLLexerT__6       = 7
	RocksQLLexerT__7       = 8
	RocksQLLexerSTRING     = 9
	RocksQLLexerNUMBER     = 10
	RocksQLLexerIDENTIFIER = 11
	RocksQLLexerWS         = 12
	RocksQLLexerCOMMENT    = 13
)
