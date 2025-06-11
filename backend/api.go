package backend

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
)

type API struct {
	Users         map[string]BackUser         `json:"users"`
	Sessions      map[string]Session          `json:"sessions"`
	Registrations map[string]BackRegistration `json:"registrations"`
	Ranks         map[uint]Rank               `json:"ranks"`
}

func NewAPI() API {
	return API{
		Users:         make(map[string]BackUser),
		Sessions:      make(map[string]Session),
		Registrations: make(map[string]BackRegistration),
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

func CheckPermission(permissions []string, permission string) bool {

	for _, p := range permissions {
		if wildcardPerm, isWildcardPerm := strings.CutSuffix(p, ".*"); isWildcardPerm {
			return strings.HasPrefix(permission, wildcardPerm)
		} else if p == permission {
			return true
		}
	}

	return false

}

func (api *API) RegisterRouteHandlers(serveMux *http.ServeMux) {

	serveMux.HandleFunc("POST /api/registrations/new/{tag}/{name}/{email}/{passhash}", api.RegistrationsNewHandler)
	serveMux.HandleFunc("GET /api/registrations/validate/{validation}", api.RegistrationsNewHandler)
	serveMux.HandleFunc("GET /api/registrations/list/{session}", api.RegistrationsListHandler)

	serveMux.HandleFunc("GET /api/sessions/info/{session}", api.SessionsInfoHandler)
	serveMux.HandleFunc("POST /api/sessions/login/{service}/{tag}/{passhash}", api.SessionsLoginHandler)
	serveMux.HandleFunc("POST /api/sessions/logout/{session}", api.SessionsLogoutHandler)

	serveMux.HandleFunc("GET /api/users/info/{tag}", api.UsersInfoHandler)
	serveMux.HandleFunc("GET /api/users/permissions/has/{tag}/{permission}", api.UsersPermissionsHasHandler)
	serveMux.HandleFunc("POST /api/users/permissions/grant/{session}/{tag}/{permission}", api.UsersPermissionsGrantHandler)
	serveMux.HandleFunc("POST /api/users/permissions/revoke/{session}/{tag}/{permission}", api.UsersPermissionsRevokeHandler)

	serveMux.HandleFunc("GET /api/ranks/list", api.RanksListHandler)
	serveMux.HandleFunc("GET /api/ranks/info/{id}", api.RanksInfoHandler)
	serveMux.HandleFunc("POST /api/ranks/new/{session}/{id}/{name}", api.RanksNewHandler)
	serveMux.HandleFunc("GET /api/ranks/permissions/has/{session}/{permission}", api.RanksPermissionsGrantHandler)
	serveMux.HandleFunc("POST /api/ranks/permissions/grant/{session}/{id}/{permission}", api.RanksPermissionsGrantHandler)
	serveMux.HandleFunc("POST /api/ranks/permissions/revoke/{session}/{id}/{permission}", api.RanksPermissionsRevokeHandler)

}
