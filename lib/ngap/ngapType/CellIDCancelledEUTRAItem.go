package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type CellIDCancelledEUTRAItem struct {
	EUTRACGI           EUTRACGI `aper:"valueExt"`
	NumberOfBroadcasts NumberOfBroadcasts
	IEExtensions       *ProtocolExtensionContainerCellIDCancelledEUTRAItemExtIEs `aper:"optional"`
}
