package shared_code

import (
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"time"
)

// GenerateDatetimeTimeStampForDB
// Generate DataBaseTimeStamp, eg '2022-02-08 17:35:04.000000'
func GenerateDatetimeTimeStampForDB() (currentTimeStampAsString string) {

	timeStampLayOut := "2006-01-02 15:04:05.000000" //milliseconds
	currentTimeStamp := time.Now()
	currentTimeStampAsString = currentTimeStamp.Format(timeStampLayOut)

	return currentTimeStampAsString
}

// GenerateDatetimeFromTimeInputForDB
// Generate DataBaseTimeStamp, eg '2022-02-08 17:35:04.000000'
func GenerateDatetimeFromTimeInputForDB(currentTime time.Time) (currentTimeStampAsString string) {

	timeStampLayOut := "2006-01-02 15:04:05.000000" //milliseconds
	currentTimeStampAsString = currentTime.Format(timeStampLayOut)

	return currentTimeStampAsString
}

// ConvertGrpcTimeStampToStringForDB
// Convert a gRPC-timestamp into a string that can be used to store in the database
func ConvertGrpcTimeStampToStringForDB(grpcTimeStamp *timestamppb.Timestamp) (grpcTimeStampAsTimeStampAsString string) {
	grpcTimeStampAsTimeStamp := grpcTimeStamp.AsTime()

	timeStampLayOut := "2006-01-02 15:04:05.000000" //milliseconds

	grpcTimeStampAsTimeStampAsString = grpcTimeStampAsTimeStamp.Format(timeStampLayOut)

	return grpcTimeStampAsTimeStampAsString
}

// GetHighestExecutionWorkerProtoFileVersion
// Get the highest GetHighestExecutionWorkerProtoFileVersion for Execution Worker
func GetHighestExecutionWorkerProtoFileVersion() int32 {

	// Check if there already is a 'highestExecutionWorkerProtoFileVersion' saved, if so use that one
	if highestExecutionWorkerProtoFileVersion != -1 {
		return highestExecutionWorkerProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixExecutionWorkerGrpcApi.CurrentFenixExecutionWorkerProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestExecutionWorkerProtoFileVersion = maxValue

	return highestExecutionWorkerProtoFileVersion
}

// GetHighestExecutionBuilderProtoFileVersion
// Get the highest GetHighestExecutionBuilderProtoFileVersion for TestCase Builder Server
func GetHighestBuilderProtoFileVersion() int32 {

	// Check if there already is a 'highestExecutionBuilderProtoFileVersion' saved, if so use that one
	if highestExecutionBuilderProtoFileVersion != -1 {
		return highestExecutionBuilderProtoFileVersion
	}

	// Find the highest value for proto-file version
	var maxValue int32
	maxValue = 0

	for _, v := range fenixTestCaseBuilderServerGrpcApi.CurrentFenixTestCaseBuilderProtoFileVersionEnum_value {
		if v > maxValue {
			maxValue = v
		}
	}

	highestExecutionBuilderProtoFileVersion = maxValue

	return highestExecutionBuilderProtoFileVersion
}

// MustGetenv
// is a helper function for getting environment variables.
// End program if environment variable is not set .
func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}
