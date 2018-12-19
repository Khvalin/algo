package main

func trafficLights1D(roadMap [][]int) (t int) {
	x := 0
	for i := 0; i < len(roadMap); i++ {
		stopInfo := roadMap[i]
		d := stopInfo[0] - x
		t += d
		r := (t % stopInfo[1])
		if r%2 == 1 {
			t += (r+1)*stopInfo[1] - t
		}
		x += d
	}

	return
}
