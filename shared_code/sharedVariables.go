package shared_code

import "github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"

// Environment variables
var (
	//RelativePathToAllowedUsersList
	// Relative path us json file with allowed users for this connector
	RelativePathToAllowedUsersList string
)

// AllowedUsers, which are loaded from a json-file
var AllowedUsersLoadFromJsonFile *TestInstructionAndTestInstuctionContainerTypes.AllowedUsersStruct

var highestExecutionWorkerProtoFileVersion int32 = -1
var highestExecutionBuilderProtoFileVersion int32 = -1

const InitialValueBeforeHashed = "HASH"
