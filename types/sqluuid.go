package types

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

// SqlUuid -> type to use as binary uuid in BBDD
type SqlUuid uuid.UUID

// StringToSqlUuid -> parse string to MYTYPE
func StringToSqlUuid(s string) (SqlUuid, error) {
	id, err := uuid.Parse(s)
	return SqlUuid(id), err
}

// New -> Creates a new SqlUuid
func (my *SqlUuid) New() {
	u := uuid.New()
	*my = SqlUuid(u)

}

// String -> String Representation of Binary16
func (my SqlUuid) String() string {
	return uuid.UUID(my).String()
}

// GormDataType -> sets type to binary(16)
func (my SqlUuid) GormDataType() string {
	return "binary(16)"
}

func (my SqlUuid) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(my)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (my *SqlUuid) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*my = SqlUuid(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (my *SqlUuid) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = SqlUuid(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (my SqlUuid) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
