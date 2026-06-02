package main

import (
	"net/http"
	"strconv"
	"time"

	"charankurva.net/blogbook/internal/data"
	"github.com/julienschmidt/httprouter"
)

func (app *Application) CreateBlog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("creating blog ......."))
}
func (app *Application) GetBlogById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("Id")
	val, err := strconv.Atoi(id)

	blog := data.Blog{
		ID:         int64(val),
		Title:      "Learning go by building Blogsite",
		Poster:     "Http.blogimage.com",
		Content:    "we are going dive in designing the blog first by following the rules",
		Author:     27,
		Upvotes:    200,
		Downvotes:  10,
		CreatedAt:  time.Now(),
		CategoryID: 1,
		SubjectID:  2,
		TopicID:    1,
	}
	err = app.WriteJson(w, r, http.StatusOK, blog, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "got an unexpexted problwm in serving you the response", http.StatusInternalServerError)
	}

}
