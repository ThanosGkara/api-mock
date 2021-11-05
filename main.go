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
func main() {

	m, e := ioutil.ReadFile("/app/mocks/response.json")
	if e != nil {
		fmt.Printf("Error reading mock file: %v\n", e)
		os.Exit(1)
	}
	searchMock = m

	fmt.Println("starting...")

	router := fasthttprouter.New()
	router.POST("/", Mock)

	log.Fatal(fasthttp.ListenAndServe(":80", router.Handler))

}
