package main

import (
	"html/template"
	"net/http"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
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

func (m *Manager) NewSessionID() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//	既存セッションの存在チェック
func (m *Manager) Exists(sessionID string) bool {
	_, r := m.database[sessionID]
	return r 
}

func (m *Manager) New(r *http.Request, cookieName string) (*Session, error) {
	//	Cookieを取得できるかどうか
	cookie, err := r.Cookie(cookieName)

	if err == nil && m.Exists(cookie.Value) {
		return nil, errors.New("sessionIDはすでに発行されています")
	}

	return nil, nil 
}

func count(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", count)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}