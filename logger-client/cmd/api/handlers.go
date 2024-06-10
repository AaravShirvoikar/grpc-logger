package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/AaravShirvoikar/logger-client/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Log(w http.ResponseWriter, r *http.Request) {
	var payload LogPayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		errorJSON(w, err)
		return
	}

	app.LogEvent(w, r, payload)
}

func (app *Config) LogEvent(w http.ResponseWriter, r *http.Request, p LogPayload) {
	conn, err := grpc.Dial("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		errorJSON(w, err)
		return
	}
	defer conn.Close()

	c := logs.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.WriteLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name: p.Name,
			Data: p.Data,
		},
	})
	if err != nil {
		errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

func errorJSON(w http.ResponseWriter, err error) {
	statusCode := http.StatusBadRequest

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}
