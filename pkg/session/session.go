package session

import (
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	uuid "github.com/satori/go.uuid"
)

//ISessionService encapulates session flash messages
type ISessionService interface {
	SetFlash(r *http.Request, w http.ResponseWriter, message string, alertType string) error
	GetFlash(r *http.Request, w http.ResponseWriter) (*Flash, error)
	CreateSession(r *http.Request, w http.ResponseWriter, userID int) (string, error)
	GetSession(r *http.Request) (*UserSession, error)
	DeleteSession(r *http.Request, w http.ResponseWriter) error
	UpdateSession(r *http.Request, w http.ResponseWriter) error
	CreateSessionValue(r *http.Request, w http.ResponseWriter, key string, value interface{}) error
	GetSessionValue(r *http.Request, key string) (interface{}, error)
	DeleteSessionValue(r *http.Request, w http.ResponseWriter, key string) error
	UpdateSessionValue(r *http.Request, w http.ResponseWriter, key string) error
}

//CookieStoreServiceImpl provides session flash structure
type CookieStoreServiceImpl struct {
	Store  sessions.Store
	Secure bool
}

//Flash wrapper for message and alert type
type Flash struct {
	Message string
	Type    string
}

//UserSession provides user session
type UserSession struct {
	UserID    int
	SessionID string
}

//ZohoToken wrapper for zoho tokens
type ZohoToken struct {
	Access  string
	Refresh string
}

//SetFlash set flash message
func (fl *CookieStoreServiceImpl) SetFlash(r *http.Request, w http.ResponseWriter, message string, alertType string) error {

	//register the structure encoding/gob knows about it
	gob.Register(&Flash{})

	session, err := fl.Store.Get(r, `goreact_session_flash`)
	if err != nil {
		return err
	}

	//save flash structure in session
	flash := &Flash{
		Message: message,
		Type:    alertType,
	}
	session.AddFlash(flash, `flash`)

	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

//GetFlash get flash message by reading session store values
//Delete the session after session flash message has been read
func (fl *CookieStoreServiceImpl) GetFlash(r *http.Request, w http.ResponseWriter) (*Flash, error) {

	//register the structure encoding/gob knows about it
	gob.Register(&Flash{})

	session, err := fl.Store.Get(r, `goreact_session_flash`)
	if err != nil {
		return nil, err
	}
	flashes := session.Flashes(`flash`)
	if len(flashes) > 0 {
		flash := flashes[0].(*Flash)

		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			return nil, err
		}
		return flash, nil
	}
	return nil, nil
}

//CreateSession creates session for a user
//Creates random uuid for a user, prepares a user session structure and save it to a session
//User will be authenticated based by user session
func (fl *CookieStoreServiceImpl) CreateSession(r *http.Request, w http.ResponseWriter, userID int) (string, error) {
	session, _ := fl.Store.Get(r, `goreact_session_user`)

	//register the structure encoding/gob knows about it
	gob.Register(&UserSession{})

	sessionID, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	//set session configurations
	session.Options.MaxAge = 1800
	// session.Options.HttpOnly = true
	// session.Options.Secure = fl.Secure

	session.Values[`session-user-auth`] = &UserSession{
		UserID:    userID,
		SessionID: sessionID.String(),
	}

	err = session.Save(r, w)
	if err != nil {
		return "", err
	}
	return sessionID.String(), nil
}

//GetSession get session for a user
//Returns user session structure with user id and session id if session is found
func (fl *CookieStoreServiceImpl) GetSession(r *http.Request) (*UserSession, error) {
	session, err := fl.Store.Get(r, `goreact_session_user`)
	if err != nil {
		return nil, errors.New(`Session Not Available`)
	}
	if len(session.Values) > 0 {
		return session.Values[`session-user-auth`].(*UserSession), nil
	}
	return nil, errors.New(`Session Not Found`)
}

//DeleteSession deletes the user's session by setting max age value to -1
func (fl *CookieStoreServiceImpl) DeleteSession(r *http.Request, w http.ResponseWriter) error {
	session, _ := fl.Store.Get(r, `goreact_session_user`)
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

//UpdateSession updates session by adding more 30 minutes to session
func (fl *CookieStoreServiceImpl) UpdateSession(r *http.Request, w http.ResponseWriter) error {
	session, err := fl.Store.Get(r, `goreact_session_user`)
	if err != nil {
		return err
	}
	session.Options.MaxAge = 1800
	session.Options.HttpOnly = true
	session.Options.Secure = fl.Secure
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

//CreateSessionValue creates session for a storing any data
func (fl *CookieStoreServiceImpl) CreateSessionValue(r *http.Request, w http.ResponseWriter, key string, value interface{}) error {
	session, _ := fl.Store.Get(r, fmt.Sprintf(`goreact_session_%s`, key))

	//register the structure encoding/gob knows about it
	gob.Register(value)

	//set session configurations
	session.Options.MaxAge = 1800
	// session.Options.HttpOnly = true
	// session.Options.Secure = fl.Secure

	session.Values[`session-user-value`] = value

	if err := session.Save(r, w); err != nil {
		return err
	}
	return nil
}

//GetSessionValue get session for a user
//Returns user session structure with user id and session id if session is found
func (fl *CookieStoreServiceImpl) GetSessionValue(r *http.Request, key string) (interface{}, error) {
	session, err := fl.Store.Get(r, fmt.Sprintf(`goreact_session_%s`, key))
	if err != nil {
		return nil, errors.New(`Session Not Available`)
	}
	if len(session.Values) > 0 {
		return session.Values[`session-user-value`], nil
	}
	return nil, errors.New(`Session Not Found`)
}

//DeleteSessionValue deletes the user's session by setting max age value to -1
func (fl *CookieStoreServiceImpl) DeleteSessionValue(r *http.Request, w http.ResponseWriter, key string) error {
	session, _ := fl.Store.Get(r, fmt.Sprintf(`goreact_session_%s`, key))
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

//UpdateSessionValue updates session by adding more 30 minutes to session
func (fl *CookieStoreServiceImpl) UpdateSessionValue(r *http.Request, w http.ResponseWriter, key string) error {
	session, err := fl.Store.Get(r, fmt.Sprintf(`goreact_session_%s`, key))
	if err != nil {
		return err
	}
	session.Options.MaxAge = 1800
	// session.Options.HttpOnly = true
	// session.Options.Secure = fl.Secure
	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}
