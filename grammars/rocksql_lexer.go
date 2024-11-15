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
		"", "'PUT'", "'GET'", "'DEL'", "'BATCH'", "'{'", "'}'", "'SNAPSHOT'",
		"'CREATE'", "'DELETE'", "'ADMIN'", "'COMPACT'", "'STATS'", "'CONFIG'",
		"'SET'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "STRING",
		"NUMBER", "IDENTIFIER", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "STRING", "NUMBER", "IDENTIFIER",
		"WS", "COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 19, 164, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 120, 8, 14, 10,
		14, 12, 14, 123, 9, 14, 1, 14, 1, 14, 1, 15, 4, 15, 128, 8, 15, 11, 15,
		12, 15, 129, 1, 15, 1, 15, 4, 15, 134, 8, 15, 11, 15, 12, 15, 135, 3, 15,
		138, 8, 15, 1, 16, 1, 16, 5, 16, 142, 8, 16, 10, 16, 12, 16, 145, 9, 16,
		1, 17, 4, 17, 148, 8, 17, 11, 17, 12, 17, 149, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 5, 18, 158, 8, 18, 10, 18, 12, 18, 161, 9, 18, 1, 18,
		1, 18, 0, 0, 19, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 1, 0, 6, 2, 0, 34, 34, 92, 92, 1, 0, 48, 57, 3, 0, 65, 90,
		95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13,
		13, 32, 32, 2, 0, 10, 10, 13, 13, 171, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 1, 39, 1, 0, 0, 0, 3, 43, 1, 0, 0, 0, 5,
		47, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0, 9, 57, 1, 0, 0, 0, 11, 59, 1, 0, 0,
		0, 13, 61, 1, 0, 0, 0, 15, 70, 1, 0, 0, 0, 17, 77, 1, 0, 0, 0, 19, 84,
		1, 0, 0, 0, 21, 90, 1, 0, 0, 0, 23, 98, 1, 0, 0, 0, 25, 104, 1, 0, 0, 0,
		27, 111, 1, 0, 0, 0, 29, 115, 1, 0, 0, 0, 31, 127, 1, 0, 0, 0, 33, 139,
		1, 0, 0, 0, 35, 147, 1, 0, 0, 0, 37, 153, 1, 0, 0, 0, 39, 40, 5, 80, 0,
		0, 40, 41, 5, 85, 0, 0, 41, 42, 5, 84, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44,
		5, 71, 0, 0, 44, 45, 5, 69, 0, 0, 45, 46, 5, 84, 0, 0, 46, 4, 1, 0, 0,
		0, 47, 48, 5, 68, 0, 0, 48, 49, 5, 69, 0, 0, 49, 50, 5, 76, 0, 0, 50, 6,
		1, 0, 0, 0, 51, 52, 5, 66, 0, 0, 52, 53, 5, 65, 0, 0, 53, 54, 5, 84, 0,
		0, 54, 55, 5, 67, 0, 0, 55, 56, 5, 72, 0, 0, 56, 8, 1, 0, 0, 0, 57, 58,
		5, 123, 0, 0, 58, 10, 1, 0, 0, 0, 59, 60, 5, 125, 0, 0, 60, 12, 1, 0, 0,
		0, 61, 62, 5, 83, 0, 0, 62, 63, 5, 78, 0, 0, 63, 64, 5, 65, 0, 0, 64, 65,
		5, 80, 0, 0, 65, 66, 5, 83, 0, 0, 66, 67, 5, 72, 0, 0, 67, 68, 5, 79, 0,
		0, 68, 69, 5, 84, 0, 0, 69, 14, 1, 0, 0, 0, 70, 71, 5, 67, 0, 0, 71, 72,
		5, 82, 0, 0, 72, 73, 5, 69, 0, 0, 73, 74, 5, 65, 0, 0, 74, 75, 5, 84, 0,
		0, 75, 76, 5, 69, 0, 0, 76, 16, 1, 0, 0, 0, 77, 78, 5, 68, 0, 0, 78, 79,
		5, 69, 0, 0, 79, 80, 5, 76, 0, 0, 80, 81, 5, 69, 0, 0, 81, 82, 5, 84, 0,
		0, 82, 83, 5, 69, 0, 0, 83, 18, 1, 0, 0, 0, 84, 85, 5, 65, 0, 0, 85, 86,
		5, 68, 0, 0, 86, 87, 5, 77, 0, 0, 87, 88, 5, 73, 0, 0, 88, 89, 5, 78, 0,
		0, 89, 20, 1, 0, 0, 0, 90, 91, 5, 67, 0, 0, 91, 92, 5, 79, 0, 0, 92, 93,
		5, 77, 0, 0, 93, 94, 5, 80, 0, 0, 94, 95, 5, 65, 0, 0, 95, 96, 5, 67, 0,
		0, 96, 97, 5, 84, 0, 0, 97, 22, 1, 0, 0, 0, 98, 99, 5, 83, 0, 0, 99, 100,
		5, 84, 0, 0, 100, 101, 5, 65, 0, 0, 101, 102, 5, 84, 0, 0, 102, 103, 5,
		83, 0, 0, 103, 24, 1, 0, 0, 0, 104, 105, 5, 67, 0, 0, 105, 106, 5, 79,
		0, 0, 106, 107, 5, 78, 0, 0, 107, 108, 5, 70, 0, 0, 108, 109, 5, 73, 0,
		0, 109, 110, 5, 71, 0, 0, 110, 26, 1, 0, 0, 0, 111, 112, 5, 83, 0, 0, 112,
		113, 5, 69, 0, 0, 113, 114, 5, 84, 0, 0, 114, 28, 1, 0, 0, 0, 115, 121,
		5, 34, 0, 0, 116, 120, 8, 0, 0, 0, 117, 118, 5, 92, 0, 0, 118, 120, 9,
		0, 0, 0, 119, 116, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 123, 1, 0, 0,
		0, 121, 119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123,
		121, 1, 0, 0, 0, 124, 125, 5, 34, 0, 0, 125, 30, 1, 0, 0, 0, 126, 128,
		7, 1, 0, 0, 127, 126, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 127, 1, 0,
		0, 0, 129, 130, 1, 0, 0, 0, 130, 137, 1, 0, 0, 0, 131, 133, 5, 46, 0, 0,
		132, 134, 7, 1, 0, 0, 133, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135,
		133, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136, 138, 1, 0, 0, 0, 137, 131,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 32, 1, 0, 0, 0, 139, 143, 7, 2,
		0, 0, 140, 142, 7, 3, 0, 0, 141, 140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0,
		143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 34, 1, 0, 0, 0, 145, 143,
		1, 0, 0, 0, 146, 148, 7, 4, 0, 0, 147, 146, 1, 0, 0, 0, 148, 149, 1, 0,
		0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0,
		151, 152, 6, 17, 0, 0, 152, 36, 1, 0, 0, 0, 153, 154, 5, 47, 0, 0, 154,
		155, 5, 47, 0, 0, 155, 159, 1, 0, 0, 0, 156, 158, 8, 5, 0, 0, 157, 156,
		1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0,
		0, 0, 160, 162, 1, 0, 0, 0, 161, 159, 1, 0, 0, 0, 162, 163, 6, 18, 0, 0,
		163, 38, 1, 0, 0, 0, 9, 0, 119, 121, 129, 135, 137, 143, 149, 159, 1, 6,
		0, 0,
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
	RocksQLLexerT__8       = 9
	RocksQLLexerT__9       = 10
	RocksQLLexerT__10      = 11
	RocksQLLexerT__11      = 12
	RocksQLLexerT__12      = 13
	RocksQLLexerT__13      = 14
	RocksQLLexerSTRING     = 15
	RocksQLLexerNUMBER     = 16
	RocksQLLexerIDENTIFIER = 17
	RocksQLLexerWS         = 18
	RocksQLLexerCOMMENT    = 19
)
