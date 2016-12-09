package canvas2d

import "time"
import "math/rand"

var seed bool = true

func Random(min, max int) int {
	if seed {
		rand.Seed(time.Now().Unix())
		seed = false
	}
	max++
	return min + rand.Intn(max-min)
}