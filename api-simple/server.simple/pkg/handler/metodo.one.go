package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func MetodoOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	vars := mux.Vars(r)
	if len(vars["id"]) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"Error Id é obrigatorio!"}`))
		return
	}

	jsonserver := `
	{"widget": {
	    "debug": "on",
	    "window": {
	        "title": "Sample Konfabulator Widget",
	        "name": "main_window",
	        "width": 500,
	        "height": 500
	    },
	    "image": { 
	        "src": "Images/Sun.png",
	        "name": "sun1",
	        "hOffset": 250,
	        "vOffset": 250,
	        "alignment": "center"
	    },
	    "text": {
	        "data": "Click Here",
	        "size": 36,
	        "style": "bold",
	        "name": "text1",
	        "hOffset": 250,
	        "vOffset": 100,
	        "alignment": "center",
	        "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
	    }
	}}`

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonserver))
	return
}
