package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func amazonCheckmate(king, amazon string) []int {
	kr, kc := int(king[0]-'a'), int(king[1]-'1')
	ar, ac := int(amazon[0]-'a'), int(amazon[1]-'1')

	isUnderAttack := func(r, c int) bool {
		res := (r == kr && abs(c-kc) < 2) || (c == kc && abs(r-kr) < 2) || (abs(c-kc) == 1 && abs(r-kr) == 1) || (abs(r-ar)+abs(c-ac) == 3) //
		if res {
			return res
		}

		for x := ar - 1; x >= 0; x-- {
			if x == kr && kc == ac {
				break
			}
			if x == r && c == ac {
				res = true
			}
		}

		for x := ar + 1; x < 9; x++ {
			if x == kr && kc == ac {
				break
			}
			if x == r && c == ac {
				res = true
			}
		}

		for y := ac - 1; y >= 0; y-- {
			if y == kc && kr == r {
				break
			}
			if c == ac && y == ar {
				res = true
			}
		}

		for y := ac + 1; y < 8; y++ {
			if y == kc && kr == r {
				break
			}
			if c == ac && y == ar {
				res = true
			}
		}

		x, y := ar-1, ac-1
		for x >= 0 && y >= 0 && x < 8 && y < 8 {
			if x == kr && y == kc {
				break
			}
			res = res || (x == r && y == c)

			x--
			y--
		}

		x, y = ar-1, ac+1
		for x >= 0 && y >= 0 && x < 8 && y < 8 {
			if x == kr && y == kc {
				break
			}
			res = res || (x == r && y == c)

			x--
			y++
		}

		x, y = ar+1, ac-1
		for x >= 0 && y >= 0 && x < 8 && y < 8 {
			if x == kr && y == kc {
				break
			}
			res = res || (x == r && y == c)

			x++
			y--
		}

		x, y = ar+1, ac+1
		for x >= 0 && y >= 0 && x < 8 && y < 8 {
			if x == kr && y == kc {
				break
			}
			res = res || (x == r && y == c)

			x++
			y++
		}

		return res
	}

	// checkmate positions, check positions, stalemate positions, safe positions
	var mate, check, stale, safe int
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if r == kr && c == kc || r == ar && c == ac {
				continue
			}

			canMove := false
			for dr := -1; dr < 2; dr++ {
				for dc := -1; dc < 2; dc++ {
					if dr == 0 && dc == 0 || r+dr < 0 || r+dr > 7 || c+dc < 0 || c+dc > 7 {
						continue
					}

					canMove = canMove || !isUnderAttack(r+dr, c+dc)
				}
			}

			if isUnderAttack(r, c) {
				if canMove {
					check++
				} else {
					mate++
				}
			} else {
				if canMove {
					safe++
					fmt.Print(string('a'+r)+string('1'+c), " ")
				} else {
					stale++
				}
			}

			//fmt.Print(string('a'+r)+string('1'+c), " ")
		}
	}

	return []int{mate, check, stale, safe}
}

func main() {
	fmt.Println(amazonCheckmate("d3", "e4"))
}
