// PKGPATH: gno.land/r/test
package test

type MyStruct struct{}

func (ms MyStruct) Panic() {
	panic("panic")
}

func main() {
	ms := MyStruct{}
	ms.Panic()
}

// Stacktrace:
// panic: panic
// ms<VPBlock(1,0)>.Panic()
//     gno.land/r/test/main.gno:7
// main<VPBlock(1,1)>()
//     gno.land/r/test/main.gno:12
