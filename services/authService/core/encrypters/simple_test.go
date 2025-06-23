package encrypters

import (
	"log"
	"testing"

	"github.com/JuanGQCadavid/now-project/services/authService/core/core/domain"
	"github.com/JuanGQCadavid/now-project/services/authService/core/core/ports"
	"github.com/stretchr/testify/assert"
)

var (
	alteredToken     = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzbW9uIjoiZWIxNGZjZGEtMzU0YS00OTgzLWFhZmUtODYwYjQ3NDYzNGNmIiwidXNlcklkIjoiYjM0MzlmOTItNWQzZS00NThlLThlMDItZDIyMjAyNDVjNTI0IiwidXNlck5hbWUiOiJKdWFuR1FDYWRhdmlkIiwidXNlclBob25lIjoiKzM3MjUzOTU2NTgxIn0.KyZ2d3S55DVSD4oDMP4rAtRjWBxbH5QQ3GuMAAaCtJQ"
	validToken       = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uIjoiZWIxNGZjZGEtMzU0YS00OTgzLWFhZmUtODYwYjQ3NDYzNGNmIiwidXNlcklkIjoiYjM0MzlmOTItNWQzZS00NThlLThlMDItZDIyMjAyNDVjNTI0IiwidXNlck5hbWUiOiJKdWFuR1FDYWRhdmlkIiwidXNlclBob25lIjoiKzM3MjUzOTU2NTgxIn0.KyZ2d3S55DVSD4oDMP4rAtRjWBxbH5QQ3GuMAAaCtJQ"
	tokenUserDetails = &domain.UserDetails{
		SessionHash: "eb14fcda-354a-4983-aafe-860b474634cf",
		UserID:      "b3439f92-5d3e-458e-8e02-d2220245c524",
		Name:        "JuanGQCadavid",
		PhoneNumber: "+37253956581",
	}
	jwtKey = "DEFAULT"
)

func TestJWTClaims(t *testing.T) {
	encryp := NewSimpleEncrypt([]byte(jwtKey))
	userDetaild, err := encryp.DecodeJWTToken(validToken)
	assert.Nil(t, err)
	assert.EqualValuesf(t, userDetaild, tokenUserDetails, "Details are not equals")
	log.Println(userDetaild)
}

func TestAlteredToken(t *testing.T) {
	encryp := NewSimpleEncrypt([]byte(jwtKey))
	_, err := encryp.DecodeJWTToken(alteredToken)

	assert.Equal(t, err, ports.ErrBadFormatToken)
}
