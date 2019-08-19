package inner

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gospeak/auth-service/api/server"
	"github.com/gospeak/auth-service/postgres"
)

// GetUser return user
func GetUser(ctx server.Context, w http.ResponseWriter, r *http.Request) {
	// if you want use this header, wrap this method in token middleware(refactor this plz)
	t := r.Header.Get("token")
	f := &postgres.UserFilter{}
	f.ByToken(t)
	user, err := ctx.DB.User.Get(f)
	if err != nil {
		log.Println("database error:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
