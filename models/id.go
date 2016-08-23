package models

type Id string

func (id Id) String() string {
	return string(id)
}
