package backend

import "net/http"

type BackUser struct {
	UUID        string   `json:"uuid"`
	Tag         string   `json:"tag"`
	Name        string   `json:"name"`
	Rank        int      `json:"rank"`
	EMail       string   `json:"email"`
	PassHash    string   `json:"passhash"`
	Permissions []string `json:"permissions"`
}

type FrontUser struct {
	UUID        string   `json:"uuid"`
	Tag         string   `json:"tag"`
	Name        string   `json:"name"`
	Rank        int      `json:"rank"`
	Permissions []string `json:"permissions"`
}

func (user *BackUser) ToFrontUser() FrontUser {
	return FrontUser{
		UUID:        user.UUID,
		Name:        user.Name,
		Rank:        user.Rank,
		Permissions: user.Permissions,
	}
}

func (api *API) UsersInfoHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) UsersPermissionsHasHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) UsersPermissionsGrantHandler(w http.ResponseWriter, r *http.Request) {

}

func (api *API) UsersPermissionsRevokeHandler(w http.ResponseWriter, r *http.Request) {

}
