package pkg

import (
	"encoding/json"
	"io"
)

func DisallowUnkownJSONProps(body io.Reader, schema any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(schema); err != nil {
		return err
	}

	return nil
}
