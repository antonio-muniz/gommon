package jwt

import "strings"

type PackSettings struct {
	SignatureKey  string
	EncryptionKey string
}

type UnpackSettings struct {
	SignatureKey  string
	DecryptionKey string
}

func Pack(token Token, settings PackSettings) (string, error) {
	encodedToken, err := Serialize(token)
	if err != nil {
		return "", err
	}
	signedToken, err := Sign(encodedToken, settings.SignatureKey)
	if err != nil {
		return "", err
	}
	encryptedToken, err := Encrypt(signedToken, settings.EncryptionKey)
	if err != nil {
		return "", err
	}
	return encryptedToken, nil
}

func Unpack(packedToken string, settings UnpackSettings) (Token, error) {
	decryptedToken, err := Decrypt(packedToken, settings.DecryptionKey)
	if err != nil {
		return Token{}, err
	}
	err = CheckSignature(decryptedToken, settings.SignatureKey)
	if err != nil {
		return Token{}, err
	}
	separationIndex := strings.LastIndex(decryptedToken, ".")
	unsignedToken := decryptedToken[:separationIndex]
	token, err := Deserialize(unsignedToken)
	if err != nil {
		return Token{}, err
	}
	return token, nil
}
