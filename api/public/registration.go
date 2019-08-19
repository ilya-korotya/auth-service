package public

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/gospeak/auth-service/model"
	"net/http"

	"github.com/gospeak/auth-service/api/server"
)

func RegistrationUser(ctx server.Context, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user.Password = hash(user.Password)
	u, err := ctx.DB.User.Set(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	d, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(d)
}

func hash(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return hex.EncodeToString(h.Sum(nil))
}
