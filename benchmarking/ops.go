package benchmarking

const (
	OpStoreGetObject     byte = 0x01 // get value from store
	OpStoreSetObject     byte = 0x02 // set value in store
	OpStoreDeleteObject  byte = 0x03 // delete value from store
	OpStoreGetPackage    byte = 0x04 // get package from store
	OpStoreGetType       byte = 0x05 // get type from store
	OpStoreSetType       byte = 0x06 // set type in store
	OpStoreGetBlockNode  byte = 0x07 // get block node from store
	OpStoreSetBlockNode  byte = 0x08 // set block node in store
	OpStoreAddMemPackage byte = 0x09 // add mempackage to store
	OpStoreGetMemPackage byte = 0x0A // get mempackage from store
	OpFinalizeTx         byte = 0x0B // finalize realm transaction

	invalidStorageOp string = "OpStoreInvalid"
)

var opCodeNames = []string{
	invalidStorageOp,
	"OpStoreGetObject",
	"OpStoreSetObject",
	"OpStoreDeleteObject",
	"OpStoreGetPackage",
	"OpStoreGetType",
	"OpStoreSetType",
	"OpStoreGetBlockNode",
	"OpStoreSetBlockNode",
	"OpStoreAddMemPackage",
	"OpStoreGetMemPackage",
	"OpFinalizeTx",
}

type OpCode [2]byte

func VMOpCode(op byte) OpCode {
	return [2]byte{op, 0x00}
}

func StorageOpCode(op byte) OpCode {
	return [2]byte{0x00, op}
}

func OpCodeString(op byte) string {
	if int(op) >= len(opCodeNames) {
		return invalidStorageOp
	}
	return opCodeNames[op]
}
