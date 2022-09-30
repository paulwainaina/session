package session

import (
	"testing"
	"time"
)



func TestSessionCreation(t *testing.T){
	sessionManager :=NewSessionManager()
	x:=sessionManager.NumberofActiveSessions()
	if x!=0{
		t.Errorf("expected 0 sessions received %v ",x)
	}
	session,_:=sessionManager.CreateSession(1000,"test")
	x=sessionManager.NumberofActiveSessions()
	if x!=1{
		t.Errorf("expected 1 sessions received %v ",x)
	}
	if session.UserID!="test"{
		t.Errorf("expected \"test\" sessions received %v ",session.UserID)
	}
	_,err:=sessionManager.CreateSession(1000,"test")
	if err==nil{
		t.Errorf("expected error due to another session")
	}
}

func TestSessionDelete(t *testing.T){
	sessionManager :=NewSessionManager()
	sessionManager.CreateSession(1000,"test1")
	sessionManager.CreateSession(10000,"test2")
	sessionManager.CreateSession(2000,"test3")
	x:=sessionManager.NumberofActiveSessions()
	if x!=3{
		t.Errorf("expected 3 sessions received %v ",x)
	}
	time.Sleep(time.Second* time.Duration(3000))
	
	go sessionManager.DeleteSession()
	
	x=sessionManager.NumberofActiveSessions()
	if x!=1{
		t.Errorf("expected 1 sessions received %v ",x)
	}
	time.Sleep(time.Second* time.Duration(7000))
	x=sessionManager.NumberofActiveSessions()
	if x!=0{
		t.Errorf("expected 0 sessions received %v ",x)
	}
}

func TestSessionDeleteByID(t *testing.T){
	sessionManager :=NewSessionManager()
	sessionManager.CreateSession(1000,"test1")
	sessionManager.CreateSession(1000,"test2")
	sessionManager.CreateSession(1000,"test1")
	x:=sessionManager.NumberofActiveSessions()
	if x!=2{
		t.Errorf("expected 2 sessions received %v ",x)
	}
	sess,err:=sessionManager.UserActiveSession("test1")
	if err!=nil{
		t.Errorf("expected an active session")
	}
	_,err=sessionManager.SessionExist(sess.SessionID)
	if err!=nil{
		t.Errorf("expected an session to exist")
	}
}