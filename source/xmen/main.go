package main

import (
	//"fmt"
	//"log"
	"encoding/json"
	"net/http"
	"./xmenrecluting"
	"./person"
)



func Server(w http.ResponseWriter, r *http.Request) {

	var p person.Person
	err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }


	js := xmenrecluting.IsMutant(p)
	

	if js {
		w.WriteHeader(200)
		w.Write([]byte("HTTP 200-OK"))
		HTTP 200-OK
	}else{
		w.WriteHeader(http.StatusForbidden)
        w.Write([]byte("403 Forbidden"))
	}


	//w.Header().Set("Content-Type", "application/json")
  	//w.Write([]byte(stringResponse))
	
}

func main() {
	http.HandleFunc("/", Server)
	http.ListenAndServe(":3009", nil)
}
