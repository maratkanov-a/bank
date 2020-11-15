package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/maratkanov-a/bank/internal/pkg/config"
	"github.com/pkg/errors"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"github.com/utrack/clay/v2/transport"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type optionsListener struct {
	HTTPPort int
	RPCPort  int
	// DebugPort is a port where an app listens for debug/pprof connections; default is 32000
	DebugPort int
}

type OptionsHTTP struct {
	// AllowOriginHosts sets Access-Control-Allow-Origin hostname list headers for browsers.
	AllowOriginHosts []string
}

type OptionsGRPC struct {
	UnaryInterceptor grpc.UnaryServerInterceptor
}

type Listeners struct {
	HTTP    net.Listener
	RPC     net.Listener
	HTTPDev net.Listener
}

const defaultDebugPort = 32000

// newListeners creates and binds listeners for a server.
func newListeners(opts optionsListener) (*Listeners, error) {
	if opts.DebugPort == 0 {
		opts.DebugPort = defaultDebugPort
	}
	dbg, err := net.Listen("tcp", fmt.Sprintf(":%v", opts.DebugPort))
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get DebugPort listener")
	}
	http, err := net.Listen("tcp", fmt.Sprintf(":%v", opts.HTTPPort))
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get HTTP listener")
	}
	rpc, err := net.Listen("tcp", fmt.Sprintf(":%v", opts.RPCPort))
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get RPC listener")
	}
	return &Listeners{HTTPDev: dbg, HTTP: http, RPC: rpc}, nil
}

func newDebugHTTPMux(mux chi.Router) http.Handler {
	//mux.Handle("/metrics")
	mux.Mount("/debug", chimiddleware.Profiler())
	return mux
}

func runDebug(l *Listeners) error {
	errs := make(chan error, 1)
	logrus.Infof("running debug server on %v", l.HTTPDev.Addr().String())
	go func() {
		hmux := chi.NewRouter()
		errs <- http.Serve(l.HTTPDev, newDebugHTTPMux(hmux))
	}()
	select {
	case err := <-errs:
		return errors.Wrap(err, "couldn't start debug handler")
	case <-time.After(time.Millisecond * 100):
		return nil
	}
}

func nilOrErr(c chan<- error, f func() error) {
	err := f()
	if err != nil {
		c <- err
	}
}

func Run(gs *grpc.Server, hmux *chi.Mux, l *Listeners) error {
	errs := make(chan error, 2)
	httpSrv := &http.Server{Handler: hmux}

	// Run listeners in parallel, report errors
	go nilOrErr(errs, func() error {
		logrus.Infof("running RPC server on %v", l.RPC.Addr().String())
		return gs.Serve(l.RPC)
	})
	go nilOrErr(errs, func() error {
		logrus.Infof("running HTTP server on %v", l.HTTP.Addr().String())
		return httpSrv.Serve(l.HTTP)
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Wait on error or SIGINT
	select {
	case err := <-errs:
		return errors.Wrap(err, "error returned from the listener")
	case <-stop:
		break
	}

	// INT caught, stop gracefully

	logrus.Info("caught SIGINT, shutting down...")
	// in 10s
	waitCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	exErrs := make(chan error, 2)
	go func() {
		exErrs <- httpSrv.Shutdown(waitCtx)
	}()

	for i := 0; i < 1; i++ {
		err := <-exErrs
		if err != nil {
			return errors.Wrap(err, "couldn't stop server gracefully")
		}
	}
	return nil
}

func RunAll(compound *transport.CompoundServiceDesc, cfg *config.Env) error {
	staticFS, err := fs.New()
	if err != nil {
		return err
	}

	hmux := chi.NewRouter()
	gs := grpc.NewServer()

	hmux.Mount("/", http.FileServer(staticFS))
	hmux.MethodFunc("GET", "/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		_, _ = w.Write(compound.SwaggerDef())
	})

	compound.RegisterHTTP(hmux)
	compound.RegisterGRPC(gs)

	lis, err := newListeners(
		optionsListener{
			HTTPPort: cfg.PortHTTP,
			RPCPort:  cfg.PortGRPC,
		})

	if err != nil {
		return errors.Wrap(err, "couldn't assign ports: ")
	}

	if err := runDebug(lis); err != nil {
		return errors.Wrap(err, "couldn't run debug profile server: ")
	}

	return Run(gs, hmux, lis)
}
