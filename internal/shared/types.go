package shared

type State uint8

const (
	UNKNOWN State = iota
	DOWN    State = iota
	UP      State = iota
	ERROR   State = iota
)
