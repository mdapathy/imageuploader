package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mdapathy/imageuploader/pkg/api/httpapi"
	"github.com/mdapathy/imageuploader/pkg/config"
	"github.com/mdapathy/imageuploader/pkg/domain/query"
	"github.com/mdapathy/imageuploader/pkg/domain/service"
	mongostore "github.com/mdapathy/imageuploader/pkg/store"
	"golang.org/x/sync/errgroup"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	cfg := flag.String("config", "config/config.yaml", "Path to the configuration file")
	port := flag.Int("port", 8080, "Port of the web server")

	flag.Parse()

	conf, err := config.New(*cfg)
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("connecting to mongodb: %s", conf.Mongo.URI())
	storeClient, err := mongostore.NewClient(context.Background(), conf.Mongo.URI())
	if err != nil {
		log.Fatalf("failed connect to mongo DB: %v", err)
	}

	store, err := mongostore.NewStore(storeClient, conf.Mongo.Database)
	if err != nil {
		log.Fatalf("failed create mongo store: %v", err)
	}

	defer func() {
		if err := storeClient.Disconnect(ctx); err != nil {
			log.Fatalf("can't disconnect mongo client: %s", err)
		}
	}()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		addr := fmt.Sprintf("0.0.0.0:%d", *port)
		filterer := mongostore.NewFilterer()
		apiServer := httpapi.NewServer(
			addr,
			&conf.Server,
			service.New(store.Image(), filterer),
			query.NewFactory(mongostore.NewSorter(), filterer),
		)

		log.Printf("http server start listening on: %s", addr)
		return apiServer.Start(gCtx)
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("failed: %s", err)
	}
}
