package endpoint

import (
	"context"

	pb "bitbucket.org/jdbergmann/protobuf-go/ticket/ticket"
	"github.com/go-godin/example-structure/internal/ticket"
	"github.com/go-kit/kit/endpoint"
)

func CreateEndpoint(svc ticket.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.CreateTicketRequest)

		// CONNECT THE DOMAIN LAYER HERE!
		t := svc.Create(ctx, req.Title)

		resp := &pb.CreateTicketResponse{
			Ticket: &pb.Ticket{
				Title: t.Title,
			},
		}

		return resp, nil
	}
}
