package main

import (
	"facade/pkg/bassist"
	"facade/pkg/drummer"
	"facade/pkg/facade"
	"facade/pkg/guitarist"
	"facade/pkg/vocalist"
)

func main() {
	bassist := bassist.NewBassist()
	drummer := drummer.NewDrummer()
	guitarist := guitarist.NewGuitarist()
	vocalist := vocalist.NewVocalist()
	f := facade.RockBand(bassist, drummer, guitarist, vocalist)
	f.PlayCoolSong()
}
