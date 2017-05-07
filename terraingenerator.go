package main

import (
	"fmt"
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
	TerrainTypes := map[int]string{
		0: "Water",
		1: "Water",
		2: "Field",
		3: "Field",
		4: "Mountain",
	}

	terrain := initialise()                  // Constructs a terrain instance
	world := generate(terrain, TerrainTypes) // Fills out the world array
	draw(world, terrain)                     // Draws the world array
}

func generate(terrain *Terrain, terrainTypes map[int]string) [][]int {
	log.Println("Generating world.")

	// Generate the 2d slice of sizeY rows, and sizeX columns
	// Traverse the 2D array and set random numbers
	// Confused by the X,Y coordinates? Rows are Y. Columns X
	world := make([][]int, terrain.SizeY)
	for y := 0; y < terrain.SizeY; y++ {
		world[y] = make([]int, terrain.SizeX)
		for x := 0; x < terrain.SizeX; x++ {
			world[y][x] = rand.Intn(len(terrainTypes)) // Sets each element to one of the terrain types.
		}
	}
	return world
}

// draw should print the world to the console.
func draw(world [][]int, terrain *Terrain) {
	log.Println("Drawing world.")
	for y := range world {
		log.Println(world[y])
	}
	for y := range world {
		fmt.Println(world[y])
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
