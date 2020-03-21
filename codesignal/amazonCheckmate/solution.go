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

	illegal := [8][8]bool{}
	setIllegal := func(r, c int) {
		if r >= 0 && c >= 0 && r < 8 && c < 8 {
			illegal[r][c] = true
		}
	}

	isAttackedByKing := func(r, c int) bool {
		return (abs(r-kr) <= 1 && abs(c-kc) <= 1)
	}

	// king
	for r := kr - 1; r <= kr+1; r++ {
		for c := kc - 1; c <= kc+1; c++ {
			setIllegal(r, c)
		}
	}

	// knight
	for r := ar - 2; r <= ar+2; r++ {
		for c := ac - 2; c <= ac+2; c++ {
			if abs(r-ar)+abs(c-ac) == 3 {
				setIllegal(r, c)
			}
		}
	}

	//queen
	for k := -1; k <= 1; k += 2 {
		//rows
		for r := ar; r >= 0 && r < 8; r += k {
			if r == kr && ac == kc {
				break
			}
			setIllegal(r, ac)
		}

		//columns
		for c := ac; c >= 0 && c < 8; c += k {
			if c == kc && ar == kr {
				break
			}
			setIllegal(ar, c)
		}
	}

	for i := -1; i <= 1; i += 2 {
		for j := -1; j <= 1; j += 2 {
			for r, c := ar, ac; r >= 0 && c >= 0 && r < 8 && c < 8; r, c = i+r, j+c {
				if c == kc && r == kr {
					break
				}
				setIllegal(r, c)
			}
		}
	}

	// checkmate positions, check positions, stalemate positions, safe positions
	var mate, check, stale, safe int
	for r := 0; r < 8; r++ {
		for c := 0; c < 8; c++ {
			if isAttackedByKing(r, c) || r == ar && c == ac {
				continue
			}

			canMove := false
			for dr := -1; dr < 2; dr++ {
				for dc := -1; dc < 2; dc++ {
					r1, c1 := r+dr, c+dc
					if dr == 0 && dc == 0 || r1 < 0 || r1 > 7 || c1 < 0 || c1 > 7 {
						continue
					}

					canMove = canMove || !(illegal[r1][c1] && (isAttackedByKing(r1, c1) || r1 != ar || c1 != ac))
				}
			}

			if illegal[r][c] {
				if canMove {
					check++
				} else {
					mate++
					//fmt.Print(string('a'+r)+string('1'+c), " ")
				}
			} else {
				if canMove {
					safe++
					//fmt.Print(string('a'+r)+string('1'+c), " ")
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
	fmt.Println(amazonCheckmate("a1", "g5"))
}
