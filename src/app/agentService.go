package main

import (
	"app/enmu"
	"app/session"
	"bufio"
	"github.com/golang/glog"
	"net"
	"os"
)

var closed = make(chan struct{}, 1)

func agentRun() {
	lestener, err := net.Listen("tcp", enmu.ServerHost+":"+enmu.ServerPort)
	if err != nil {
		glog.Info("listen error:", err)
		os.Exit(1)
	}
	defer lestener.Close()
	glog.Info("listening on " + enmu.ServerHost + ":" + enmu.ServerPort)
	for {
		conn, err := lestener.Accept()
		if err != nil {
			glog.Info("accept error:", err)
			os.Exit(1)
		}
		glog.Infof("message %s->%s\n", conn.RemoteAddr(), conn.LocalAddr())
		sess := session.NewSession()
		go handleRequest(conn, sess)
		go handWriteResp(conn, sess)
	}
}

func handleRequest(conn net.Conn, sess *session.Session) {
	ip := conn.RemoteAddr().String()
	defer func() {
		glog.Info("disconnect:" + ip)
		sess.AddDieChan()
		closed <- struct{}{}
		conn.Close()
	}()
	in := make(chan []byte, 16)
	sess.EvaluationReciveChan(in)
	reader := bufio.NewReader(conn)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			glog.Info(err)
			return
		}
		in <- bytes
	}
}

func handWriteResp(conn net.Conn, sess *session.Session) {
	ch := make(chan []byte, 16)
	sess.EvaluationSendChan(ch)
	writer := bufio.NewWriter(conn)
	for {
		select {
		case msg := <-ch:
			writer.Write(msg)
			writer.Write([]byte("\n"))
			writer.Flush()
		case <-closed:
			glog.Info("closed resp connect")
			return
		default:
		}
	}
}
