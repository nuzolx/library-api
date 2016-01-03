package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nuzolx/library-api/src/Godeps/_workspace/src/github.com/gorilla/mux"

	"github.com/nuzolx/library-api/src/app"
	"github.com/nuzolx/library-api/src/db"
	"github.com/nuzolx/library-api/src/models"
	"github.com/nuzolx/library-api/src/utils"
)

func (s *Server) CreateUser(res http.ResponseWriter, req *http.Request) error {

	user := &models.User{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(user)
	if err != nil {
		return app.ErrBadRequest
	}

	if err := db.CreateUser(s.Context, user); err != nil {
		return err
	}
	return nil
}

func (s *Server) GetUser(res http.ResponseWriter, req *http.Request) error {

	vars := mux.Vars(req)

	user, err := db.GetUser(s.Context, vars["userName"])
	if err != nil {
		return err
	}

	return utils.Write(res, user)
}
