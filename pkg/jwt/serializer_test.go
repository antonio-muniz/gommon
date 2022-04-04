package jwt_test

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/antonio-muniz/gommon/pkg/jwt"
)

func TestSerialize(t *testing.T) {
	scenarios := []struct {
		description     string
		token           jwt.Token
		expectedHeader  map[string]interface{}
		expectedPayload map[string]interface{}
	}{
		{
			description: "serializes_a_token_to_string",
			token: jwt.Token{
				Issuer:   "alph",
				Audience: "example.org",
				Subject:  "someone@example.org",
				ExpirationTime: jwt.Timestamp(
					time.Date(2020, time.May, 24, 20, 35, 37, 0, time.UTC),
				),
			},
			expectedHeader: map[string]interface{}{
				"alg": "HS256",
				"typ": "JWT",
			},
			expectedPayload: map[string]interface{}{
				"iss": "alph",
				"aud": "example.org",
				"sub": "someone@example.org",
				"exp": float64(1590352537),
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.description, func(t *testing.T) {
			serializedToken, err := jwt.Serialize(scenario.token)
			require.NoError(t, err)

			tokenComponents := strings.SplitN(serializedToken, ".", 2)
			require.Len(t, tokenComponents, 2)

			header := deserializeToMap(t, tokenComponents[0])
			require.Equal(t, scenario.expectedHeader, header)

			payload := deserializeToMap(t, tokenComponents[1])
			require.Equal(t, scenario.expectedPayload, payload)

			deserializedToken, err := jwt.Deserialize(serializedToken)
			require.NoError(t, err)
			require.Equal(t, scenario.token, deserializedToken)
		})
	}
}

func deserializeToMap(t *testing.T, serializedComponent string) map[string]interface{} {
	componentJSON, err := base64.RawURLEncoding.DecodeString(serializedComponent)
	require.NoError(t, err)
	var component map[string]interface{}
	err = json.Unmarshal([]byte(componentJSON), &component)
	require.NoError(t, err)
	return component
}
