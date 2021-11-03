package tracing

import (
	"io"
	"os"

	"git.xenonstack.com/util/test-portal/src/tracing"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

func GrpcServerOptions(service string) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor, ot.Tracer, io.Closer) {
	trace := tracing.New()
	trace.ServiceName = service
	trace.JaegerHost = os.Getenv("JAEGER_AGENT_HOST")
	trace.JeagerPort = os.Getenv("JAEGER_AGENT_PORT")

	// creating new jaeger tracer
	tracer, closer := trace.InitJaeger()

	// return unary server interceptor and stream server interceptor, tracer, closer
	return grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)), grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)), tracer, closer
}
