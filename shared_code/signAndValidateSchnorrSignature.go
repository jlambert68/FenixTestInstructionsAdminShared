package shared_code

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/schnorr"
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
)

// SignMessageUsingSchnorrSignature
// Signs a Message. Used by the Connector to sign the content of what is sent to TestCaseBuilderServer regarding
// TestInstructions,Users and TestData
func SignMessageUsingSchnorrSignature(
	messageToSign string) (
	signatureToVerifyAsBase64String string,
	err error) {

	//var privKeyAsString string
	//privKeyAsString = "dncjhBsRB8zFrI9KUSHY3tOkKqd6kPTzKfPvIRXFO/w="

	// ***** Private Key *****
	// Load the Private key from an environment variable
	if privateKeyAsBase64String == "" {
		privateKeyAsBase64String = environmentVariables.
			ExtractEnvironmentVariableOrInjectedEnvironmentVariable("PrivateKey")
	}

	// Convert Private key from Base64 string into "a proper" private key
	var privateKeyAsByteArray []byte
	privateKeyAsByteArray, err = base64.StdEncoding.DecodeString(privateKeyAsBase64String)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when decoding private key as string: %v", err))

		return "", err
	}
	var privateKey *secp256k1.PrivateKey
	privateKey = secp256k1.PrivKeyFromBytes(privateKeyAsByteArray)

	// 	// ***** The Message *****
	// Convert message into byte array
	var messageToSignAsByteArray []byte
	messageToSignAsByteArray = []byte(messageToSign)

	// Hash the message using SHA-256.
	var messageHashAsByteArray [32]byte
	messageHashAsByteArray = sha256.Sum256(messageToSignAsByteArray)

	// ***** Sign Message *****
	// Sign the message hash using the private key.
	var schnorrSignature *schnorr.Signature
	schnorrSignature, err = schnorr.Sign(privateKey, messageHashAsByteArray[:])
	if err != nil {
		err = errors.New(fmt.Sprintf("error when signing the message: %v", err))

		return "", err
	}

	// Serialize the signature into a byte array
	var signatureAsByteArray []byte
	signatureAsByteArray = schnorrSignature.Serialize()

	// Convert the signature into a Base64 string
	signatureToVerifyAsBase64String = base64.StdEncoding.EncodeToString(signatureAsByteArray)

	return signatureToVerifyAsBase64String, err

}

// VerifySchnorrSignature
// Verifies a Schnorr signature. Used to validate that the correct Connector did send the
// TestInstructions,Users and TestData
func VerifySchnorrSignature(
	messageToVerify string,
	publicKeyAsBase64String string,
	signatureToVerifyAsBase64String string) (
	err error) {

	// ***** Public Key *****
	// Convert the public key (as string) into a real public key
	var publicKeyAsByteArray []byte
	publicKeyAsByteArray, err = base64.StdEncoding.DecodeString(publicKeyAsBase64String)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when doing base64 decoding of the public key: %v", err))

		return err
	}

	var publicKey *secp256k1.PublicKey
	publicKey, err = secp256k1.ParsePubKey(publicKeyAsByteArray)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when parsing the public key: %v", err))

		return err
	}

	// ***** Signature *****
	// Convert the signature (as string) into a real Schnorr signature
	var signatureAsByteArray []byte
	signatureAsByteArray, err = base64.StdEncoding.DecodeString(signatureToVerifyAsBase64String)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when doing base64 decoding of the signature: %v", err))

		return err
	}

	var schnorrSignature *schnorr.Signature
	schnorrSignature, err = schnorr.ParseSignature(signatureAsByteArray)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when parsing the signature: %v", err))

		return err
	}

	// 	***** The Message *****
	// Convert message into byte array
	var messageToVerifyAsByteArray []byte
	messageToVerifyAsByteArray = []byte(messageToVerify)

	// Hash the message using SHA-256.
	var messageHashAsByteArray [32]byte
	messageHashAsByteArray = sha256.Sum256(messageToVerifyAsByteArray)

	// ***** Verify Signature *****
	// Verify the signature using the public key and the message hash.
	var signatureIsValid bool
	signatureIsValid = schnorrSignature.Verify(messageHashAsByteArray[:], publicKey)
	if signatureIsValid == false {
		err = errors.New(fmt.Sprintf("signature verification failed"))

		return err
	}

	return err

}

// GenerateNewPrivateKeyAsBase64String
// Generates a new Private key which can be used when setting up a new private-public key par
func GenerateNewPrivateKeyAsBase64String() (newPrivateKeyAsBase64String string, err error) {

	// Generate a new private key.
	var newPrivateKey *secp256k1.PrivateKey
	newPrivateKey, err = secp256k1.GeneratePrivateKey()

	if err != nil {
		err = errors.New(fmt.Sprintf("error generating private key: %v", err))

		return "", err
	}

	// Convert private key into a byte array
	var privateKeyAsByteArray []byte
	privateKeyAsByteArray = newPrivateKey.Serialize()
	privateKeyAsBase64String = base64.StdEncoding.EncodeToString(privateKeyAsByteArray)

	return privateKeyAsBase64String, err
}

// GeneratePublicKeyAsBase64StringFromPrivateKey
// Generate the Public Key from the Private Key
func GeneratePublicKeyAsBase64StringFromPrivateKey() (
	publicKeyUnCompressedAsString string, err error) {

	// ***** Private Key *****
	// Load the Private key from an environment variable
	if privateKeyAsBase64String == "" {
		privateKeyAsBase64String = environmentVariables.
			ExtractEnvironmentVariableOrInjectedEnvironmentVariable("PrivateKey")
	}

	publicKeyUnCompressedAsString, err = GeneratePublicKeyAsBase64StringFromPrivateKeyInput(privateKeyAsBase64String)

	return publicKeyUnCompressedAsString, err

}

// GeneratePublicKeyAsBase64StringFromPrivateKeyInput
// Generate the Public Key from the Private Key as the input parameter
func GeneratePublicKeyAsBase64StringFromPrivateKeyInput(
	privateKeyAsBase64String string) (
	publicKeyUnCompressedAsString string, err error) {

	// Convert Private key from Base64 string into "a proper" private key
	var privateKeyAsByteArray []byte
	privateKeyAsByteArray, err = base64.StdEncoding.DecodeString(privateKeyAsBase64String)
	if err != nil {
		err = errors.New(fmt.Sprintf("error when decoding private key as string: %v", err))

		return "", err
	}
	var privateKey *secp256k1.PrivateKey
	privateKey = secp256k1.PrivKeyFromBytes(privateKeyAsByteArray)

	// ****** Public Key *****
	// Generate Public Key from Private Key
	var publicKey *secp256k1.PublicKey
	publicKey = privateKey.PubKey()

	// Convert the private key into a byte array
	var publicKeyUnCompressedAsByteArray []byte
	publicKeyUnCompressedAsByteArray = publicKey.SerializeUncompressed()

	// Convert the public key into a base64 encoded string
	publicKeyUnCompressedAsString = base64.StdEncoding.EncodeToString(publicKeyUnCompressedAsByteArray)

	return publicKeyUnCompressedAsString, err
}
