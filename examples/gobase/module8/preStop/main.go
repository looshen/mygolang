package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	httpServer()
}

func httpServer() {
	go SignalTERM()
	portPtr := flag.String("port", "8080", "port to listen on")
	fmt.Printf("httpServer portPtr:%+v\n", *portPtr)
	flag.Parse()
	os.Setenv("VERSION", "1.2.3")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
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
	log.Fatal(http.ListenAndServe(":"+*portPtr, nil))
}

func SignalTERM() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signals, os.Interrupt,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	pid := os.Getpid()
	go func() {
		p, err := os.FindProcess(pid)
		if err != nil {
			fmt.Printf("Error Getpid:%+v, err:%+v\n", pid, err)
			return
		}
		fmt.Printf("Getpid:%+v, err:%+v, p:%+v\n", pid, err, p)
		// test start
		fmt.Println("stop Service func...")
		time.Sleep(time.Second * 10)
		// err = p.Signal(syscall.SIGKILL)
		err = p.Signal(syscall.SIGTERM)
		// err = p.Signal(syscall.SIGQUIT)
		// test end
		sig := <-signals
		fmt.Printf("Signal:%+v\n", sig)
		done <- true
	}()
	fmt.Println("Running...")
	<-done
	fmt.Printf("pid %d exit\n", pid)
	os.Exit(0)
}
