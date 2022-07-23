package main

import (
	"context"
	"fmt"
	"learn/common"
	"learn/config"
	"learn/entity"
	"learn/handler"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.Load()
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	var db = common.GetDB()
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Project{})
	srv := handler.Run(config.Get().Port)
	for {
		sig := <-sigs
		fmt.Println("received sig: ", sig)
		ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("error when shutdown server.", err.Error())
			os.Exit(1)
		} else {
			fmt.Println("server was shutdown.")
			os.Exit(0)
		}
	}

}
