package sessions

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"errors"
)

//InvalidSessionID represents an empty, invalid session ID
const InvalidSessionID SessionID = ""

//idLength is the length of the ID portion
const idLength = 32

//signedLength is the full length of the signed session ID
//(ID portion plus signature)
const signedLength = idLength + sha256.Size

//SessionID represents a valid, digitally-signed session ID.
//This is a base64 URL encoded string created from a byte slice
//where the first `idLength` bytes are crytographically random
//bytes representing the unique session ID, and the remaining bytes
//are an HMAC hash of those ID bytes (i.e., a digital signature).
//The byte slice layout is like so:
//+-----------------------------------------------------+
//|...32 crypto random bytes...|HMAC hash of those bytes|
//+-----------------------------------------------------+
type SessionID string

//ErrInvalidID is returned when an invalid session id is passed to ValidateID()
var ErrInvalidID = errors.New("Invalid Session ID")

//NewSessionID creates and returns a new digitally-signed session ID,
//using `signingKey` as the HMAC signing key. An error is returned only
//if there was an error generating random bytes for the session ID
func NewSessionID(signingKey string) (SessionID, error) {
	if len(signingKey) == 0 {
		return InvalidSessionID, ErrInvalidID
	}
	message := make([]byte, idLength)
	rand.Read(message)
	key := signingKey
	h := hmac.New(sha256.New, []byte(key))
	h.Write(message)
	signature := h.Sum(nil)
	hmac := append(message, signature...)
	return SessionID(b64.URLEncoding.EncodeToString(hmac)), nil
}

//ValidateID validates the string in the `id` parameter
//using the `signingKey` as the HMAC signing key
//and returns an error if invalid, or a SessionID if valid
func ValidateID(id string, signingKey string) (SessionID, error) {
	uDec, err := b64.URLEncoding.DecodeString(id)
	if err != nil {
		return InvalidSessionID, ErrInvalidID
	}
	message := uDec[0:idLength]
	hmacHash := uDec[idLength:]

	key := signingKey
	h := hmac.New(sha256.New, []byte(key))
	h.Write(message)
	signature := h.Sum(nil)

	if bytes.Equal(hmacHash, signature) {
		return SessionID(id), nil
	}
	return InvalidSessionID, ErrInvalidID
}

//String returns a string representation of the sessionID
func (sid SessionID) String() string {
	return string(sid)
}
