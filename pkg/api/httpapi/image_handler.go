package httpapi

import (
	"encoding/json"
	"github.com/gorilla/mux"
	command "github.com/mdapathy/imageuploader/pkg/domain/cmd"
	"github.com/mdapathy/imageuploader/pkg/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type imagesHandler struct {
	*resources
}

func (h *imagesHandler) List() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		query, err := h.queries.NewList(NewQueryList(r, 5))
		if err != nil {
			return handleErr(err)
		}

		images, err := h.images.List(r.Context(), query)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(images)
	}
}

func (h *imagesHandler) Details() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {

		id, ok := mux.Vars(r)["id"]
		if !ok {
			return handleErr(model.ErrIDInvalid)
		}
		queryDTO := QueryDetail{
			ID:     id,
			UserID: UserIDFromContext(r.Context()),
		}

		detail, err := h.queries.NewDetail(&queryDTO)
		if err != nil {
			return handleErr(err)
		}
		image, err := h.images.Details(r.Context(), detail)
		if err != nil {
			return handleErr(err)
		}

		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(image)
	}
}

func (h *imagesHandler) Create() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req CreateImageRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return ErrInvalidRequestBody
		}

		cmd := command.Create{
			Content: req.Content,
			UserID:  UserIDFromContext(r.Context()),
		}

		if err := h.images.Create(r.Context(), &cmd); err != nil {
			return handleErr(err)
		}

		w.WriteHeader(http.StatusCreated)
		return nil
	}
}

func (h *imagesHandler) Delete() HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id, ok := mux.Vars(r)["id"]
		if !ok {
			return handleErr(model.ErrIDRequired)
		}
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return handleErr(model.ErrIDInvalid)
		}

		cmd := command.Delete{
			ID:     objectID,
			UserID: UserIDFromContext(r.Context()),
		}

		if err = h.images.Delete(r.Context(), &cmd); err != nil {
			return handleErr(err)
		}

		w.WriteHeader(http.StatusOK)
		return nil
	}
}
