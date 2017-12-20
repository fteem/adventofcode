package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
	Z int
}

func (c Coordinate) equal(o Coordinate) bool {
	return c.X == o.X && c.Y == o.Y && c.Z == o.Z
}

type Particle struct {
	p     Coordinate
	v     Coordinate
	a     Coordinate
	alive bool
}

func (p *Particle) tick() {
	// Increase the X velocity by the X acceleration
	p.v.X = p.v.X + p.a.X
	// Increase the Y velocity by the Y acceleration
	p.v.Y = p.v.Y + p.a.Y
	// Increase the Z velocity by the Z acceleration
	p.v.Z = p.v.Z + p.a.Z

	// Increase the X position by the X velocity
	p.p.X = p.p.X + p.v.X
	// Increase the Y position by the Y velocity
	p.p.Y = p.p.Y + p.v.Y
	// Increase the Z position by the Z velocity
	p.p.Z = p.p.Z + p.v.Z
}

func (particle *Particle) distance() int {
	return int(math.Abs(float64(particle.p.X)) + math.Abs(float64(particle.p.Y)) + math.Abs(float64(particle.p.Z)))
}

func closestToCenter(particles []*Particle) int {
	closest := 0
	for i, p := range particles {
		if particles[closest].distance() > p.distance() {
			closest = i
		}
	}

	return closest
}

func part1(particles []*Particle) int {
	closest := 0
	for i := 0; i < 1000; i++ {
		for _, p := range particles {
			p.tick()
		}
		closest = closestToCenter(particles)
	}
	return closest
}

func part2(particles []*Particle) int {
	prevLen := len(particles)
	for i := 0; ; i++ {
		for _, p := range particles {
			p.tick()
		}
		sort.Slice(particles, func(i, j int) bool { return particles[i].distance() < particles[j].distance() })
		for i := 1; i < len(particles); i++ {
			if particles[i].p.equal(particles[i-1].p) {
				particles[i].alive = false
				particles[i-1].alive = false
			}
		}
		for i := 0; i < len(particles); i++ {
			if !particles[i].alive {
				particles = append(particles[:i], particles[i+1:]...)
			}
		}
		if i > 0 && i%10000 == 0 {
			if prevLen == len(particles) {
				break
			}
			prevLen = len(particles)
		}
	}
	return len(particles)
}

func main() {
	checkForInputFile()
	p1 := buildParticles(os.Args[1])
	fmt.Println("Part 1:", part1(p1))
	p2 := buildParticles(os.Args[1])
	fmt.Println("Part 2:", part2(p2))
}

func buildParticles(path string) []*Particle {
	f, err := os.Open(path)
	checkError(err)

	var particles []*Particle
	s := bufio.NewScanner(f)

	for s.Scan() {
		parts := strings.Fields(s.Text())
		pX, pY, pZ := extractCoordinates(parts[0])
		p := Coordinate{
			X: pX,
			Y: pY,
			Z: pZ,
		}
		vX, vY, vZ := extractCoordinates(parts[1])
		v := Coordinate{
			X: vX,
			Y: vY,
			Z: vZ,
		}
		aX, aY, aZ := extractCoordinates(parts[2])
		a := Coordinate{
			X: aX,
			Y: aY,
			Z: aZ,
		}

		particles = append(particles, &Particle{a: a, v: v, p: p, alive: true})
	}

	return particles
}

func extractCoordinates(s string) (int, int, int) {
	rawInput := strings.Trim(strings.Trim(s, " "), ",")
	input := rawInput[3 : len(rawInput)-1]
	coordinates := strings.Split(input, ",")
	x, err1 := strconv.Atoi(coordinates[0])
	checkError(err1)
	y, err2 := strconv.Atoi(coordinates[1])
	checkError(err2)
	z, err3 := strconv.Atoi(coordinates[2])
	checkError(err3)

	return x, y, z
}

// Util functions

func checkForInputFile() {
	if len(os.Args) == 1 {
		checkError(errors.New("Missing input file as first argument. Exiting."))
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
