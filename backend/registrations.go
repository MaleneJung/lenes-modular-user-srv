package backend

import "net/http"

type BackRegistration struct {
	Tag             string `json:"tag"`
	Name            string `json:"name"`
	EMail           string `json:"email"`
	PassHash        string `json:"passhash"`
	ValidationToken string `json:"validation"`
}

type FrontRegistration struct {
	Tag             string `json:"tag"`
	Name            string `json:"name"`
	EMail           string `json:"email"`
	ValidationToken string `json:"validation"`
}

func (reg *BackRegistration) ToFrontRegistration() FrontRegistration {
	return FrontRegistration{
		Tag:             reg.Tag,
		Name:            reg.Name,
		EMail:           reg.EMail,
		ValidationToken: reg.ValidationToken,
	}
}

func (api *API) RegistrationsNewHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) RegistrationsValidateHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) RegistrationsListHandler(w http.ResponseWriter, r *http.Request) {

}
