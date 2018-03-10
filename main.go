package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db.AutoMigrate(&Task{})

	err := logan.open()
	if err != nil {
		panic("open log file err" + err.Error())
	}
	err = ginLogan.open()
	if err != nil {
		panic("open log file err" + err.Error())
	}
	// clear upt
	defer logan.close()
	defer ginLogan.close()
	defer os.RemoveAll(tempDir)
	defer db.Close()

	logan.Info("Use temp file", tempDir)

	r := engine()
	logan.Info("http Server")
	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		logan.Info("listen and server:", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			logan.Error("ListenAndServe err:", err)
		}
	}()

	// wait for os interrupt signal to stop.
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	<-quit

	logan.Info("shuting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logan.Error("shut down err", err)
	}
	logan.Info("server exiting")
}
