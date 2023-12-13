package shared_code

import (
	"encoding/json"
	"fmt"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/LocalExecutionMethods"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
)

// CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes
// Calculates the hashes for the test instructions, test instruction containers, and allowed users in the given message
func CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct) (err error) {

	// Used for converting before hashing and when hashing
	var byteSlice []byte
	var byteSliceAsString string
	var hashedValue string

	// Loop TestInstruction
	var testInstructionInstancesHashesSlice []string
	for _, tempTestInstruction := range testInstructionsAndTestInstructionContainersMessage.TestInstructions.TestInstructionsMap {

		// For each TestInstruction loop TestInstructionVersions
		var testInstructionVersionsHashesSlice []string
		for _, tempTestInstructionVersion := range tempTestInstruction.TestInstructionVersions {

			// Temporary set Local Execution Methods to nil due to they shouldn't be included in Hash
			var tempLocalExecutionMethods *LocalExecutionMethods.MethodsForLocalExecutionsStruct

			// Save reference copy LocalExecution-object
			tempLocalExecutionMethods = tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods

			// Clear LocalExecution before hashing
			tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods = nil

			// Convert TestInstructionVersion to byte-string and then Hash message
			byteSlice, err = json.Marshal(&tempTestInstructionVersion)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return err
			}

			// Repopulate LocalExecution-object after Hashing
			tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods = tempLocalExecutionMethods

			// Convert byteSlice into string
			byteSliceAsString = string(byteSlice)

			// Hash the json-string
			hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

			// Add hash to the specific TestInstructionInstanceVersion
			tempTestInstructionVersion.TestInstructionInstanceVersionHash = hashedValue

			// Add the hash to slice of Hashes for TestInstInstructionVersions
			testInstructionVersionsHashesSlice = append(testInstructionVersionsHashesSlice, hashedValue)

		}
		// Hash all values in slice with hashes for TestInstInstructionVersions
		hashedValue = fenixSyncShared.HashValues(testInstructionVersionsHashesSlice, false)

		// Add hash to the TestInstructionInstance,that have all versions
		tempTestInstruction.TestInstructionVersionsHash = hashedValue

		// hash for TestInstructionInstance to slice of hashes for all TestInstructionInstances
		testInstructionInstancesHashesSlice = append(testInstructionInstancesHashesSlice, hashedValue)

	}

	// Hash all values in slice with hashes for TestInstructionInstances
	hashedValue = fenixSyncShared.HashValues(testInstructionInstancesHashesSlice, false)

	// Add hash for all TestInstructionInstances
	testInstructionsAndTestInstructionContainersMessage.TestInstructions.TestInstructionsHash = hashedValue

	// Loop TestInstructionContainer
	var TestInstructionContainerInstancesHashesSlice []string
	for _, tempTestInstructionContainer := range testInstructionsAndTestInstructionContainersMessage.TestInstructionContainers.TestInstructionContainersMap {

		// For each TestInstructionContainer loop TestInstructionContainerVersions
		var TestInstructionContainerVersionsHashesSlice []string
		for _, tempTestInstructionContainerVersion := range tempTestInstructionContainer.TestInstructionContainerVersions {

			// Convert TestInstructionContainerVersion to byte-string and then Hash message
			byteSlice, err = json.Marshal(&tempTestInstructionContainerVersion)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return err
			}

			// Convert byteSlice into string
			byteSliceAsString = string(byteSlice)

			// Hash the json-string
			hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

			// Add hash to the specific TestInstructionContainerInstanceVersion
			tempTestInstructionContainerVersion.TestInstructionContainerInstanceHash = hashedValue

			// Add the hash to slice of Hashes for TestInstInstructionVersions
			TestInstructionContainerVersionsHashesSlice = append(TestInstructionContainerVersionsHashesSlice, hashedValue)

		}
		// Hash all values in slice with hashes for TestInstInstructionVersions
		hashedValue = fenixSyncShared.HashValues(TestInstructionContainerVersionsHashesSlice, false)

		// Add hash to the TestInstructionContainerInstance,that have all versions
		tempTestInstructionContainer.TestInstructionContainerVersionsHash = hashedValue

		// hash for TestInstructionContainerInstance to slice of hashes for all TestInstructionContainerInstances
		TestInstructionContainerInstancesHashesSlice = append(TestInstructionContainerInstancesHashesSlice, hashedValue)

	}

	// Hash all values in slice with hashes for TestInstructionContainerInstances
	hashedValue = fenixSyncShared.HashValues(TestInstructionContainerInstancesHashesSlice, false)

	// Add hash for all TestInstructionContainerInstances
	testInstructionsAndTestInstructionContainersMessage.TestInstructionContainers.TestInstructionContainersHash = hashedValue

	// Loop Allowed Users
	var allowedUsersHashesSlice []string
	for _, tempAllowedUsers := range testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllowedUsers {

		// Convert AllowedUser to byte-string and then Hash message
		byteSlice, err = json.Marshal(&tempAllowedUsers)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return err
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

	// Add hash for all AllowedUsers-message
	testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllowedUsersHash = hashedValue

	// Create the full Message Hash
	var messageHash []string

	// Append TestInstructions-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessage.TestInstructions.TestInstructionsHash)

	// Append TestInstructionContainers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessage.TestInstructionContainers.TestInstructionContainersHash)

	// Append AllowedUsers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllowedUsersHash)

	// Calculate message Hash
	hashedValue = fenixSyncShared.HashValues(messageHash, false)

	// Add message hash to message
	testInstructionsAndTestInstructionContainersMessage.TestInstructionsAndTestInstructionsContainersAndUsersMessageHash = hashedValue

	return err

}
