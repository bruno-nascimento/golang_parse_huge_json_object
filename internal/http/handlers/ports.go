package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"test/internal/models"
	"test/internal/service"
)

type Ports struct {
	service service.PortService
}

func NewPortsHandler() *Ports {
	return &Ports{service: service.NewPort()}
}

func (p Ports) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	// https://stackoverflow.com/a/67326917
	dec := json.NewDecoder(req.Body)
	t, err := dec.Token()
	if err != nil {
		log.Printf("error : %s", err.Error())
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}
	if t != json.Delim('{') {
		msg := fmt.Sprintf("expected {, got %v", t)
		log.Printf(msg)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(msg))
		return
	}
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			log.Printf("error : %s", err.Error())
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
		key := t.(string)
		var m map[string]interface{}
		if err := dec.Decode(&m); err != nil {
			log.Printf("error : %s", err.Error())
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}

		port := models.NewPortFromMap(m)
		port.ID = key
		err = p.service.Add(port)
		if err != nil {
			log.Printf("error : %s", err.Error())
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
	}
	err = p.service.PortsBufferFlush()
	if err != nil {
		log.Printf("error : %s", err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusNoContent)
}
