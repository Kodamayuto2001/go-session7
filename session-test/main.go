package main

import (
	"html/template"
	"net/http"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"fmt"
)

type Manager struct {
	database map[string]interface{}
}

type Session struct {
	cookieName	string
	ID 			string
	manager 	*Manager
	request 	*http.Request
	writer 		http.ResponseWriter
	Values		map[string]interface{}
}

var mg Manager

func NewManager() *Manager {
	return &mg
}

//	セッションIDの発行
func (m *Manager) NewSessionID() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//	セッションの開始
func NewSession(manager *Manager, cookieName string) *Session {
	return &Session{
		cookieName:	cookieName,
		ID:			"",
		manager:	manager,
		request:	nil,
		writer:		nil,
		Values:		map[string]interface{}{},
	}
}

//	既存セッションの存在チェック
func (m *Manager) Exists(sessionID string) bool {
	_, r := m.database[sessionID]
	return r 
}

//	新規セッションの開始
func (m *Manager) New(r *http.Request, cookieName string) (*Session, error) {
	//	Cookieを取得できるかどうか
	cookie, err := r.Cookie(cookieName)

	if err == nil && m.Exists(cookie.Value) {
		return nil, errors.New("sessionIDはすでに発行されています")
	}

	session := NewSession(m, cookieName)
	session.ID = m.NewSessionID()
	session.request = r 
	return session, nil 
}

//	セッションの保存
func (s *Session) Save() error {
	return s.manager.Save(s.request, s.writer, s)
}

//	セッション名の取得
func (s *Session) Name() string {
	return s.cookieName
}

//	セッション変数値の取得
func (s *Session) Get(key string) (interface{}, bool) {
	ret, exists := s.Values[key]
	return ret, exists 
}

//	セッション変数値のセット
func (s *Session) Set(key string, val interface{}) {
	s.Values[key] = val 
}

//	セッション変数の削除
func (s *Session) Delete(key string) {
	delete(s.Values, key)
}

//	セッションの削除
func (s *Session) Terminate() {
	s.manager.Destroy(s.ID)
}

//	セッション情報の保存
func (m *Manager) Save(r *http.Request, w http.ResponseWriter, session *Session) error {
	m.database[session.ID] = session

	c := &http.Cookie {
		Name:	session.Name(),
		Value:	session.ID,
		Path:	"/",
	}

	http.SetCookie(session.writer, c)

	return nil 
}

//	既存セッションの取得
func (m *Manager) Get(r *http.Request, cookieName string) (*Session, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err 
	}

	sessionID := cookie.Value 

	buffer, exists := m.database[sessionID]
	if !exists {
		return nil, errors.New("無効なセッションIDです")
	}

	session := buffer.(*Session)
	session.request = r 
	return session, nil 
}

//	セッションの破棄
func (m *Manager) Destroy(sessionID string) {
	delete(m.database, sessionID)
}

func count(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	
	manager := NewManager()
	var session *Session 
	var err error 

	cookieName := "counter"

	//	既存のセッションの取得
	session, err = manager.Get(r, cookieName)	
	if err != nil {
		//	新規セッション作成
		session, err = manager.New(r, cookieName)
		if err != nil {
			fmt.Println(err.Error())
		}
		session.Set("countnum", 1)
	} else {
		cnt, _ := session.Get("countnum")
		session.Set("countnum", cnt.(int) + 1)
	}

	session.writer = w 
	session.request = r 
	session.manager = manager
	session.manager.database = map[string]interface{}{}
	session.Save()
	tpl.Execute(w, session)
}

func main() {
	http.HandleFunc("/", count)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}