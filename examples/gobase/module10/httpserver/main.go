package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"

	"mygolang/examples/gobase/module10/httpserver/prometheus"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	httpServer()
}

func httpServer() {
	portPtr := flag.String("port", "8080", "port to listen on")
	fmt.Printf("httpServer portPtr:%+v\n", *portPtr)
	flag.Parse()
	os.Setenv("VERSION", "1.2.3")
	prometheus.Register()
	mux := http.NewServeMux()
	// fmt.Printf("httpServer 2 portPtr:%+v\n", *portPtr)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer()
		defer timer.ObserveTotal()
		delay := randInt(20, 3000)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		for key, values := range r.Header {
			fmt.Printf("httpServer key:%+v, values:%+v\n", key, values)
			for _, value := range values {
				w.Header().Add(key, value)
				fmt.Fprintf(w, fmt.Sprintf("key:%+v, value:%+v\n", key, value))
			}
		}
		version := os.Getenv("VERSION")
		w.Header().Set("VERSION", version)
		fmt.Fprintf(w, fmt.Sprintf("version:%+v\n", version))
		fmt.Fprintf(w, fmt.Sprintf("r.URL.Path:%+v\n", r.URL.Path))
		fmt.Fprintf(w, fmt.Sprintf("r.Host:%+v\n", r.Host))
		fmt.Fprintf(w, fmt.Sprintf("Client IP: %s | HTTP Status Code: %d\n", r.RemoteAddr, http.StatusOK))
		fmt.Printf("httpServer version:%+v, Path:%+v\n", version, r.URL.Path)
		log.Printf("Client IP: %s | http Status Code: %d\n", r.RemoteAddr, http.StatusOK)
		if strings.Contains(r.Host, "localhost") && strings.EqualFold(r.URL.Path, "/healthz") {
			fmt.Printf("httpServer healthz:%+v", http.StatusOK)
			fmt.Fprintf(w, fmt.Sprintf("localhost/healthz  :%+v\n", http.StatusOK))
			w.WriteHeader(http.StatusOK)
		}
	})
	mux.Handle("/metrics", promhttp.Handler())
	server := &http.Server{
		Addr:    ":" + *portPtr,
		Handler: mux,
	}
	server.ListenAndServe()
	// log.Fatal(http.ListenAndServe(":"+*portPtr, nil))
}

func httpServer2() {
	portPtr := flag.String("port", "8080", "port to listen on")
	fmt.Printf("httpServer portPtr:%+v\n", *portPtr)
	flag.Parse()
	os.Setenv("VERSION", "1.2.3")
	prometheus.Register()
	// mux := http.NewServeMux()
	// fmt.Printf("httpServer 2 portPtr:%+v\n", *portPtr)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer()
		defer timer.ObserveTotal()
		delay := randInt(20, 3000)
		time.Sleep(time.Duration(delay) * time.Millisecond)
		for key, values := range r.Header {
			fmt.Printf("httpServer key:%+v, values:%+v\n", key, values)
			for _, value := range values {
				w.Header().Add(key, value)
				fmt.Fprintf(w, fmt.Sprintf("key:%+v, value:%+v\n", key, value))
			}
		}
		version := os.Getenv("VERSION")
		w.Header().Set("VERSION", version)
		fmt.Fprintf(w, fmt.Sprintf("version:%+v\n", version))
		fmt.Fprintf(w, fmt.Sprintf("r.URL.Path:%+v\n", r.URL.Path))
		fmt.Fprintf(w, fmt.Sprintf("r.Host:%+v\n", r.Host))
		fmt.Fprintf(w, fmt.Sprintf("Client IP: %s | HTTP Status Code: %d\n", r.RemoteAddr, http.StatusOK))
		fmt.Printf("httpServer version:%+v, Path:%+v\n", version, r.URL.Path)
		log.Printf("Client IP: %s | http Status Code: %d\n", r.RemoteAddr, http.StatusOK)
		if strings.Contains(r.Host, "localhost") && strings.EqualFold(r.URL.Path, "/healthz") {
			fmt.Printf("httpServer healthz:%+v", http.StatusOK)
			fmt.Fprintf(w, fmt.Sprintf("localhost/healthz  :%+v\n", http.StatusOK))
			w.WriteHeader(http.StatusOK)
		}
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":"+*portPtr, nil))
}

func randInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randomMillisecond := rand.Intn(max-min) + min
	return randomMillisecond
}
