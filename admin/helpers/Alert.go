package helpers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func SetAlert(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return err
	}
}
