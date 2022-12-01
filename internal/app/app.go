package app

import (
	"flag"
	"fmt"
	"github.com/farzanehshahi/user-kit/internal/config"
	"github.com/farzanehshahi/user-kit/internal/user"
	"github.com/farzanehshahi/user-kit/pkg/postgres"
	"github.com/farzanehshahi/user-kit/pkg/validator"
	kitPrometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	stdPrometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	var flagConfig = flag.String("config", configPath, "path to the config file")
	flag.Parse()

	// create logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "user",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	// setup configuration
	cfg, err := config.Load(*flagConfig)
	if err != nil {
		fmt.Println(err)
		level.Error(logger).Log("msg", "failed to load application configuration", "error:", err)
		os.Exit(-1)
	}

	// connect to postgres database
	//DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
	//	"postgres", "5431", "farzaneh", "ukit", "3971231050")
	db, err := postgres.GetDB(cfg.PostgresDSN)
	if err != nil {
		level.Error(logger).Log("msg", "failed to connect to database", "error:", err)
		os.Exit(-1)
	}

	// setup v
	// todo: validator name
	v := validator.NewValidator()

	fieldKeys := []string{"method"}
	httpLogger := log.With(logger, "component", "http")

	// setup services
	// setup user service
	userService := user.NewService(
		user.NewRepository(db, logger),
		v,
		logger,
	)
	userService = user.NewLoggingService(
		log.With(logger, "component", "user"),
		userService,
	)
	userService = user.NewInstrumentingService(
		kitPrometheus.NewCounterFrom(stdPrometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "user_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "user_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		userService,
	)

	router := mux.NewRouter()
	router.Use(commonMiddleware)

	endpoints := user.MakeEndpoints(userService)
	user.MakeHandlers(router, endpoints, httpLogger)
	//

	router.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong\n")
	}))
	router.Handle("/metrics", promhttp.Handler())

	bindAddress := fmt.Sprintf(":%d", cfg.ServerPort)
	srv := http.Server{
		Addr:    bindAddress,
		Handler: router,
	}

	errs := make(chan error, 2)
	go func() {
		level.Info(logger).Log("msg", "server started...", "address", bindAddress)
		errs <- srv.ListenAndServe()

		//errs <- http.ListenAndServe(bindAddress, nil)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)

}

func pong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
