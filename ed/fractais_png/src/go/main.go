package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func circulo(pen Pen, raio float64) {
	if raio < 1{
		return
	}
	pen.DrawCircle(raio)

	for range 6{
		pen.Up()
		pen.Walk(raio)
		pen.Down()
		circulo(pen, raio*0.3)
		pen.Walk(-raio)
		pen.Right(60)
	}
}

/*func arvere(pen *Pen, dist float64) {
	if dist < 10 {
		if ri(0, 50) == 0 {
			pen.SetRGB(200, 0, 0)
			pen.FillCircle(10)
		}
		return
	}
	ang_dir := ri(10, 40)
	ang_esq := ri(10, 40)

	pen.SetLineWidth(dist / 5)
	pen.SetRGB(80, 50, 35)
	pen.Walk(dist)
	pen.Right(ang_dir)
	arvere(pen, dist*(ri(80, 85)/100))
	pen.Left(ang_dir + ang_esq)
	arvere(pen, dist*(ri(80, 85)/100))
	pen.Right(ang_esq)
	pen.SetRGB(20, 90, 35)
	pen.Walk(-dist)
}*/

func main() {
	pen := NewPen(600, 600)
	pen.SetRGB(0, 0, 0)
	pen.SetHeading(90)
	pen.SetPosition(300, 300)

	/*dist := 300.0
	for dist > 10 {
		pen.Walk(dist)
		dist -= 3
		pen.Right(90)

		for i := 0; i < 10; i++ {
			r := max(0, min(255, rand.Float64()*dist+ri(-10, 10)))
			g := max(0, min(255, rand.Float64()*dist+ri(-10, 10)))
			b := max(0, min(255, rand.Float64()*dist+ri(-10, 10)))
			pen.SetRGB(r, g, b)
		}
		/*r := dist + 65
		g := dist + 87
		b := dist + 63

		pen.SetRGB(r, g, b)
	}*/

	/*pen.SetHeading(0)
	pen.Walk(100)
	pen.Right(90)
	pen.SetRGB(200, 0, 0)
	pen.Walk(200)*/
	//arvere(pen, 80)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
