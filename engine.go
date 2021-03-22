package battleshipgameng

func StartGame() map[string][]string {
	clearBoard()
	generateShips()

	s := make(map[string][]string)

	for ship, locations := range ships {
		for _, loc := range locations {
			s[ship] = append(s[ship], loc.getStringLoc())
		}
	}

	return s
}

func Shoot() {

}
