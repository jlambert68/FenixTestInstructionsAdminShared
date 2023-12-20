package shared_code

import (
	"encoding/json"
	"fmt"
	"github.com/jlambert68/FenixSyncShared"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
)

// VerifyTestInstructionAndTestInstructionContainerAndUsersMessageHashes
// Verifies the hashes for the test instructions, test instruction containers, and allowed users in the
// given gRPC-message and compare to calculates Hashes
func VerifyTestInstructionAndTestInstructionContainerAndUsersMessageHashes(
	testInstructionsAndTestInstructionContainersMessageToCheck *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) (errorSlice []error) {

	// Used for converting before hashing and when hashing
	var byteSlice []byte
	var byteSliceAsString string
	var hashedValue string

	// Errors that will be created when comparing calculated hash to already existing Hash sare stored in this slice
	var err error
	var wrongHashesSlice []error

	// Loop TestInstruction
	var testInstructionInstancesHashesSlice []string
	for _, tempTestInstruction := range testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructions.TestInstructionsMap {

		// For each TestInstruction loop TestInstructionVersions
		var testInstructionVersionsHashesSlice []string
		for _, tempTestInstructionVersion := range tempTestInstruction.TestInstructionVersions {

			// Temporary set 'Hash' to a standard value to be able to recreate Hash-value
			var tempHashValue string

			// Save Hash-value before hashing
			tempHashValue = tempTestInstructionVersion.TestInstructionInstanceVersionHash

			// Set 'Hash' to a standard value
			tempTestInstructionVersion.TestInstructionInstanceVersionHash = InitialValueBeforeHashed

			// Convert TestInstructionVersion to byte-string and then Hash message
			byteSlice, err = json.Marshal(&tempTestInstructionVersion)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return []error{err}
			}

			// Repopulate Hash-value after Hashing
			tempTestInstructionVersion.TestInstructionInstanceVersionHash = tempHashValue

			// Convert byteSlice into string
			byteSliceAsString = string(byteSlice)

			// Hash the json-string
			hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

			// Verify if recalculated hash is the same that was received via gRPC-message for specific TestInstructionInstanceVersion
			if tempTestInstructionVersion.TestInstructionInstanceVersionHash != hashedValue {
				var newHashError error
				newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for TestInstruction "+
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got Hash=%s but recalculated Hash=%s. ",
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionUUID,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.TestInstructionName,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MajorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstance.TestInstruction.MinorVersionNumber,
					tempTestInstructionVersion.TestInstructionInstanceVersionHash,
					hashedValue)

				// Append Error to slice with Errors
				wrongHashesSlice = append(wrongHashesSlice, newHashError)
			}

			// Set the new Hash
			tempTestInstructionVersion.TestInstructionInstanceVersionHash = hashedValue

			// Add the hash to slice of Hashes for TestInstInstructionVersions
			testInstructionVersionsHashesSlice = append(testInstructionVersionsHashesSlice, hashedValue)

		}
		// Hash all values in slice with hashes for TestInstInstructionVersions
		hashedValue = fenixSyncShared.HashValues(testInstructionVersionsHashesSlice, false)

		// Verify if recalculated hash is the same that was received via gRPC-message for the TestInstructionInstance,that have all versions
		if tempTestInstruction.TestInstructionVersionsHash != hashedValue {
			var newHashError error
			newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for TestInstructionInstance, "+
				"with all its versions having UUID=%s, with Name=%s. Got Hash=%s but recalculated Hash=%s. ",
				tempTestInstruction.TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionUUID,
				tempTestInstruction.TestInstructionVersions[0].TestInstructionInstance.TestInstruction.TestInstructionName,
				tempTestInstruction.TestInstructionVersionsHash,
				hashedValue)

			// Append Error to slice with Errors
			wrongHashesSlice = append(wrongHashesSlice, newHashError)
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
			"Got Hash=%s but recalculated Hash=%s. ",
			testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructions.TestInstructionsHash,
			hashedValue)

		// Append Error to slice with Errors
		wrongHashesSlice = append(wrongHashesSlice, newHashError)
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
					"with UUID=%s, with Name=%s, having MajorVersion=%d and MinorVersion=%d. Got Hash=%s but recalculated Hash=%s. ",
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MajorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstance.TestInstructionContainer.MinorVersionNumber,
					tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash,
					hashedValue)

				// Append Error to slice with Errors
				wrongHashesSlice = append(wrongHashesSlice, newHashError)
			}

			// Set the new Hash
			tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash = hashedValue

			// Add the hash to slice of Hashes for TestInstInstructionVersions
			TestInstructionContainerVersionsHashesSlice = append(TestInstructionContainerVersionsHashesSlice, hashedValue)

		}
		// Hash all values in slice with hashes for TestInstInstructionVersions
		hashedValue = fenixSyncShared.HashValues(TestInstructionContainerVersionsHashesSlice, false)

		// Verify if recalculated hash is the same that was received via gRPC-message for the TestInstructionContainerInstance,that have all versions
		if tempTestInstructionContainer.TestInstructionContainerVersionsHash != hashedValue {
			var newHashError error
			newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for TestInstructionContainerInstance, "+
				"with all its versions having UUID=%s, with Name=%s. Got Hash=%s but recalculated Hash=%s. ",
				tempTestInstructionContainer.TestInstructionContainerVersions[0].TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerUUID,
				tempTestInstructionContainer.TestInstructionContainerVersions[0].TestInstructionContainerInstance.TestInstructionContainer.TestInstructionContainerName,
				tempTestInstructionContainer.TestInstructionContainerVersionsHash,
				hashedValue)

			// Append Error to slice with Errors
			wrongHashesSlice = append(wrongHashesSlice, newHashError)
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
			"Got Hash=%s but recalculated Hash=%s. ",
			testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionContainers.TestInstructionContainersHash,
			hashedValue)

		// Append Error to slice with Errors
		wrongHashesSlice = append(wrongHashesSlice, newHashError)
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
	hashedValue = fenixSyncShared.HashValues(allowedUsersHashesSlice, false)

	// Verify if recalculated hash is the same that was received via gRPC-message for all AllowedUsers-message
	if testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsersHash != hashedValue {
		var newHashError error
		newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for all AllowedUsers, "+
			"Got Hash=%s but recalculated Hash=%s. ",
			testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsersHash,
			hashedValue)

		// Append Error to slice with Errors
		wrongHashesSlice = append(wrongHashesSlice, newHashError)
	}

	// Set the new Hash
	testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsersHash = hashedValue

	// Create Hash for Connectors Domain-information
	var connectorsDomainInformationSlice []string
	connectorsDomainInformationSlice = append(connectorsDomainInformationSlice, testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainUUID)
	connectorsDomainInformationSlice = append(connectorsDomainInformationSlice, testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainName)

	// Hash all values in slice with value for Domain belongings for the Connector
	hashedValue = fenixSyncShared.HashValues(connectorsDomainInformationSlice, true)

	// Set Hash for ConnectorsDomain-information
	testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainHash = hashedValue

	// Create the full Message Hash
	var messageHash []string

	// Append TestInstructions-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructions.TestInstructionsHash)

	// Append TestInstructionContainers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionContainers.TestInstructionContainersHash)

	// Append AllowedUsers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.AllowedUsers.AllowedUsersHash)

	// Append Connector-Domain-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessageToCheck.ConnectorsDomain.ConnectorsDomainHash)

	// Calculate message Hash
	hashedValue = fenixSyncShared.HashValues(messageHash, false)

	// *Verify if recalculated hash is the same that was received via gRPC-message for full message
	if testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash != hashedValue {
		var newHashError error
		newHashError = fmt.Errorf("Recalculated Hash is not the same as received Hash for the full message, "+
			"Got Hash=%s but recalculated Hash=%s. ",
			testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash,
			hashedValue)

		// Append Error to slice with Errors
		wrongHashesSlice = append(wrongHashesSlice, newHashError)
	}

	// Set the new Hash
	testInstructionsAndTestInstructionContainersMessageToCheck.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash = hashedValue

	return wrongHashesSlice

}
