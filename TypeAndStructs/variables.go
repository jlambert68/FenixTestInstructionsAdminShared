package TypeAndStructs

type DomainUUIDType string
type DomainNameType string
type OriginalElementUUIDType string
type OriginalElementNameType string
type TestInstructionNameType string
type TestInstructionTypeUUIDType string
type TestInstructionTypeNameType string
type UpdatedTimeStampType string
type ColorType string
type TCRuleDeletionType string
type TCRuleSwapType string
type DropZoneUUIDType string
type DropZoneNameType string
type TestInstructionAttributeTypeType string
type TestInstructionAttributeUUIDType string
type TestInstructionAttributeNameType string
type AttributeValueAsStringType string
type AttributeValueUUIDType string
type AttributeActionCommandType int
type TestInstructionAttributeInputMaskType string
type TestInstructionAttributeTypeUUIDType string
type TestInstructionAttributeTypeNameType string

type TestInstructionContainerNameType string
type TestInstructionContainerTypeUUIDType string
type TestInstructionContainerTypeNameType string
type TestInstructionContainerExecutionTypeType string //'SERIAL_PROCESSED', 'PARALLELLED_PROCESSED'

type BondUUIDType string
type BondNameType string
type BondDescriptionType string
type BondMouseOverTextType string
type BondDeprecatedType bool
type BondEnabledType bool
type BondVisibleType bool
type BondColorType string
type BondCanBeDeletedType bool
type BondCanBeSwappedOutType bool
type BondShowBondAttributesType bool
type BondTCRuleDeletionType string
type BondTCRuleSwapType string
type BondUpdatedTimeStampType string

// type BondTestCaseModelElementTypeType string
// type BondTestCaseModelElementTypeGrpcMappingIDType int

type TestCaseModelElementTypeType string
type TestCaseModelElementDescriptionType string
type TestCaseModelElementGrpcMappingIdType int

type TestInstructions []TestInstructionStruct
type TestInstructionStruct struct {
	DomainUUID                   DomainUUIDType              `json:"DomainUuid"`
	DomainName                   DomainNameType              `json:"DomainName"`
	ExecutionDomainUUID          DomainUUIDType              `json:"ExecutionDomainUUID"`
	ExecutionDomainName          DomainNameType              `json:"ExecutionDomainName"`
	TestInstructionUUID          OriginalElementUUIDType     `json:"TestInstructionUuid"`
	TestInstructionName          TestInstructionNameType     `json:"TestInstructionName"`
	TestInstructionTypeUUID      TestInstructionTypeUUIDType `json:"TestInstructionTypeUuid"`
	TestInstructionTypeName      TestInstructionTypeNameType `json:"TestInstructionTypeName"`
	TestInstructionDescription   string                      `json:"TestInstructionDescription"`
	TestInstructionMouseOverText string                      `json:"TestInstructionMouseOverText"`
	Deprecated                   bool                        `json:"Deprecated"`
	Enabled                      bool                        `json:"Enabled"`
	MajorVersionNumber           int                         `json:"MajorVersionNumber"`
	MinorVersionNumber           int                         `json:"MinorVersionNumber"`
	UpdatedTimeStamp             UpdatedTimeStampType        `json:"UpdatedTimeStamp"`
}

type BasicTestInstructionsInformation []BasicTestInstructionInformationStruct
type BasicTestInstructionInformationStruct struct {
	DomainUUID                   DomainUUIDType              `json:"DomainUuid"`
	DomainName                   DomainNameType              `json:"DomainName"`
	ExecutionDomainUUID          DomainUUIDType              `json:"ExecutionDomainUUID"`
	ExecutionDomainName          DomainNameType              `json:"ExecutionDomainName"`
	TestInstructionUUID          OriginalElementUUIDType     `json:"TestInstructionUuid"`
	TestInstructionName          TestInstructionNameType     `json:"TestInstructionName"`
	TestInstructionTypeUUID      TestInstructionTypeUUIDType `json:"TestInstructionTypeUuid"`
	TestInstructionTypeName      TestInstructionTypeNameType `json:"TestInstructionTypeName"`
	Deprecated                   bool                        `json:"Deprecated"`
	MajorVersionNumber           int                         `json:"MajorVersionNumber"`
	MinorVersionNumber           int                         `json:"MinorVersionNumber"`
	UpdatedTimeStamp             UpdatedTimeStampType        `json:"UpdatedTimeStamp"`
	TestInstructionColor         ColorType                   `json:"TestInstructionColor"`
	TCRuleDeletion               TCRuleDeletionType          `json:"TCRuleDeletion"`
	TCRuleSwap                   TCRuleSwapType              `json:"TCRuleSwap"`
	TestInstructionDescription   string                      `json:"TestInstructionDescription"`
	TestInstructionMouseOverText string                      `json:"TestInstructionMouseOverText"`
	Enabled                      bool                        `json:"Enabled"`
}

type ImmatureTestInstructionsInformation []ImmatureTestInstructionInformationStruct
type ImmatureTestInstructionInformationStruct struct {
	DomainUUID                   DomainUUIDType                   `json:"DomainUuid"`
	DomainName                   DomainNameType                   `json:"DomainName"`
	TestInstructionUUID          OriginalElementUUIDType          `json:"TestInstructionUuid"`
	TestInstructionName          TestInstructionNameType          `json:"TestInstructionName"`
	DropZoneUUID                 DropZoneUUIDType                 `json:"DropZoneUuid"`
	DropZoneName                 DropZoneNameType                 `json:"DropZoneName"`
	DropZoneDescription          string                           `json:"DropZoneDescription"`
	DropZoneMouseOver            string                           `json:"DropZoneMouseOver"`
	DropZoneColor                ColorType                        `json:"DropZoneColor"`
	TestInstructionAttributeType TestInstructionAttributeTypeType `json:"TestInstructionAttributeType"` //TEXTBOX...
	TestInstructionAttributeUUID TestInstructionAttributeUUIDType `json:"TestInstructionAttributeUuid"`
	TestInstructionAttributeName TestInstructionAttributeNameType `json:"TestInstructionAttributeName"`
	AttributeValueAsString       AttributeValueAsStringType       `json:"AttributeValueAsString"`
	AttributeValueUUID           AttributeValueUUIDType           `json:"AttributeValueUuid"`
	FirstImmatureElementUUID     OriginalElementUUIDType          `json:"FirstImmatureElementUuid"`
	AttributeActionCommand       AttributeActionCommandType       `json:"AttributeActionCommand"`
}

type TestInstructionAttributes []TestInstructionAttributeStruct
type TestInstructionAttributeStruct struct {
	DomainUUID                                    DomainUUIDType                        `json:"DomainUuid"`
	DomainName                                    DomainNameType                        `json:"DomainName"`
	TestInstructionUUID                           OriginalElementUUIDType               `json:"TestInstructionUuid"`
	TestInstructionName                           TestInstructionNameType               `json:"TestInstructionName"`
	TestInstructionAttributeUUID                  TestInstructionAttributeUUIDType      `json:"TestInstructionAttributeUuid"`
	TestInstructionAttributeName                  TestInstructionAttributeNameType      `json:"TestInstructionAttributeName"`
	TestInstructionAttributeDescription           string                                `json:"TestInstructionAttributeDescription"`
	TestInstructionAttributeMouseOver             string                                `json:"TestInstructionAttributeMouseOver"`
	TestInstructionAttributeTypeUUID              TestInstructionAttributeTypeUUIDType  `json:"TestInstructionAttributeTypeUuid"`
	TestInstructionAttributeTypeName              TestInstructionAttributeTypeNameType  `json:"TestInstructionAttributeTypeName"`
	TestInstructionAttributeValueAsString         AttributeValueAsStringType            `json:"TestInstructionAttributeValueAsString"`
	TestInstructionAttributeValueUUID             AttributeValueUUIDType                `json:"TestInstructionAttributeValueUuid"`
	TestInstructionAttributeVisible               bool                                  `json:"TestInstructionAttributeVisible"`
	TestInstructionAttributeEnabled               bool                                  `json:"TestInstructionAttributeEnabled"`
	TestInstructionAttributeMandatory             bool                                  `json:"TestInstructionAttributeMandatory"`
	TestInstructionAttributeVisibleInTestCaseArea bool                                  `json:"TestInstructionAttributeVisibleInTestCaseArea"`
	TestInstructionAttributeIsDeprecated          bool                                  `json:"TestInstructionAttributeIsDeprecated"`
	TestInstructionAttributeInputMask             TestInstructionAttributeInputMaskType `json:"TestInstructionAttributeInputMask"`
	TestInstructionAttributeType                  TestInstructionAttributeTypeType      `json:"TestInstructionAttributeType"` // TEXTBOX...
}

type TestInstructionContainers []TestInstructionContainerStruct
type TestInstructionContainerStruct struct {
	DomainUUID                            DomainUUIDType                       `json:"DomainUuid"`
	DomainName                            DomainNameType                       `json:"DomainName"`
	TestInstructionContainerUUID          OriginalElementUUIDType              `json:"TestInstructionContainerUuid"`
	TestInstructionContainerName          TestInstructionContainerNameType     `json:"TestInstructionContainerName"`
	TestInstructionContainerTypeUUID      TestInstructionContainerTypeUUIDType `json:"TestInstructionContainerTypeUuid"`
	TestInstructionContainerTypeName      TestInstructionContainerTypeNameType `json:"TestInstructionContainerTypeName"`
	TestInstructionContainerDescription   string                               `json:"TestInstructionContainerDescription"`
	TestInstructionContainerMouseOverText string                               `json:"TestInstructionContainerMouseOverText"`
	Deprecated                            bool                                 `json:"Deprecated"`
	Enabled                               bool                                 `json:"Enabled"`
	MajorVersionNumber                    int                                  `json:"MajorVersionNumber"`
	MinorVersionNumber                    int                                  `json:"MinorVersionNumber"`
	UpdatedTimeStamp                      UpdatedTimeStampType                 `json:"UpdatedTimeStamp"`
	ChildrenIsParallelProcessed           bool                                 `json:"ChildrenIsParallelProcessed"`
}

type BasicTestInstructionContainersInformation []BasicTestInstructionContainerInformationStruct
type BasicTestInstructionContainerInformationStruct struct {
	DomainUUID                            DomainUUIDType                            `json:"DomainUuid"`
	DomainName                            DomainNameType                            `json:"DomainName"`
	TestInstructionContainerUUID          OriginalElementUUIDType                   `json:"TestInstructionContainerUuid"`
	TestInstructionContainerName          TestInstructionContainerNameType          `json:"TestInstructionContainerName"`
	TestInstructionContainerTypeUUID      TestInstructionContainerTypeUUIDType      `json:"TestInstructionContainerTypeUuid"`
	TestInstructionContainerTypeName      TestInstructionContainerTypeNameType      `json:"TestInstructionContainerTypeName"`
	Deprecated                            bool                                      `json:"Deprecated"`
	MajorVersionNumber                    int                                       `json:"MajorVersionNumber"`
	MinorVersionNumber                    int                                       `json:"MinorVersionNumber"`
	UpdatedTimeStamp                      UpdatedTimeStampType                      `json:"UpdatedTimeStamp"`
	TestInstructionContainerColor         ColorType                                 `json:"TestInstructionContainerColor"`
	TCRuleDeletion                        TCRuleDeletionType                        `json:"TCRuleDeletion"`
	TCRuleSwap                            TCRuleSwapType                            `json:"TCRuleSwap"`
	TestInstructionContainerDescription   string                                    `json:"TestInstructionContainerDescription"`
	TestInstructionContainerMouseOverText string                                    `json:"TestInstructionContainerMouseOverText"`
	Enabled                               bool                                      `json:"Enabled"`
	TestInstructionContainerExecutionType TestInstructionContainerExecutionTypeType `json:"TestInstructionContainerExecutionType"`
}

type ImmatureTestInstructionContainersMessage []ImmatureTestInstructionContainerMessageStruct
type ImmatureTestInstructionContainerMessageStruct struct {
	DomainUUID                   DomainUUIDType                   `json:"DomainUuid"`
	DomainName                   DomainNameType                   `json:"DomainName"`
	TestInstructionContainerUUID OriginalElementUUIDType          `json:"TestInstructionContainerUuid"`
	TestInstructionContainerName TestInstructionContainerNameType `json:"TestInstructionContainerName"`
	DropZoneUUID                 DropZoneUUIDType                 `json:"DropZoneUuid"`
	DropZoneName                 DropZoneNameType                 `json:"DropZoneName"`
	DropZoneDescription          string                           `json:"DropZoneDescription"`
	DropZoneMouseOver            string                           `json:"DropZoneMouseOver"`
	DropZoneColor                ColorType                        `json:"DropZoneColor"`
	TestInstructionAttributeType TestInstructionAttributeTypeType `json:"TestInstructionAttributeType"` //TEXTBOX
	TestInstructionAttributeUUID TestInstructionAttributeUUIDType `json:"TestInstructionAttributeUuid"`
	TestInstructionAttributeName TestInstructionAttributeNameType `json:"TestInstructionAttributeName"`
	AttributeValueAsString       AttributeValueAsStringType       `json:"AttributeValueAsString"`
	AttributeValueUUID           AttributeValueUUIDType           `json:"AttributeValueUuid"`
	FirstImmatureElementUUID     OriginalElementUUIDType          `json:"FirstImmatureElementUuid"`
	AttributeActionCommand       AttributeActionCommandType       `json:"AttributeActionCommand"`
}

type ImmatureElementModelsMessage []ImmatureElementModelMessageStruct
type ImmatureElementModelMessageStruct struct {
	DomainUUID               DomainUUIDType               `json:"DomainUuid"`
	DomainName               DomainNameType               `json:"DomainName"`
	ImmatureElementUUID      OriginalElementUUIDType      `json:"ImmatureElementUuid"`
	ImmatureElementName      OriginalElementNameType      `json:"ImmatureElementName"`
	PreviousElementUUID      OriginalElementUUIDType      `json:"PreviousElementUuid"`
	NextElementUUID          OriginalElementUUIDType      `json:"NextElementUuid"`
	FirstChildElementUUID    OriginalElementUUIDType      `json:"FirstChildElementUuid"`
	ParentElementUUID        OriginalElementUUIDType      `json:"ParentElementUuid"`
	TestCaseModelElementType TestCaseModelElementTypeType `json:"TestCaseModelElementType"`
	OriginalElementUUID      OriginalElementUUIDType      `json:"OriginalElementUuid"`
	TopImmatureElementUUID   OriginalElementUUIDType      `json:"TopImmatureElementUuid"`
	IsTopElement             bool                         `json:"IsTopElement"`
}

type ImmatureBondStruct struct {
	BondUUID                              OriginalElementUUIDType               `json:"BondUuid"`
	BondName                              BondNameType                          `json:"BondName"`
	BondDescription                       BondDescriptionType                   `json:"BondDescription"`
	BondMouseOverText                     BondMouseOverTextType                 `json:"BondMouseOverText"`
	Deprecated                            BondDeprecatedType                    `json:"Deprecated"`
	Enabled                               BondEnabledType                       `json:"Enabled"`
	Visible                               BondVisibleType                       `json:"Visible"`
	BondColor                             BondColorType                         `json:"BondColor"`
	CanBeDeleted                          BondCanBeDeletedType                  `json:"CanBeDeleted"`
	CanBeSwappedOut                       BondCanBeSwappedOutType               `json:"CanBeSwappedOut"`
	ShowBondAttributes                    BondShowBondAttributesType            `json:"ShowBondAttributes"`
	TCRuleDeletion                        BondTCRuleDeletionType                `json:"TCRuleDeletion"`
	TCRuleSwap                            BondTCRuleSwapType                    `json:"TCRuleSwap"`
	UpdatedTimeStamp                      UpdatedTimeStampType                  `json:"UpdatedTimeStamp"`
	TestCaseModelElementType              TestCaseModelElementTypeType          `json:"TestCaseModelElementType"`
	TestCaseModelElementTypeGrpcMappingID TestCaseModelElementGrpcMappingIdType `json:"TestCaseModelElementTypeGrpcMappingId"`
}
