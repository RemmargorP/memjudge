package models

type Rights struct {
	Rights byte `bson:"rights"`
}

const (
	R_Read = 1 << iota
	R_Write
	R_SetRights
)
