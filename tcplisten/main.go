package main

import (
	"fmt"
	"net"
	"net/http"
	"syscall"
	"time"
)

type tcpResizeRlimits struct {
	*net.TCPListener
}

func (ln tcpResizeRlimits) Accept() (net.Conn, error) {
	conn, err := ln.AcceptTCP()
	if err != nil {
		if e, ok := err.(syscall.Errno); ok {
			switch e {
			case syscall.EMFILE:
				fallthrough
			case syscall.ENFILE:
				rlim := &syscall.Rlimit{}
				serr := syscall.Getrlimit(syscall.RLIMIT_NOFILE, rlim)
				if serr != nil {
					return nil, fmt.Errorf("%v: get limit: %v", err, serr)
				}
				cur, max := rlim.Cur*rlim.Cur, rlim.Max*rlim.Max
				serr = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{cur, max})
				if serr != nil {
					return nil, fmt.Errorf("%v: set limit: %v", err, serr)
				}
				conn, err = ln.AcceptTCP()
				if err != nil {
					return nil, err
				}
				goto noerr
			}
		}
		return nil, err
	}
noerr:
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(time.Minute)
	return conn, nil
}

func main() {
	srv := &http.Server{
		Handler: http.HandlerFunc(hello),
	}
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Run server port 8080..")
	err = srv.Serve(tcpResizeRlimits{l.(*net.TCPListener)})
	panic(err)
}

func hello(w http.ResponseWriter, t *http.Request) {
	w.Write([]byte("hello listen here..."))
}
