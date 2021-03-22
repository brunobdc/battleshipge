package battleshipgameng

import (
	"math/rand"
	"time"
)

var board [8][8]string

func clearBoard() {
	board = [8][8]string{
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", " "},
	}
}

var ships map[string][]location

func newShips() {
	ships = map[string][]location{
		"Battleship": {location{}, location{}, location{}, location{}},
		"Cruiser":    {location{}, location{}, location{}},
		"Destroyer":  {location{}, location{}},
	}
}

func generateShips() {
	newShips()

	r := rand.Rand{}
	var x, y, loc int
	var wcolision bool

	for _, locations := range ships {

		for wcolision = false; wcolision; {
			r = *rand.New(rand.NewSource(time.Now().UnixNano()))

			//0 representa horizontal <-> e 1 vertical ↕
			if r.Intn(1) == 0 {
				//7 é o tamanho máximo da horizontal do tabuleiro pois começa contando do 0
				x = r.Intn(7 - len(locations))
				y = r.Intn(7)

				//agora é necessário verificar se tem colisão com outros navios(ships)
				//se sim, retorna-rá diretamente para o loop para tentar o mesmo navio com novos valores
				for i := x; i <= x+len(locations) || wcolision; i++ {
					_, wcolision = colision(y, i)
				}

				//se não ouve colisões coloca o navio(ship) no tabuleiro(board)
				if !wcolision {
					loc = 0
					for i := x; i <= x+len(locations); i++ {
						locations[loc].y = y
						locations[loc].x = i
						loc++
					}
				}

			} else {
				//7 é o tamanho máximo da vertical do tabuleiro pois começa contando do 0
				y = r.Intn(7 - len(locations))
				x = r.Intn(7)

				//agora é necessário verificar se tem colisão com outros navios(ships)
				//se sim, retorna-rá diretamente para o loop para tentar o mesmo navio com novos valores
				for i := y; i <= x+len(locations) || wcolision; i++ {
					_, wcolision = colision(i, x)
				}

				//se não ouve colisões coloca o navio(ship) no tabuleiro(board)
				if !wcolision {
					loc = 0
					for i := y; i <= x+len(locations); i++ {
						locations[loc].y = i
						locations[loc].x = x
						loc++
					}
				}
			}
		}
	}
}

func colision(y, x int) (ship string, hit bool) {
	for ship, locations := range ships {
		for _, loc := range locations {
			if loc.y == y && loc.x == x {
				return ship, false
			}
		}
	}

	return "", true
}
