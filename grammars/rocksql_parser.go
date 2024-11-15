// Code generated from grammars/RocksQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammars // RocksQL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RocksQLParser struct {
	*antlr.BaseParser
}

var RocksQLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rocksqlParserInit() {
	staticData := &RocksQLParserStaticData
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
		"command", "operation", "putOperation", "getOperation", "deleteOperation",
		"batchOperation", "batchEntry", "snapshotOperation", "adminCommand",
		"configCommand", "key", "value", "identifier", "option",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 86, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 4, 5, 53, 8, 5, 11, 5,
		12, 5, 54, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 61, 8, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 71, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 0, 0, 14,
		0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 3, 1, 0, 8, 9, 2,
		0, 15, 15, 17, 17, 1, 0, 15, 16, 80, 0, 28, 1, 0, 0, 0, 2, 37, 1, 0, 0,
		0, 4, 39, 1, 0, 0, 0, 6, 43, 1, 0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 49, 1,
		0, 0, 0, 12, 60, 1, 0, 0, 0, 14, 62, 1, 0, 0, 0, 16, 66, 1, 0, 0, 0, 18,
		72, 1, 0, 0, 0, 20, 77, 1, 0, 0, 0, 22, 79, 1, 0, 0, 0, 24, 81, 1, 0, 0,
		0, 26, 83, 1, 0, 0, 0, 28, 29, 3, 2, 1, 0, 29, 30, 5, 0, 0, 1, 30, 1, 1,
		0, 0, 0, 31, 38, 3, 4, 2, 0, 32, 38, 3, 6, 3, 0, 33, 38, 3, 8, 4, 0, 34,
		38, 3, 10, 5, 0, 35, 38, 3, 14, 7, 0, 36, 38, 3, 16, 8, 0, 37, 31, 1, 0,
		0, 0, 37, 32, 1, 0, 0, 0, 37, 33, 1, 0, 0, 0, 37, 34, 1, 0, 0, 0, 37, 35,
		1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 3, 1, 0, 0, 0, 39, 40, 5, 1, 0, 0,
		40, 41, 3, 20, 10, 0, 41, 42, 3, 22, 11, 0, 42, 5, 1, 0, 0, 0, 43, 44,
		5, 2, 0, 0, 44, 45, 3, 20, 10, 0, 45, 7, 1, 0, 0, 0, 46, 47, 5, 3, 0, 0,
		47, 48, 3, 20, 10, 0, 48, 9, 1, 0, 0, 0, 49, 50, 5, 4, 0, 0, 50, 52, 5,
		5, 0, 0, 51, 53, 3, 12, 6, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54,
		52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 5, 6, 0,
		0, 57, 11, 1, 0, 0, 0, 58, 61, 3, 4, 2, 0, 59, 61, 3, 8, 4, 0, 60, 58,
		1, 0, 0, 0, 60, 59, 1, 0, 0, 0, 61, 13, 1, 0, 0, 0, 62, 63, 5, 7, 0, 0,
		63, 64, 7, 0, 0, 0, 64, 65, 3, 24, 12, 0, 65, 15, 1, 0, 0, 0, 66, 70, 5,
		10, 0, 0, 67, 71, 5, 11, 0, 0, 68, 71, 5, 12, 0, 0, 69, 71, 3, 18, 9, 0,
		70, 67, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 17, 1,
		0, 0, 0, 72, 73, 5, 13, 0, 0, 73, 74, 5, 14, 0, 0, 74, 75, 3, 26, 13, 0,
		75, 76, 3, 22, 11, 0, 76, 19, 1, 0, 0, 0, 77, 78, 7, 1, 0, 0, 78, 21, 1,
		0, 0, 0, 79, 80, 7, 2, 0, 0, 80, 23, 1, 0, 0, 0, 81, 82, 5, 17, 0, 0, 82,
		25, 1, 0, 0, 0, 83, 84, 5, 17, 0, 0, 84, 27, 1, 0, 0, 0, 4, 37, 54, 60,
		70,
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

// RocksQLParserInit initializes any static state used to implement RocksQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRocksQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RocksQLParserInit() {
	staticData := &RocksQLParserStaticData
	staticData.once.Do(rocksqlParserInit)
}

// NewRocksQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRocksQLParser(input antlr.TokenStream) *RocksQLParser {
	RocksQLParserInit()
	this := new(RocksQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &RocksQLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "RocksQL.g4"

	return this
}

// RocksQLParser tokens.
const (
	RocksQLParserEOF        = antlr.TokenEOF
	RocksQLParserT__0       = 1
	RocksQLParserT__1       = 2
	RocksQLParserT__2       = 3
	RocksQLParserT__3       = 4
	RocksQLParserT__4       = 5
	RocksQLParserT__5       = 6
	RocksQLParserT__6       = 7
	RocksQLParserT__7       = 8
	RocksQLParserT__8       = 9
	RocksQLParserT__9       = 10
	RocksQLParserT__10      = 11
	RocksQLParserT__11      = 12
	RocksQLParserT__12      = 13
	RocksQLParserT__13      = 14
	RocksQLParserSTRING     = 15
	RocksQLParserNUMBER     = 16
	RocksQLParserIDENTIFIER = 17
	RocksQLParserWS         = 18
	RocksQLParserCOMMENT    = 19
)

// RocksQLParser rules.
const (
	RocksQLParserRULE_command           = 0
	RocksQLParserRULE_operation         = 1
	RocksQLParserRULE_putOperation      = 2
	RocksQLParserRULE_getOperation      = 3
	RocksQLParserRULE_deleteOperation   = 4
	RocksQLParserRULE_batchOperation    = 5
	RocksQLParserRULE_batchEntry        = 6
	RocksQLParserRULE_snapshotOperation = 7
	RocksQLParserRULE_adminCommand      = 8
	RocksQLParserRULE_configCommand     = 9
	RocksQLParserRULE_key               = 10
	RocksQLParserRULE_value             = 11
	RocksQLParserRULE_identifier        = 12
	RocksQLParserRULE_option            = 13
)

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operation() IOperationContext
	EOF() antlr.TerminalNode

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Operation() IOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperationContext)
}

func (s *CommandContext) EOF() antlr.TerminalNode {
	return s.GetToken(RocksQLParserEOF, 0)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *RocksQLParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RocksQLParserRULE_command)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Operation()
	}
	{
		p.SetState(29)
		p.Match(RocksQLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperationContext is an interface to support dynamic dispatch.
type IOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PutOperation() IPutOperationContext
	GetOperation() IGetOperationContext
	DeleteOperation() IDeleteOperationContext
	BatchOperation() IBatchOperationContext
	SnapshotOperation() ISnapshotOperationContext
	AdminCommand() IAdminCommandContext

	// IsOperationContext differentiates from other interfaces.
	IsOperationContext()
}

type OperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationContext() *OperationContext {
	var p = new(OperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_operation
	return p
}

func InitEmptyOperationContext(p *OperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_operation
}

func (*OperationContext) IsOperationContext() {}

func NewOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationContext {
	var p = new(OperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_operation

	return p
}

func (s *OperationContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationContext) PutOperation() IPutOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPutOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPutOperationContext)
}

func (s *OperationContext) GetOperation() IGetOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetOperationContext)
}

func (s *OperationContext) DeleteOperation() IDeleteOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteOperationContext)
}

func (s *OperationContext) BatchOperation() IBatchOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBatchOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBatchOperationContext)
}

func (s *OperationContext) SnapshotOperation() ISnapshotOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISnapshotOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISnapshotOperationContext)
}

func (s *OperationContext) AdminCommand() IAdminCommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdminCommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdminCommandContext)
}

func (s *OperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterOperation(s)
	}
}

func (s *OperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitOperation(s)
	}
}

func (p *RocksQLParser) Operation() (localctx IOperationContext) {
	localctx = NewOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RocksQLParserRULE_operation)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RocksQLParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.PutOperation()
		}

	case RocksQLParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.GetOperation()
		}

	case RocksQLParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.DeleteOperation()
		}

	case RocksQLParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(34)
			p.BatchOperation()
		}

	case RocksQLParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(35)
			p.SnapshotOperation()
		}

	case RocksQLParserT__9:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(36)
			p.AdminCommand()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPutOperationContext is an interface to support dynamic dispatch.
type IPutOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	Value() IValueContext

	// IsPutOperationContext differentiates from other interfaces.
	IsPutOperationContext()
}

type PutOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPutOperationContext() *PutOperationContext {
	var p = new(PutOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_putOperation
	return p
}

func InitEmptyPutOperationContext(p *PutOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_putOperation
}

func (*PutOperationContext) IsPutOperationContext() {}

func NewPutOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PutOperationContext {
	var p = new(PutOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_putOperation

	return p
}

func (s *PutOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *PutOperationContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *PutOperationContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PutOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PutOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterPutOperation(s)
	}
}

func (s *PutOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitPutOperation(s)
	}
}

func (p *RocksQLParser) PutOperation() (localctx IPutOperationContext) {
	localctx = NewPutOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RocksQLParserRULE_putOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(RocksQLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Key()
	}
	{
		p.SetState(41)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGetOperationContext is an interface to support dynamic dispatch.
type IGetOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext

	// IsGetOperationContext differentiates from other interfaces.
	IsGetOperationContext()
}

type GetOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetOperationContext() *GetOperationContext {
	var p = new(GetOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_getOperation
	return p
}

func InitEmptyGetOperationContext(p *GetOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_getOperation
}

func (*GetOperationContext) IsGetOperationContext() {}

func NewGetOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetOperationContext {
	var p = new(GetOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_getOperation

	return p
}

func (s *GetOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *GetOperationContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *GetOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterGetOperation(s)
	}
}

func (s *GetOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitGetOperation(s)
	}
}

func (p *RocksQLParser) GetOperation() (localctx IGetOperationContext) {
	localctx = NewGetOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RocksQLParserRULE_getOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(RocksQLParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Key()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeleteOperationContext is an interface to support dynamic dispatch.
type IDeleteOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext

	// IsDeleteOperationContext differentiates from other interfaces.
	IsDeleteOperationContext()
}

type DeleteOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteOperationContext() *DeleteOperationContext {
	var p = new(DeleteOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_deleteOperation
	return p
}

func InitEmptyDeleteOperationContext(p *DeleteOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_deleteOperation
}

func (*DeleteOperationContext) IsDeleteOperationContext() {}

func NewDeleteOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteOperationContext {
	var p = new(DeleteOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_deleteOperation

	return p
}

func (s *DeleteOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteOperationContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *DeleteOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterDeleteOperation(s)
	}
}

func (s *DeleteOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitDeleteOperation(s)
	}
}

func (p *RocksQLParser) DeleteOperation() (localctx IDeleteOperationContext) {
	localctx = NewDeleteOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RocksQLParserRULE_deleteOperation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(RocksQLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Key()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBatchOperationContext is an interface to support dynamic dispatch.
type IBatchOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBatchEntry() []IBatchEntryContext
	BatchEntry(i int) IBatchEntryContext

	// IsBatchOperationContext differentiates from other interfaces.
	IsBatchOperationContext()
}

type BatchOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatchOperationContext() *BatchOperationContext {
	var p = new(BatchOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_batchOperation
	return p
}

func InitEmptyBatchOperationContext(p *BatchOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_batchOperation
}

func (*BatchOperationContext) IsBatchOperationContext() {}

func NewBatchOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BatchOperationContext {
	var p = new(BatchOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_batchOperation

	return p
}

func (s *BatchOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *BatchOperationContext) AllBatchEntry() []IBatchEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBatchEntryContext); ok {
			len++
		}
	}

	tst := make([]IBatchEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBatchEntryContext); ok {
			tst[i] = t.(IBatchEntryContext)
			i++
		}
	}

	return tst
}

func (s *BatchOperationContext) BatchEntry(i int) IBatchEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBatchEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBatchEntryContext)
}

func (s *BatchOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BatchOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BatchOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterBatchOperation(s)
	}
}

func (s *BatchOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitBatchOperation(s)
	}
}

func (p *RocksQLParser) BatchOperation() (localctx IBatchOperationContext) {
	localctx = NewBatchOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RocksQLParserRULE_batchOperation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(RocksQLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(RocksQLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == RocksQLParserT__0 || _la == RocksQLParserT__2 {
		{
			p.SetState(51)
			p.BatchEntry()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(56)
		p.Match(RocksQLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBatchEntryContext is an interface to support dynamic dispatch.
type IBatchEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PutOperation() IPutOperationContext
	DeleteOperation() IDeleteOperationContext

	// IsBatchEntryContext differentiates from other interfaces.
	IsBatchEntryContext()
}

type BatchEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatchEntryContext() *BatchEntryContext {
	var p = new(BatchEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_batchEntry
	return p
}

func InitEmptyBatchEntryContext(p *BatchEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_batchEntry
}

func (*BatchEntryContext) IsBatchEntryContext() {}

func NewBatchEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BatchEntryContext {
	var p = new(BatchEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_batchEntry

	return p
}

func (s *BatchEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *BatchEntryContext) PutOperation() IPutOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPutOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPutOperationContext)
}

func (s *BatchEntryContext) DeleteOperation() IDeleteOperationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteOperationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteOperationContext)
}

func (s *BatchEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BatchEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BatchEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterBatchEntry(s)
	}
}

func (s *BatchEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitBatchEntry(s)
	}
}

func (p *RocksQLParser) BatchEntry() (localctx IBatchEntryContext) {
	localctx = NewBatchEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RocksQLParserRULE_batchEntry)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RocksQLParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.PutOperation()
		}

	case RocksQLParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.DeleteOperation()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISnapshotOperationContext is an interface to support dynamic dispatch.
type ISnapshotOperationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext

	// IsSnapshotOperationContext differentiates from other interfaces.
	IsSnapshotOperationContext()
}

type SnapshotOperationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySnapshotOperationContext() *SnapshotOperationContext {
	var p = new(SnapshotOperationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_snapshotOperation
	return p
}

func InitEmptySnapshotOperationContext(p *SnapshotOperationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_snapshotOperation
}

func (*SnapshotOperationContext) IsSnapshotOperationContext() {}

func NewSnapshotOperationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SnapshotOperationContext {
	var p = new(SnapshotOperationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_snapshotOperation

	return p
}

func (s *SnapshotOperationContext) GetParser() antlr.Parser { return s.parser }

func (s *SnapshotOperationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SnapshotOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SnapshotOperationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SnapshotOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterSnapshotOperation(s)
	}
}

func (s *SnapshotOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitSnapshotOperation(s)
	}
}

func (p *RocksQLParser) SnapshotOperation() (localctx ISnapshotOperationContext) {
	localctx = NewSnapshotOperationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RocksQLParserRULE_snapshotOperation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Match(RocksQLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RocksQLParserT__7 || _la == RocksQLParserT__8) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(64)
		p.Identifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdminCommandContext is an interface to support dynamic dispatch.
type IAdminCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConfigCommand() IConfigCommandContext

	// IsAdminCommandContext differentiates from other interfaces.
	IsAdminCommandContext()
}

type AdminCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdminCommandContext() *AdminCommandContext {
	var p = new(AdminCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_adminCommand
	return p
}

func InitEmptyAdminCommandContext(p *AdminCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_adminCommand
}

func (*AdminCommandContext) IsAdminCommandContext() {}

func NewAdminCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdminCommandContext {
	var p = new(AdminCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_adminCommand

	return p
}

func (s *AdminCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *AdminCommandContext) ConfigCommand() IConfigCommandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConfigCommandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConfigCommandContext)
}

func (s *AdminCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdminCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdminCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterAdminCommand(s)
	}
}

func (s *AdminCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitAdminCommand(s)
	}
}

func (p *RocksQLParser) AdminCommand() (localctx IAdminCommandContext) {
	localctx = NewAdminCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RocksQLParserRULE_adminCommand)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(RocksQLParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case RocksQLParserT__10:
		{
			p.SetState(67)
			p.Match(RocksQLParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RocksQLParserT__11:
		{
			p.SetState(68)
			p.Match(RocksQLParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case RocksQLParserT__12:
		{
			p.SetState(69)
			p.ConfigCommand()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConfigCommandContext is an interface to support dynamic dispatch.
type IConfigCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Option() IOptionContext
	Value() IValueContext

	// IsConfigCommandContext differentiates from other interfaces.
	IsConfigCommandContext()
}

type ConfigCommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigCommandContext() *ConfigCommandContext {
	var p = new(ConfigCommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_configCommand
	return p
}

func InitEmptyConfigCommandContext(p *ConfigCommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_configCommand
}

func (*ConfigCommandContext) IsConfigCommandContext() {}

func NewConfigCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigCommandContext {
	var p = new(ConfigCommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_configCommand

	return p
}

func (s *ConfigCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigCommandContext) Option() IOptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionContext)
}

func (s *ConfigCommandContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ConfigCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterConfigCommand(s)
	}
}

func (s *ConfigCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitConfigCommand(s)
	}
}

func (p *RocksQLParser) ConfigCommand() (localctx IConfigCommandContext) {
	localctx = NewConfigCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RocksQLParserRULE_configCommand)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(RocksQLParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(RocksQLParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Option()
	}
	{
		p.SetState(75)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) STRING() antlr.TerminalNode {
	return s.GetToken(RocksQLParserSTRING, 0)
}

func (s *KeyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RocksQLParserIDENTIFIER, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *RocksQLParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RocksQLParserRULE_key)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RocksQLParserSTRING || _la == RocksQLParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(RocksQLParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(RocksQLParserNUMBER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *RocksQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RocksQLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !(_la == RocksQLParserSTRING || _la == RocksQLParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RocksQLParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *RocksQLParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RocksQLParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(RocksQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionContext is an interface to support dynamic dispatch.
type IOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsOptionContext differentiates from other interfaces.
	IsOptionContext()
}

type OptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionContext() *OptionContext {
	var p = new(OptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_option
	return p
}

func InitEmptyOptionContext(p *OptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = RocksQLParserRULE_option
}

func (*OptionContext) IsOptionContext() {}

func NewOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionContext {
	var p = new(OptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = RocksQLParserRULE_option

	return p
}

func (s *OptionContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(RocksQLParserIDENTIFIER, 0)
}

func (s *OptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.EnterOption(s)
	}
}

func (s *OptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RocksQLListener); ok {
		listenerT.ExitOption(s)
	}
}

func (p *RocksQLParser) Option() (localctx IOptionContext) {
	localctx = NewOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RocksQLParserRULE_option)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(RocksQLParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}