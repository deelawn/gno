package gnolang

import (
	"fmt"
	"strings"
	"unicode"
	"unsafe"

	"github.com/gnolang/gno/tm2/pkg/crypto"
)

//----------------------------------------
// Misc.

func cp(bz []byte) (ret []byte) {
	ret = make([]byte, len(bz))
	copy(ret, bz)
	return ret
}

// Returns the associated machine operation for binary AST operations.  TODO:
// to make this faster and inlineable, remove the switch statement and create a
// mathematical mapping between them.
func word2BinaryOp(w Word) Op {
	switch w {
	case ADD:
		return OpAdd
	case SUB:
		return OpSub
	case MUL:
		return OpMul
	case QUO:
		return OpQuo
	case REM:
		return OpRem
	case BAND:
		return OpBand
	case BOR:
		return OpBor
	case XOR:
		return OpXor
	case SHL:
		return OpShl
	case SHR:
		return OpShr
	case BAND_NOT:
		return OpBandn
	case LAND:
		return OpLand
	case LOR:
		return OpLor
	case EQL:
		return OpEql
	case LSS:
		return OpLss
	case GTR:
		return OpGtr
	case NEQ:
		return OpNeq
	case LEQ:
		return OpLeq
	case GEQ:
		return OpGeq
	default:
		panic(fmt.Sprintf("unexpected binary operation word %v", w.String()))
	}
}

func word2UnaryOp(w Word) Op {
	switch w {
	case ADD:
		return OpUpos
	case SUB:
		return OpUneg
	case NOT:
		return OpUnot
	case XOR:
		return OpUxor
	case MUL:
		panic("unexpected unary operation * - use StarExpr instead")
	case BAND:
		panic("unexpected unary operation & - use RefExpr instead")
	case ARROW:
		return OpUrecv
	default:
		panic("unexpected unary operation")
	}
}

func toString(n Node) string {
	if n == nil {
		return "<nil>"
	}
	return n.String()
}

// true if the first rune is uppercase.
func isUpper(s string) bool {
	var first rune
	for _, c := range s {
		first = c
		break
	}
	return unicode.IsUpper(first)
}

//----------------------------------------
// converting uintptr to bytes.

const sizeOfUintPtr = unsafe.Sizeof(uintptr(0))

func uintptrToBytes(u *uintptr) []byte {
	return (*[sizeOfUintPtr]byte)(unsafe.Pointer(u))[:]
}

func defaultPkgName(gopkgPath string) Name {
	parts := strings.Split(gopkgPath, "/")
	last := parts[len(parts)-1]
	parts = strings.Split(last, "-")
	name := parts[len(parts)-1]
	name = strings.ToLower(name)
	return Name(name)
}

//----------------------------------------
// value convenience

func toTypeValue(t Type) TypeValue {
	return TypeValue{
		Type: t,
	}
}

//----------------------------------------
// reserved & uverse names

var reservedNames = map[Name]struct{}{
	"break": {}, "default": {}, "func": {}, "interface": {}, "select": {},
	"case": {}, "defer": {}, "go": {}, "map": {}, "struct": {},
	"chan": {}, "else": {}, "goto": {}, "package": {}, "switch": {},
	"const": {}, "fallthrough": {}, "if": {}, "range": {}, "type": {},
	"continue": {}, "for": {}, "import": {}, "return": {}, "var": {},
}

// if true, caller should generally panic.
func isReservedName(n Name) bool {
	_, ok := reservedNames[n]
	return ok
}

// scans uverse static node for blocknames. (slow)
func isUverseName(n Name) bool {
	uverseNames := UverseNode().GetBlockNames()
	for _, name := range uverseNames {
		if name == n {
			return true
		}
	}
	return false
}

//----------------------------------------
// other

// For keeping record of package & realm coins.
func DerivePkgAddr(pkgPath string) crypto.Address {
	// NOTE: must not collide with pubkey addrs.
	return crypto.AddressFromPreimage([]byte("pkgPath:" + pkgPath))
}

var opString = map[Op]string{
	/* Control operators */
	OpInvalid:             "OpInvalid",
	OpHalt:                "OpHalt",
	OpNoop:                "OpNoop",
	OpExec:                "OpExec",
	OpPrecall:             "OpPrecall",
	OpCall:                "OpCall",
	OpCallNativeBody:      "OpCallNativeBody",
	OpReturn:              "OpReturn",
	OpReturnFromBlock:     "OpReturnFromBlock",
	OpReturnToBlock:       "OpReturnToBlock",
	OpDefer:               "OpDefer",
	OpCallDeferNativeBody: "OpCallDeferNativeBody",
	OpGo:                  "OpGo",
	OpSelect:              "OpSelect",
	OpSwitchClause:        "OpSwitchClause",
	OpSwitchClauseCase:    "OpSwitchClauseCase",
	OpTypeSwitch:          "OpTypeSwitch",
	OpIfCond:              "OpIfCond",
	OpPopValue:            "OpPopValue",
	OpPopResults:          "OpPopResults",
	OpPopBlock:            "OpPopBlock",
	OpPopFrameAndReset:    "OpPopFrameAndReset",
	OpPanic1:              "OpPanic1",
	OpPanic2:              "OpPanic2",

	/* Unary & binary operators */
	OpUpos:  "OpUpos",
	OpUneg:  "OpUneg",
	OpUnot:  "OpUnot",
	OpUxor:  "OpUxor",
	OpUrecv: "OpUrecv",
	OpLor:   "OpLor",
	OpLand:  "OpLand",
	OpEql:   "OpEql",
	OpNeq:   "OpNeq",
	OpLss:   "OpLss",
	OpLeq:   "OpLeq",
	OpGtr:   "OpGtr",
	OpGeq:   "OpGeq",
	OpAdd:   "OpAdd",
	OpSub:   "OpSub",
	OpBor:   "OpBor",
	OpXor:   "OpXor",
	OpMul:   "OpMul",
	OpQuo:   "OpQuo",
	OpRem:   "OpRem",
	OpShl:   "OpShl",
	OpShr:   "OpShr",
	OpBand:  "OpBand",
	OpBandn: "OpBandn",

	/* Other expression operators */
	OpEval:         "OpEval",
	OpBinary1:      "OpBinary1",
	OpIndex1:       "OpIndex1",
	OpIndex2:       "OpIndex2",
	OpSelector:     "OpSelector",
	OpSlice:        "OpSlice",
	OpStar:         "OpStar",
	OpRef:          "OpRef",
	OpTypeAssert1:  "OpTypeAssert1",
	OpTypeAssert2:  "OpTypeAssert2",
	OpStaticTypeOf: "OpStaticTypeOf",
	OpCompositeLit: "OpCompositeLit",
	OpArrayLit:     "OpArrayLit",
	OpSliceLit:     "OpSliceLit",
	OpSliceLit2:    "OpSliceLit2",
	OpMapLit:       "OpMapLit",
	OpStructLit:    "OpStructLit",
	OpFuncLit:      "OpFuncLit",
	OpConvert:      "OpConvert",

	/* Native operators */
	OpArrayLitGoNative:  "OpArrayLitGoNative",
	OpSliceLitGoNative:  "OpSliceLitGoNative",
	OpStructLitGoNative: "OpStructLitGoNative",
	OpCallGoNative:      "OpCallGoNative",

	/* Type operators */
	OpFieldType:       "OpFieldType",
	OpArrayType:       "OpArrayType",
	OpSliceType:       "OpSliceType",
	OpPointerType:     "OpPointerType",
	OpInterfaceType:   "OpInterfaceType",
	OpChanType:        "OpChanType",
	OpFuncType:        "OpFuncType",
	OpMapType:         "OpMapType",
	OpStructType:      "OpStructType",
	OpMaybeNativeType: "OpMaybeNativeType",

	/* Statement operators */
	OpAssign:      "OpAssign",
	OpAddAssign:   "OpAddAssign",
	OpSubAssign:   "OpSubAssign",
	OpMulAssign:   "OpMulAssign",
	OpQuoAssign:   "OpQuoAssign",
	OpRemAssign:   "OpRemAssign",
	OpBandAssign:  "OpBandAssign",
	OpBandnAssign: "OpBandnAssign",
	OpBorAssign:   "OpBorAssign",
	OpXorAssign:   "OpXorAssign",
	OpShlAssign:   "OpShlAssign",
	OpShrAssign:   "OpShrAssign",
	OpDefine:      "OpDefine",
	OpInc:         "OpInc",
	OpDec:         "OpDec",

	/* Decl operators */
	OpValueDecl: "OpValueDecl",
	OpTypeDecl:  "OpTypeDecl",

	/* Loop (sticky) operators (>= 0xD0) */
	OpSticky:            "OpSticky",
	OpBody:              "OpBody",
	OpForLoop:           "OpForLoop",
	OpRangeIter:         "OpRangeIter",
	OpRangeIterString:   "OpRangeIterString",
	OpRangeIterMap:      "OpRangeIterMap",
	OpRangeIterArrayPtr: "OpRangeIterArrayPtr",
	OpReturnCallDefers:  "OpReturnCallDefers",
}

var opCPU = map[Op]int64{
	/* Control operators */
	OpInvalid:             1,
	OpHalt:                1,
	OpNoop:                1,
	OpExec:                1,
	OpPrecall:             1,
	OpCall:                1,
	OpCallNativeBody:      1,
	OpReturn:              1,
	OpReturnFromBlock:     1,
	OpReturnToBlock:       1,
	OpDefer:               1,
	OpCallDeferNativeBody: 1,
	OpGo:                  1,
	OpSelect:              1,
	OpSwitchClause:        1,
	OpSwitchClauseCase:    1,
	OpTypeSwitch:          1,
	OpIfCond:              1,
	OpPopValue:            1,
	OpPopResults:          1,
	OpPopBlock:            1,
	OpPopFrameAndReset:    1,
	OpPanic1:              1,
	OpPanic2:              1,

	/* Unary & binary operators */
	OpUpos:  1,
	OpUneg:  1,
	OpUnot:  1,
	OpUxor:  1,
	OpUrecv: 1,
	OpLor:   1,
	OpLand:  1,
	OpEql:   1,
	OpNeq:   1,
	OpLss:   1,
	OpLeq:   1,
	OpGtr:   1,
	OpGeq:   1,
	OpAdd:   1,
	OpSub:   1,
	OpBor:   1,
	OpXor:   1,
	OpMul:   1,
	OpQuo:   1,
	OpRem:   1,
	OpShl:   1,
	OpShr:   1,
	OpBand:  1,
	OpBandn: 1,

	/* Other expression operators */
	OpEval:         1,
	OpBinary1:      1,
	OpIndex1:       1,
	OpIndex2:       1,
	OpSelector:     1,
	OpSlice:        1,
	OpStar:         1,
	OpRef:          1,
	OpTypeAssert1:  1,
	OpTypeAssert2:  1,
	OpStaticTypeOf: 1,
	OpCompositeLit: 1,
	OpArrayLit:     1,
	OpSliceLit:     1,
	OpSliceLit2:    1,
	OpMapLit:       1,
	OpStructLit:    1,
	OpFuncLit:      1,
	OpConvert:      1,

	/* Native operators */
	OpArrayLitGoNative:  1,
	OpSliceLitGoNative:  1,
	OpStructLitGoNative: 1,
	OpCallGoNative:      1,

	/* Type operators */
	OpFieldType:       1,
	OpArrayType:       1,
	OpSliceType:       1,
	OpPointerType:     1,
	OpInterfaceType:   1,
	OpChanType:        1,
	OpFuncType:        1,
	OpMapType:         1,
	OpStructType:      1,
	OpMaybeNativeType: 1,

	/* Statement operators */
	OpAssign:      1,
	OpAddAssign:   1,
	OpSubAssign:   1,
	OpMulAssign:   1,
	OpQuoAssign:   1,
	OpRemAssign:   1,
	OpBandAssign:  1,
	OpBandnAssign: 1,
	OpBorAssign:   1,
	OpXorAssign:   1,
	OpShlAssign:   1,
	OpShrAssign:   1,
	OpDefine:      1,
	OpInc:         1,
	OpDec:         1,

	/* Decl operators */
	OpValueDecl: 1,
	OpTypeDecl:  1,

	/* Loop (sticky) operators (>= 0xD0) */
	OpSticky:            1,
	OpBody:              1,
	OpForLoop:           1,
	OpRangeIter:         1,
	OpRangeIterString:   1,
	OpRangeIterMap:      1,
	OpRangeIterArrayPtr: 1,
	OpReturnCallDefers:  1,
}
