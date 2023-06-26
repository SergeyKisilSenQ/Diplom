package server

import (
	"Diplom/go-final-dpo/conf"
	"Diplom/go-final-dpo/pkg/repo_result"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func Run() {
	conf.GetConf()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.MarshalIndent(repo_result.GetResult(), "", " ")
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Write(res)
	})
	fmt.Println("start server...")
	http.ListenAndServe("localhost:8282", r)

}
