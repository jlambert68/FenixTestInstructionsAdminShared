package shared_code

import (
	iam_credentials "cloud.google.com/go/iam/credentials/apiv1"
	"context"
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	iam_credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"
)

// SignMessageToProveIdentityToBuilderServer
// Sign Message to be sent to BuilderServer
func SignMessageToProveIdentityToBuilderServer(
	messageToBeSigned string,
	serviceAccountUsedForSigning string) (
	hashOfSignature string,
	hashedKeyId string,
	err error) {

	ctx := context.Background()

	// Initialize the client
	credsClient, err := iam_credentials.NewIamCredentialsClient(ctx)
	if err != nil {
		return "", "", err
	}

	defer credsClient.Close()

	// The data to be signed
	data := []byte(messageToBeSigned)

	// Request to sign a byte array with the service account's private key
	req := &iam_credentialspb.SignBlobRequest{
		Name:    serviceAccountUsedForSigning,
		Payload: data,
	}

	// Call the API to sign the data
	var signResponse *iam_credentialspb.SignBlobResponse
	signResponse, err = credsClient.SignBlob(ctx, req)
	if err != nil {
		return "", "", err
	}

	var signedMessage []byte
	signedMessage = signResponse.SignedBlob

	// Hash the signature
	hashOfSignature = fenixSyncShared.HashSingleValue(string(signedMessage))

	// Extract KeyId used when signing
	var keyId string
	keyId = signResponse.GetKeyId()

	// Hash KeyId
	hashedKeyId = fenixSyncShared.HashSingleValue(keyId)

	// Return result
	return hashOfSignature, hashedKeyId, err
}

// VerifySignatureFromSignedMessageToProveIdentityToBuilderServer
// Verify signature received from Worker
func VerifySignatureFromSignedMessageToProveIdentityToBuilderServer(
	signedMessageByServiceAccountMessage *fenixTestCaseBuilderServerGrpcApi.SignedMessageByWorkerServiceAccountMessage,
	serviceAccountUsedForSigning string) (
	verificationOfSignatureSucceeded bool,
	err error) {

	// Recreate signature information
	var reCreatedHashOfSignature string
	var reCreatedHashedKeyId string
	reCreatedHashOfSignature, reCreatedHashedKeyId, err = SignMessageToProveIdentityToBuilderServer(
		signedMessageByServiceAccountMessage.MessageToBeSigned,
		serviceAccountUsedForSigning)

	// Got some problem when signing the message
	if err != nil {
		return false, err
	}

	// Verify recreated signature with signature produced by Worker
	if reCreatedHashOfSignature != signedMessageByServiceAccountMessage.HashOfSignature {
		return false, err
	}

	// Verify recreated KeyId with KeyId produced by Worker
	if reCreatedHashedKeyId != signedMessageByServiceAccountMessage.HashedKeyId {
		return false, err
	}

	// Success in signature verification
	return true, err
}
