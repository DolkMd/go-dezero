package sysutil

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func Stats() (memStats runtime.MemStats) {
	runtime.ReadMemStats(&memStats)
	return
}

func RunViewMemoryLocal() {
	go func() {
		log.Println("http://localhost:6060/debug/pprof")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
}
