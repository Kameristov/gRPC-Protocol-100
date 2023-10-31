package main

import (
	"fmt"
	"gRPC-Protocol-100/internal/grpcserver"
	"gRPC-Protocol-100/internal/grpcserver/stream"
	"gRPC-Protocol-100/pkg/config"
	"gRPC-Protocol-100/pkg/logger"
	"gRPC-Protocol-100/pkg/port"
	"gRPC-Protocol-100/pkg/protocol100r5"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	var err error

	// Config
	cfg, err := config.New()
	if err != nil {
		return err
	}

	// log
	l := logger.New(cfg.LogLevel)
	l.Info("Config: %+v", cfg)

	// Создание порта
	port := port.New(cfg.Scaler.SerialPort, cfg.Scaler.Type, cfg.Scaler.Protocol)
	if err = port.Conect(); err != nil {
		return fmt.Errorf("error connect to port: %v", err)
	}

	// Создание протокола
	protocol100 := protocol100r5.New(port)

	// Создание GRPC
	listener, err := net.Listen("tcp", cfg.GRPC.Address)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	stream.RegisterApiCallerScaleServer(s, &grpcserver.Server{L: *l, C: protocol100, P: port})

	// Перехват сигнала прерывания
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Завершение работы по прерыванию
	go func() {
		signal := <-interrupt
		l.Info("interrupt signal: OS " + signal.String())
		s.GracefulStop()
		port.Disconect()
		wg.Done()
	}()

	// Запуск gRPC сервера
	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("error start grpc server: %v", err)
	}

	wg.Wait()

	return nil
}
