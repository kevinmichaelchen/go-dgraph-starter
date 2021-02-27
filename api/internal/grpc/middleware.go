package grpc

import (
	"context"
	"strings"

	"github.com/MyOrg/go-dgraph-starter/internal/kontext"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/rs/zerolog/log"

	"github.com/MyOrg/go-dgraph-starter/internal/configuration"

	grpczerolog "github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/tags"
	grpcGo "google.golang.org/grpc"
)

func newServer() *grpcGo.Server {
	// Logger is used, allowing pre-definition of certain fields by the user.
	logger := configuration.GetLogger()
	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []logging.Option{
		logging.WithDecider(func(fullMethodName string, err error) bool {
			// Don't log health gRPC endpoint
			if strings.HasSuffix(fullMethodName, "/Check") {
				return false
			}
			return true
		}),
	}

	return grpcGo.NewServer(
		middleware.WithUnaryServerChain(
			kontext.UnaryServerInterceptor,
			tags.UnaryServerInterceptor(),
			// TODO wait until they use OpenTelemetry
			//opentracing.UnaryServerInterceptor(),
			//prometheus.UnaryServerInterceptor,
			logging.UnaryServerInterceptor(grpczerolog.InterceptorLogger(logger), opts...),
			requestLoggingInterceptor,
			recovery.UnaryServerInterceptor(),
		),
		middleware.WithStreamServerChain(
			tags.StreamServerInterceptor(),
			// TODO wait until they use OpenTelemetry
			//opentracing.StreamServerInterceptor(),
			//prometheus.StreamServerInterceptor,
			logging.StreamServerInterceptor(grpczerolog.InterceptorLogger(logger), opts...),
			recovery.StreamServerInterceptor(),
		),
	)
}

func requestLoggingInterceptor(ctx context.Context, req interface{}, info *grpcGo.UnaryServerInfo, handler grpcGo.UnaryHandler) (resp interface{}, err error) {
	if !strings.HasSuffix(info.FullMethod, "/Check") {
		m := jsonpb.Marshaler{EmitDefaults: true}

		if pb, ok := req.(proto.Message); ok {
			s, err := m.MarshalToString(pb)
			if err != nil {
				log.Warn().Err(err).Msg("failed to marshal grpc request")
			} else {
				log.Info().Msgf("🐒 ⚡ found grpc request: %s %s", info.FullMethod, s)
			}
		} else {
			log.Warn().Err(err).Msg("empty interface was not type-assertable to proto.Message")
		}
	}

	return handler(ctx, req)
}
