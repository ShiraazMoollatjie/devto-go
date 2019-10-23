package devtogo

import "time"

// Empty time is a type where the time can be an empty string.
type emptyTime struct {
	*time.Time
}

// UnmarshalJSON is a copy of the time.Time UnmarshalJSON function
// with an extra check for an empty string
func (m *emptyTime) UnmarshalJSON(data []byte) error {

	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = emptyTime{&tt}
	return err
}
