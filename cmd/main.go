package main

import (
	"fmt"
	"github.com/alpha-omega-corp/api-gateway/pkg/docker"
	"github.com/alpha-omega-corp/api-gateway/pkg/jwt"
	"github.com/alpha-omega-corp/services/config"
	"github.com/alpha-omega-corp/services/httputils"

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
	c := config.Get("dev")

	router := bunrouter.New(
		bunrouter.WithMiddleware(bunrouterotel.NewMiddleware()),
		bunrouter.WithMiddleware(reqlog.NewMiddleware(
			reqlog.WithEnabled(true),
			reqlog.WithVerbose(true),
		)))

	authentication := *jwt.RegisterRoutes(router, c)
	docker.RegisterRoutes(router, c, &authentication)

	listenAndServe(router, c.HOST)
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
