package sessions

import (
	"errors"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"
const paramAuthorization = "auth"
const schemeBearer = "Bearer "

//ErrNoSessionID is used when no session ID was found in the Authorization header
var ErrNoSessionID = errors.New("no session ID found in " + headerAuthorization + " header")

//ErrInvalidScheme is used when the authorization scheme is not supported
var ErrInvalidScheme = errors.New("authorization scheme not supported")

//BeginSession creates a new SessionID, saves the `sessionState` to the store, adds an
//Authorization header to the response with the SessionID, and returns the new SessionID
func BeginSession(signingKey string, store Store, sessionState interface{}, w http.ResponseWriter) (SessionID, error) {
	sid, err := NewSessionID(signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	err = Store.Save(store, sid, sessionState)
	if err != nil {
		return InvalidSessionID, err
	}
	w.Header().Add(headerAuthorization, schemeBearer+string(sid))
	return sid, nil
}

//GetSessionID extracts and validates the SessionID from the request headers
func GetSessionID(r *http.Request, signingKey string) (SessionID, error) {
	header := r.Header.Get(headerAuthorization)
	if header == "" {
		header = r.URL.Query().Get("auth")
	}
	if header == "" {
		return InvalidSessionID, ErrNoSessionID
	}
	headerSlice := strings.Split(header, " ")
	if headerSlice[0] != "Bearer" {
		return InvalidSessionID, ErrInvalidScheme
	}
	sessionID, err := ValidateID(headerSlice[1], signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	return sessionID, nil
}

//GetState extracts the SessionID from the request,
//gets the associated state from the provided store into
//the `sessionState` parameter, and returns the SessionID
func GetState(r *http.Request, signingKey string, store Store, sessionState interface{}) (SessionID, error) {
	sessionID, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	err = Store.Get(store, sessionID, sessionState)
	if err != nil {
		return InvalidSessionID, err
	}
	return sessionID, nil
}

//EndSession extracts the SessionID from the request,
//and deletes the associated data in the provided store, returning
//the extracted SessionID.
func EndSession(r *http.Request, signingKey string, store Store) (SessionID, error) {
	sessionID, err := GetSessionID(r, signingKey)
	if err != nil {
		return InvalidSessionID, err
	}
	err = Store.Delete(store, sessionID)
	if err != nil {
		return InvalidSessionID, err
	}
	return sessionID, nil
}
