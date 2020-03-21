package facade

type bandBassist interface {
	FollowTheDrums(name string)
	ChangeRhythm(name string)
	StopPlaying(name string)
}

type bandDrummer interface {
	StartPlaying(name string)
	StopPlaying(name string)
}

type bandGuitarist interface {
	PlayCoolOpening(name string)
	PlayCoolRiffs(name string)
	PlayAnotherCoolRiffs(name string)
	PlayIncrediblyCoolSolo(name string)
	PlayFinalAccord(name string)
}

type bandVocalist interface {
	SingCouplet(name string)
	SingChorus(name string)
}

type Band interface {
	PlayCoolSong()
}

type band struct {
	bassist   bandBassist
	drummer   bandDrummer
	guitarist bandGuitarist
	vocalist  bandVocalist
}

func (b *band) PlayCoolSong() {
	b.guitarist.PlayCoolOpening("john")
	b.drummer.StartPlaying("Vitalik")
	b.vocalist.SingCouplet("Boris")
	b.bassist.FollowTheDrums("Summer")
}

// RockBand creating
func RockBand(bassist bandBassist, drummer bandDrummer, guitarist bandGuitarist, vocalist bandVocalist) Band {
	return &band{
		bassist:   bassist,
		drummer:   drummer,
		guitarist: guitarist,
		vocalist:  vocalist,
	}
}
