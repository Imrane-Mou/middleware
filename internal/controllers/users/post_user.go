package users

import (
	"encoding/json"
	"net/http"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	svc "middleware/example/internal/services/users"
)

type createUserInput struct {
	Name string `json:"name"`
}

func PostUser(w http.ResponseWriter, r *http.Request) {

	var in createUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
			Message: "invalid JSON body",
		})
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	if in.Name == "" {
		body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
			Message: "name is required",
		})
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	user, err := svc.CreateUser(in.Name)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
}
