package main

import (
	"context"
	"net/http"
	"time"

	"github.com/AaravShirvoikar/logger-client/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) Log(w http.ResponseWriter, r *http.Request) {
	var payload LogPayload

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	app.LogEvent(w, r, payload)
}

func (app *Config) LogEvent(w http.ResponseWriter, r *http.Request, p LogPayload) {
	conn, err := grpc.Dial("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		app.errorJSON(w, err)
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
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "logged"

	app.writeJSON(w, http.StatusAccepted, payload)
}
