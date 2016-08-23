package models

type Rights struct {
	Rights byte
}

const (
	R_Read = 1
	R_Write
	R_SetRights
)
