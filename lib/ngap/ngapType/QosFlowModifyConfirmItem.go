package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type QosFlowModifyConfirmItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowModifyConfirmItemExtIEs `aper:"optional"`
}
