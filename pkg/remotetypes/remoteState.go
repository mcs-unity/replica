package remotetypes

import "time"

type RemoteState struct {
	Online    bool
	Timestamp time.Time
}
