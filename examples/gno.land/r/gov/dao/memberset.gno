package govdao

import (
	"std"

	pproposal "gno.land/p/gov/proposal"
)

const daoPkgPath = "gno.land/r/gov/dao"

const (
	errNoChangesProposed = "no set changes proposed"
	errNotGovDAO         = "caller not govdao executor"
)

func NewPropExecutor(changesFn func() []std.Address) pproposal.Executor {
	if changesFn == nil {
		panic(errNoChangesProposed)
	}

	callback := func() error {
		// Make sure the GovDAO executor runs the valset changes
		assertGovDAOCaller()

		for _, addr := range changesFn() {
			members = append(members, addr)
		}

		return nil
	}

	return pproposal.NewExecutor(callback)
}

// assertGovDAOCaller verifies the caller is the GovDAO executor
func assertGovDAOCaller() {
	if std.CurrentRealm().PkgPath() != daoPkgPath {
		panic(errNotGovDAO)
	}
}
