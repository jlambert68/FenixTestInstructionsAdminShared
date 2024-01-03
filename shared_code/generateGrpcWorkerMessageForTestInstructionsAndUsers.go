package shared_code

import (
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GenerateTestInstructionAndTestInstructionContainerAndUserGrpcWorkerMessage
// Generates the gRPC message to be sent to Worker for Supported  TestInstructions, TestInstructionContainers and Allowed Users
func GenerateTestInstructionAndTestInstructionContainerAndUserGrpcWorkerMessage(
	domainUuid string,
	testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) (
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage *fenixExecutionWorkerGrpcApi.SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage,
	err error) {

	// Convert TestInstructions into gRPC-message-version
	// Generate gRPC-version of TestInstructionInstanceVersionsMessageMap
	var testInstructionInstanceVersionsMessageMap map[string]*fenixExecutionWorkerGrpcApi.SupportedTestInstructionInstanceVersionsMessage
	testInstructionInstanceVersionsMessageMap = make(map[string]*fenixExecutionWorkerGrpcApi.SupportedTestInstructionInstanceVersionsMessage)

	// Loop all TestInstructionInstanceVersionsMessages
	for originalElementUUIDType, testInstructionInstanceVersionsMessage := range testInstructionsAndTestInstructionContainersMessage.TestInstructions.TestInstructionsMap {

		// Loop TestInstructionVersions and create gRPC-version
		var testInstructionInstanceVersionMessagesGrpc []*fenixExecutionWorkerGrpcApi.SupportedTestInstructionInstanceVersionMessage
		for _, testInstructionInstanceVersionMessage := range testInstructionInstanceVersionsMessage.TestInstructionVersions {

			// Loop and Create 'ImmatureTestInstructionInformations' for 'testInstructionInstanceVersionMessageGrpc'
			var immatureTestInstructionInformationMessagesGrpc []*fenixExecutionWorkerGrpcApi.SupportedImmatureTestInstructionInformationMessage
			for _, immatureTestInstructionInformationMessage := range testInstructionInstanceVersionMessage.TestInstructionInstance.ImmatureTestInstructionInformation {

				// Create the gRPC-version of the message
				var immatureTestInstructionInformationMessageGrpc *fenixExecutionWorkerGrpcApi.SupportedImmatureTestInstructionInformationMessage
				immatureTestInstructionInformationMessageGrpc = &fenixExecutionWorkerGrpcApi.SupportedImmatureTestInstructionInformationMessage{
					DomainUUID:                   string(immatureTestInstructionInformationMessage.DomainUUID),
					DomainName:                   string(immatureTestInstructionInformationMessage.DomainName),
					TestInstructionUUID:          string(immatureTestInstructionInformationMessage.TestInstructionUUID),
					TestInstructionName:          string(immatureTestInstructionInformationMessage.TestInstructionName),
					DropZoneUUID:                 string(immatureTestInstructionInformationMessage.DropZoneUUID),
					DropZoneName:                 string(immatureTestInstructionInformationMessage.DropZoneName),
					DropZoneDescription:          immatureTestInstructionInformationMessage.DropZoneDescription,
					DropZoneMouseOver:            immatureTestInstructionInformationMessage.DropZoneMouseOver,
					DropZoneColor:                string(immatureTestInstructionInformationMessage.DropZoneColor),
					TestInstructionAttributeType: string(immatureTestInstructionInformationMessage.TestInstructionAttributeType),
					TestInstructionAttributeUUID: string(immatureTestInstructionInformationMessage.TestInstructionAttributeUUID),
					TestInstructionAttributeName: string(immatureTestInstructionInformationMessage.TestInstructionAttributeName),
					AttributeValueAsString:       string(immatureTestInstructionInformationMessage.AttributeValueAsString),
					AttributeValueUUID:           string(immatureTestInstructionInformationMessage.AttributeValueUUID),
					FirstImmatureElementUUID:     string(immatureTestInstructionInformationMessage.FirstImmatureElementUUID),
					AttributeActionCommand:       int32(immatureTestInstructionInformationMessage.AttributeActionCommand),
				}

				// Append to slice of messages
				immatureTestInstructionInformationMessagesGrpc = append(immatureTestInstructionInformationMessagesGrpc, immatureTestInstructionInformationMessageGrpc)
			}

			// Loop and Create 'TestInstructionAttributes' for 'testInstructionInstanceVersionMessageGrpc'
			var testInstructionAttributesGrpc []*fenixExecutionWorkerGrpcApi.SupportedTestInstructionAttributeMessage
			for _, testInstructionAttribute := range testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstructionAttribute {

				// Create the gRPC-version of the message
				var testInstructionAttributeMessageGrpc *fenixExecutionWorkerGrpcApi.SupportedTestInstructionAttributeMessage
				testInstructionAttributeMessageGrpc = &fenixExecutionWorkerGrpcApi.SupportedTestInstructionAttributeMessage{
					DomainUUID:                                    string(testInstructionAttribute.DomainUUID),
					DomainName:                                    string(testInstructionAttribute.DomainName),
					TestInstructionUUID:                           string(testInstructionAttribute.TestInstructionUUID),
					TestInstructionName:                           string(testInstructionAttribute.TestInstructionName),
					TestInstructionAttributeUUID:                  string(testInstructionAttribute.TestInstructionAttributeUUID),
					TestInstructionAttributeName:                  string(testInstructionAttribute.TestInstructionAttributeName),
					TestInstructionAttributeDescription:           testInstructionAttribute.TestInstructionAttributeDescription,
					TestInstructionAttributeMouseOver:             testInstructionAttribute.TestInstructionAttributeMouseOver,
					TestInstructionAttributeTypeUUID:              string(testInstructionAttribute.TestInstructionAttributeTypeUUID),
					TestInstructionAttributeTypeName:              string(testInstructionAttribute.TestInstructionAttributeTypeName),
					TestInstructionAttributeValueAsString:         string(testInstructionAttribute.TestInstructionAttributeValueAsString),
					TestInstructionAttributeValueUUID:             string(testInstructionAttribute.TestInstructionAttributeValueUUID),
					TestInstructionAttributeVisible:               testInstructionAttribute.TestInstructionAttributeVisible,
					TestInstructionAttributeEnabled:               testInstructionAttribute.TestInstructionAttributeEnabled,
					TestInstructionAttributeMandatory:             testInstructionAttribute.TestInstructionAttributeMandatory,
					TestInstructionAttributeVisibleInTestCaseArea: testInstructionAttribute.TestInstructionAttributeVisibleInTestCaseArea,
					TestInstructionAttributeIsDeprecated:          testInstructionAttribute.TestInstructionAttributeIsDeprecated,
					TestInstructionAttributeInputMask:             string(testInstructionAttribute.TestInstructionAttributeInputMask),
					TestInstructionAttributeType:                  string(testInstructionAttribute.TestInstructionAttributeType),
				}

				// Append to slice of messages
				testInstructionAttributesGrpc = append(testInstructionAttributesGrpc, testInstructionAttributeMessageGrpc)
			}

			// Loop and Create 'ImmatureElementModel' for 'testInstructionInstanceVersionMessageGrpc'
			var immatureElementModelMessagesGrpc []*fenixExecutionWorkerGrpcApi.SupportedImmatureElementModelMessage
			for _, immatureElementModel := range testInstructionInstanceVersionMessage.TestInstructionInstance.ImmatureElementModel {

				// Create the gRPC-version of the message
				var immatureElementModelMessageGrpc *fenixExecutionWorkerGrpcApi.SupportedImmatureElementModelMessage
				immatureElementModelMessageGrpc = &fenixExecutionWorkerGrpcApi.SupportedImmatureElementModelMessage{
					DomainUUID:               string(immatureElementModel.DomainUUID),
					DomainName:               string(immatureElementModel.DomainName),
					ImmatureElementUUID:      string(immatureElementModel.ImmatureElementUUID),
					ImmatureElementName:      string(immatureElementModel.ImmatureElementName),
					PreviousElementUUID:      string(immatureElementModel.PreviousElementUUID),
					NextElementUUID:          string(immatureElementModel.NextElementUUID),
					FirstChildElementUUID:    string(immatureElementModel.FirstChildElementUUID),
					ParentElementUUID:        string(immatureElementModel.ParentElementUUID),
					TestCaseModelElementType: string(immatureElementModel.TestCaseModelElementType),
					OriginalElementUUID:      string(immatureElementModel.OriginalElementUUID),
					TopImmatureElementUUID:   string(immatureElementModel.TopImmatureElementUUID),
					IsTopElement:             immatureElementModel.IsTopElement,
				}

				// Append to slice of messages
				immatureElementModelMessagesGrpc = append(immatureElementModelMessagesGrpc, immatureElementModelMessageGrpc)
			}

			// Create TestInstructionInstance to be converted into a gRPC-version
			var testInstructionInstanceVersionMessageGrpc *fenixExecutionWorkerGrpcApi.SupportedTestInstructionInstanceVersionMessage
			testInstructionInstanceVersionMessageGrpc = &fenixExecutionWorkerGrpcApi.SupportedTestInstructionInstanceVersionMessage{
				TestInstructionInstance: &fenixExecutionWorkerGrpcApi.SupportedTestInstructionMessage{
					TestInstruction: &fenixExecutionWorkerGrpcApi.SupportedTestInstructionBaseMessage{
						DomainUUID:                   string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.DomainUUID),
						DomainName:                   string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.DomainName),
						TestInstructionUUID:          string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.TestInstructionUUID),
						TestInstructionName:          string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.TestInstructionName),
						TestInstructionTypeUUID:      string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.TestInstructionTypeUUID),
						TestInstructionTypeName:      string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.TestInstructionTypeName),
						TestInstructionDescription:   string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.TestInstructionDescription),
						TestInstructionMouseOverText: string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.TestInstructionMouseOverText),
						Deprecated:                   testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.Deprecated,
						Enabled:                      testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.Enabled,
						MajorVersionNumber:           int32(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.MajorVersionNumber),
						MinorVersionNumber:           int32(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.MinorVersionNumber),
						UpdatedTimeStamp:             string(testInstructionInstanceVersionMessage.TestInstructionInstance.TestInstruction.UpdatedTimeStamp),
					},
					BasicTestInstructionInformation: &fenixExecutionWorkerGrpcApi.SupportedBasicTestInstructionInformationMessage{
						DomainUUID:                   string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.DomainUUID),
						DomainName:                   string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.DomainName),
						TestInstructionUUID:          string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TestInstructionUUID),
						TestInstructionName:          string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TestInstructionName),
						TestInstructionTypeUUID:      string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TestInstructionTypeUUID),
						TestInstructionTypeName:      string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TestInstructionTypeName),
						Deprecated:                   testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.Deprecated,
						MajorVersionNumber:           int32(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.MajorVersionNumber),
						MinorVersionNumber:           int32(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.MinorVersionNumber),
						UpdatedTimeStamp:             string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.UpdatedTimeStamp),
						TestInstructionColor:         string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TestInstructionColor),
						TCRuleDeletion:               string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TCRuleDeletion),
						TCRuleSwap:                   string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TCRuleSwap),
						TestInstructionDescription:   string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TestInstructionDescription),
						TestInstructionMouseOverText: string(testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.TestInstructionMouseOverText),
						Enabled:                      testInstructionInstanceVersionMessage.TestInstructionInstance.BasicTestInstructionInformation.Enabled,
					},
					ImmatureTestInstructionInformations: immatureTestInstructionInformationMessagesGrpc,
					TestInstructionAttributes:           testInstructionAttributesGrpc,
					ImmatureElementModel:                immatureElementModelMessagesGrpc,
				},
				TestInstructionInstanceMajorVersion: int32(testInstructionInstanceVersionMessage.TestInstructionInstanceMajorVersion),
				TestInstructionInstanceMinorVersion: int32(testInstructionInstanceVersionMessage.TestInstructionInstanceMinorVersion),
				Deprecated:                          testInstructionInstanceVersionMessage.Deprecated,
				Enabled:                             testInstructionInstanceVersionMessage.Enabled,
				TestInstructionInstanceVersionHash:  testInstructionInstanceVersionMessage.TestInstructionInstanceVersionHash,
			}

			testInstructionInstanceVersionMessagesGrpc = append(testInstructionInstanceVersionMessagesGrpc, testInstructionInstanceVersionMessageGrpc)
		}

		// Create the object to be saved in Map
		var testInstructionInstanceVersionsGrpc *fenixExecutionWorkerGrpcApi.SupportedTestInstructionInstanceVersionsMessage
		testInstructionInstanceVersionsGrpc = &fenixExecutionWorkerGrpcApi.SupportedTestInstructionInstanceVersionsMessage{
			TestInstructionVersions:     testInstructionInstanceVersionMessagesGrpc,
			TestInstructionVersionsHash: testInstructionInstanceVersionsMessage.TestInstructionVersionsHash,
		}

		// Add to Map
		testInstructionInstanceVersionsMessageMap[string(originalElementUUIDType)] = testInstructionInstanceVersionsGrpc

	}

	// Convert TestInstructionContainers into gRPC-message-version
	// Generate gRPC-version of TestInstructionContainerInstanceVersionsMessageMap
	var testInstructionContainerInstanceVersionsMessageMap map[string]*fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerInstanceVersionsMessage
	testInstructionContainerInstanceVersionsMessageMap = make(map[string]*fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerInstanceVersionsMessage)

	// Loop all TestInstructionContainerInstanceVersionsMessages
	for originalElementUUIDType, testInstructionContainerInstanceVersionsMessage := range testInstructionsAndTestInstructionContainersMessage.TestInstructionContainers.TestInstructionContainersMap {

		// Loop TestInstructionContainerVersions and create gRPC-version
		var testInstructionContainerInstanceVersionMessagesGrpc []*fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerInstanceVersionMessage
		for _, testInstructionContainerInstanceVersionMessage := range testInstructionContainerInstanceVersionsMessage.TestInstructionContainerVersions {

			// Loop and Create 'ImmatureTestInstructionContainerInformations' for 'testInstructionContainerInstanceVersionMessageGrpc'
			var immatureTestInstructionContainerInformationMessagesGrpc []*fenixExecutionWorkerGrpcApi.SupportedImmatureTestInstructionContainerInformationMessage
			for _, immatureTestInstructionContainerInformationMessage := range testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.ImmatureTestInstructionContainer {

				// Create the gRPC-version of the message
				var immatureTestInstructionContainerInformationMessageGrpc *fenixExecutionWorkerGrpcApi.SupportedImmatureTestInstructionContainerInformationMessage
				immatureTestInstructionContainerInformationMessageGrpc = &fenixExecutionWorkerGrpcApi.SupportedImmatureTestInstructionContainerInformationMessage{
					DomainUUID:                            string(immatureTestInstructionContainerInformationMessage.DomainUUID),
					DomainName:                            string(immatureTestInstructionContainerInformationMessage.DomainName),
					TestInstructionContainerUUID:          string(immatureTestInstructionContainerInformationMessage.TestInstructionContainerUUID),
					TestInstructionContainerName:          string(immatureTestInstructionContainerInformationMessage.TestInstructionContainerName),
					DropZoneUUID:                          string(immatureTestInstructionContainerInformationMessage.DropZoneUUID),
					DropZoneName:                          string(immatureTestInstructionContainerInformationMessage.DropZoneName),
					DropZoneDescription:                   immatureTestInstructionContainerInformationMessage.DropZoneDescription,
					DropZoneMouseOver:                     immatureTestInstructionContainerInformationMessage.DropZoneMouseOver,
					DropZoneColor:                         string(immatureTestInstructionContainerInformationMessage.DropZoneColor),
					TestInstructionContainerAttributeType: string(immatureTestInstructionContainerInformationMessage.TestInstructionAttributeType),
					TestInstructionContainerAttributeUUID: string(immatureTestInstructionContainerInformationMessage.TestInstructionAttributeUUID),
					TestInstructionContainerAttributeName: string(immatureTestInstructionContainerInformationMessage.TestInstructionAttributeName),
					AttributeValueAsString:                string(immatureTestInstructionContainerInformationMessage.AttributeValueAsString),
					AttributeValueUUID:                    string(immatureTestInstructionContainerInformationMessage.AttributeValueUUID),
					FirstImmatureElementUUID:              string(immatureTestInstructionContainerInformationMessage.FirstImmatureElementUUID),
					AttributeActionCommand:                int32(immatureTestInstructionContainerInformationMessage.AttributeActionCommand),
				}

				// Append to slice of messages
				immatureTestInstructionContainerInformationMessagesGrpc = append(immatureTestInstructionContainerInformationMessagesGrpc, immatureTestInstructionContainerInformationMessageGrpc)
			}

			// Loop and Create 'ImmatureElementModel' for 'testInstructionContainerInstanceVersionMessageGrpc'
			var immatureElementModelMessagesGrpc []*fenixExecutionWorkerGrpcApi.SupportedImmatureElementModelMessage
			for _, immatureElementModel := range testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.ImmatureElementModel {

				// Create the gRPC-version of the message
				var immatureElementModelMessageGrpc *fenixExecutionWorkerGrpcApi.SupportedImmatureElementModelMessage
				immatureElementModelMessageGrpc = &fenixExecutionWorkerGrpcApi.SupportedImmatureElementModelMessage{
					DomainUUID:               string(immatureElementModel.DomainUUID),
					DomainName:               string(immatureElementModel.DomainName),
					ImmatureElementUUID:      string(immatureElementModel.ImmatureElementUUID),
					ImmatureElementName:      string(immatureElementModel.ImmatureElementName),
					PreviousElementUUID:      string(immatureElementModel.PreviousElementUUID),
					NextElementUUID:          string(immatureElementModel.NextElementUUID),
					FirstChildElementUUID:    string(immatureElementModel.FirstChildElementUUID),
					ParentElementUUID:        string(immatureElementModel.ParentElementUUID),
					TestCaseModelElementType: string(immatureElementModel.TestCaseModelElementType),
					OriginalElementUUID:      string(immatureElementModel.OriginalElementUUID),
					TopImmatureElementUUID:   string(immatureElementModel.TopImmatureElementUUID),
					IsTopElement:             immatureElementModel.IsTopElement,
				}

				// Append to slice of messages
				immatureElementModelMessagesGrpc = append(immatureElementModelMessagesGrpc, immatureElementModelMessageGrpc)
			}

			// Create TestInstructionContainerInstance to be converted into a gRPC-version
			var testInstructionContainerInstanceVersionMessageGrpc *fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerInstanceVersionMessage
			testInstructionContainerInstanceVersionMessageGrpc = &fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerInstanceVersionMessage{
				TestInstructionContainerInstance: &fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerMessage{
					TestInstructionContainer: &fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerBaseMessage{
						DomainUUID:                            string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.DomainUUID),
						DomainName:                            string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.DomainName),
						TestInstructionContainerUUID:          string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID),
						TestInstructionContainerName:          string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName),
						TestInstructionContainerTypeUUID:      string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerTypeUUID),
						TestInstructionContainerTypeName:      string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerTypeName),
						TestInstructionContainerDescription:   string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerDescription),
						TestInstructionContainerMouseOverText: string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerMouseOverText),
						Deprecated:                            testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.Deprecated,
						Enabled:                               testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.Enabled,
						MajorVersionNumber:                    int32(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.MajorVersionNumber),
						MinorVersionNumber:                    int32(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.MinorVersionNumber),
						UpdatedTimeStamp:                      string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.UpdatedTimeStamp),
						ChildrenIsParallelProcessed:           testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.TestInstructionContainer.ChildrenIsParallelProcessed,
					},
					BasicTestInstructionContainerInformation: &fenixExecutionWorkerGrpcApi.SupportedBasicTestInstructionContainerInformationMessage{
						DomainUUID:                            string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.DomainUUID),
						DomainName:                            string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.DomainName),
						TestInstructionContainerUUID:          string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerUUID),
						TestInstructionContainerName:          string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerName),
						TestInstructionContainerTypeUUID:      string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerTypeUUID),
						TestInstructionContainerTypeName:      string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerTypeName),
						Deprecated:                            testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.Deprecated,
						MajorVersionNumber:                    int32(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.MajorVersionNumber),
						MinorVersionNumber:                    int32(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.MinorVersionNumber),
						UpdatedTimeStamp:                      string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.UpdatedTimeStamp),
						TestInstructionContainerColor:         string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerColor),
						TCRuleDeletion:                        string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TCRuleDeletion),
						TCRuleSwap:                            string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TCRuleSwap),
						TestInstructionContainerDescription:   string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerDescription),
						TestInstructionContainerMouseOverText: string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerMouseOverText),
						Enabled:                               testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.Enabled,
						TestInstructionContainerExecutionType: string(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.TestInstructionContainerExecutionType),
					},
					ImmatureTestInstructionContainerInformations: immatureTestInstructionContainerInformationMessagesGrpc,
					ImmatureElementModel:                         immatureElementModelMessagesGrpc,
				},
				TestInstructionContainerInstanceMajorVersion: int32(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstanceMajorVersion),
				TestInstructionContainerInstanceMinorVersion: int32(testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstanceMinorVersion),
				Deprecated: testInstructionContainerInstanceVersionMessage.Deprecated,
				Enabled:    testInstructionContainerInstanceVersionMessage.Enabled,
				TestInstructionContainerInstanceVersionHash: testInstructionContainerInstanceVersionMessage.TestInstructionContainerInstanceHash,
			}

			// Append to slice of messages
			testInstructionContainerInstanceVersionMessagesGrpc = append(testInstructionContainerInstanceVersionMessagesGrpc, testInstructionContainerInstanceVersionMessageGrpc)
		}

		// Create the object to be saved in Map
		var testInstructionContainerInstanceVersionsGrpc *fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerInstanceVersionsMessage
		testInstructionContainerInstanceVersionsGrpc = &fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainerInstanceVersionsMessage{
			TestInstructionContainerVersions: testInstructionContainerInstanceVersionMessagesGrpc,
			TestInstructionVersionsHash:      testInstructionContainerInstanceVersionsMessage.TestInstructionContainerVersionsHash,
		}

		// Add to Map
		testInstructionContainerInstanceVersionsMessageMap[string(originalElementUUIDType)] = testInstructionContainerInstanceVersionsGrpc

	}

	// Convert Allowed Users into gRPC-message-version
	// Loop all TestInstructionContainerInstanceVersionsMessages
	var allowedUsersGrpc []*fenixExecutionWorkerGrpcApi.SupportedAllowedUserMessage
	for _, allowedUser := range testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllowedUsers {

		var allowedUserGrpc *fenixExecutionWorkerGrpcApi.SupportedAllowedUserMessage
		allowedUserGrpc = &fenixExecutionWorkerGrpcApi.SupportedAllowedUserMessage{
			UserIdOnComputer:     allowedUser.UserIdOnComputer,
			GCPAuthenticatedUser: allowedUser.GCPAuthenticatedUser,
			UserEmail:            allowedUser.UserEmail,
			UserFirstName:        allowedUser.UserFirstName,
			UserLastName:         allowedUser.UserLastName,
			UserAuthorizationRights: &fenixExecutionWorkerGrpcApi.UserAuthorizationRightsMessage{
				CanListAndViewTestCaseOwnedByThisDomain:                    allowedUser.UserAuthorizationRights.CanListAndViewTestCaseOwnedByThisDomain,
				CanBuildAndSaveTestCaseOwnedByThisDomain:                   allowedUser.UserAuthorizationRights.CanBuildAndSaveTestCaseOwnedByThisDomain,
				CanListAndViewTestCaseHavingTIandTICFromThisDomain:         allowedUser.UserAuthorizationRights.CanListAndViewTestCaseHavingTIandTICFromThisDomain,
				CanListAndViewTestCaseHavingTIandTICFromThisDomainExtended: allowedUser.UserAuthorizationRights.CanListAndViewTestCaseHavingTIandTICFromThisDomainExtended,
				CanBuildAndSaveTestCaseHavingTIandTICFromThisDomain:        allowedUser.UserAuthorizationRights.CanBuildAndSaveTestCaseHavingTIandTICFromThisDomain,
			},
		}

		// Append to slice of messages
		allowedUsersGrpc = append(allowedUsersGrpc, allowedUserGrpc)
	}
	// Convert Connector-Domain-info into gRPC-message-version
	var connectorsDomainGrpc *fenixExecutionWorkerGrpcApi.SupportedConnectorDomainMessage
	connectorsDomainGrpc = &fenixExecutionWorkerGrpcApi.SupportedConnectorDomainMessage{
		ConnectorsDomainUUID: string(testInstructionsAndTestInstructionContainersMessage.ConnectorsDomain.ConnectorsDomainUUID),
		ConnectorsDomainName: string(testInstructionsAndTestInstructionContainersMessage.ConnectorsDomain.ConnectorsDomainName),
		ConnectorsDomainHash: testInstructionsAndTestInstructionContainersMessage.ConnectorsDomain.ConnectorsDomainHash,
	}

	// Create the full gRPC message for all supported TestInstructions, TestInstructionContainers and Allowed Users
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage = &fenixExecutionWorkerGrpcApi.SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage{
		ClientSystemIdentification: &fenixExecutionWorkerGrpcApi.ClientSystemIdentificationMessage{
			DomainUuid:                   domainUuid,
			ProtoFileVersionUsedByClient: fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum(GetHighestExecutionWorkerProtoFileVersion()),
		},
		TestInstructions: &fenixExecutionWorkerGrpcApi.SupportedTestInstructionsMessage{
			TestInstructionsMap:  testInstructionInstanceVersionsMessageMap,
			TestInstructionsHash: testInstructionsAndTestInstructionContainersMessage.TestInstructions.TestInstructionsHash,
		},
		TestInstructionContainers: &fenixExecutionWorkerGrpcApi.SupportedTestInstructionContainersMessage{
			TestInstructionsMap:           testInstructionContainerInstanceVersionsMessageMap,
			TestInstructionContainersHash: testInstructionsAndTestInstructionContainersMessage.TestInstructionContainers.TestInstructionContainersHash,
		},
		AllowedUsers: &fenixExecutionWorkerGrpcApi.SupportedAllowedUsersMessage{
			AllowedUsers:     allowedUsersGrpc,
			AllowedUsersHash: testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllowedUsersHash,
		},
		MessageCreationTimeStamp:                                        timestamppb.New(testInstructionsAndTestInstructionContainersMessage.MessageCreationTimeStamp),
		TestInstructionsAndTestInstructionsContainersMessageHash:        testInstructionsAndTestInstructionContainersMessage.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash,
		ForceNewBaseLineForTestInstructionsAndTestInstructionContainers: testInstructionsAndTestInstructionContainersMessage.ForceNewBaseLineForTestInstructionsAndTestInstructionContainers,
		ConnectorDomain: connectorsDomainGrpc,
	}

	return supportedTestInstructionsAndTestInstructionContainersAndAllowedUsersMessage, err
}
