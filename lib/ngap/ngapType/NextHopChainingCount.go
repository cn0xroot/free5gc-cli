package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type NextHopChainingCount struct {
	Value int64 `aper:"valueLB:0,valueUB:7"`
}
