package main

import (
	"context"
	"fmt"
	conf "lecture/WBA-BC-Project-04/conf"
	ctl "lecture/WBA-BC-Project-04/contorller"
	lg "lecture/WBA-BC-Project-04/logger"
	md "lecture/WBA-BC-Project-04/model"
	rt "lecture/WBA-BC-Project-04/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	cf := conf.GetConfig("./conf/config.toml")
	fmt.Println()
	fmt.Println(cf.Server.Port)
	fmt.Println(cf.DB["account"]["user"])
	for _, w := range cf.Work {
		fmt.Println(w.Desc)
	}

	// 로그 초기화
	if err := lg.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	lg.Debug("ready server....")

	//model 모듈 선언
	if mod, err := md.NewModel(cf); err != nil {
		panic(err)
	} else if controller, err := ctl.NewCTL(mod); err != nil {
		panic(fmt.Errorf("controller.New > %v", err))
	} else if rt, err := rt.NewRouter(controller); err != nil {
		panic(fmt.Errorf("router.NewRouter > %v", err))
	} else {
		mapi := &http.Server{
			Addr:           ":8080",
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			fmt.Println("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			fmt.Println("timeout of 5 seconds.")
		}

		fmt.Println("Server exiting")
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
