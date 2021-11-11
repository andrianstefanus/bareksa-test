package delivery

import (
	fx "bareksa-test/function"
	md "bareksa-test/model"
	st "bareksa-test/struck"
	"net/http"

	"github.com/gorilla/mux"
)

type TagsHandler struct {
	TagsUsecase md.TagsUsecase
}

func InitiateTagsHandler(route *mux.Router, tagsUsecase md.TagsUsecase) {
	handler := &TagsHandler{
		TagsUsecase: tagsUsecase,
	}

	route.HandleFunc("/tags/add", handler.Add).Methods("POST")
	route.HandleFunc("/tags/remove", handler.Remove).Methods("POST")
	route.HandleFunc("/tags/list", handler.List).Methods("POST")
}

func (h *TagsHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var
	var logs []st.Log
	var tagInput md.Tag

	// get params
	fx.GetBody(r.Body, &tagInput)

	// add
	response, _ := h.TagsUsecase.Add(ctx, tagInput, &logs)

	// response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fx.Marshal(response)))
}

func (h *TagsHandler) Remove(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var
	var logs []st.Log
	var tagInput md.Tag

	// get params
	fx.GetBody(r.Body, &tagInput)

	// remove
	response, _ := h.TagsUsecase.Remove(ctx, tagInput, &logs)

	// response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fx.Marshal(response)))
}

func (h *TagsHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var
	var logs []st.Log
	var tagInput md.Tag

	// get params
	fx.GetBody(r.Body, &tagInput)

	// add
	response, _ := h.TagsUsecase.List(ctx, tagInput, &logs)

	// response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fx.Marshal(response)))
}
