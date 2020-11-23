package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type PDUSessionResourceReleasedItemRelRes struct {
	PDUSessionID                              PDUSessionID
	PDUSessionResourceReleaseResponseTransfer aper.OctetString
	IEExtensions                              *ProtocolExtensionContainerPDUSessionResourceReleasedItemRelResExtIEs `aper:"optional"`
}
