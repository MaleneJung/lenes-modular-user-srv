package backend

import "net/http"

type Rank struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (api *API) RanksListHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) RanksInfoHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) RanksNewHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) RanksPermissionsGrantHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) RanksPermissionsRevokeHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) RanksPermissionsHasHandler(w http.ResponseWriter, r *http.Request) {

}
