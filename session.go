package session

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Session struct{
	UserID interface{}
	SessionID string
	ExpiryTime time.Time
}

func NewSession(duration uint64 ,user interface{})(*Session){
	return &Session{
		ExpiryTime: time.Now().Add(time.Second*time.Duration(duration)),
		SessionID:  uuid.NewString(),
		UserID: user,
	}
}

func (session *Session)SessionExpired()(bool){
	return !session.ExpiryTime.Before(time.Now())
}

type SessionManager struct{
	Sessions []*Session
}

func NewSessionManager()(*SessionManager){
	return &SessionManager{
		Sessions:[]*Session{},
	}
}

func (sessionManager * SessionManager) CreateSession(dur uint64 ,user interface{})(Session,error){
	sess,err:=sessionManager.UserActiveSession(user)
	if err==nil{
		return Session{},fmt.Errorf("user has an active session %v ",sess.SessionID)
	}
	session :=NewSession(dur,user)
	sessionManager.Sessions=append(sessionManager.Sessions,session)
	return *session,nil
}

func (sessionManager * SessionManager)NumberofActiveSessions()(int){
	return len(sessionManager.Sessions)
}

func (sessionManager * SessionManager)UserActiveSession(user interface{})( Session,error){
	for _,se := range sessionManager.Sessions{
		if se.UserID==user{
			return *se,nil
		}
	}
	return Session{},fmt.Errorf("user %v has no active session ",user)
}

func (sessionManager * SessionManager)SessionExist(sess string)(Session,error){
	for _,se := range sessionManager.Sessions{
		if se.SessionID==sess{
			return *se,nil
		}
	}
	return Session{},fmt.Errorf("session id %v not found ",sess)
}

func (sessionManager * SessionManager) DeleteSession(){
	for{
		for i,se := range sessionManager.Sessions{
			if se.SessionExpired(){
				sessionManager.Sessions=append(sessionManager.Sessions[:i],sessionManager.Sessions[i+1:]... )
			}
		}
	}
}

func (sessionManager * SessionManager) DeleteSessionByID(sess string)(Session,error){
	for _,se := range sessionManager.Sessions{
		if se.SessionID==sess{
			return *se,nil
		}
	}
	return Session{},fmt.Errorf("session id %v not found ",sess)	
}