package kata

func Cake(recipe, available map[string]int) int {
	cakes := 0
	counter := 1

	for r := range recipe {
		totalCakes := 0

		for available[r] >= recipe[r] {
			totalCakes++
			available[r] -= recipe[r]
		}

		if counter == 1 {
			cakes = totalCakes
		}

		if totalCakes < cakes {
			cakes = totalCakes
		}

		counter++
	}

	return cakes
}

/*
// must return 2
cakes({flour: 500, sugar: 200, eggs: 1}, {flour: 1200, sugar: 1200, eggs: 5, milk: 200});
// must return 0
cakes({apples: 3, flour: 300, sugar: 150, milk: 100, oil: 100}, {sugar: 500, flour: 2000, milk: 2000});

NEED flour = 500
AVAILABLE flour = 1200
1200 > 500 ? true
1200 - 500 = 700
700 > 500 = true
700 - 500 = 200
200 > 500 = false
STOP

NEED apple = 3
AVAILABLE apple = 0
0 > 3 ? false
STOP
*/
