package encryption_test

import (
	"testing"

	"github.com/antonio-muniz/gommon/pkg/encryption"
	fixtures "github.com/antonio-muniz/gommon/test/fixtures/encryption"
	"github.com/stretchr/testify/require"
)

func TestRSAEncryptDecrypt(t *testing.T) {
	scenarios := []struct {
		description           string
		message               string
		encryptionKey         string
		expectedDecryptionKey string
	}{
		{
			description:           "encrypts_then_decrypts_a_message_using_RSA_algorithm",
			message:               "shhh-this-is-secret",
			encryptionKey:         fixtures.PublicKey(),
			expectedDecryptionKey: fixtures.PrivateKey(),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			encryptedMessage, err := encryption.RSAEncrypt(scenario.message, scenario.encryptionKey)
			require.NoError(t, err)
			decryptedMessage, err := encryption.RSADecrypt(encryptedMessage, scenario.expectedDecryptionKey)
			require.NoError(t, err)
			require.Equal(t, scenario.message, decryptedMessage)
		})
	}
}
