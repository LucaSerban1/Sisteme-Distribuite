package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Robot struct {
	Position Point
	Activ    bool
}

func rotunjire(val float64) float64 {
	return math.Round(val*100) / 100
}

func getDistance(p1, p2 Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

func computeNewPositions(roboti []Robot, raza float64) []Point {
	newPositions := make([]Point, len(roboti))
	for i := 0; i < len(roboti); i++ {
		if !roboti[i].Activ {
			continue
		}

		sumX, sumY := roboti[i].Position.X, roboti[i].Position.Y
		count := 1.0

		for j := 0; j < len(roboti); j++ {
			if i != j && roboti[j].Activ {
				distance := getDistance(roboti[i].Position, roboti[j].Position)

				if distance <= raza {
					sumX += roboti[j].Position.X
					sumY += roboti[j].Position.Y
					count++
				}
			}
		}
		newPositions[i] = Point{
			X: sumX / count,
			Y: sumY / count,
		}
	}
	return newPositions
}

func updatePositions(roboti []Robot, newPositions []Point) {
	for i := 0; i < len(roboti); i++ {
		if roboti[i].Activ {
			roboti[i].Position = newPositions[i]
		}
	}
}

func updateRobots(roboti []Robot) {
	for i := 0; i < len(roboti)-1; i++ {
		if !roboti[i].Activ {
			continue
		}

		x1 := rotunjire(roboti[i].Position.X)
		y1 := rotunjire(roboti[i].Position.Y)

		for j := i + 1; j < len(roboti); j++ {
			if !roboti[j].Activ {
				continue
			}

			x2 := rotunjire(roboti[j].Position.X)
			y2 := rotunjire(roboti[j].Position.Y)

			if x1 == x2 && y1 == y2 {
				roboti[j].Activ = false
			}
		}
	}
	
}

func verificaRobotiActivi(roboti []Robot) (int, Point) {
	robotiActivi := 0
	var ultimaPozitie Point

	for _, r := range roboti {
		if r.Activ {
			robotiActivi++
			ultimaPozitie = r.Position
		}
	}
	return robotiActivi, ultimaPozitie
}

func lcmAveraging(roboti []Robot, raza float64) {
	for i := 0; i < 50; i++ {
		newPositions := computeNewPositions(roboti, raza)
		updatePositions(roboti, newPositions)
		updateRobots(roboti)

		robotiActivi, ultimaPozitie := verificaRobotiActivi(roboti)

		fmt.Printf("Iteratia %d: %d roboti activi\n", i+1, robotiActivi)

		if robotiActivi <= 1 {
			fmt.Printf("Roboți rămași: %d\n", robotiActivi)
			fmt.Printf("Ultima poziție validă: (%.4f, %.4f)\n", ultimaPozitie.X, ultimaPozitie.Y)
			break
		}
	}
}

func main() {
	razaTest := 15.0

	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println("Eroare la citirea numărului de roboți")
		return
	}

	roboti := make([]Robot, 0, n)
	for i := 0; i < n; i++ {
		var x, y float64
		if _, err := fmt.Scan(&x, &y); err != nil {
			fmt.Println("Eroare la citirea coordonatelor")
			return
		}
		roboti = append(roboti, Robot{
			Position: Point{X: x, Y: y},
			Activ:    true,
		})
	}
	lcmAveraging(roboti, razaTest)
}