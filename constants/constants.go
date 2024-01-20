package constants

const (
	EOL           = "\r\n"
	SIMPLE_STRING = "+"

	//lANGUAGE SPECIFIC
	// - error
	// + simple string
	// * array
	// $ bulkString
	// : int
	ERROR = "-"
	ARRAY = "*"
	NULL  = "_"
	BULK  = "$"
	SIZE  = "$"

	//commands
	CLI_CMD_PING = "PING"
)
