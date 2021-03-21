package battleshipgameng

func StartGame() (ships map[string][]location) {
	clearBoard()
	generateShips()

	return ships
}

func Shoot() {

}
