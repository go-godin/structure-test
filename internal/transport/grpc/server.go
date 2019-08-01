package grpc

import (
	"context"

	"github.com/go-godin/example-structure/internal/endpoint"

	pb "bitbucket.org/jdbergmann/protobuf-go/ticket/ticket"
	kitGrpc "github.com/go-kit/kit/transport/grpc"
)

type Server struct {
	CreateHandler kitGrpc.Handler
}

func NewServer(endpoints endpoint.Set) pb.TicketServiceServer {
	return &Server{
		CreateHandler: kitGrpc.NewServer(
			endpoints.CreateEndpoint,
			DecodeCreateRequest,
			EncodeCreateResponse,
			nil),
	}
}

func DecodeCreateRequest(ctx context.Context, pbRequest interface{}) (request interface{}, err error) {
	request = pbRequest.(endpoint.CreateTicketRequest)
	return request, nil
}

func EncodeCreateResponse(ctx context.Context, response interface{}) (pbResponse interface{}, err error) {
	pbResponse = response.(*pb.CreateTicketResponse)
	return pbResponse, nil
}

func (s *Server) Create(ctx context.Context, req *pb.CreateTicketRequest) (*pb.CreateTicketResponse, error) {
	_, resp, err := s.CreateHandler.ServeGRPC(ctx, req)
	if err != nil {
	    return nil, err // TODO: encode err
	}
	return resp.(*pb.CreateTicketResponse), nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetTicketRequest) (*pb.GetTicketResponse, error) {
	return nil, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteTicketRequest) (*pb.DeleteTicketResponse, error) {
	return nil, nil
}

func (s *Server) List(context.Context, *pb.ListTicketsRequest) (*pb.ListTicketsResponse, error) {
	panic("implement me")
}

func (s *Server) SetAssignee(context.Context, *pb.SetAssigneeRequest) (*pb.SetAssigneeResponse, error) {
	panic("implement me")
}

func (s *Server) AddPerson(context.Context, *pb.AddPersonRequest) (*pb.AddPersonResponse, error) {
	panic("implement me")
}

func (s *Server) RemovePerson(context.Context, *pb.RemovePersonRequest) (*pb.RemovePersonResponse, error) {
	panic("implement me")
}

func (s *Server) Update(context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	panic("implement me")
}

func (s *Server) Close(context.Context, *pb.CloseRequest) (*pb.CloseResponse, error) {
	panic("implement me")
}

func (s *Server) UpdateReference(context.Context, *pb.UpdateReferenceRequest) (*pb.UpdateReferenceResponse, error) {
	panic("implement me")
}

func (s *Server) UpdateType(context.Context, *pb.UpdateTypeRequest) (*pb.UpdateTypeResponse, error) {
	panic("implement me")
}

func (s *Server) UpdateStatus(context.Context, *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	panic("implement me")
}
