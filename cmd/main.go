package main

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/middlewares"
	"github.com/alpha-omega-corp/api-gateway/pkg/github"
	"github.com/alpha-omega-corp/api-gateway/pkg/user"
	"github.com/alpha-omega-corp/services/config"
	"github.com/alpha-omega-corp/services/httputils"
	_ "github.com/spf13/viper/remote"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/bunrouterotel"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := bunrouter.New(
		bunrouter.WithMiddleware(bunrouterotel.NewMiddleware()),
		bunrouter.WithMiddleware(middlewares.NewCorsMiddleware()),
		bunrouter.WithMiddleware(middlewares.NewErrorHandler),
		bunrouter.WithMiddleware(reqlog.NewMiddleware(
			reqlog.WithEnabled(true),
			reqlog.WithVerbose(true),
		)))

	authClient := user.RegisterRoutes(router)
	github.RegisterRoutes(router, authClient)

	env, err := config.NewHandler().Environment("gateway")
	if err != nil {
		panic(err)
	}

	listenAndServe(router, env.Host.Url)
}

func listenAndServe(r *bunrouter.Router, host string) {
	var handler http.Handler
	handler = httputils.ExitOnPanicHandler{Next: r}

	srv := &http.Server{
		Addr:         host,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !isServerClosed(err) {
			log.Printf("ListenAndServe failed: %s", err)
		}
	}()

	fmt.Printf("listening on http://%s\n", srv.Addr)
	awaitSignal()
}

func isServerClosed(err error) bool {
	return err.Error() == "http: Server closed"
}

func awaitSignal() os.Signal {
	ch := make(chan os.Signal, 3)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-ch
}
