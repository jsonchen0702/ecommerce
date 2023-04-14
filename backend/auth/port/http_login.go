package port

import (
	"encoding/json"
	"net/http"

	"github.com/bkielbasa/go-ecommerce/backend/internal/https"
)

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Router       /auth/login [post]
// @Accept       json
// @Produce      json
// @Param user  body Client true "Client"
// @Failure      500  {object}  https.ErrorResponse
// @Failure      404  {object}  https.ErrorResponse
func (h HTTP) Login(w http.ResponseWriter, r *http.Request) {
	var c Client
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		https.BadRequest(w, "serialization-error", err.Error())
		return
	}

	sess, err := h.auth.Login(r.Context(), c.Username, c.Password)
	if err != nil {
		https.BadRequest(w, "serialization-error", err.Error())
		return
	}

	// TODO: add configuration of cookie expiration time
	cookie := http.Cookie{Name: "session_id", Value: sess.ID(), Expires: sess.ExpiresAt()}

	http.SetCookie(w, &cookie)
	https.NoContent(w)
}
