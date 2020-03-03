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

const ( version = "v1.0.0" )

func headers(w http.ResponseWriter, req *http.Request) {
    log.Println("Requested Header API")
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
            log.Println(fmt.Sprintf("%v: %v", name, h))
        }
    }
}


func main() {

	log.Println("Starting Factorial Application...")

        http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, version)
                                                                                   log.Println("Requested Version API")
                                                                                 })
        http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200)
                                                                                  log.Println("Requested Health API")
                                                                                })
        http.HandleFunc("/header", headers )

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

            pp := r.URL.Path[1:len(r.URL.Path)]
            k, err := strconv.Atoi(pp)
                
            if ( (len(pp) < 1) || (err != nil) ) { 
                fmt.Fprint(w,"Try http://127.0.0.1:3000/5 to calculate 5! = 120") 
                // log.Println("User did not provide valid input")
                return
              } else {
            var f big.Int
            f.MulRange(1,int64(k))
            answer := f.String() 
            log.Println("Using math/big library to calculate answer " + pp + "! = " + answer)
            fmt.Fprint(w,"Calculating Factorial: " + pp + "! = " + answer)	
              }
	})

	var port string
	port = os.Getenv("PORT")
        if ( len(port) == 0 ) {
		port = "3000"
	}

        log.Println("Using port"+port)

	log.Fatal(http.ListenAndServe(":"+port, nil))

	s := http.Server{Addr: ":" + port }

	go func() { log.Fatal(s.ListenAndServe()) }()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	log.Println("Shutdown signal received, exiting...")

	s.Shutdown(context.Background())
}
