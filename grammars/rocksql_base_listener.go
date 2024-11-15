// Code generated from grammars/RocksQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammars // RocksQL
import "github.com/antlr4-go/antlr/v4"

// BaseRocksQLListener is a complete listener for a parse tree produced by RocksQLParser.
type BaseRocksQLListener struct{}

var _ RocksQLListener = &BaseRocksQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRocksQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRocksQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRocksQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRocksQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseRocksQLListener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseRocksQLListener) ExitCommand(ctx *CommandContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseRocksQLListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseRocksQLListener) ExitOperation(ctx *OperationContext) {}

// EnterPutOperation is called when production putOperation is entered.
func (s *BaseRocksQLListener) EnterPutOperation(ctx *PutOperationContext) {}

// ExitPutOperation is called when production putOperation is exited.
func (s *BaseRocksQLListener) ExitPutOperation(ctx *PutOperationContext) {}

// EnterGetOperation is called when production getOperation is entered.
func (s *BaseRocksQLListener) EnterGetOperation(ctx *GetOperationContext) {}

// ExitGetOperation is called when production getOperation is exited.
func (s *BaseRocksQLListener) ExitGetOperation(ctx *GetOperationContext) {}

// EnterDeleteOperation is called when production deleteOperation is entered.
func (s *BaseRocksQLListener) EnterDeleteOperation(ctx *DeleteOperationContext) {}

// ExitDeleteOperation is called when production deleteOperation is exited.
func (s *BaseRocksQLListener) ExitDeleteOperation(ctx *DeleteOperationContext) {}

// EnterBatchOperation is called when production batchOperation is entered.
func (s *BaseRocksQLListener) EnterBatchOperation(ctx *BatchOperationContext) {}

// ExitBatchOperation is called when production batchOperation is exited.
func (s *BaseRocksQLListener) ExitBatchOperation(ctx *BatchOperationContext) {}

// EnterBatchEntry is called when production batchEntry is entered.
func (s *BaseRocksQLListener) EnterBatchEntry(ctx *BatchEntryContext) {}

// ExitBatchEntry is called when production batchEntry is exited.
func (s *BaseRocksQLListener) ExitBatchEntry(ctx *BatchEntryContext) {}

// EnterSnapshotOperation is called when production snapshotOperation is entered.
func (s *BaseRocksQLListener) EnterSnapshotOperation(ctx *SnapshotOperationContext) {}

// ExitSnapshotOperation is called when production snapshotOperation is exited.
func (s *BaseRocksQLListener) ExitSnapshotOperation(ctx *SnapshotOperationContext) {}

// EnterAdminCommand is called when production adminCommand is entered.
func (s *BaseRocksQLListener) EnterAdminCommand(ctx *AdminCommandContext) {}

// ExitAdminCommand is called when production adminCommand is exited.
func (s *BaseRocksQLListener) ExitAdminCommand(ctx *AdminCommandContext) {}

// EnterConfigCommand is called when production configCommand is entered.
func (s *BaseRocksQLListener) EnterConfigCommand(ctx *ConfigCommandContext) {}

// ExitConfigCommand is called when production configCommand is exited.
func (s *BaseRocksQLListener) ExitConfigCommand(ctx *ConfigCommandContext) {}

// EnterKey is called when production key is entered.
func (s *BaseRocksQLListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseRocksQLListener) ExitKey(ctx *KeyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseRocksQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseRocksQLListener) ExitValue(ctx *ValueContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseRocksQLListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseRocksQLListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterOption is called when production option is entered.
func (s *BaseRocksQLListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseRocksQLListener) ExitOption(ctx *OptionContext) {}
