const (

	/* Control operators */
	OpInvalid             Op = 0x00 // invalid
//	OpHalt                Op = 0x01 // halt (e.g. last statement)
	OpNoop                Op = 0x02 // no-op
//	OpExec                Op = 0x03 // exec next statement
	OpPrecall             Op = 0x04 // sets X (func) to frame
	OpCall                Op = 0x05 // call(Frame.Func, [...])
	OpCallNativeBody      Op = 0x06 // call body is native
	OpReturn              Op = 0x07 // return ...
//	OpReturnFromBlock     Op = 0x08 // return results (after defers)
	OpReturnToBlock       Op = 0x09 // copy results to block (before defer)
	OpDefer               Op = 0x0A // defer call(X, [...])
	OpCallDeferNativeBody Op = 0x0B // call body is native
	OpGo                  Op = 0x0C // go call(X, [...])
	OpSelect              Op = 0x0D // exec next select case
	OpSwitchClause        Op = 0x0E // exec next switch clause
	OpSwitchClauseCase    Op = 0x0F // exec next switch clause case
	OpTypeSwitch          Op = 0x10 // exec type switch clauses (all)
	OpIfCond              Op = 0x11 // eval cond
	OpPopValue            Op = 0x12 // pop X
	OpPopResults          Op = 0x13 // pop n call results
	OpPopBlock            Op = 0x14 // pop block NOTE breaks certain invariants.
	OpPopFrameAndReset    Op = 0x15 // pop frame and reset.
	OpPanic1              Op = 0x16 // pop exception and pop call frames.
	OpPanic2              Op = 0x17 // pop call frames.

	/* Unary & binary operators */
	OpUpos  Op = 0x20 // + (unary)
	OpUneg  Op = 0x21 // - (unary)
	OpUnot  Op = 0x22 // ! (unary)
	OpUxor  Op = 0x23 // ^ (unary)
	OpUrecv Op = 0x25 // <- (unary) // TODO make expr
	OpLor   Op = 0x26 // ||
	OpLand  Op = 0x27 // &&
	OpEql   Op = 0x28 // ==
	OpNeq   Op = 0x29 // !=
	OpLss   Op = 0x2A // <
	OpLeq   Op = 0x2B // <=
	OpGtr   Op = 0x2C // >
	OpGeq   Op = 0x2D // >=
	OpAdd   Op = 0x2E // +
	OpSub   Op = 0x2F // -
	OpBor   Op = 0x30 // |
	OpXor   Op = 0x31 // ^
	OpMul   Op = 0x32 // *
	OpQuo   Op = 0x33 // /
	OpRem   Op = 0x34 // %
	OpShl   Op = 0x35 // <<
	OpShr   Op = 0x36 // >>
	OpBand  Op = 0x37 // &
	OpBandn Op = 0x38 // &^

	/* Other expression operators */
	OpEval         Op = 0x40 // eval next expression
	OpBinary1      Op = 0x41 // X op ?
	OpIndex1       Op = 0x42 // X[Y]
	OpIndex2       Op = 0x43 // (_, ok :=) X[Y]
	OpSelector     Op = 0x44 // X.Y
	OpSlice        Op = 0x45 // X[Low:High:Max]
	OpStar         Op = 0x46 // *X (deref or pointer-to)
	OpRef          Op = 0x47 // &X
	OpTypeAssert1  Op = 0x48 // X.(Type)
	OpTypeAssert2  Op = 0x49 // (_, ok :=) X.(Type)
	OpStaticTypeOf Op = 0x4A // static type of X
	OpCompositeLit Op = 0x4B // X{???}
	OpArrayLit     Op = 0x4C // [Len]{...}
	OpSliceLit     Op = 0x4D // []{value,...}
	OpSliceLit2    Op = 0x4E // []{key:value,...}
	OpMapLit       Op = 0x4F // X{...}
	OpStructLit    Op = 0x50 // X{...}
	OpFuncLit      Op = 0x51 // func(T){Body}
	OpConvert      Op = 0x52 // Y(X)

	/* Native operators */
	OpArrayLitGoNative  Op = 0x60
	OpSliceLitGoNative  Op = 0x61
	OpStructLitGoNative Op = 0x62
	OpCallGoNative      Op = 0x63

	/* Type operators */
	OpFieldType       Op = 0x70 // Name: X `tag`
	OpArrayType       Op = 0x71 // [X]Y{}
	OpSliceType       Op = 0x72 // []X{}
	OpPointerType     Op = 0x73 // *X
	OpInterfaceType   Op = 0x74 // interface{...}
	OpChanType        Op = 0x75 // [<-]chan[<-]X
	OpFuncType        Op = 0x76 // func(params...)results...
	OpMapType         Op = 0x77 // map[X]Y
	OpStructType      Op = 0x78 // struct{...}
	OpMaybeNativeType Op = 0x79 // maybenative{X}

	/* Statement operators */
	OpAssign      Op = 0x80 // Lhs = Rhs
	OpAddAssign   Op = 0x81 // Lhs += Rhs
	OpSubAssign   Op = 0x82 // Lhs -= Rhs
	OpMulAssign   Op = 0x83 // Lhs *= Rhs
	OpQuoAssign   Op = 0x84 // Lhs /= Rhs
	OpRemAssign   Op = 0x85 // Lhs %= Rhs
	OpBandAssign  Op = 0x86 // Lhs &= Rhs
	OpBandnAssign Op = 0x87 // Lhs &^= Rhs
	OpBorAssign   Op = 0x88 // Lhs |= Rhs
	OpXorAssign   Op = 0x89 // Lhs ^= Rhs
	OpShlAssign   Op = 0x8A // Lhs <<= Rhs
	OpShrAssign   Op = 0x8B // Lhs >>= Rhs
	OpDefine      Op = 0x8C // X... := Y...
	OpInc         Op = 0x8D // X++
	OpDec         Op = 0x8E // X--

	/* Decl operators */
//	OpValueDecl Op = 0x90 // var/const ...
//	OpTypeDecl  Op = 0x91 // type ...

	/* Loop (sticky) operators (>= 0xD0) */
	OpSticky            Op = 0xD0 // not a real op.
//	OpBody              Op = 0xD1 // if/block/switch/select.
	OpForLoop           Op = 0xD2
	OpRangeIter         Op = 0xD3
	OpRangeIterString   Op = 0xD4
	OpRangeIterMap      Op = 0xD5
	OpRangeIterArrayPtr Op = 0xD6
	OpReturnCallDefers  Op = 0xD7 // TODO rename?
)
// 124-18 = 106
