package domain

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (dh *DomainHandler) ShowDomainHandlerFunc(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		dh.responser.Error(w, "Id is required", http.StatusBadRequest)
	}

	uid, uuidErr := uuid.Parse(id)

	if uuidErr != nil {
		dh.responser.Error(w, uuidErr.Error(), http.StatusBadRequest)

		return
	}

	domain, repoErr := dh.repo.FindByID(uid)

	if repoErr != nil {
		dh.responser.Error(w, repoErr.Error(), http.StatusBadRequest)

		return
	}

	if err := dh.responser.JSON(w, domain, http.StatusOK); err != nil {
		dh.responser.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
