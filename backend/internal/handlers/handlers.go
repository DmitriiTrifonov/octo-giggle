package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/DmitriiTrifonov/octo-giggle/backend/internal/model"
)

// Command stores a command number.
var Command uint64 //nolint:gochecknoglobals // no need to check globals for now

// DataHandler is a struct for POST request handling.
type DataHandler struct{}

// ServeHTTP implements Handler interface.
func (dh *DataHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		rw.Header().Add("Allow", http.MethodPost)
		http.Error(rw, "Method "+request.Method+" is not allowed", http.StatusMethodNotAllowed)
		return
	}
	bytes, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var req model.Request
	err = json.Unmarshal(bytes, &req)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	atomic.AddUint64(&Command, 1)
	resp := model.Response{
		Config:  model.Config{},
		Command: atomic.LoadUint64(&Command),
	}
	out, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = rw.Write(out)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
