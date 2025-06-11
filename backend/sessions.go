package backend

import "net/http"

type Session struct {
	Token   string `json:"token"`
	UUID    string `json:"uuid"`
	Service string `json:"service"`
}

func (api *API) SessionsInfoHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) SessionsLoginHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) SessionsLogoutHandler(w http.ResponseWriter, r *http.Request) {

}
