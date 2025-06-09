package decoder

import (
	"encoding/json"
	"io"
)

/*
expects a reader with an inbound JSON payload
that will be decoded into the target struct pointer
*/
func Decode(r io.Reader, target any) error {
	d := json.NewDecoder(r)
	if err := d.Decode(target); err != nil {
		return err
	}
	return nil
}

func Encode(w io.Writer, payload any) error {
	d := json.NewEncoder(w)
	if err := d.Encode(payload); err != nil {
		return err
	}
	return nil
}
