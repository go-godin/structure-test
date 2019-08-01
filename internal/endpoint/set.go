package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	ticket2 "bitbucket.org/jdbergmann/protobuf-go/ticket/ticket"
	"github.com/go-godin/example-structure/internal/ticket"
)

type Set struct {
	CreateEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}

func (Set) Get(context.Context, *ticket2.GetTicketRequest) (*ticket2.GetTicketResponse, error) {
	panic("implement me")
}

func (s Set) Create(ctx context.Context, req *ticket2.CreateTicketRequest) (*ticket2.CreateTicketResponse, error) {
	resp, err := s.CreateEndpoint(ctx, req)
	return resp.(*ticket2.CreateTicketResponse), err
}

func (Set) Delete(context.Context, *ticket2.DeleteTicketRequest) (*ticket2.DeleteTicketResponse, error) {
	panic("implement me")
}

func (Set) List(context.Context, *ticket2.ListTicketsRequest) (*ticket2.ListTicketsResponse, error) {
	panic("implement me")
}

func (Set) SetAssignee(context.Context, *ticket2.SetAssigneeRequest) (*ticket2.SetAssigneeResponse, error) {
	panic("implement me")
}

func (Set) AddPerson(context.Context, *ticket2.AddPersonRequest) (*ticket2.AddPersonResponse, error) {
	panic("implement me")
}

func (Set) RemovePerson(context.Context, *ticket2.RemovePersonRequest) (*ticket2.RemovePersonResponse, error) {
	panic("implement me")
}

func (Set) Update(context.Context, *ticket2.UpdateRequest) (*ticket2.UpdateResponse, error) {
	panic("implement me")
}

func (Set) Close(context.Context, *ticket2.CloseRequest) (*ticket2.CloseResponse, error) {
	panic("implement me")
}

func (Set) UpdateReference(context.Context, *ticket2.UpdateReferenceRequest) (*ticket2.UpdateReferenceResponse, error) {
	panic("implement me")
}

func (Set) UpdateType(context.Context, *ticket2.UpdateTypeRequest) (*ticket2.UpdateTypeResponse, error) {
	panic("implement me")
}

func (Set) UpdateStatus(context.Context, *ticket2.UpdateStatusRequest) (*ticket2.UpdateStatusResponse, error) {
	panic("implement me")
}

func NewEndpointSet(svc ticket.Service) Set {

	var create endpoint.Endpoint
	{
		create = CreateEndpoint(svc)
	}

	return Set{
		CreateEndpoint: create,
	}
}
