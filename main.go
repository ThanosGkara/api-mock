package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var searchMock []byte
var searchMock1 []byte
var searchMock2 []byte

func Mock(ctx *fasthttp.RequestCtx) {
	fmt.Println("mocked!")
	mock_num := os.Getenv("MOCK")
	fmt.Println(mock_num)
	fmt.Println(ctx.Request.Header.String())
	fmt.Println(string(ctx.PostBody()))
	ctx.Response.Header.Set("Content-Type", "application/json")
	time.Sleep(100 * time.Millisecond)
	ctx.Write(searchMock)
}

func Mock1(ctx *fasthttp.RequestCtx) {
	fmt.Println("mocked!")
	mock_num := os.Getenv("MOCK")
	fmt.Println(mock_num)
	fmt.Println(ctx.Request.Header.String())
	fmt.Println(string(ctx.PostBody()))
	ctx.Response.Header.Set("Content-Type", "application/json")
	time.Sleep(100 * time.Millisecond)
	ctx.Write(searchMock1)
}

func Mock2(ctx *fasthttp.RequestCtx) {
	fmt.Println("mocked!")
	mock_num := os.Getenv("MOCK")
	fmt.Println(mock_num)
	fmt.Println(ctx.Request.Header.String())
	fmt.Println(string(ctx.PostBody()))
	ctx.Response.Header.Set("Content-Type", "application/json")
	time.Sleep(100 * time.Millisecond)
	ctx.Write(searchMock2)
}

func main() {

	m, e := ioutil.ReadFile("/app/mocks/response.json")
	m1, e1 := ioutil.ReadFile("/app/mocks/response1.json")
	m2, e2 := ioutil.ReadFile("/app/mocks/response2.json")

	if e != nil || e1 != nil || e2 != nil {
		fmt.Printf("Error reading mock file: %v\n", e)
		os.Exit(1)
	}
	searchMock = m
	searchMock1 = m1
	searchMock2 = m2

	fmt.Println("starting...")

	router := fasthttprouter.New()
	router.POST("/", Mock)
	router.POST("/api1", Mock1)
	router.GET("/api2", Mock2)

	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))

}
