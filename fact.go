package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
        "strconv"
        "math/big"
)

const (
	version = "v1.0.0"
)

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

	log.Println("Starting Factorial Application...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

            pp := r.URL.Path[1:len(r.URL.Path)]
            s, err := strconv.Atoi(pp)
                
            if ( len(pp) < 1 || err!=nil ) { fmt.Fprint(w,"Try http://127.0.0.1:3000/5 to calculate 5! = 120") 
                                             log.Println("User did not provide valid input")
                                             return
                                           }
            var f big.Int
            f.MulRange(1,int64(s))
            answer := f.String() 
            log.Println("Using math/big library to calculate answer " + pp + "! = " + answer)
            fmt.Fprint(w,"Calculating Factorial: " + pp + "! = " + answer)	
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, version) })
         http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })
         http.HandleFunc("/header", headers )

	s := http.Server{Addr: ":3000"}

	go func() { log.Fatal(s.ListenAndServe()) }()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
