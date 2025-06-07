package shared

type State uint8

const (
	DOWN    State = iota
	UP      State = iota
	UNKNOWN State = iota
)
