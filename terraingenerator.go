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

	// Constructs a terrain instance
	terrain := initialise()

	// Generate the 2d slice of sizeY rows, and sizeX columns
	world := make([][]int32, terrain.SizeY)
	for y := 0; y < terrain.SizeY; y++ {
		world[y] = make([]int32, terrain.SizeX)
	}

	// Traverse the 2D array and set random numbers
	for y := 0; y < terrain.SizeY; y++ {
		for x := 0; x < terrain.SizeX; x++ {
			c := rand.Int31n(100)
			world[y][x] = c
		}
	}

	fmt.Println("Generate.")
	fmt.Println("Draw.")
	fmt.Println(terrain)
	fmt.Println(world)
}

func generate() {

}

func draw() {

}

// initialise returns an instance of Terrain
func initialise() *Terrain {
	fmt.Println("Initialising.")
	rand.Seed(time.Now().UTC().UnixNano())

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

	return &Terrain{
		Name:  name,
		SizeX: sizeX,
		SizeY: sizeY,
	}
}
