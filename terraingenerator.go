package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fogleman/gg"
)

// Terrain is just the world meta? data
type Terrain struct {
	Name  string
	SizeX int
	SizeY int
}

func main() {
	defer timeTrack(time.Now(), "main")
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
	defer timeTrack(time.Now(), "generate")
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
	defer timeTrack(time.Now(), "draw")
	log.Println("Drawing world.")
	dc := gg.NewContext(800, 800)
	for y := 0; y < terrain.SizeY; y++ {
		for x := 0; x < terrain.SizeX; x++ {
			dc.Push()
			switch world[y][x] {
			case 0: // water
				dc.SetRGB255(0, 0, 255)
			case 1: // water
				dc.SetRGB255(0, 50, 255)
			case 2: // field
				dc.SetRGB255(0, 255, 50)
			case 3: // field
				dc.SetRGB255(10, 220, 10)
			case 4: // mountain
				dc.SetRGB255(60, 60, 60)
			}
			dc.DrawRectangle(float64(x*10), float64(y*10), 10.0, 10.0)
			dc.Fill()
			dc.Pop()

		}
	}
	t := time.Now()
	f := t.String()
	filename := "output/" + f + ".png"
	log.Println(filename)
	dc.SavePNG(filename)

}

// initialise returns an instance of Terrain
func initialise() *Terrain {
	defer timeTrack(time.Now(), "initialise")
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

// timeTrack taken from stathat.com
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("function %s took %s", name, elapsed)
}
