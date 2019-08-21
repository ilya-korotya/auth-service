package public

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/gospeak/auth-service/model"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gospeak/auth-service/api/server"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RegistrationUser(ctx server.Context, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// registration user
	user.Password = hash(user.Password)
	user.Token.Content = generateToken(63)
	user.Token.ExpiredAt = time.Now().Add(100 * time.Hour)
	u, err := ctx.DB.User.Set(user)
	if err != nil {
		log.Println("cannot save user to database:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// save user to cache not require
	ctx.Cache.Set(user.Token.Content, user.Token.ID, user.Token.ExpiredAt)
	d, err := json.Marshal(u)
	if err != nil {
		log.Println("cannot generate response:", err)
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

func generateToken(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
