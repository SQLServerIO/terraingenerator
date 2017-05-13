package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"runtime"

	"github.com/fogleman/gg"
)

// Vertex represents a point in the terrain.
type Vertex struct {
	X, Y int
}

// Chart is the world's representation. Should multiple Charts make up a world?
type Chart struct {
	Name          string
	SizeX         int
	SizeY         int
	NumberOfPeaks int
	TerrainTypes  map[string]int
	PeakLocations []Vertex
	Graph         [][]int
}

var threadCounter = 0
var waitGroup sync.WaitGroup

func main() {
	defer timeTrack(time.Now(), "main")
	runtime.GOMAXPROCS(runtime.NumCPU())
	var chart Chart
	chart.initialise() // Constructs a terrain instance. Gets all the command line arguments, etc.
	chart.generate()   // Generates the world terrain.
	chart.draw()       // Draws the world array and outputs to .png
}

func (c *Chart) generate() {
	defer timeTrack(time.Now(), "generate")
	log.Println("Generating world.")

	// Generates the 2d slice of sizeY rows, and sizeX columns
	c.Graph = make([][]int, c.SizeY)
	for y := 0; y < c.SizeY; y++ {
		c.Graph[y] = make([]int, c.SizeX)
	}
	waitGroup.Add(c.NumberOfPeaks)
	// Generate a go thread for the number of peaks
	for i := 0; i < c.NumberOfPeaks; i++ {
		go c.generatePeaks()
	}
	waitGroup.Wait()
}

// generateMountains generates a random number of mountain peaks for the map.
func (c *Chart) generatePeaks() {
	threadCounter++
	defer waitGroup.Done()
	defer timeTrack(time.Now(), strconv.Itoa(threadCounter))
	x := rand.Intn(c.SizeX)
	y := rand.Intn(c.SizeY)
	// Deference world, and set the type to mountain.
	c.Graph[y][x] = c.TerrainTypes["Peak"] // sets the mountain peak
	// Capture all of the peak locations.
	c.PeakLocations = append(c.PeakLocations, Vertex{X: x, Y: y})

	// Surround peaks with mountains. This should be randomized in the future.
	c.setTerrain(x-1, y-1, c.TerrainTypes["Mountain"])
	c.setTerrain(x-1, y, c.TerrainTypes["Mountain"])
	c.setTerrain(x-1, y+1, c.TerrainTypes["Mountain"])
	c.setTerrain(x+1, y-1, c.TerrainTypes["Mountain"])
	c.setTerrain(x+1, y, c.TerrainTypes["Mountain"])
	c.setTerrain(x+1, y+1, c.TerrainTypes["Mountain"])
	c.setTerrain(x, y+1, c.TerrainTypes["Mountain"])
	c.setTerrain(x, y-1, c.TerrainTypes["Mountain"])

}

// draw outputs the terrain to a .png file.
func (c *Chart) draw() {
	defer timeTrack(time.Now(), "draw")
	log.Println("Drawing world.")
	dc := gg.NewContext(800, 800)
	for y := 0; y < c.SizeY; y++ {
		for x := 0; x < c.SizeX; x++ {
			dc.Push()
			switch c.Graph[y][x] {
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
	dc.SavePNG(filename)
}

// initialise returns an instance of Terrain
func (c *Chart) initialise() {
	var err error
	defer timeTrack(time.Now(), "initialise")
	log.Println("Initialising.")
	rand.Seed(time.Now().UTC().UnixNano())

	// Get the input from the command line.
	c.Name = os.Args[1]

	c.SizeX, err = strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	c.SizeY, err = strconv.Atoi(os.Args[3])
	if err != nil {
		panic(err)
	}
	c.NumberOfPeaks, err = strconv.Atoi(os.Args[4]) // How to default to 1? numberOfPeaks = 1
	if err != nil {
		panic(err)
	}

	c.TerrainTypes = map[string]int{
		"DeepWater":    0,
		"ShallowWater": 1,
		"FieldLow":     2,
		"FieldHigh":    3,
		"Mountain":     4,
		"Peak":         5,
	}
}

func formMountain() {

}

// setTerrain sets the Y,X coordinate terrain type.
func (c *Chart) setTerrain(x, y, terrainType int) {
	// Check that the coordinates are within the map boundaries.
	if (y >= 0) && (y < c.SizeY) && (x >= 0) && (x < c.SizeX) {
		c.Graph[y][x] = terrainType
	}
}

// timeTrack taken from stathat.com
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("function %s took %s", name, elapsed)
}
