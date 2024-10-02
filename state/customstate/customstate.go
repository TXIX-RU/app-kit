package customstate

import (
	"github.com/txix-open/walx/state"
)

type nothing struct{}

type customstateInterface[EventType any] interface {
	Apply(EventType) error
}
type CustomState[EventType any] struct {
	name         string
	streamSuffix []byte
	mutator      state.Mutator
	state        customstateInterface[EventType]
}

func New[EventType any](state customstateInterface[EventType], name string) *CustomState[EventType] {
	cs := CustomState[EventType]{
		name:         name,
		streamSuffix: []byte(name),
		state:        state,
	}
	return &cs
}

func (cs *CustomState[EventType]) SetMutator(mutator state.Mutator) {
	cs.mutator = mutator
}

func (cs *CustomState[EventType]) StateName() string {
	return cs.name
}

func (cs *CustomState[EventType]) Apply(data []byte) (any, error) {
	var evnet EventType

	if err := state.UnmarshalEvent(data, &evnet); err != nil {
		return nothing{}, err
	}
	// fmt.Printf("\n>>>%s\n\n%+v\n", string(data), evnet)
	err := cs.state.Apply(evnet)
	if err != nil {
		return nothing{}, err
	}
	return nothing{}, nil
}

func (cs *CustomState[EventType]) Write(event EventType) error {
	_, err := state.ApplyWithStreamSuffix[nothing](cs.mutator, event, cs.streamSuffix)
	return err
}
