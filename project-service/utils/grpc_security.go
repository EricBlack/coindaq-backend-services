package utils

import (
	"google.golang.org/grpc"
	"runtime"
	"google.golang.org/grpc/codes"
	"time"
	"context"
	log1 "github.com/wothing/log"
)

const (
	MAXSTACKSIZE = 4096
)

func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			// log stack
			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]
			log1.CtxErrorf(ctx, "panic grpc invoke: %s, err=%v, stack:\n%s", info.FullMethod, r, string(stack))
			// if panic, set custom error to 'err', in order that client and sense it.
			err = grpc.Errorf(codes.Internal, "panic error: %v", r)
		}
	}()
	return handler(ctx, req)
}

func Logging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	log1.CtxInfof(ctx, "calling %s, req=%s", info.FullMethod, Marshal(req))
	resp, err = handler(ctx, req)
	log1.CtxInfof(ctx, "finished %s, took=%v, resp=%v, err=%v", info.FullMethod, time.Since(start), Marshal(resp), err)
	return resp, err
}

// UnaryInterceptorChain build the multi interceptors into one interceptor chain.
func UnaryInterceptorChain(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			chain = build(interceptors[i], chain, info)
		}
		return chain(ctx, req)
	}
}

// build is the interceptor chain helper
func build(c grpc.UnaryServerInterceptor, n grpc.UnaryHandler, info *grpc.UnaryServerInfo) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return c(ctx, req, info, n)
	}
}

func NewServer() *grpc.Server {
	return grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptorChain(Recovery, Logging)))
}