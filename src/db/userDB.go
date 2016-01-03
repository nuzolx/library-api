package db

import (
	"github.com/nuzolx/library-api/src/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"

	"github.com/nuzolx/library-api/src/app"
	"github.com/nuzolx/library-api/src/models"
)

func CreateUser(appC *app.AppContext, user *models.User) error {

	collection := appC.DB.C("users")
	if err := collection.Insert(user); err != nil {
		return err
	}
	return nil
}

func GetUser(appC *app.AppContext, userName string) (*models.User, error) {

	collection := appC.DB.C("users")

	var res *models.User
	if err := collection.Find(bson.M{"login": userName}).One(&res); err != nil {
		return nil, err
	}
	return res, nil
}
