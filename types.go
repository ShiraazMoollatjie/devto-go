package devtogo

import "time"

// Empty time is a type where the time can be an empty string.
type emptyTime struct {
	*time.Time
}

func (m *emptyTime) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" || string(data) == `""` {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = emptyTime{&tt}
	return err
}
