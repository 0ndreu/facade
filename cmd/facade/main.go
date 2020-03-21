package main

import (
	"../../pkg/bassist"
	"../../pkg/drummer"
	"../../pkg/facade"
	"../../pkg/guitarist"
	"../../pkg/vocalist"
)

func main() {
	bassist := bassist.NewBassist()
	drummer := drummer.NewDrummer()
	guitarist := guitarist.NewGuitarist()
	vocalist := vocalist.NewVocalist()
	f := facade.RockBand(bassist, drummer, guitarist, vocalist)
	f.PlayCoolSong()
}
