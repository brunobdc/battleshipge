package battleshipgameng

import (
	"math/rand"
	"time"
)

var board = [8][8]string{
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " ", " "},
}

type location struct {
	y   int
	x   int
	hit bool
}

var ships = map[string][]location{
	"Battleship": {location{}, location{}, location{}, location{}},
	"Cruiser":    {location{}, location{}, location{}},
	"Destroyer":  {location{}, location{}},
}

func generateShips() {
	r := &rand.Rand{}
	var i, line, x, y int

	for _, ship := range ships {
		i = 0
		for i < len((ship)) {
			r = rand.New(rand.NewSource(time.Now().UnixNano()))

			//0 representa horizontal e 1 vertical
			line = r.Intn(1)
			if line == 0 {
				//7 é o tamanho máximo da horizontal do tabuleiro pois começa contando do 0
				x = r.Intn(7 - len(ship))
				y = r.Intn(7)

				//agora é necessário verificar se tem colisão com outros navios
				//adiciona 1 à i se não tiver colisões
				//senão retorna só retorna para o loop
			} else {
				//7 é o tamanho máximo da vertical do tabuleiro pois começa contando do 0
				y = r.Intn(7 - len(ship))
				x = r.Intn(7)

				//agora é necessário verificar se tem colisão com outros navios
				//adiciona 1 à i se não tiver colisões
				//senão retorna só retorna para o loop
			}
		}
	}
}
