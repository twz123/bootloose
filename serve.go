package main

import (
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/k0sproject/footloose/pkg/api"
	"github.com/k0sproject/footloose/pkg/cluster"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launch a footloose server",
	RunE:  serve,
}

var serveOptions struct {
	listen       string
	keyStorePath string
	debug        bool
}

func baseURI(addr string) (string, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", err
	}
	if host == "" || host == "0.0.0.0" || host == "[::]" {
		host = "localhost"
	}
	return fmt.Sprintf("http://%s:%s", host, port), nil
}

func init() {
	serveCmd.Flags().StringVarP(&serveOptions.listen, "listen", "l", ":2444", "Cluster configuration file")
	serveCmd.Flags().StringVar(&serveOptions.keyStorePath, "keystore-path", defaultKeyStorePath, "Path of the public keys store")
	serveCmd.Flags().BoolVar(&serveOptions.debug, "debug", false, "Enable debug")
	footloose.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) error {
	opts := &serveOptions

	baseURI, err := baseURI(opts.listen)
	if err != nil {
		return errors.Wrapf(err, "invalid listen address '%s'", opts.listen)
	}

	log.Infof("Starting server on: %s\n", opts.listen)

	keyStore := cluster.NewKeyStore(opts.keyStorePath)
	if err := keyStore.Init(); err != nil {
		return errors.Wrapf(err, "could not init keystore")
	}

	log.Infof("Key store successfully initialized in path: %s\n", opts.keyStorePath)

	api := api.New(baseURI, keyStore, opts.debug)
	router := api.Router()

	err = http.ListenAndServe(opts.listen, router)
	if err != nil {
		log.Fatalf("Unable to start server: %s", err)
	}

	return nil
}