package shared_code

import (
	"encoding/json"
	"fmt"
	"github.com/jlambert68/FenixSyncShared"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"strconv"
)

// VerifyTestInstructionAndTestInstructionContainerAndUsersMessageHashesAndDomain
// Verifies the hashes for the test instructions, test instruction containers, and allowed users in the
// given gRPC-message and compare to calculates Hashes.
// This functions also verify that the same DomainUUID is used everywhere in the message
func VerifyTestInstructionAndTestInstructionContainerAndUsersMessageHashesAndDomain(
	domainUUIDToVerify TypeAndStructs.DomainUUIDType,
	testInstructionsAndTestInstructionContainersMessageToCheck *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) (errorSlice []error) {

	// Used for converting before hashing and when hashing
	var byteSlice []byte
	var byteSliceAsString string
	var hashedValue string

	// Errors that will be created when comparing calculated hash to already existing Hash sare stored in this slice
	var err error
	var wrongHashesOrDomainUUIDSlice []error

	// Loop TestInstruction
	var testInstructionInstancesHashesSlice []string
	for _, tempTestInstruction := range testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructions.TestInstructionsMap {

		// For each TestInstruction loop TestInstructionVersions
		var testInstructionVersionsHashesSlice []string
		for _, tempTestInstructionVersion := range tempTestInstruction.TestInstructionVersions {

			// Temporary set 'Hash' to a standard value to be able to recreate Hash-value
			var tempTestInstructionInstanceVersionHash string
			var tempTestInstructionInstanceVersionAndResponseVariablesHash string

			// Save Hash-value before hashing
			tempTestInstructionInstanceVersionHash = tempTestInstructionVersion.TestInstructionInstanceVersionHash
			tempTestInstructionInstanceVersionAndResponseVariablesHash = tempTestInstructionVersion.TestInstructionInstanceVersionAndResponseVariablesHash

			// Set 'Hash' to a standard value
			tempTestInstructionVersion.TestInstructionInstanceVersionHash = InitialValueBeforeHashed
			tempTestInstructionVersion.TestInstructionInstanceVersionAndResponseVariablesHash = InitialValueBeforeHashed

			// Save local copy of 'ResponseVariablesMapStructure'
			var tempLocalResponseVariablesMapStructure *TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct
			tempLocalResponseVariablesMapStructure = tempTestInstructionVersion.ResponseVariablesMapStructure

			// Clear 'ResponseVariablesMapStructure' before hashing
			tempTestInstructionVersion.ResponseVariablesMapStructure = nil

			// Convert TestInstructionVersion to byte-string and then Hash message
			byteSlice, err = json.Marshal(&tempTestInstructionVersion)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return []error{err}
			}

			// Convert byteSlice into string
			byteSliceAsString = string(byteSlice)

			// Hash the json-string
			hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

			// Verify if recalculated hash is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			if tempTestInstructionInstanceVersionHash != hashedValue {
				var newHashError error
				newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for TestInstruction "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
					tempTestInstructionInstanceVersionHash,
					hashedValue,
					"1abb307e-82b7-4a00-914b-96a83932d3a6")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
			}

			// Set the new Hash
			tempTestInstructionVersion.TestInstructionInstanceVersionHash = hashedValue

			// Repopulate Hash-value after Hashing
			tempTestInstructionVersion.TestInstructionInstanceVersionAndResponseVariablesHash = tempTestInstructionInstanceVersionAndResponseVariablesHash

			// Repopulate ResponseVariablesMapStructure-object after Hashing
			tempTestInstructionVersion.ResponseVariablesMapStructure = tempLocalResponseVariablesMapStructure

			// Create Hashes for Response variables
			var responseVariablesHashesSlice []string
			for _, tempResponseVariable := range tempTestInstructionVersion.
				ResponseVariablesMapStructure.ResponseVariablesMap {

				// Convert Response Variable to byte-string and then Hash message
				byteSlice, err = json.Marshal(&tempResponseVariable.ResponseVariable)
				if err != nil {
					fmt.Printf("Error: %s", err)
					return []error{err}
				}

				// Convert byteSlice into string
				byteSliceAsString = string(byteSlice)

				// Hash the json-string
				hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

				// Verify if recalculated hash is the same that was received via gRPC-message for specific TestInstructionContainerInstanceVersion
				if tempResponseVariable.ResponseVariableHash != hashedValue {
					var newHashError error
					newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for ResponseVariable "+
						"with UUID=%s, with Name=%s. Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
						tempResponseVariable.ResponseVariable.ResponseVariableUuid,
						tempResponseVariable.ResponseVariable.ResponseVariableName,
						tempResponseVariable.ResponseVariableHash,
						hashedValue,
						"db52c9f9-18d5-4086-a855-2196fa9e8e5b")

					// Append Error to slice with Errors
					wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
				}

				// Set the new Hash
				tempResponseVariable.ResponseVariableHash = hashedValue

				// Add the hash to slice of Hashes for Response Variables
				responseVariablesHashesSlice = append(responseVariablesHashesSlice, hashedValue)

				// Store back the Response variable in the Map
				tempTestInstructionVersion.ResponseVariablesMapStructure.
					ResponseVariablesMap[tempResponseVariable.ResponseVariable.ResponseVariableUuid] = tempResponseVariable
			}

			// Hash all values in slice with hashes for Response variables
			var hashedValueForResponseVariables string
			hashedValueForResponseVariables = fenixSyncShared.HashValues(responseVariablesHashesSlice, false)

			// Verify if recalculated hash is the same that was received via gRPC-message for final Response variables
			if tempTestInstructionVersion.ResponseVariablesMapStructure.
				ResponseVariablesMapHash != hashedValueForResponseVariables {
				var newHashError error
				newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for all ResponseVariables "+
					"Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
					tempTestInstructionVersion.ResponseVariablesMapStructure.
						ResponseVariablesMapHash,
					hashedValue,
					"e342d1ec-30c6-4f9d-aab9-344d8009c7d5")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
			}

			// Store the final Response variables Hash in the structure
			tempTestInstructionVersion.ResponseVariablesMapStructure.
				ResponseVariablesMapHash = hashedValueForResponseVariables

			// Calculate to total hash for TestInstructionInstance
			var tempTotalTestInstructionInstanceVersionHash []string

			// Append the hash for the TestInstructionInstance itself
			tempTotalTestInstructionInstanceVersionHash = append(tempTotalTestInstructionInstanceVersionHash, tempTestInstructionVersion.TestInstructionInstanceVersionHash)

			// Append the hash for the Response variables
			tempTotalTestInstructionInstanceVersionHash = append(tempTotalTestInstructionInstanceVersionHash, hashedValueForResponseVariables)

			// Create the hash to be store for the complete TestInstructionInstance
			hashedValue = fenixSyncShared.HashValues(tempTotalTestInstructionInstanceVersionHash, false)

			// Verify if recalculated hash for full TestInstructionInstanceVersion is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			if tempTestInstructionVersion.TestInstructionInstanceVersionAndResponseVariablesHash != hashedValue {
				var newHashError error
				newHashError = fmt.Errorf("Recalculated full Hash is not the same as received Hash for TestInstructionInstanceVersion "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstanceVersionAndResponseVariablesHash,
					hashedValue,
					"2a0858df-8b05-4cde-9e9f-21a452b1f5bd")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
			}

			// Set the new Hash
			tempTestInstructionVersion.TestInstructionInstanceVersionAndResponseVariablesHash = hashedValue

			// Add the hash to slice of Hashes for TestInstInstructionVersions
			testInstructionVersionsHashesSlice = append(testInstructionVersionsHashesSlice, hashedValue)

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			// TestInstruction-Struct
			if tempTestInstructionVersion.TestInstructionInstance.TestInstruction.DomainUUID != domainUUIDToVerify {
				var newDomainError error
				newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstruction(TestInstruction-Struct) "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.DomainUUID,
					domainUUIDToVerify,
					"f961187f-92c2-4389-87a5-8932d5baa95b")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
			}

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			// BasicTestInstructionInformation-struct
			if tempTestInstructionVersion.TestInstructionInstance.BasicTestInstructionInformation.DomainUUID != domainUUIDToVerify {
				var newDomainError error
				newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstruction(BasicTestInstructionInformation-struct) "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstance.BasicTestInstructionInformation.DomainUUID,
					domainUUIDToVerify,
					"df2e4dd5-2245-4838-9c81-75b2334dbc8b")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
			}

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			// Domains used within ImmatureTestInstructionInformation
			for slicePosition, tempImmatureTestInstructionInformation := range tempTestInstructionVersion.TestInstructionInstance.ImmatureTestInstructionInformation {

				if tempImmatureTestInstructionInformation.DomainUUID != domainUUIDToVerify {
					var newDomainError error
					newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstruction(ImmatureTestInstructionInformation, ArrayPosition=%d) "+
						"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
						slicePosition,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
						tempImmatureTestInstructionInformation.DomainUUID,
						domainUUIDToVerify,
						"90eb2907-cd79-41b8-b2c3-d495a99c44ac")

					// Append Error to slice with Errors
					wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
				}
			}

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			// Domains used within ImmatureElementModel
			for slicePosition, tempImmatureElementModel := range tempTestInstructionVersion.TestInstructionInstance.ImmatureElementModel {

				if tempImmatureElementModel.DomainUUID != domainUUIDToVerify {
					var newDomainError error
					newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstruction(ImmatureElementModel, ArrayPosition=%d) "+
						"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
						slicePosition,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
						tempImmatureElementModel.DomainUUID,
						domainUUIDToVerify,
						"64b845d0-db13-40b8-b6d4-873aa3343b2d")

					// Append Error to slice with Errors
					wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
				}
			}

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			// Domains used within TestInstructionAttribute
			for slicePosition, tempTestInstructionAttribute := range tempTestInstructionVersion.TestInstructionInstance.TestInstructionAttribute {

				if tempTestInstructionAttribute.DomainUUID != domainUUIDToVerify {
					var newDomainError error
					newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstruction(TestInstructionAttribute, ArrayPosition=%d) "+
						"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
						slicePosition,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
						tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
						tempTestInstructionAttribute.DomainUUID,
						domainUUIDToVerify,
						"0d074a43-f0ec-418b-94d5-978e7b881ebd")

					// Append Error to slice with Errors
					wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
				}
			}
		}
		// Hash all values in slice with hashes for TestInstInstructionVersions
		hashedValue = fenixSyncShared.HashValues(testInstructionVersionsHashesSlice, false)

		// Verify if recalculated hash is the same that was received via gRPC-message for the TestInstructionInstance,that have all versions
		if tempTestInstruction.TestInstructionVersionsHash != hashedValue {
			var newHashError error
			newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for TestInstructionInstance, "+
				"with all its versions having UUID=%s, with Name=%s. Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
				tempTestInstruction.TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
				tempTestInstruction.TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName,
				tempTestInstruction.TestInstructionVersionsHash,
				hashedValue,
				"9bf6effd-4f3e-4f1a-8ca7-f228dc8ca751")

			// Append Error to slice with Errors
			wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
		}

		// Set the new Hash
		tempTestInstruction.TestInstructionVersionsHash = hashedValue

		// hash for TestInstructionInstance to slice of hashes for all TestInstructionInstances
		testInstructionInstancesHashesSlice = append(testInstructionInstancesHashesSlice, hashedValue)

	}

	// Hash all values in slice with hashes for TestInstructionInstances
	hashedValue = fenixSyncShared.HashValues(testInstructionInstancesHashesSlice, false)

	// Verify if recalculated hash is the same that was received via gRPC-message for all TestInstructionInstances
	if testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructions.TestInstructionsHash != hashedValue {
		var newHashError error
		newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for all TestInstructions, "+
			"Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
			testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructions.TestInstructionsHash,
			hashedValue,
			"1aa93d5b-a531-4df4-a001-c88c36451c38")

		// Append Error to slice with Errors
		wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
	}

	// Set the new Hash
	testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructions.TestInstructionsHash = hashedValue

	// Loop TestInstructionContainer
	var TestInstructionContainerInstancesHashesSlice []string
	for _, tempTestInstructionContainer := range testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionContainers.TestInstructionContainersMap {

		// For each TestInstructionContainer loop TestInstructionContainerVersions
		var TestInstructionContainerVersionsHashesSlice []string
		for _, tempTestInstructionContainerVersion := range tempTestInstructionContainer.TestInstructionContainerVersions {

			// Temporary set 'Hash' to a standard value to be able to recreate Hash-value
			var tempHashValue string

			// Save Hash-value before hashing
			tempHashValue = tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash

			// Set 'Hash' to a standard value
			tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash = InitialValueBeforeHashed

			// Convert TestInstructionContainerVersion to byte-string and then Hash message
			byteSlice, err = json.Marshal(&tempTestInstructionContainerVersion)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return []error{err}
			}

			// Repopulate Hash-value after Hashing
			tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash = tempHashValue

			// Convert byteSlice into string
			byteSliceAsString = string(byteSlice)

			// Hash the json-string
			hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

			// Verify if recalculated hash is the same that was received via gRPC-message for specific TestInstructionContainerInstanceVersion
			if tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash != hashedValue {
				var newHashError error
				newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for TestInstructionContainer "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MajorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MinorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash,
					hashedValue,
					"8cdab5d4-ebcf-4ffc-b54e-1f349789f2cd")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
			}

			// Set the new Hash
			tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash = hashedValue

			// Add the hash to slice of Hashes for TestInstInstructionVersions
			TestInstructionContainerVersionsHashesSlice = append(TestInstructionContainerVersionsHashesSlice, hashedValue)

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionContainerInstanceVersion
			// TestInstructionContainer-Struct
			if tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.DomainUUID != domainUUIDToVerify {
				var newDomainError error
				newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstructionContainer(TestInstructionContainer-Struct) "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MajorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MinorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.DomainUUID,
					domainUUIDToVerify,
					"2d372f66-f5c4-4af2-a1ac-b1a52b996728")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
			}

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionContainerInstanceVersion
			// BasicTestInstructionContainerInformation-struct
			if tempTestInstructionContainerVersion.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.DomainUUID != domainUUIDToVerify {
				var newDomainError error
				newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstructionContainer(BasicTestInstructionContainerInformation-struct) "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MajorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MinorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.BasicTestInstructionContainerInformation.DomainUUID,
					domainUUIDToVerify,
					"1776ce6c-6df6-4474-baf9-2407d3c4bda0")

				// Append Error to slice with Errors
				wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
			}

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionContainerInstanceVersion
			// Domains used within ImmatureTestInstructionInformation
			for slicePosition, tempImmatureTestInstructionContainer := range tempTestInstructionContainerVersion.TestInstructionContainerInstance.ImmatureTestInstructionContainer {

				if tempImmatureTestInstructionContainer.DomainUUID != domainUUIDToVerify {
					var newDomainError error
					newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstructionContainer(ImmatureTestInstructionInformation, ArrayPosition=%d) "+
						"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
						slicePosition,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MajorVersionNumber,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MinorVersionNumber,
						tempImmatureTestInstructionContainer.DomainUUID,
						domainUUIDToVerify,
						"7c48cb3a-a865-4e6c-89e7-7c2dc484c221")

					// Append Error to slice with Errors
					wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
				}
			}

			// Verify if supported DomainUUID is the same that was received via gRPC-message for specific TestInstructionContainerInstanceVersion
			// Domains used within ImmatureElementModel
			for slicePosition, tempImmatureElementModel := range tempTestInstructionContainerVersion.TestInstructionContainerInstance.ImmatureElementModel {

				if tempImmatureElementModel.DomainUUID != domainUUIDToVerify {
					var newDomainError error
					newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for TestInstructionContainer(ImmatureElementModel, ArrayPosition=%d) "+
						"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
						slicePosition,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MajorVersionNumber,
						tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MinorVersionNumber,
						tempImmatureElementModel.DomainUUID,
						domainUUIDToVerify,
						"0adf94ec-c77e-4a43-9f17-d6a6b996cae9")

					// Append Error to slice with Errors
					wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
				}
			}
		}

		// Hash all values in slice with hashes for TestInstInstructionVersions
		hashedValue = fenixSyncShared.HashValues(TestInstructionContainerVersionsHashesSlice, false)

		// Verify if recalculated hash is the same that was received via gRPC-message for the TestInstructionContainerInstance,that have all versions
		if tempTestInstructionContainer.TestInstructionContainerVersionsHash != hashedValue {
			var newHashError error
			newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for TestInstructionContainerInstance, "+
				"with all its versions having UUID=%s, with Name=%s. Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
				tempTestInstructionContainer.TestInstructionContainerVersions[0].TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
				tempTestInstructionContainer.TestInstructionContainerVersions[0].TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
				tempTestInstructionContainer.TestInstructionContainerVersionsHash,
				hashedValue,
				"aca8992b-131f-4cc3-8f9c-90ee4a40b2b4")

			// Append Error to slice with Errors
			wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
		}

		// Set the new Hash
		tempTestInstructionContainer.TestInstructionContainerVersionsHash = hashedValue

		// hash for TestInstructionContainerInstance to slice of hashes for all TestInstructionContainerInstances
		TestInstructionContainerInstancesHashesSlice = append(TestInstructionContainerInstancesHashesSlice, hashedValue)

	}

	// Hash all values in slice with hashes for TestInstructionContainerInstances
	hashedValue = fenixSyncShared.HashValues(TestInstructionContainerInstancesHashesSlice, false)

	// Verify if recalculated hash is the same that was received via gRPC-message for all TestInstructionContainerInstances
	if testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionContainers.TestInstructionContainersHash != hashedValue {
		var newHashError error
		newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for all TestInstructionContainers, "+
			"Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
			testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionContainers.TestInstructionContainersHash,
			hashedValue,
			"4722ed85-927c-459a-91dd-7690ca888de5")

		// Append Error to slice with Errors
		wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
	}

	// Set the new Hash
	testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionContainers.TestInstructionContainersHash = hashedValue

	// Loop Allowed Users
	var allowedUsersHashesSlice []string
	for _, tempAllowedUsers := range testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsers {

		// Convert AllowedUser to byte-string and then Hash message
		byteSlice, err = json.Marshal(&tempAllowedUsers)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return []error{err}
		}

		// Convert byteSlice into string
		byteSliceAsString = string(byteSlice)

		// Hash the json-string
		hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

		// Add the hash to slice of Hashes for Allowed Users
		allowedUsersHashesSlice = append(allowedUsersHashesSlice, hashedValue)
	}

	// Hash all values in slice with hashes for Allowed Users
	var hashedValueForAllowedUsers string
	hashedValueForAllowedUsers = fenixSyncShared.HashValues(allowedUsersHashesSlice, false)

	// Create Hash for AllUsersAuthorizationRights-message
	var allUsersAuthorizationRightsSlice []string
	var allUsersCanListAndViewTestCaseHavingTIandTICFromThisDomainAsString string
	var allUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomainAsString string
	allUsersCanListAndViewTestCaseHavingTIandTICFromThisDomainAsString = strconv.FormatBool(
		testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllUsersAuthorizationRights.
			AllUsersCanListAndViewTestCaseHavingTIandTICFromThisDomain)
	allUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomainAsString = strconv.FormatBool(
		testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllUsersAuthorizationRights.
			AllUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomain)

	allUsersAuthorizationRightsSlice = append(allUsersAuthorizationRightsSlice,
		allUsersCanListAndViewTestCaseHavingTIandTICFromThisDomainAsString)
	allUsersAuthorizationRightsSlice = append(allUsersAuthorizationRightsSlice,
		allUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomainAsString)

	// Hash all values in slice with value for AllUsersAuthorizationRights-message
	hashedValue = fenixSyncShared.HashValues(allUsersAuthorizationRightsSlice, true)

	// Combine hashed from AllowedUsers and AllUsersAuthorizationRights
	var combindUserSlice []string
	combindUserSlice = append(combindUserSlice, hashedValueForAllowedUsers)
	combindUserSlice = append(combindUserSlice, hashedValue)

	hashedValue = fenixSyncShared.HashValues(combindUserSlice, false)

	// Verify if recalculated hash is the same that was received via gRPC-message for all AllowedUsers-message
	if testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsersHash != hashedValue {
		var newHashError error
		newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for all AllowedUsers, "+
			"Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
			testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsersHash,
			hashedValue,
			"c420b6f6-31fe-49af-b80a-45c24dc8d61d")

		// Append Error to slice with Errors
		wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
	}

	// Set the new Hash
	testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsersHash = hashedValue

	// Create Hash for Connectors Domain-information
	var connectorsDomainInformationSlice []string
	var connectorsDomainUUIDAsString string
	var connectorsDomainNameAsString string
	connectorsDomainUUIDAsString = string(testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainUUID)
	connectorsDomainNameAsString = string(testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainName)
	connectorsDomainInformationSlice = append(connectorsDomainInformationSlice, connectorsDomainUUIDAsString)
	connectorsDomainInformationSlice = append(connectorsDomainInformationSlice, connectorsDomainNameAsString)

	// Hash all values in slice with value for Domain belongings for the Connector
	hashedValue = fenixSyncShared.HashValues(connectorsDomainInformationSlice, true)

	// Set Hash for ConnectorsDomain-information
	testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainHash = hashedValue

	// Create the full Message Hash
	var messageHash []string

	// Append TestInstructions-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.
		TestInstructions.TestInstructionsHash)

	// Append TestInstructionContainers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.
		TestInstructionContainers.TestInstructionContainersHash)

	// Append AllowedUsers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.
		AllowedUsers.AllowedUsersHash)

	// Append Connector-Domain-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.
		ConnectorsDomain.ConnectorsDomainHash)

	// Calculate message Hash
	hashedValue = fenixSyncShared.HashValues(messageHash, false)

	// *Verify if recalculated hash is the same that was received via gRPC-message for full message
	if testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash != hashedValue {
		var newHashError error
		newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for the full message, "+
			"Got Hash=%s but recalculated Hash=%s. [ErrorID=%s]",
			testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash,
			hashedValue,
			"abd112ec-eaa8-4e5c-afa1-61aeba2a5b6c")

		// Append Error to slice with Errors
		wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newHashError)
	}

	// Set the new Hash
	testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash = hashedValue

	// Verify if supported DomainUUID is the same that was received via gRPC-message for ConnectorDomain
	if testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainUUID != domainUUIDToVerify {
		var newDomainError error
		newDomainError = fmt.Errorf("Supported DomainUUID is is not the same as received DomainUUID for ConnectorDomain. "+
			"Got DomainUUID=%s but supported DomainUUID=%s. [ErrorID=%s]",
			testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainUUID,
			domainUUIDToVerify,
			"8a4788ab-1f57-477b-92f6-6abd9f71a12e")

		// Append Error to slice with Errors
		wrongHashesOrDomainUUIDSlice = append(wrongHashesOrDomainUUIDSlice, newDomainError)
	}

	return wrongHashesOrDomainUUIDSlice

}
