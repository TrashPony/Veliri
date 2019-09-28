package debug

import (
	"sync"
	"time"
)

var Store = newStore()

func newStore() *MessageStore {
	return &MessageStore{
		messages: make([]*Message, 0),
		Move:     true,
		MoveInit: false,

		MoveEndPoint: false,

		AStartNeighbours: false,
		AStartResult:     false,

		RegionFindDebug: false,
		RegionResult:    false,

		HandAlgorithm: false,

		SearchCollisionLine: false,
	}
}

type MessageStore struct {
	messages            []*Message
	mx                  sync.Mutex
	Move                bool
	MoveInit            bool
	AStartNeighbours    bool
	AStartResult        bool
	RegionFindDebug     bool
	RegionResult        bool
	HandAlgorithm       bool
	SearchCollisionLine bool
	MoveEndPoint        bool
}

type Message struct {
	Type  string
	Color string
	X     int
	Y     int
	ToX   int
	ToY   int
	Size  int
	MapID int
	MS    int64
}

func (s *MessageStore) AddMessage(msgType, color string, x, y, toX, toY, size, mpId int, ms int64) {
	s.mx.Lock()

	s.messages = append(s.messages, &Message{
		msgType,
		color,
		x,
		y,
		toX,
		toY,
		size,
		mpId,
		ms})

	s.mx.Unlock()

	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func (s *MessageStore) GetAllMessages() []*Message {
	s.mx.Lock()
	defer s.mx.Unlock()

	result := s.messages
	s.messages = make([]*Message, 0)

	return result
}
