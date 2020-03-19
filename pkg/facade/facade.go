package facade

type band struct {
	bassist   *bassist
	drummer   *drummer
	guitarist *guitarist
	vocalist  *vocalist
}

func (b *band) PlayCoolSong() {
	b.guitarist.playCoolOpening(b.guitarist.name)
	b.drummer.startPlaying(b.drummer.name)
	b.vocalist.singCouplet(b.vocalist.name)
	b.bassist.followTheDrums(b.bassist.name)
}

// RockBand creating
func RockBand() *band {
	return &band{
		bassist:   &bassist{"Джонни"},
		drummer:   &drummer{"Вилли"},
		guitarist: &guitarist{"Марк"},
		vocalist:  &vocalist{"Виталик"},
	}
}
