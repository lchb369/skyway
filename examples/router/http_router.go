package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/skyhttprouter"
	"log"
)

// Index is the index handler
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

// Hello is the Hello handler
func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func HelloTest(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "HelloTest, %s!\n", ctx.UserValue("name"))
}

// MultiParams is the multi params handler
func MultiParams(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hi, %s, %s!\n", ctx.UserValue("name"), ctx.UserValue("word"))
}

// QueryArgs is used for uri query args test #11:
// if the req uri is /ping?name=foo, output: Pong! foo
// if the req uri is /piNg?name=foo, redirect to /ping, output: Pong!
func QueryArgs(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	fmt.Fprintf(ctx, "Pong! %s\n", string(name))
}

func main() {

	router := skyhttprouter.New()
	router.GET("/", Index)
	//router.GET("/hello/:name", Hello)
	router.GET("/hello/test", HelloTest)
	router.GET("/hello/test1", HelloTest)
	router.GET("/multi/:name/:word", MultiParams)
	router.GET("/ping", QueryArgs)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}