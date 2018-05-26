package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	repo "github.com/roshniashok/crud/repo"
	model "github.com/roshniashok/crud/model"
)


func GetRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK) //setting status
	//json encoding (marshaling)
	if err := json.NewEncoder(w).Encode(repo.Callrecords); err != nil {
		panic(err)
	}
}

func PostRecord(w http.ResponseWriter, r *http.Request) {
	var callrecord model.CallRecord
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	//reading body, unmarshaled to todo struct
	if err := json.Unmarshal(body, &callrecord); err != nil {
		//  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//  w.WriteHeader(422) // unprocessable entity
		//encoding json to post data via w
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}

	}

	t := repo.RepoCreateRecord(callrecord)
	//  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//  w.WriteHeader(http.StatusCreated)
	ro := json.NewEncoder(w).Encode(t) //??
	fmt.Println(ro)

	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
