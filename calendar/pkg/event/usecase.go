package event

import (
	"context"
)

type UseCase interface {
	CreateEvent(ctx context.Context, e *Event) error
	GetEventByID(ctx context.Context, id ID) (*Event, error)
	UpdateEvent(ctx context.Context, e *Event) error
	DeleteEvent(ctx context.Context, id ID) error
	ListEvents(ctx context.Context) ([]*Event, error)
}
