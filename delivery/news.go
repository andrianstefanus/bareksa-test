package delivery

import (
	fx "bareksa-test/function"
	md "bareksa-test/model"
	st "bareksa-test/struck"
	"net/http"

	"github.com/gorilla/mux"
)

type NewsHandler struct {
	NewsUsecase md.NewsUsecase
}

func InitiateNewsHandler(route *mux.Router, newsUsecase md.NewsUsecase) {
	handler := &NewsHandler{
		NewsUsecase: newsUsecase,
	}

	route.HandleFunc("/news/add", handler.Add).Methods("POST")
	route.HandleFunc("/news/update", handler.Update).Methods("POST")
	route.HandleFunc("/news/remove", handler.Remove).Methods("POST")
	route.HandleFunc("/news/list", handler.List).Methods("POST")
}

func (h *NewsHandler) Add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var
	var logs []st.Log
	var newsInput md.News

	// get params
	fx.GetBody(r.Body, &newsInput)

	// add
	response, _ := h.NewsUsecase.Add(ctx, newsInput, &logs)

	// response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fx.Marshal(response)))
}

func (h *NewsHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var
	var logs []st.Log
	var newsInput md.News

	// get params
	fx.GetBody(r.Body, &newsInput)

	// update
	response, _ := h.NewsUsecase.Update(ctx, newsInput, &logs)

	// response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fx.Marshal(response)))
}

func (h *NewsHandler) Remove(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var
	var logs []st.Log
	var newsInput md.News

	// get params
	fx.GetBody(r.Body, &newsInput)

	// remove
	response, _ := h.NewsUsecase.Remove(ctx, newsInput, &logs)

	// response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fx.Marshal(response)))
}

func (h *NewsHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// var
	var logs []st.Log
	var newsInput md.NewsFilter

	// get params
	fx.GetBody(r.Body, &newsInput)

	// add
	response, _ := h.NewsUsecase.List(ctx, newsInput, &logs)

	// response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fx.Marshal(response)))
}
