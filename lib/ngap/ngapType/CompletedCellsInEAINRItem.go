package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type CompletedCellsInEAINRItem struct {
	NRCGI        NRCGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAINRItemExtIEs `aper:"optional"`
}
