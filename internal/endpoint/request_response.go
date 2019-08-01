package endpoint

// request_response introduces intermediary data types in order to encode domain-errors and make them available in
// the endpoints for proper transport-layer error encoding :)

import (
	pb "bitbucket.org/jdbergmann/protobuf-go/ticket/ticket"
)

type CreateTicketRequest struct {
	AccountId       int64
	Reference       pb.Reference
	Source          pb.Source
	Type            pb.Type
	Status          pb.Status
	Title           string
	Description     string
	MediaIds        []string
	AppointmentIds  []string
	Priority        pb.Priority
	DueDate         int64
	Assignee        pb.Assignee
	InvolvedPersons []pb.InvolvedPerson
}

type UpdateRequest struct {
	Id          string
	Title       string
	Description string
}

type CloseRequest struct {
	Id string
}

type UpdateReferenceRequest struct {
	Id        string
	Reference pb.Reference
}

type DeleteTicketRequest struct {
	Id string
}

type SetAssigneeRequest struct {
	Id       string
	Assignee pb.Assignee
}

type AddPersonRequest struct {
	Id     string
	Person pb.InvolvedPerson
}

type RemovePersonRequest struct {
	Id     string
	Person pb.InvolvedPerson
}

type UpdateTypeRequest struct {
	Id   string
	Type pb.Type
}

type UpdateStatusRequest struct {
	Id     string
	Status pb.Status
}

type GetTicketResponse struct {
	Ticket pb.Ticket
	Err    error
}

type ListTicketsResponse struct {
	Tickets []pb.Ticket
	Err     error
}

type CreateTicketResponse struct {
	Ticket pb.Ticket
	Err    error
}

type UpdateResponse struct {
	Ticket pb.Ticket
	Err    error
}

type CloseResponse struct {
	Err error
}

type UpdateReferenceResponse struct {
	Ticket pb.Ticket
	Err    error
}

type DeleteTicketResponse struct {
	Err error
}

type SetAssigneeResponse struct {
	Ticket pb.Ticket
	Err    error
}

type AddPersonResponse struct {
	Ticket pb.Ticket
	Err    error
}

type RemovePersonResponse struct {
	Ticket pb.Ticket
	Err    error
}

type UpdateTypeResponse struct {
	Ticket pb.Ticket
	Err    error
}

type UpdateStatusResponse struct {
	Ticket pb.Ticket
	Err    error
}
