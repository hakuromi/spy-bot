package models

type Game struct {
	Players []Player
	SpyID   int64
	Hero    string
	Active  bool
	Mode    string
}
