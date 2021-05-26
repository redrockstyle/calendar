package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// responseJSON кодирует переданный в аргументе v объект в json и возвращает созданный json клиенту с заданным кодом состояния
func responseJSON(ctx *fasthttp.RequestCtx, v interface{}, statusCode int) error {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)

	encoder := json.NewEncoder(ctx)
	err := encoder.Encode(v)

	return err
}

// response возвращает содержимое аргумента content клиенту с заданным кодом состояния
func response(ctx *fasthttp.RequestCtx, content string, statusCode int) {
	ctx.SetStatusCode(statusCode)
	// nolint: errcheck
	fmt.Fprint(ctx, content)
}