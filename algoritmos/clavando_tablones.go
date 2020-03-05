package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func crear_array_tablones(n int, m int) ([]int, []int) {
	a := make([]int, n)
	b := make([]int, n)
	for k := range a {
		random_1 := rand.Intn(2*m) + 1
		random_2 := rand.Intn(2*m) + 1
		if random_1 < random_2 {
			a[k] = random_1
			b[k] = random_2
		} else {
			a[k] = random_2
			b[k] = random_1
		}
	}
	return a, b
}

func crear_array_clavos(n int, m int) []int {
	c := make([]int, n)
	for k := range c {
		c[k] = rand.Intn(2*m) + 1
	}
	return c
}

func solucion(a, b, c []int, n, m int) int {
	leer_siguiente_clavo := true
	cantidad_tablones_clavados := 0
	j := 0 // clavos indice
	i := 0 // tablones indice
	tablones_clavados := ""
	for leer_siguiente_clavo {
		leer_siguiente_tablon := true
		for leer_siguiente_tablon {
			if a[i] <= c[j] && c[j] <= b[i] {
				tablones_clavados += "(" + strconv.Itoa(a[i]) + "," + strconv.Itoa(b[i]) + ") "
				cantidad_tablones_clavados++
				i++
				if i == n {
					fmt.Println("Usando", j+1, "clavo/s se clavan los tablones: ", tablones_clavados)
					return j + 1
				}
			} else {
				fmt.Println("Usando", j+1, "clavo/s se clavan los tablones:", tablones_clavados)
				leer_siguiente_tablon = false
			}
		}
		j++
		if j == m {
			leer_siguiente_clavo = false
		}
	}
	return -1
}

func main() {

	rand.Seed(time.Now().UnixNano())
	var n = rand.Intn(10-1) + 1 //uso 10 pero deberia ser 30000
	var m = rand.Intn(10-1) + 1 //uso 10 pero deberia ser 30000
	fmt.Println("N es: ", n)
	fmt.Println("M es: ", m)
	a, b := crear_array_tablones(n, m)
	c := crear_array_clavos(m, m)
	fmt.Println("A es: ", a)
	fmt.Println("B es: ", b)
	fmt.Println("C es: ", c)

	// Ejemplo en el examen:
	//n := 4
	//m := 5
	//a := []int{1, 4, 5, 8}
	//b := []int{4, 5, 9, 10}
	//c := []int{4, 6, 7, 10, 2}

	solucion := solucion(a, b, c, n, m)
	if solucion != -1 {
		fmt.Println("Número minimo de clavos para clavar todos los tablones: ", solucion)
	} else {
		fmt.Println("No es posible clavar todos los tablones.")
	}
}
