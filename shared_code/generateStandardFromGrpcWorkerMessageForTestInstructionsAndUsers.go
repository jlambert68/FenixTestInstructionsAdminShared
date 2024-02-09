package shared_code

import (
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
)

// GenerateStandardFromGrpcWorkerMessageForTestInstructionsAndUsers
// Generate the original message structure from Worker gRPC-version
func GenerateStandardFromGrpcWorkerMessageForTestInstructionsAndUsers(
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage *fenixExecutionWorkerGrpcApi.SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage) (
	testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct,
	err error) {

	// Convert TestInstructions from gRPC-message-version
	// Generate original-version of TestInstructionInstanceVersionsMessageMap
	var testInstructionInstanceVersionsMessageMap map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct
	testInstructionInstanceVersionsMessageMap = make(map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct)

	// Loop all TestInstructionInstanceVersionsMessages
	for originalElementUUIDTypeGrpc, testInstructionInstanceVersionsMessageGrpc := range supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.TestInstructions.TestInstructionsMap {

		// Loop TestInstructionVersions and create original-version
		var testInstructionInstanceVersionMessages []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct
		for _, testInstructionInstanceVersionMessageGrpc := range testInstructionInstanceVersionsMessageGrpc.TestInstructionVersions {

			// Loop and Create 'ImmatureTestInstructionsInformation' from 'testInstructionInstanceVersionMessageGrpc'
			var immatureTestInstructionInformationMessages []*TypeAndStructs.ImmatureTestInstructionInformationStruct
			for _, immatureTestInstructionInformationMessageGrpc := range testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.ImmatureTestInstructionInformation {

				// Create the gRPC-version of the message
				var immatureTestInstructionInformationMessage *TypeAndStructs.ImmatureTestInstructionInformationStruct
				immatureTestInstructionInformationMessage = &TypeAndStructs.ImmatureTestInstructionInformationStruct{
					DomainUUID:                   TypeAndStructs.DomainUUIDType(immatureTestInstructionInformationMessageGrpc.GetDomainUUID()),
					DomainName:                   TypeAndStructs.DomainNameType(immatureTestInstructionInformationMessageGrpc.GetDomainName()),
					TestInstructionUUID:          TypeAndStructs.OriginalElementUUIDType(immatureTestInstructionInformationMessageGrpc.GetTestInstructionUUID()),
					TestInstructionName:          TypeAndStructs.TestInstructionNameType(immatureTestInstructionInformationMessageGrpc.GetTestInstructionName()),
					DropZoneUUID:                 TypeAndStructs.DropZoneUUIDType(immatureTestInstructionInformationMessageGrpc.GetDropZoneUUID()),
					DropZoneName:                 TypeAndStructs.DropZoneNameType(immatureTestInstructionInformationMessageGrpc.GetDropZoneName()),
					DropZoneDescription:          immatureTestInstructionInformationMessageGrpc.GetDropZoneDescription(),
					DropZoneMouseOver:            immatureTestInstructionInformationMessageGrpc.GetDropZoneMouseOver(),
					DropZoneColor:                TypeAndStructs.ColorType(immatureTestInstructionInformationMessageGrpc.GetDropZoneColor()),
					TestInstructionAttributeType: TypeAndStructs.TestInstructionAttributeTypeType(immatureTestInstructionInformationMessageGrpc.GetTestInstructionAttributeType()),
					TestInstructionAttributeUUID: TypeAndStructs.TestInstructionAttributeUUIDType(immatureTestInstructionInformationMessageGrpc.GetTestInstructionAttributeUUID()),
					TestInstructionAttributeName: TypeAndStructs.TestInstructionAttributeNameType(immatureTestInstructionInformationMessageGrpc.GetTestInstructionAttributeName()),
					AttributeValueAsString:       TypeAndStructs.AttributeValueAsStringType(immatureTestInstructionInformationMessageGrpc.GetAttributeValueAsString()),
					AttributeValueUUID:           TypeAndStructs.AttributeValueUUIDType(immatureTestInstructionInformationMessageGrpc.GetAttributeValueUUID()),
					FirstImmatureElementUUID:     TypeAndStructs.OriginalElementUUIDType(immatureTestInstructionInformationMessageGrpc.GetFirstImmatureElementUUID()),
					AttributeActionCommand:       TypeAndStructs.AttributeActionCommandType(immatureTestInstructionInformationMessageGrpc.GetAttributeActionCommand()),
				}

				// Append to slice of messages
				immatureTestInstructionInformationMessages = append(immatureTestInstructionInformationMessages, immatureTestInstructionInformationMessage)
			}

			// Loop and Create 'TestInstructionAttributes' from 'testInstructionInstanceVersionMessageGrpc'
			var testInstructionAttributes []*TypeAndStructs.TestInstructionAttributeStruct
			for _, testInstructionAttributeGrpc := range testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstructionAttributes {

				// Create the gRPC-version of the message
				var testInstructionAttributeMessage *TypeAndStructs.TestInstructionAttributeStruct
				testInstructionAttributeMessage = &TypeAndStructs.TestInstructionAttributeStruct{
					DomainUUID:                                    TypeAndStructs.DomainUUIDType(testInstructionAttributeGrpc.GetDomainUUID()),
					DomainName:                                    TypeAndStructs.DomainNameType(testInstructionAttributeGrpc.GetDomainName()),
					TestInstructionUUID:                           TypeAndStructs.OriginalElementUUIDType(testInstructionAttributeGrpc.GetTestInstructionUUID()),
					TestInstructionName:                           TypeAndStructs.TestInstructionNameType(testInstructionAttributeGrpc.GetTestInstructionName()),
					TestInstructionAttributeUUID:                  TypeAndStructs.TestInstructionAttributeUUIDType(testInstructionAttributeGrpc.GetTestInstructionAttributeUUID()),
					TestInstructionAttributeName:                  TypeAndStructs.TestInstructionAttributeNameType(testInstructionAttributeGrpc.GetTestInstructionAttributeName()),
					TestInstructionAttributeDescription:           testInstructionAttributeGrpc.GetTestInstructionAttributeDescription(),
					TestInstructionAttributeMouseOver:             testInstructionAttributeGrpc.GetTestInstructionAttributeMouseOver(),
					TestInstructionAttributeTypeUUID:              TypeAndStructs.TestInstructionAttributeTypeUUIDType(testInstructionAttributeGrpc.GetTestInstructionAttributeTypeUUID()),
					TestInstructionAttributeTypeName:              TypeAndStructs.TestInstructionAttributeTypeNameType(testInstructionAttributeGrpc.GetTestInstructionAttributeTypeName()),
					TestInstructionAttributeValueAsString:         TypeAndStructs.AttributeValueAsStringType(testInstructionAttributeGrpc.GetTestInstructionAttributeValueAsString()),
					TestInstructionAttributeValueUUID:             TypeAndStructs.AttributeValueUUIDType(testInstructionAttributeGrpc.GetTestInstructionAttributeValueUUID()),
					TestInstructionAttributeVisible:               testInstructionAttributeGrpc.GetTestInstructionAttributeVisible(),
					TestInstructionAttributeEnabled:               testInstructionAttributeGrpc.GetTestInstructionAttributeEnabled(),
					TestInstructionAttributeMandatory:             testInstructionAttributeGrpc.GetTestInstructionAttributeMandatory(),
					TestInstructionAttributeVisibleInTestCaseArea: testInstructionAttributeGrpc.GetTestInstructionAttributeVisibleInTestCaseArea(),
					TestInstructionAttributeIsDeprecated:          testInstructionAttributeGrpc.GetTestInstructionAttributeIsDeprecated(),
					TestInstructionAttributeInputMask:             TypeAndStructs.TestInstructionAttributeInputMaskType(testInstructionAttributeGrpc.GetTestInstructionAttributeInputMask()),
					TestInstructionAttributeType:                  TypeAndStructs.TestInstructionAttributeTypeType(testInstructionAttributeGrpc.GetTestInstructionAttributeType()),
				}

				// Append to slice of messages
				testInstructionAttributes = append(testInstructionAttributes, testInstructionAttributeMessage)
			}

			// Loop and Create 'ImmatureElementModel' from 'testInstructionInstanceVersionMessageGrpc'
			var immatureElementModelMessages []*TypeAndStructs.ImmatureElementModelMessageStruct
			for _, immatureElementModelGrpc := range testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.ImmatureElementModel {

				// Create the gRPC-version of the message
				var immatureElementModelMessage *TypeAndStructs.ImmatureElementModelMessageStruct
				immatureElementModelMessage = &TypeAndStructs.ImmatureElementModelMessageStruct{
					DomainUUID:               TypeAndStructs.DomainUUIDType(immatureElementModelGrpc.GetDomainUUID()),
					DomainName:               TypeAndStructs.DomainNameType(immatureElementModelGrpc.GetDomainName()),
					ImmatureElementUUID:      TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetImmatureElementUUID()),
					ImmatureElementName:      TypeAndStructs.OriginalElementNameType(immatureElementModelGrpc.GetImmatureElementName()),
					PreviousElementUUID:      TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetPreviousElementUUID()),
					NextElementUUID:          TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetNextElementUUID()),
					FirstChildElementUUID:    TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetFirstChildElementUUID()),
					ParentElementUUID:        TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetParentElementUUID()),
					TestCaseModelElementType: TypeAndStructs.TestCaseModelElementTypeType(immatureElementModelGrpc.GetTestCaseModelElementType()),
					OriginalElementUUID:      TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetOriginalElementUUID()),
					TopImmatureElementUUID:   TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetTopImmatureElementUUID()),
					IsTopElement:             immatureElementModelGrpc.GetIsTopElement(),
				}

				// Append to slice of messages
				immatureElementModelMessages = append(immatureElementModelMessages, immatureElementModelMessage)
			}

			// Create ResponseVariablesMapStructure to be converted from a gRPC-version
			var responseVariablesMapStructureMessage *TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct
			responseVariablesMapStructureMessage = &TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct{
				ResponseVariablesMap:     nil,
				ResponseVariablesMapHash: testInstructionInstanceVersionMessageGrpc.ResponseVariablesMapStructure.GetResponseVariablesMapHash(),
			}

			// Create 'ResponseVariablesMap'  to be converted from a gRPC-version
			var responseVariablesMap map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct
			responseVariablesMap = make(map[TypeAndStructs.ResponseVariableUuidType]*TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct)

			for responseVariableUuid, responseVariable := range testInstructionInstanceVersionMessageGrpc.
				ResponseVariablesMapStructure.ResponseVariablesMap {

				var tempResponseVariable *TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct
				tempResponseVariable = &TestInstructionAndTestInstuctionContainerTypes.ResponseVariableStructureStruct{
					ResponseVariable: TypeAndStructs.ResponseVariableStruct{
						ResponseVariableUuid:        TypeAndStructs.ResponseVariableUuidType(responseVariable.GetResponseVariable().GetResponseVariableUuid()),
						ResponseVariableName:        TypeAndStructs.ResponseVariableNameType(responseVariable.GetResponseVariable().GetResponseVariableName()),
						ResponseVariableDescription: TypeAndStructs.ResponseVariableDescriptionType(responseVariable.GetResponseVariable().GetResponseVariableDescription()),
						ResponseVariableIsMandatory: TypeAndStructs.ResponseVariableIsMandatoryType(responseVariable.GetResponseVariable().GetResponseVariableIsMandatory()),
					},
					ResponseVariableHash: responseVariable.ResponseVariableHash,
				}

				responseVariablesMap[TypeAndStructs.ResponseVariableUuidType(responseVariableUuid)] = tempResponseVariable
			}

			responseVariablesMapStructureMessage.ResponseVariablesMap = responseVariablesMap

			// Create TestInstructionInstance to be converted from a gRPC-version
			var testInstructionInstanceVersionMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct
			testInstructionInstanceVersionMessage = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionStruct{
				TestInstructionInstance: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionStruct{
					TestInstruction: &TypeAndStructs.TestInstructionStruct{
						DomainUUID:                   TypeAndStructs.DomainUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetDomainUUID()),
						DomainName:                   TypeAndStructs.DomainNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetDomainName()),
						ExecutionDomainUUID:          TypeAndStructs.DomainUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetExecutionDomainUuid()),
						ExecutionDomainName:          TypeAndStructs.DomainNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetExecutionDomainName()),
						TestInstructionUUID:          TypeAndStructs.OriginalElementUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetTestInstructionUUID()),
						TestInstructionName:          TypeAndStructs.TestInstructionNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetTestInstructionName()),
						TestInstructionTypeUUID:      TypeAndStructs.TestInstructionTypeUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetTestInstructionTypeUUID()),
						TestInstructionTypeName:      TypeAndStructs.TestInstructionTypeNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetTestInstructionTypeName()),
						TestInstructionDescription:   testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetTestInstructionDescription(),
						TestInstructionMouseOverText: testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetTestInstructionMouseOverText(),
						Deprecated:                   testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetDeprecated(),
						Enabled:                      testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetEnabled(),
						MajorVersionNumber:           int(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetMajorVersionNumber()),
						MinorVersionNumber:           int(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetMinorVersionNumber()),
						UpdatedTimeStamp:             TypeAndStructs.UpdatedTimeStampType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.TestInstruction.GetUpdatedTimeStamp()),
					},
					BasicTestInstructionInformation: &TypeAndStructs.BasicTestInstructionInformationStruct{
						DomainUUID:                   TypeAndStructs.DomainUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetDomainUUID()),
						DomainName:                   TypeAndStructs.DomainNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetDomainName()),
						ExecutionDomainUUID:          TypeAndStructs.DomainUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetExecutionDomainUuid()),
						ExecutionDomainName:          TypeAndStructs.DomainNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetExecutionDomainName()),
						TestInstructionUUID:          TypeAndStructs.OriginalElementUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTestInstructionUUID()),
						TestInstructionName:          TypeAndStructs.TestInstructionNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTestInstructionName()),
						TestInstructionTypeUUID:      TypeAndStructs.TestInstructionTypeUUIDType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTestInstructionTypeUUID()),
						TestInstructionTypeName:      TypeAndStructs.TestInstructionTypeNameType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTestInstructionTypeName()),
						Deprecated:                   testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetDeprecated(),
						MajorVersionNumber:           int(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetMajorVersionNumber()),
						MinorVersionNumber:           int(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetMinorVersionNumber()),
						UpdatedTimeStamp:             TypeAndStructs.UpdatedTimeStampType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetUpdatedTimeStamp()),
						TestInstructionColor:         TypeAndStructs.ColorType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTestInstructionColor()),
						TCRuleDeletion:               TypeAndStructs.TCRuleDeletionType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTCRuleDeletion()),
						TCRuleSwap:                   TypeAndStructs.TCRuleSwapType(testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTCRuleSwap()),
						TestInstructionDescription:   testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTestInstructionDescription(),
						TestInstructionMouseOverText: testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetTestInstructionMouseOverText(),
						Enabled:                      testInstructionInstanceVersionMessageGrpc.TestInstructionInstance.BasicTestInstructionInformation.GetEnabled(),
					},
					ImmatureTestInstructionInformation: immatureTestInstructionInformationMessages,
					TestInstructionAttribute:           testInstructionAttributes,
					ImmatureElementModel:               immatureElementModelMessages,
					LocalExecutionMethods:              TestInstructionAndTestInstuctionContainerTypes.AnyType{Value: nil}, // Don't include when calculating Hash
				},
				TestInstructionInstanceMajorVersion: int(testInstructionInstanceVersionMessageGrpc.GetTestInstructionInstanceMajorVersion()),
				TestInstructionInstanceMinorVersion: int(testInstructionInstanceVersionMessageGrpc.GetTestInstructionInstanceMinorVersion()),
				Deprecated:                          testInstructionInstanceVersionMessageGrpc.GetDeprecated(),
				Enabled:                             testInstructionInstanceVersionMessageGrpc.GetEnabled(),
				TestInstructionInstanceVersionHash:  testInstructionInstanceVersionMessageGrpc.GetTestInstructionInstanceVersionHash(),
				ResponseVariablesMapStructure:       responseVariablesMapStructureMessage,
				TestInstructionInstanceVersionAndResponseVariablesHash: testInstructionInstanceVersionMessageGrpc.GetTestInstructionInstanceVersionAndResponseVariablesHash(),
			}

			testInstructionInstanceVersionMessages = append(testInstructionInstanceVersionMessages, testInstructionInstanceVersionMessage)
		}

		// Create the object to be saved in Map
		var testInstructionInstanceVersions *TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct
		testInstructionInstanceVersions = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionInstanceVersionsStruct{
			TestInstructionVersions:     testInstructionInstanceVersionMessages,
			TestInstructionVersionsHash: testInstructionInstanceVersionsMessageGrpc.GetTestInstructionVersionsHash(),
		}

		// Add to Map
		testInstructionInstanceVersionsMessageMap[TypeAndStructs.OriginalElementUUIDType(originalElementUUIDTypeGrpc)] = testInstructionInstanceVersions

	}

	// Convert TestInstructionContainers from gRPC-message-version
	// Generate original-version of TestInstructionContainerInstanceVersionsMessageMap
	var testInstructionContainerInstanceVersionsMessageMap map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct
	testInstructionContainerInstanceVersionsMessageMap = make(map[TypeAndStructs.OriginalElementUUIDType]*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct)

	// Loop all TestInstructionContainerInstanceVersionsMessages
	for originalElementUUIDTypeGrpc, testInstructionContainerInstanceVersionsMessageGrpc := range supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.TestInstructionContainers.TestInstructionsMap {

		// Loop TestInstructionContainerVersions and create original-version
		var testInstructionContainerInstanceVersionMessages []*TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct
		for _, testInstructionContainerInstanceVersionMessageGrpc := range testInstructionContainerInstanceVersionsMessageGrpc.TestInstructionContainerVersions {

			// Loop and Create 'ImmatureTestInstructionContainersInformation' from 'testInstructionContainerInstanceVersionMessageGrpc'
			var immatureTestInstructionContainerInformationMessages []*TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
			for _, immatureTestInstructionContainerInformationMessageGrpc := range testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.ImmatureTestInstructionContainerInformations {

				// Create the gRPC-version of the message
				var immatureTestInstructionContainerInformationMessage *TypeAndStructs.ImmatureTestInstructionContainerMessageStruct
				immatureTestInstructionContainerInformationMessage = &TypeAndStructs.ImmatureTestInstructionContainerMessageStruct{
					DomainUUID:                   TypeAndStructs.DomainUUIDType(immatureTestInstructionContainerInformationMessageGrpc.GetDomainUUID()),
					DomainName:                   TypeAndStructs.DomainNameType(immatureTestInstructionContainerInformationMessageGrpc.GetDomainName()),
					TestInstructionContainerUUID: TypeAndStructs.OriginalElementUUIDType(immatureTestInstructionContainerInformationMessageGrpc.GetTestInstructionContainerUUID()),
					TestInstructionContainerName: TypeAndStructs.TestInstructionContainerNameType(immatureTestInstructionContainerInformationMessageGrpc.GetTestInstructionContainerName()),
					DropZoneUUID:                 TypeAndStructs.DropZoneUUIDType(immatureTestInstructionContainerInformationMessageGrpc.GetDropZoneUUID()),
					DropZoneName:                 TypeAndStructs.DropZoneNameType(immatureTestInstructionContainerInformationMessageGrpc.GetDropZoneName()),
					DropZoneDescription:          immatureTestInstructionContainerInformationMessageGrpc.GetDropZoneDescription(),
					DropZoneMouseOver:            immatureTestInstructionContainerInformationMessageGrpc.GetDropZoneMouseOver(),
					DropZoneColor:                TypeAndStructs.ColorType(immatureTestInstructionContainerInformationMessageGrpc.GetDropZoneColor()),
					TestInstructionAttributeType: TypeAndStructs.TestInstructionAttributeTypeType(immatureTestInstructionContainerInformationMessageGrpc.GetTestInstructionContainerAttributeType()),
					TestInstructionAttributeUUID: TypeAndStructs.TestInstructionAttributeUUIDType(immatureTestInstructionContainerInformationMessageGrpc.GetTestInstructionContainerAttributeUUID()),
					TestInstructionAttributeName: TypeAndStructs.TestInstructionAttributeNameType(immatureTestInstructionContainerInformationMessageGrpc.GetTestInstructionContainerAttributeName()),
					AttributeValueAsString:       TypeAndStructs.AttributeValueAsStringType(immatureTestInstructionContainerInformationMessageGrpc.GetAttributeValueAsString()),
					AttributeValueUUID:           TypeAndStructs.AttributeValueUUIDType(immatureTestInstructionContainerInformationMessageGrpc.GetAttributeValueUUID()),
					FirstImmatureElementUUID:     TypeAndStructs.OriginalElementUUIDType(immatureTestInstructionContainerInformationMessageGrpc.GetFirstImmatureElementUUID()),
					AttributeActionCommand:       TypeAndStructs.AttributeActionCommandType(immatureTestInstructionContainerInformationMessageGrpc.GetAttributeActionCommand()),
				}

				// Append to slice of messages
				immatureTestInstructionContainerInformationMessages = append(immatureTestInstructionContainerInformationMessages, immatureTestInstructionContainerInformationMessage)
			}

			// Loop and Create 'ImmatureElementModel' from 'testInstructionContainerInstanceVersionMessageGrpc'
			var immatureElementModelMessages []*TypeAndStructs.ImmatureElementModelMessageStruct
			for _, immatureElementModelGrpc := range testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.ImmatureElementModel {

				// Create the gRPC-version of the message
				var immatureElementModelMessage *TypeAndStructs.ImmatureElementModelMessageStruct
				immatureElementModelMessage = &TypeAndStructs.ImmatureElementModelMessageStruct{
					DomainUUID:               TypeAndStructs.DomainUUIDType(immatureElementModelGrpc.GetDomainUUID()),
					DomainName:               TypeAndStructs.DomainNameType(immatureElementModelGrpc.GetDomainName()),
					ImmatureElementUUID:      TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetImmatureElementUUID()),
					ImmatureElementName:      TypeAndStructs.OriginalElementNameType(immatureElementModelGrpc.GetImmatureElementName()),
					PreviousElementUUID:      TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetPreviousElementUUID()),
					NextElementUUID:          TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetNextElementUUID()),
					FirstChildElementUUID:    TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetFirstChildElementUUID()),
					ParentElementUUID:        TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetParentElementUUID()),
					TestCaseModelElementType: TypeAndStructs.TestCaseModelElementTypeType(immatureElementModelGrpc.GetTestCaseModelElementType()),
					OriginalElementUUID:      TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetOriginalElementUUID()),
					TopImmatureElementUUID:   TypeAndStructs.OriginalElementUUIDType(immatureElementModelGrpc.GetTopImmatureElementUUID()),
					IsTopElement:             immatureElementModelGrpc.GetIsTopElement(),
				}

				// Append to slice of messages
				immatureElementModelMessages = append(immatureElementModelMessages, immatureElementModelMessage)
			}

			// Create TestInstructionContainerInstance to be converted from a gRPC-version
			var testInstructionContainerInstanceVersionMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct
			testInstructionContainerInstanceVersionMessage = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionStruct{
				TestInstructionContainerInstance: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerStruct{
					TestInstructionContainer: &TypeAndStructs.TestInstructionContainerStruct{
						DomainUUID:                            TypeAndStructs.DomainUUIDType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetDomainUUID()),
						DomainName:                            TypeAndStructs.DomainNameType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetDomainName()),
						TestInstructionContainerUUID:          TypeAndStructs.OriginalElementUUIDType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetTestInstructionContainerUUID()),
						TestInstructionContainerName:          TypeAndStructs.TestInstructionContainerNameType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetTestInstructionContainerName()),
						TestInstructionContainerTypeUUID:      TypeAndStructs.TestInstructionContainerTypeUUIDType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetTestInstructionContainerTypeUUID()),
						TestInstructionContainerTypeName:      TypeAndStructs.TestInstructionContainerTypeNameType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetTestInstructionContainerTypeName()),
						TestInstructionContainerDescription:   testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetTestInstructionContainerDescription(),
						TestInstructionContainerMouseOverText: testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetTestInstructionContainerMouseOverText(),
						Deprecated:                            testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetDeprecated(),
						Enabled:                               testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetEnabled(),
						MajorVersionNumber:                    int(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetMajorVersionNumber()),
						MinorVersionNumber:                    int(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetMinorVersionNumber()),
						UpdatedTimeStamp:                      TypeAndStructs.UpdatedTimeStampType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetUpdatedTimeStamp()),
						ChildrenIsParallelProcessed:           testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.TestInstructionContainer.GetChildrenIsParallelProcessed(),
					},
					BasicTestInstructionContainerInformation: &TypeAndStructs.BasicTestInstructionContainerInformationStruct{
						DomainUUID:                            TypeAndStructs.DomainUUIDType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetDomainUUID()),
						DomainName:                            TypeAndStructs.DomainNameType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetDomainName()),
						TestInstructionContainerUUID:          TypeAndStructs.OriginalElementUUIDType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerUUID()),
						TestInstructionContainerName:          TypeAndStructs.TestInstructionContainerNameType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerName()),
						TestInstructionContainerTypeUUID:      TypeAndStructs.TestInstructionContainerTypeUUIDType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerTypeUUID()),
						TestInstructionContainerTypeName:      TypeAndStructs.TestInstructionContainerTypeNameType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerTypeName()),
						Deprecated:                            testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetDeprecated(),
						MajorVersionNumber:                    int(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetMajorVersionNumber()),
						MinorVersionNumber:                    int(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetMinorVersionNumber()),
						UpdatedTimeStamp:                      TypeAndStructs.UpdatedTimeStampType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetUpdatedTimeStamp()),
						TestInstructionContainerColor:         TypeAndStructs.ColorType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerColor()),
						TCRuleDeletion:                        TypeAndStructs.TCRuleDeletionType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTCRuleDeletion()),
						TCRuleSwap:                            TypeAndStructs.TCRuleSwapType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTCRuleSwap()),
						TestInstructionContainerDescription:   testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerDescription(),
						TestInstructionContainerMouseOverText: testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerMouseOverText(),
						Enabled:                               testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetEnabled(),
						TestInstructionContainerExecutionType: TypeAndStructs.TestInstructionContainerExecutionTypeType(testInstructionContainerInstanceVersionMessageGrpc.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.GetTestInstructionContainerExecutionType()),
					},
					ImmatureTestInstructionContainer: immatureTestInstructionContainerInformationMessages,
					ImmatureElementModel:             immatureElementModelMessages,
				},
				TestInstructionContainerInstanceMajorVersion: int(testInstructionContainerInstanceVersionMessageGrpc.GetTestInstructionContainerInstanceMajorVersion()),
				TestInstructionContainerInstanceMinorVersion: int(testInstructionContainerInstanceVersionMessageGrpc.GetTestInstructionContainerInstanceMinorVersion()),
				Deprecated:                           testInstructionContainerInstanceVersionMessageGrpc.GetDeprecated(),
				Enabled:                              testInstructionContainerInstanceVersionMessageGrpc.GetEnabled(),
				TestInstructionContainerInstanceHash: testInstructionContainerInstanceVersionMessageGrpc.GetTestInstructionContainerInstanceVersionHash(),
			}

			testInstructionContainerInstanceVersionMessages = append(testInstructionContainerInstanceVersionMessages, testInstructionContainerInstanceVersionMessage)
		}

		// Create the object to be saved in Map
		var testInstructionContainerInstanceVersions *TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct
		testInstructionContainerInstanceVersions = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainerInstanceVersionsStruct{
			TestInstructionContainerVersions:     testInstructionContainerInstanceVersionMessages,
			TestInstructionContainerVersionsHash: testInstructionContainerInstanceVersionsMessageGrpc.GetTestInstructionVersionsHash(),
		}

		// Add to Map
		testInstructionContainerInstanceVersionsMessageMap[TypeAndStructs.OriginalElementUUIDType(originalElementUUIDTypeGrpc)] = testInstructionContainerInstanceVersions

	}

	// Convert Allowed Users from gRPC-message-version
	// Loop all TestInstructionContainerInstanceVersionsMessages
	var allowedUsers []*TestInstructionAndTestInstuctionContainerTypes.AllowedUserStruct
	for _, allowedUserGrPc := range supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.AllowedUsers.AllowedUsers {

		var allowedUser *TestInstructionAndTestInstuctionContainerTypes.AllowedUserStruct
		allowedUser = &TestInstructionAndTestInstuctionContainerTypes.AllowedUserStruct{
			UserIdOnComputer:     allowedUserGrPc.UserIdOnComputer,
			GCPAuthenticatedUser: allowedUserGrPc.GCPAuthenticatedUser,
			UserEmail:            allowedUserGrPc.UserEmail,
			UserFirstName:        allowedUserGrPc.UserFirstName,
			UserLastName:         allowedUserGrPc.UserLastName,
			UserAuthorizationRights: &TestInstructionAndTestInstuctionContainerTypes.UserAuthorizationRightsStruct{
				CanListAndViewTestCaseOwnedByThisDomain:                    allowedUserGrPc.UserAuthorizationRights.CanListAndViewTestCaseOwnedByThisDomain,
				CanBuildAndSaveTestCaseOwnedByThisDomain:                   allowedUserGrPc.UserAuthorizationRights.CanBuildAndSaveTestCaseOwnedByThisDomain,
				CanListAndViewTestCaseHavingTIandTICFromThisDomain:         allowedUserGrPc.UserAuthorizationRights.CanListAndViewTestCaseHavingTIandTICFromThisDomain,
				CanListAndViewTestCaseHavingTIandTICFromThisDomainExtended: allowedUserGrPc.UserAuthorizationRights.CanListAndViewTestCaseHavingTIandTICFromThisDomainExtended,
				CanBuildAndSaveTestCaseHavingTIandTICFromThisDomain:        allowedUserGrPc.UserAuthorizationRights.CanBuildAndSaveTestCaseHavingTIandTICFromThisDomain,
			},
		}

		// Append to slice of messages
		allowedUsers = append(allowedUsers, allowedUser)
	}

	// Create the full message for all supported TestInstructions, TestInstructionContainers and Allowed Users from the gRPC Message
	testInstructionsAndTestInstructionContainersMessage = &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct{
		TestInstructions: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionsStruct{
			TestInstructionsMap:  testInstructionInstanceVersionsMessageMap,
			TestInstructionsHash: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetTestInstructions().GetTestInstructionsHash(),
		},
		TestInstructionContainers: &TestInstructionAndTestInstuctionContainerTypes.TestInstructionContainersStruct{
			TestInstructionContainersMap:  testInstructionContainerInstanceVersionsMessageMap,
			TestInstructionContainersHash: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetTestInstructionContainers().TestInstructionContainersHash,
		},
		AllowedUsers: &TestInstructionAndTestInstuctionContainerTypes.AllowedUsersStruct{
			AllowedUsers: allowedUsers,
			AllUsersAuthorizationRights: &TestInstructionAndTestInstuctionContainerTypes.AllUsersAuthorizationRightsStruct{
				AllUsersCanListAndViewTestCaseHavingTIandTICFromThisDomain: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetAllowedUsers().
					AllUsersAuthorizationRights.GetAllUsersCanListAndViewTestCaseHavingTIandTICFromThisDomain(),
				AllUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomain: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetAllowedUsers().
					AllUsersAuthorizationRights.GetAllUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomain(),
			},
			AllowedUsersHash: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetAllowedUsers().GetAllowedUsersHash(),
		},
		MessageCreationTimeStamp: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetMessageCreationTimeStamp().AsTime(),
		TestInstructionsAndTestInstructionsContainersAndUsersMessageHash: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetTestInstructionsAndTestInstructionsContainersMessageHash(),
		ForceNewBaseLineForTestInstructionsAndTestInstructionContainers:  supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetForceNewBaseLineForTestInstructionsAndTestInstructionContainers(),
		ConnectorsDomain: TestInstructionAndTestInstuctionContainerTypes.ConnectorsDomainStruct{
			ConnectorsDomainUUID: TypeAndStructs.DomainUUIDType(supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetConnectorDomain().ConnectorsDomainUUID),
			ConnectorsDomainName: TypeAndStructs.DomainNameType(supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetConnectorDomain().ConnectorsDomainName),
			ConnectorsDomainHash: supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage.GetConnectorDomain().ConnectorsDomainHash,
		},
	}

	return testInstructionsAndTestInstructionContainersMessage, err
}
