package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/fogleman/gg"
)

// Vertex represents a point in the terrain.
type Vertex struct {
	X, Y int
}

// Terrain is just the world metadata
type Terrain struct {
	Name          string
	SizeX         int
	SizeY         int
	NumberOfPeaks int
	TerrainTypes  map[string]int
	PeakLocations []Vertex
}

var threadCounter = 0
var waitGroup sync.WaitGroup

func main() {
	defer timeTrack(time.Now(), "main")

	terrain := initialise()    // Constructs a terrain instance. Gets all the command line arguments, etc.
	world := generate(terrain) // Generates the world terrain.
	draw(world, terrain)       // Draws the world array and outputs to .png
}

func generate(terrain *Terrain) [][]int {
	defer timeTrack(time.Now(), "generate")
	log.Println("Generating world.")

	// Generates the 2d slice of sizeY rows, and sizeX columns
	world := make([][]int, terrain.SizeY)
	for y := 0; y < terrain.SizeY; y++ {
		world[y] = make([]int, terrain.SizeX)
	}
	waitGroup.Add(terrain.NumberOfPeaks)
	// Generate a go thread for the number of peaks
	for i := 0; i < terrain.NumberOfPeaks; i++ {
		go generatePeaks(&world, terrain)
	}
	waitGroup.Wait()
	return world
}

// generateMountains generates a random number of mountain peaks for the map.
func generatePeaks(world *[][]int, terrain *Terrain) {
	threadCounter++
	defer waitGroup.Done()
	defer timeTrack(time.Now(), strconv.Itoa(threadCounter))
	x := rand.Intn(terrain.SizeX)
	y := rand.Intn(terrain.SizeY)
	// Deference world, and set the type to mountain.
	(*world)[y][x] = terrain.TerrainTypes["Peak"] // sets the mountain peak
	// Capture all of the peak locations.
	terrain.PeakLocations = append(terrain.PeakLocations, Vertex{X: x, Y: y})

	// surround peak with mountain
	setTerrain(world, x-1, y-1, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)
	setTerrain(world, x-1, y, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)
	setTerrain(world, x-1, y+1, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)
	setTerrain(world, x+1, y-1, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)
	setTerrain(world, x+1, y, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)
	setTerrain(world, x+1, y+1, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)
	setTerrain(world, x, y+1, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)
	setTerrain(world, x, y-1, terrain.TerrainTypes["Mountain"], terrain.SizeX, terrain.SizeY)

}

// draw outputs the terrain to a .png file.
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
			case 5: // Peak / Snow
				dc.SetRGB255(240, 240, 240)
			}
			dc.DrawRectangle(float64(x*10), float64(y*10), 10.0, 10.0)
			dc.Fill()
			dc.Pop()
		}
	}

	t := time.Now().Format("240405")
	filename := "output/" + t + ".png"
	log.Println(filename)
	// writing a file takes about 20ms
	dc.SavePNG(filename)
	fmt.Println(terrain)
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
	numberOfPeaks, err := strconv.Atoi(os.Args[4]) // How to default to 1? numberOfPeaks = 1
	if err != nil {
		panic(err)
	}

	return &Terrain{
		Name:          name,
		SizeX:         sizeX,
		SizeY:         sizeY,
		NumberOfPeaks: numberOfPeaks,
		TerrainTypes: map[string]int{
			"DeepWater":    0,
			"ShallowWater": 1,
			"FieldLow":     2,
			"FieldHigh":    3,
			"Mountain":     4,
			"Peak":         5,
		},
	}
}

func formMountain() {

}

// setTerrain sets the Y,X coordinate terrain type.
func setTerrain(world *[][]int, x, y, terrainType, maxX, maxY int) {
	// Check that the coordinates are within the map boundaries.
	if (y >= 0) && (y < maxY) && (x >= 0) && (x < maxX) {
		(*world)[y][x] = terrainType
	}
}

// timeTrack taken from stathat.com
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("function %s took %s", name, elapsed)
}
