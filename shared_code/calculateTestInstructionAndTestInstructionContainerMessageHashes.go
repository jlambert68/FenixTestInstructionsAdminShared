package shared_code

import (
	"encoding/json"
	"fmt"
	"github.com/jlambert68/FenixSyncShared"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"strconv"
)

// CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes
// Calculates the hashes for the test instructions, test instruction containers, and allowed users in the given message
func CalculateTestInstructionAndTestInstructionContainerAndUsersMessageHashes(
	testInstructionsAndTestInstructionContainersMessage *TestInstructionAndTestInstuctionContainerTypes.TestInstructionsAndTestInstructionsContainersStruct,
	// pushToTempStoreFunction PushToTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType],
	// PullFromTempStoreFunction PullFromTempStoreFunctionType[*TestInstructionAndTestInstuctionContainerTypes.AnyType]
) (err error) {

	// Used for converting before hashing and when hashing
	var byteSlice []byte
	var byteSliceAsString string
	var hashedValue string

	// Loop TestInstruction
	var testInstructionInstancesHashesSlice []string
	for _, tempTestInstruction := range testInstructionsAndTestInstructionContainersMessage.TestInstructions.TestInstructionsMap {

		// For each TestInstruction loop TestInstructionVersions
		var testInstructionVersionsHashesSlice []string
		var testInstructionVersionsHash string
		for _, tempTestInstructionVersion := range tempTestInstruction.TestInstructionVersions {

			// Temporary set Local Execution Methods to nil due to they shouldn't be included in Hash
			//var tempLocalExecutionMethods *LocalExecutionMethods.MethodsForLocalExecutionsStruct
			var tempLocalExecutionMethods TestInstructionAndTestInstuctionContainerTypes.AnyType

			// Save reference copy LocalExecution-object
			//tempLocalExecutionMethods = tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods
			//pushToTempStoreFunction(tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods.Value)
			tempLocalExecutionMethods = tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods

			// Clear LocalExecution before hashing
			tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods.Value = nil

			// Save local copy of 'ResponseVariablesMapStructure'
			var tempLocalResponseVariablesMapStructure *TestInstructionAndTestInstuctionContainerTypes.ResponseVariablesMapStructureStruct
			tempLocalResponseVariablesMapStructure = tempTestInstructionVersion.ResponseVariablesMapStructure

			// Clear 'ResponseVariablesMapStructure' before hashing
			tempTestInstructionVersion.ResponseVariablesMapStructure = nil

			// Convert TestInstructionVersion to byte-string and then Hash message
			byteSlice, err = json.Marshal(&tempTestInstructionVersion)
			if err != nil {
				fmt.Printf("Error: %s", err)
				return err
			}

			// Repopulate LocalExecution-object after Hashing
			//tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods = PullFromTempStoreFunction() //tempLocalExecutionMethods
			tempTestInstructionVersion.TestInstructionInstance.LocalExecutionMethods = tempLocalExecutionMethods

			// Repopulate ResponseVariablesMapStructure-object after Hashing
			tempTestInstructionVersion.ResponseVariablesMapStructure = tempLocalResponseVariablesMapStructure

			// Convert byteSlice into string
			byteSliceAsString = string(byteSlice)

			// Hash the json-string
			testInstructionVersionsHash = fenixSyncShared.HashSingleValue(byteSliceAsString)

			// Add hash to the specific TestInstructionInstanceVersion
			tempTestInstructionVersion.TestInstructionInstanceVersionHash = testInstructionVersionsHash

			// Create Hashes for Response variables
			var responseVariablesHashesSlice []string
			for _, tempResponseVariable := range tempTestInstructionVersion.
				ResponseVariablesMapStructure.ResponseVariablesMap {

				// Convert Response Variable to byte-string and then Hash message
				byteSlice, err = json.Marshal(&tempResponseVariable.ResponseVariable)
				if err != nil {
					fmt.Printf("Error: %s", err)
					return err
				}

				// Convert byteSlice into string
				byteSliceAsString = string(byteSlice)

				// Hash the json-string
				hashedValue = fenixSyncShared.HashSingleValue(byteSliceAsString)

				// Add the hash to slice of Hashes for Allowed Users
				responseVariablesHashesSlice = append(responseVariablesHashesSlice, hashedValue)

				// Set Hash for this Response Variable
				tempResponseVariable.ResponseVariableHash = hashedValue

				// Store back the Response variable in the Map
				tempTestInstructionVersion.ResponseVariablesMapStructure.
					ResponseVariablesMap[tempResponseVariable.ResponseVariable.ResponseVariableUuid] = tempResponseVariable
			}

			// Hash all values in slice with hashes for Response variables
			var hashedValueForResponseVariables string
			hashedValueForResponseVariables = fenixSyncShared.HashValues(responseVariablesHashesSlice, false)

			// Store the final Response variables Hash in the structure
			tempTestInstructionVersion.ResponseVariablesMapStructure.ResponseVariablesMapHash = hashedValueForResponseVariables

			// Calculate to total hash for TestInstructionInstance
			var tempTotalTestInstructionInstanceVersionHash []string

			// Append the hash for the TestInstructionInstance itself
			tempTotalTestInstructionInstanceVersionHash = append(tempTotalTestInstructionInstanceVersionHash, tempTestInstructionVersion.TestInstructionInstanceVersionHash)

			// Append the hash for the Response variables
			tempTotalTestInstructionInstanceVersionHash = append(tempTotalTestInstructionInstanceVersionHash, hashedValueForResponseVariables)

			// Create the hash to be store for the complete TestInstructionInstance
			hashedValue = fenixSyncShared.HashValues(tempTotalTestInstructionInstanceVersionHash, false)

			// Add hash to the specific TestInstructionInstanceVersion
			tempTestInstructionVersion.TestInstructionInstanceVersionAndResponseVariablesHash = hashedValue

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
	var hashedValueForAllowedUsers string
	hashedValueForAllowedUsers = fenixSyncShared.HashValues(allowedUsersHashesSlice, false)

	// Create Hash for AllUsersAuthorizationRights-message
	var allUsersAuthorizationRightsSlice []string
	var allUsersCanListAndViewTestCaseHavingTIandTICFromThisDomainAsString string
	var allUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomainAsString string
	allUsersCanListAndViewTestCaseHavingTIandTICFromThisDomainAsString = strconv.FormatBool(
		testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllUsersAuthorizationRights.
			AllUsersCanListAndViewTestCaseHavingTIandTICFromThisDomain)
	allUsersCanBuildAndSaveTestCaseHavingTIandTICFromThisDomainAsString = strconv.FormatBool(
		testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllUsersAuthorizationRights.
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

	// Add hash for all AllowedUsers-message
	testInstructionsAndTestInstructionContainersMessage.AllowedUsers.AllowedUsersHash = hashedValue

	// Create Hash for Connectors Domain-information
	var connectorsDomainInformationSlice []string
	var connectorsDomainUUIDAsString string
	var connectorsDomainNameAsString string
	connectorsDomainUUIDAsString = string(testInstructionsAndTestInstructionContainersMessage.ConnectorsDomain.ConnectorsDomainUUID)
	connectorsDomainNameAsString = string(testInstructionsAndTestInstructionContainersMessage.ConnectorsDomain.ConnectorsDomainName)
	connectorsDomainInformationSlice = append(connectorsDomainInformationSlice, connectorsDomainUUIDAsString)
	connectorsDomainInformationSlice = append(connectorsDomainInformationSlice, connectorsDomainNameAsString)

	// Hash all values in slice with value for Domain belongings for the Connector
	hashedValue = fenixSyncShared.HashValues(connectorsDomainInformationSlice, true)

	// Set Hash for ConnectorsDomain-information
	testInstructionsAndTestInstructionContainersMessage.ConnectorsDomain.ConnectorsDomainHash = hashedValue

	// Create the full Message Hash
	var messageHash []string

	// Append TestInstructions-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessage.
		TestInstructions.TestInstructionsHash)

	// Append TestInstructionContainers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessage.
		TestInstructionContainers.TestInstructionContainersHash)

	// Append AllowedUsers-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessage.
		AllowedUsers.AllowedUsersHash)

	// Append Connector-Domain-hash
	messageHash = append(messageHash, testInstructionsAndTestInstructionContainersMessage.
		ConnectorsDomain.ConnectorsDomainHash)

	// Calculate message Hash
	hashedValue = fenixSyncShared.HashValues(messageHash, false)

	// Add message hash to message
	testInstructionsAndTestInstructionContainersMessage.
		TestInstructionsAndTestInstructionsContainersAndUsersMessageHash = hashedValue

	return err

}
