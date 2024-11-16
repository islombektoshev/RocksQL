grammar RocksQL;

// Starting point of parsing
command: operation EOF;

// Main operations for RocksDB
operation
    : putOperation
    | getOperation
    | deleteOperation
    | adminCommand
    ;

// Put operation (e.g., PUT key value)
putOperation: 'PUT' key value;

// Get operation (e.g., GET key)
getOperation: 'GET' key;

// Delete operation (e.g., DELETE key)
deleteOperation: 'DEL' key;


batchEntry
    : putOperation
    | deleteOperation
    ;

// Administrative commands (e.g., ADMIN COMPACT, ADMIN CONFIG SET option value)
adminCommand
    : 'ADMIN' ('COMPACT' | 'STATS' | configCommand)
    ;

configCommand: 'CONFIG' 'SET' option value;

// Basic data types
key: STRING | IDENTIFIER;
value: STRING | NUMBER;
identifier: IDENTIFIER;
option: IDENTIFIER;

// Lexer rules
STRING: '"' (~["\\] | '\\' .)* '"';
NUMBER: [0-9]+ ('.' [0-9]+)?;
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

// Whitespace and comments
WS: [ \t\r\n]+ -> skip;
COMMENT: '//' ~[\r\n]* -> skip;
