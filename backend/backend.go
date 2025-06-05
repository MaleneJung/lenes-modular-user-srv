package backend

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Session struct {
	Token string     `json:"token"`
	User  ClientUser `json:"user"`
}

type ServerUser struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Rank     string `json:"rank"`
	EMail    string `json:"email"`
	PassHash string `json:"passhash"`
}

type ClientUser struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Rank string `json:"rank"`
}

func (user *ServerUser) ToClientUser() ClientUser {
	return ClientUser{
		UUID: user.UUID,
		Name: user.Name,
		Rank: user.Rank,
	}
}

type ServerRegistration struct {
	User            ServerUser `json:"user"`
	ValidationToken string     `json:"validation"`
}

type ClientRegistration struct {
	User            ClientUser `json:"user"`
	EMail           string     `json:"email"`
	ValidationToken string     `json:"validation"`
}

func (reg *ServerRegistration) ToClientRegistration() ClientRegistration {
	return ClientRegistration{
		User:            reg.User.ToClientUser(),
		EMail:           reg.User.EMail,
		ValidationToken: reg.ValidationToken,
	}
}

type API struct {
	Users         map[string]ServerUser         `json:"users"`
	Sessions      map[string]Session            `json:"sessions"`
	Registrations map[string]ServerRegistration `json:"registrations"`
}

func NewAPI() API {
	return API{
		Users:         make(map[string]ServerUser),
		Sessions:      make(map[string]Session),
		Registrations: make(map[string]ServerRegistration),
	}
}

func (api *API) LoadFromFile() error {

	_, err := os.Stat("backend.json")
	if errors.Is(err, os.ErrNotExist) {
		log.Println("Backend wasn't created, yet!")
		return nil
	}

	fileBuffer, err := os.ReadFile("backend.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileBuffer, api); err != nil {
		return err
	}

	return nil

}

func (api *API) SaveToFile() error {

	fileBuffer, err := json.Marshal(api)
	if err != nil {
		return err
	}

	if err := os.WriteFile("backend.json", fileBuffer, 0644); err != nil {
		return err
	}

	return nil

}

func (api *API) RegisterRouteHandlers(serveMux *http.ServeMux) {

}
