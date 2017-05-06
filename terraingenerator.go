package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Terrain is the world
type Terrain struct {
	Name  string
	SizeX int
	SizeY int
}

func main() {
	fmt.Println("Init.")

	initialise()

	// Get the input from the command line.
	var name = os.Args[1]

	sizeX, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	sizeY, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}

	// Construct a terrain instance
	terrain := NewTerrain(name, sizeX, sizeY)
	var x = terrain.SizeX
	var y = terrain.SizeY

	// Generate the 2d slice of sizeX sizeY
	world := make([][]int32, y)
	for i := 0; i < y; i++ {
		world[i] = make([]int32, x)
	}

	// Traverse the 2D array and set random numbers
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			c := rand.Int31n(100)
			world[i][j] = c
		}
	}

	fmt.Println("Generate.")
	fmt.Println("Draw.")
	fmt.Println(terrain)
	fmt.Println(world)
}

func initialise() {
	fmt.Println("Initialising.")
	rand.Seed(time.Now().UTC().UnixNano())
}

func generate() {

}

func draw() {

}

// NewTerrain returns an instance of Terrain
func NewTerrain(name string, sizeX int, sizeY int) *Terrain {
	return &Terrain{
		Name:  name,
		SizeX: sizeX,
		SizeY: sizeY,
	}
}
