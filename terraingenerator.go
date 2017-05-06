package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Terrain is just the world meta? data
type Terrain struct {
	Name  string
	SizeX int
	SizeY int
}

func main() {
	terrain := initialise()    // Constructs a terrain instance
	world := generate(terrain) // Fills out the world array
	draw(world, terrain)       // Draws the world array
}

func generate(terrain *Terrain) [][]int32 {
	log.Println("Generating world.")
	// Generate the 2d slice of sizeY rows, and sizeX columns
	// Traverse the 2D array and set random numbers
	// Confused by the X,Y coordinates? Rows are Y. Columns X
	world := make([][]int32, terrain.SizeY)
	for y := 0; y < terrain.SizeY; y++ {
		world[y] = make([]int32, terrain.SizeX)
		for x := 0; x < terrain.SizeX; x++ {
			world[y][x] = rand.Int31n(100)
		}
	}
	return world
}

// draw should print the world to the console.
func draw(world [][]int32, terrain *Terrain) {
	log.Println("Drawing world.")
	for y := 0; y < terrain.SizeY; y++ {
		log.Println(world[y])
	}
}

// initialise returns an instance of Terrain
func initialise() *Terrain {
	log.Println("Initialising.")
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
