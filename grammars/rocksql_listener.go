// Code generated from grammars/RocksQL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammars // RocksQL
import "github.com/antlr4-go/antlr/v4"

// RocksQLListener is a complete listener for a parse tree produced by RocksQLParser.
type RocksQLListener interface {
	antlr.ParseTreeListener

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterPutOperation is called when entering the putOperation production.
	EnterPutOperation(c *PutOperationContext)

	// EnterGetOperation is called when entering the getOperation production.
	EnterGetOperation(c *GetOperationContext)

	// EnterDeleteOperation is called when entering the deleteOperation production.
	EnterDeleteOperation(c *DeleteOperationContext)

	// EnterBatchEntry is called when entering the batchEntry production.
	EnterBatchEntry(c *BatchEntryContext)

	// EnterAdminCommand is called when entering the adminCommand production.
	EnterAdminCommand(c *AdminCommandContext)

	// EnterConfigCommand is called when entering the configCommand production.
	EnterConfigCommand(c *ConfigCommandContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitPutOperation is called when exiting the putOperation production.
	ExitPutOperation(c *PutOperationContext)

	// ExitGetOperation is called when exiting the getOperation production.
	ExitGetOperation(c *GetOperationContext)

	// ExitDeleteOperation is called when exiting the deleteOperation production.
	ExitDeleteOperation(c *DeleteOperationContext)

	// ExitBatchEntry is called when exiting the batchEntry production.
	ExitBatchEntry(c *BatchEntryContext)

	// ExitAdminCommand is called when exiting the adminCommand production.
	ExitAdminCommand(c *AdminCommandContext)

	// ExitConfigCommand is called when exiting the configCommand production.
	ExitConfigCommand(c *ConfigCommandContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)
}
