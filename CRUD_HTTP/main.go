package main

import (
	"encoding/json"
	"examples/httpServer/message"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	addr     string
	messages *message.PackageMessage
}

func NewServer(host, port string) *Server {
	messages := message.New()

	return &Server{
		addr:     host + ":" + port,
		messages: messages,
	}
}

func main() {
	router := mux.NewRouter()
	server := NewServer("localhost", "8181")
	router.HandleFunc("/create", server.createHandler).Methods("POST")
	router.HandleFunc("/update", server.updateHandler).Methods("POST")
	router.HandleFunc("/get", server.getHandler).Methods("GET")
	router.HandleFunc("/delete", server.deleteHandler).Methods("DELETE")
	router.HandleFunc("/", server.mainHandler).Methods("GET")

	err := http.ListenAndServe("localhost:8181", router)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) mainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	w.Write([]byte("Hello World!"))
}

// curl -iX POST localhost:8181/create -d '{"id":123, "dt":"2021-10-19T20:00:00.00+03:00", "msg":"tst"}'
func (s *Server) createHandler(w http.ResponseWriter, req *http.Request) {
	record := message.Message{}
	err := json.NewDecoder(req.Body).Decode(&record)
	if err != nil {
		fmt.Println(err)
	}
	replay := []string{}
	if record.Msg == "" {
		replay = append(replay, "message cannot be empty")
	}
	if (record.DT == time.Time{}) {
		replay = append(replay, "time cannot be empty")
	}

	if len(replay) == 0 {
		s.messages.Create(record.Msg, record.DT)
		replay = append(replay, "message added to database")
	}
	w.Write([]byte(strings.Join(replay, "\n") + "\n"))
}

func (s *Server) updateHandler(w http.ResponseWriter, req *http.Request) {
	record := message.Message{}
	err := json.NewDecoder(req.Body).Decode(&record)
	if err != nil {
		fmt.Println(err)
	}
	replay := []string{}
	if record.Id == 0 {
		replay = append(replay, "id cannot be empty")
	}
	if record.Msg == "" {
		replay = append(replay, "message cannot be empty")
	}
	if (record.DT == time.Time{}) {
		replay = append(replay, "time cannot be empty")
	}

	if len(replay) == 0 {
		s.messages.Update(record.Id, record.Msg, record.DT)
		replay = append(replay, "message updated in database")
	}

	w.Write([]byte(strings.Join(replay, "\n") + "\n"))
}

func (s *Server) getHandler(w http.ResponseWriter, req *http.Request) {
	record := message.Message{}
	err := json.NewDecoder(req.Body).Decode(&record)
	if err != nil {
		fmt.Println(err)
	}
	replay := []string{}
	if record.Id == 0 {
		replay = append(replay, "id cannot be empty")
	}

	if len(replay) == 0 {
		m, err := s.messages.Get(record.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			replay = append(replay, err.Error())
		} else {
			replay = append(replay, m.DT.String()+" "+m.Msg)
		}
	}

	w.Write([]byte(strings.Join(replay, "\n") + "\n"))
}

func (s *Server) deleteHandler(w http.ResponseWriter, req *http.Request) {
	record := message.Message{}
	err := json.NewDecoder(req.Body).Decode(&record)
	if err != nil {
		fmt.Println(err)
	}
	replay := []string{}
	if record.Id == 0 {
		replay = append(replay, "id cannot be empty")
	}

	if len(replay) == 0 {
		err := s.messages.Delete(record.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			replay = append(replay, err.Error())
		}
	}

	w.Write([]byte(strings.Join(replay, "\n") + "\n"))
}
