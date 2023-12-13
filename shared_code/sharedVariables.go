package shared_code

import "FenixTestInstructionsDataAdmin/TestInstructionAndTestInstuctionContainerTypes"

// Environment variables
var (
	//RelativePathToAllowedUsersList
	// Relative path us json file with allowed users for this connector
	RelativePathToAllowedUsersList string
)

// AllowedUsers, which are loaded from a json-file
var AllowedUsersLoadFronJsonFile *TestInstructionAndTestInstuctionContainerTypes.AllowedUsersStruct

var highestExecutionWorkerProtoFileVersion int32 = -1

const InitialValueBeforeHashed = "HASH"
