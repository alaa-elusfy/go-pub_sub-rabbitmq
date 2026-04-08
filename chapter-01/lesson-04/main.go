package main

func (u user) doBattles(subCh <-chan move) []piece {
	piecesInBattle := []piece{}
	for mv := range subCh {
		for _, p := range u.pieces {
			if p.location == mv.piece.location {
				piecesInBattle = append(piecesInBattle, p)
			}
		}
	}
	return piecesInBattle
}

type user struct {
	name   string
	pieces []piece
}

type move struct {
	userName string
	piece    piece
}

type piece struct {
	location string
	name     string
}

func (u user) march(p piece, publishCh chan<- move) {
	publishCh <- move{
		userName: u.name,
		piece:    p,
	}
}

func distributeBattles(publishCh <-chan move, subChans []chan move) {
	for mv := range publishCh {
		for _, subCh := range subChans {
			subCh <- mv
		}
	}
}
