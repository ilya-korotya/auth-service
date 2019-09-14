package public

import (
	"encoding/json"
	"github.com/gospeak/auth-service/api/server"
	"github.com/gospeak/auth-service/model"
	"github.com/gospeak/auth-service/postgres"
	"log"
	"net/http"
	"time"
)

func LoginUser(ctx server.Context, w http.ResponseWriter, r *http.Request) {
	u := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	f := &postgres.UserFilter{}
	f.ByName(u.Name)
	storeUser, err := ctx.DB.User.Get(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if storeUser.Password != hash(u.Password) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// generate new token for user
	t := &model.Token{}
	t.Content = generateToken(63)
	t.UserID = storeUser.ID
	t.ExpiredAt = time.Now().Add(100 * time.Hour)
	// save new token to database
	t, err = ctx.DB.Token.Set(t)
	if err != nil {
		log.Println("cannot insert new token to database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// set new token to cache not require
	ctx.Cache.Set(t.Content, t.UserID, t.ExpiredAt)
	d, err := json.Marshal(t)
	if err != nil {
		log.Println("cannot create response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(d)
}
