package benchmarking

import "os"

const (
	KEEPER_CALL   = "call"
	KEEPER_INIT   = "init"
	KEEPER_ADDPKG = "addpkg"
)

// There are two control points to isolate benchmarking scope.
// - Keeper entry points at Init, Msg_Call, Msg_AddPkg
var (
	Entry string
	// - We set Start bencharking for true after an OpCode executed
	Start bool
	// we only turn OpCodeDetails on to understand the OpCode in benchmarking call flow. We turn it off for accurate measurement timing
	OpCodeDetails bool
)

var enabled bool

func Enabled() bool {
	return enabled
}

func Init(filepath string) {
	enabled = true
	initExporter(filepath)
	initStack()
	if os.Getenv("OPCODE_DETAILS") == "true" {
		OpCodeDetails = true
	}
}
