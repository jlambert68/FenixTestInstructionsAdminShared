package shared_code

import (
	iam_credentials "cloud.google.com/go/iam/credentials/apiv1"
	"context"
	"crypto/tls"
	"fmt"
	fenixTestCaseBuilderServerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/go_grpc_api"
	fenixSyncShared "github.com/jlambert68/FenixSyncShared"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	iam_credentialspb "google.golang.org/genproto/googleapis/iam/credentials/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"log"
)

// SignMessageToProveIdentityToBuilderServer
// Sign Message to be sent to BuilderServer
func SignMessageToProveIdentityToBuilderServer(
	messageToBeSigned string,
	serviceAccountUsedForSigning string,
	signerIsRunningInGCP bool) (
	hashOfSignature string,
	hashedKeyId string,
	err error) {

	ctx := context.Background()

	// Initialize the client
	var credsClient *iam_credentials.IamCredentialsClient
	if signerIsRunningInGCP == true {
		// Caller is running in GCP

		// Set up the custom TLS configuration
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
		}

		// Create a new gRPC client connection with the custom TLS settings
		var conn *grpc.ClientConn
		conn, err = grpc.DialContext(ctx, "iamcredentials.googleapis.com:443", grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
		if err != nil {
			log.Fatalf("Failed to dial IAM Credentials API: %v", err)
		}
		defer conn.Close()

		credsClient, err = iam_credentials.NewIamCredentialsClient(ctx, option.WithGRPCConn(conn))

		// Add credentials
		ctx, err = attachCredentials(ctx)
		if err != nil {

			log.Fatalln(fmt.Sprintf("Problem getting the credentials token to be able to sign a message. "+
				"ErrorId=%s, Error-message=%s",
				"1e46ea03-6a67-4ee1-853d-408d60b440d5",
				err.Error()))
		}

	} else {
		// Caller is running locally
		credsClient, err = iam_credentials.NewIamCredentialsClient(ctx)
	}
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
		serviceAccountUsedForSigning,
		true)

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

// Add credentials when running in GCP
func attachCredentials(ctx context.Context) (context.Context, error) {
	// Get default credentials (this works in GCP environments like GCE, GKE, Cloud Run, etc.)
	creds, err := google.FindDefaultCredentials(ctx, iam_credentials.DefaultAuthScopes()...)
	if err != nil {
		return nil, err
	}

	// Use the TokenSource from the credentials
	tokenSource := creds.TokenSource

	// Retrieve an OAuth 2.0 token
	token, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid() {
		return nil, fmt.Errorf("token is invalid")
	}

	// Create a new context with the token attached
	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "Bearer "+token.AccessToken)

	return newCtx, nil
}
