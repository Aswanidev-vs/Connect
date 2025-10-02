package dashboard

import (
	"net/http"

	"github.com/Aswanidev-vs/Connect/handler"
	tpl "github.com/Aswanidev-vs/Connect/templates"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	claims, ok := handler.GetUserFromContext(r)
	if !ok {
		// This should rarely happen because middleware already redirects
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	email, _ := claims["email"].(string)
	username, _ := claims["username"].(string) // Optional if you store username in claims

	data := struct {
		Email    string
		Username string
	}{
		Email:    email,
		Username: username,
	}

	// Render dashboard template
	if tpl.Templates["dashboard"] != nil {
		err := tpl.Templates["dashboard"].Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Fallback: simple HTML if template not loaded
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Welcome to your dashboard, " + email + "</h1>"))
}
