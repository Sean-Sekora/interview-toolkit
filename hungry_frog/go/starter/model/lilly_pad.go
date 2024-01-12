package model

type LillyPad struct {
	Id          int
	IsThereFood bool
	Paths       []*LillyPad
}
