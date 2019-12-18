package mask

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Timestamp extension time
type Timestamp struct {
	Time time.Time
}

const (
	jsonLayout = "2006-01-02 15:04:05"
)

// Now returns the current local time.
func Now() Timestamp {
	return Timestamp{
		Time: time.Now(),
	}
}

// UnmarshalBSON unmarshal bson
func (t *Timestamp) UnmarshalBSON(data []byte) (err error) {
	var d bson.D
	err = bson.Unmarshal(data, &d)
	if err != nil {
		return err
	}
	if v, ok := d.Map()["t"]; ok {
		t.Time = time.Time{}
		return t.Time.UnmarshalText([]byte(v.(string)))
	}
	return fmt.Errorf("key 't' missing")
}

// MarshalBSON marshal bson
func (t Timestamp) MarshalBSON() ([]byte, error) {
	txt, err := t.Time.MarshalText()
	if err != nil {
		return nil, err
	}
	b, err := bson.Marshal(map[string]string{"t": string(txt)})
	return b, err
}

// UnmarshalJSON unmarshal json
func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 0 || string(data) == "" || string(data) == `""` {
		return
	}
	now, err := time.ParseInLocation(`"`+jsonLayout+`"`, string(data), time.Local)
	*t = Timestamp{
		Time: now,
	}
	return
}

// MarshalJSON marshal json
func (t Timestamp) MarshalJSON() ([]byte, error) {

	b := make([]byte, 0, len(jsonLayout)+2)
	b = append(b, '"')
	b = time.Time(t.Time).AppendFormat(b, jsonLayout)
	b = append(b, '"')
	return b, nil
}
