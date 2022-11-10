package main

import (
	"Diplom/go-final-dpo/pkg/conf"
	"Diplom/go-final-dpo/pkg/repo_country"
	"Diplom/go-final-dpo/pkg/repo_email"
	"Diplom/go-final-dpo/pkg/repo_result"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	conf.GetConf()
	ED := repo_email.NewStorageEmail()
	ED.ReadFileEmail()
	NSC := repo_country.CreateNewCountryStorage()
	repo_country.ReadFile(NSC)
	for i := range ED {
		fmt.Println(ED[i])
	}
	fmt.Println(repo_email.SortedEmailData(ED))
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
	http.ListenAndServe("localhost:8282", r)
}
