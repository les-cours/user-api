package services

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/les-cours/user-api/graph"
	"github.com/les-cours/user-api/graph/resolvers"
	bookApi "github.com/les-cours/user-api/protobuf/book"
	"google.golang.org/grpc"
)

const HttpPort = "1112"

func Start() {
	var err error

	conn, err := grpc.Dial("user-service:1113", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to book domain service: %v", err)
	}
	defer conn.Close()
	bookClient := bookApi.NewBookServiceClient(conn)

	var r = resolvers.Resolver{
		BookClient: bookClient,
	}
	var srv = handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &r},
		),
	)

	http.Handle("/api", srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/api"))

	log.Printf("Starting http server on port " + HttpPort)

	err = http.ListenAndServe(":"+HttpPort, nil)
	if err != nil {
		log.Fatalf("Error http server on port %v: %v", HttpPort, err)
	}
}
