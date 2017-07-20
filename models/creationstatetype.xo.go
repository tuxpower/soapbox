// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// CreationStateType is the 'creation_state_type' enum type from schema 'public'.
type CreationStateType uint16

const (
	// CreationStateTypeCreateInfrastructureWait is the 'CREATE_INFRASTRUCTURE_WAIT' CreationStateType.
	CreationStateTypeCreateInfrastructureWait = CreationStateType(1)

	// CreationStateTypeSucceeded is the 'SUCCEEDED' CreationStateType.
	CreationStateTypeSucceeded = CreationStateType(2)

	// CreationStateTypeFailed is the 'FAILED' CreationStateType.
	CreationStateTypeFailed = CreationStateType(3)
)

// String returns the string value of the CreationStateType.
func (cst CreationStateType) String() string {
	var enumVal string

	switch cst {
	case CreationStateTypeCreateInfrastructureWait:
		enumVal = "CREATE_INFRASTRUCTURE_WAIT"

	case CreationStateTypeSucceeded:
		enumVal = "SUCCEEDED"

	case CreationStateTypeFailed:
		enumVal = "FAILED"
	}

	return enumVal
}

// MarshalText marshals CreationStateType into text.
func (cst CreationStateType) MarshalText() ([]byte, error) {
	return []byte(cst.String()), nil
}

// UnmarshalText unmarshals CreationStateType from text.
func (cst *CreationStateType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "CREATE_INFRASTRUCTURE_WAIT":
		*cst = CreationStateTypeCreateInfrastructureWait

	case "SUCCEEDED":
		*cst = CreationStateTypeSucceeded

	case "FAILED":
		*cst = CreationStateTypeFailed

	default:
		return errors.New("invalid CreationStateType")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for CreationStateType.
func (cst CreationStateType) Value() (driver.Value, error) {
	return cst.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for CreationStateType.
func (cst *CreationStateType) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid CreationStateType")
	}

	return cst.UnmarshalText(buf)
}
