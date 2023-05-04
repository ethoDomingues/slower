package main

import (
	"fmt"

	"github.com/ethoDomingues/slow"
)

var (
	db = map[string]any{}
)

func main() {
	app := slow.NewApp()

	app.Get("/todos", get)
	app.Put("/todos", put)
	app.Post("/todos", post)
	app.Delete("/todos", del)

	app.Listen()
}

func get(ctx *slow.Ctx) {
	response := ctx.Response
	response.JSON(db, 200)
}

func put(ctx *slow.Ctx) {
	response := ctx.Response
	request := ctx.Request
	for k, v := range request.Form {
		db[k] = v
	}
	response.JSON(db, 200)
}

func post(ctx *slow.Ctx) {
	response := ctx.Response
	request := ctx.Request
	for k, v := range request.Form {
		db[k] = v
	}
	response.JSON(db, 201)
}

func del(ctx *slow.Ctx) {
	response := ctx.Response
	request := ctx.Request
	fields, ok := request.Form["fields"].([]any)
	if !ok {
		panic(slow.TypeOf(request.Form["fields"]))
	}
	for _, key := range fields {
		delete(db, fmt.Sprint(key))
	}
	response.JSON(nil, 204)
}
