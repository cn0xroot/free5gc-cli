package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct CancelledCellsInTAI_NR */
/* CancelledCellsInTAINRItem */
type CancelledCellsInTAINR struct {
	List []CancelledCellsInTAINRItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
