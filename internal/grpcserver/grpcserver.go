package grpcserver

import (
	"context"
	"fmt"
	"gRPC-Protocol-100/internal/grpcserver/stream"
	"gRPC-Protocol-100/pkg/logger"
	"gRPC-Protocol-100/pkg/port"
	"gRPC-Protocol-100/pkg/protocol100r5"
	"io"
)

type Server struct {
	L logger.Logger
	C protocol100r5.Commander
	P port.Porter
	stream.UnimplementedApiCallerScaleServer
}

// Канал обмена данными с весами
func (s *Server) ScalesMessageOutChannel(st stream.ApiCallerScale_ScalesMessageOutChannelServer) error {
	s.L.Debug("Start gRPC ScalesMessageOutChannel")
	ctx := st.Context()
	for {
		// Контекст
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		// Прием данных
		in, err := st.Recv()
		if err == io.EOF {
			s.L.Error("ScalesMessageOutChannel: %v", err)
			return nil
		}
		if err != nil {
			continue
		}

		smessage := ""
		stype := ""
		ssubtype := ""
		serror := "nil"

		recv, err := s.P.Communicate([]byte(in.Message))
		if err != nil {
			serror = "Error: " + err.Error()
		} else {
			smessage = fmt.Sprintf("% 02X", recv)
		}

		out := &stream.ResponseScale{Message: smessage, Type: stype, Subtype: ssubtype, Error: serror}
		// Отправка данных
		if err := st.Send(out); err != nil {
			s.L.Error("ScalesMessageOutChannel: %v", err)
		}
	}
}

// Установить тару
func (s *Server) SetTare(ctx context.Context, in *stream.Empty) (*stream.ResponseSetScale, error) {
	s.L.Debug("Start gRPC SetTare")
	errorMessage := "nil"
	err := s.C.SetTare()
	if err != nil {
		s.L.Error("SetTare - %v", err)
		errorMessage = "Error SetTare: " + err.Error()
	}
	return &stream.ResponseSetScale{Error: errorMessage}, nil
}

// Установить вес тары
func (s *Server) SetTareValue(ctx context.Context, in *stream.RequestTareValue) (*stream.ResponseSetScale, error) {
	s.L.Debug("Start gRPC SetTareValue")
	errorMessage := "nil"
	err := s.C.SetTareValue(in.Message)
	if err != nil {
		s.L.Error("SetTareValue - %v", err)
		errorMessage = "Error SetTareValue: " + err.Error()
	}
	return &stream.ResponseSetScale{Error: errorMessage}, nil
}

// Установить нулевое значение
func (s *Server) SetZero(ctx context.Context, in *stream.Empty) (*stream.ResponseSetScale, error) {
	s.L.Debug("Start gRPC SetZero")
	errorMessage := "nil"
	err := s.C.SetZero()
	if err != nil {
		s.L.Error("SetZero - %v", err)
		errorMessage = "Error SetZero: " + err.Error()
	}
	return &stream.ResponseSetScale{Error: errorMessage}, nil
}

// Получить вес
func (s *Server) GetInstantWeight(ctx context.Context, in *stream.Empty) (*stream.ResponseInstantWeight, error) {
	s.L.Debug("Start gRPC GetInstantWeight")
	errorMessage := "nil"
	weight, err := s.C.GetMassa()
	if err != nil {
		s.L.Error("SetZero - %v", err)
		errorMessage = "Error GetInstantWeight: " + err.Error()
	}

	return &stream.ResponseInstantWeight{Message: weight, Error: errorMessage}, nil
}

// Получить статус весов
func (s *Server) GetState(ctx context.Context, in *stream.Empty) (*stream.ResponseScale, error) {
	s.L.Debug("Start gRPC GetState")
	errorMessage := "nil"
	state, err := s.C.GetScalePar()
	if err != nil {
		s.L.Error("GetState - %v", err)
		errorMessage = "Error GetState: " + err.Error()
	}

	return &stream.ResponseScale{Message: string(state), Error: errorMessage}, nil
}
