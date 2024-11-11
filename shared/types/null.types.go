package types

import (
	"database/sql"
	"encoding/json"
)

// * has to implement interfaces
// *   json.Marshaller &
// *   json.UnMarshaller

type SNullString struct {
	sql.NullString
}

func (ns *SNullString) UnmarshalJSON(b []byte) error {
	// err := json.Unmarshal(b, &ns.String)
	// ns.Valid = (err == nil)
	// return err
	var s *string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s != nil {
		ns.Valid = true
		ns.String = *s
	} else {
		ns.Valid = false
	}
	return nil
}

func (s SNullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(""), nil
}

type SNullInt32 struct {
	sql.NullInt32
}

func (ns *SNullInt32) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.Int32)
	ns.Valid = (err == nil)
	return err
}

func (s SNullInt32) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.Int32)
	}
	return []byte(""), nil
}
