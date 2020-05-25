package server

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)
//criando um novo servidor http
//addr é o endereço onde o servidor será rodado
func NewHTTPServer(addr string) *http.Server{
	//instancia um novo servidor HTTP
	httprsv := newHTTPServer()
	//cria manipulador r
	r := mux.NewRouter()
	//marca o manipulador de produção com método POST que adiciona as gravações ao log
	r.HandleFunc("/", httprsv.handleProduce).Methods("POST")
	//marca o manipulador de consumo com método GET QUE retira as gravações do log
	r.HandleFunc("/", httprsv.handleProduce).Methods("GET"))
	return &http.Server{
		Addr: addr,
		Handler: r, 
	}
}

type httpServer struct {
	Log *Log 
}

func newHTTPServer() *httpServer{
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record Record `json: "record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct{
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

func (s *httpServer) handleProduce (w http.ResponseWriter, r *http.Request) {
	var req ProduceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	off, error := s.Log.Append(req.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ProduceResponse{Offset: off}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(),  http.StatusInternalServerError)
		return 
	}
}

func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request){
	var req ConsumeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	record, err := Log.Read(req.Offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	res := ConsumeResponse{Record: record}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}