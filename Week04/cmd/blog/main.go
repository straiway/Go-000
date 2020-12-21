package main

import (
	blog "Week04/api/blog/v1"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	port = ":8301"
)

type server struct {
	blog.UnimplementedBlogServer
}

func (s *server) Publish(ctx context.Context, reply *blog.PublishReply) (*blog.PublishReply, error) {
	return &blog.PublishReply{ArticleID: 111}, nil
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return listenSignal(ctx)
	})

	g.Go(func() error {
		return startServer(ctx)
	})

	if err := g.Wait(); err != nil {
		log.Print(err)
	}
}

func listenSignal(ctx context.Context) error {
	exitSignals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT}
	stop := make(chan os.Signal, len(exitSignals))
	signal.Notify(stop)

	select {
	case s := <-stop:
		return fmt.Errorf("stopped by signal: %v", s)
	case <-ctx.Done():
		return nil
	}
}

func startServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	blog.RegisterBlogServer(s, &server{})

	done := make(chan error)
	go func() {
		if err := s.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			done <- fmt.Errorf("failed to serve: %v", err)
		}
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		s.GracefulStop()
		return nil
	}
}
