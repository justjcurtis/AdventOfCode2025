package models

import (
	"AdventOfCode2025/utils"
	"math"
)

type Vec struct {
	X int
	Y int
}

func (v Vec) Add(other Vec) Vec {
	return Vec{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec) Sub(other Vec) Vec {
	return Vec{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec) Equals(other Vec) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vec) ManhattanDistance(other Vec) int {
	return utils.IntAbs(v.X-other.X) + utils.IntAbs(v.Y-other.Y)
}

func (v Vec) Copy() Vec {
	return Vec{
		X: v.X,
		Y: v.Y,
	}
}

func (v Vec) Distance(other Vec) float64 {
	dx := float64(v.X - other.X)
	dy := float64(v.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func (v Vec) FastDistance(o Vec) float64 {
	dx := float64(v.X - o.X)
	dy := float64(v.Y - o.Y)
	return dx*dx + dy*dy
}

type Vec3 struct {
	X int
	Y int
	Z int
}

func (v Vec3) Add(other Vec3) Vec3 {
	return Vec3{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vec3) Sub(other Vec3) Vec3 {
	return Vec3{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vec3) Equals(other Vec3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

func (v Vec3) ManhattanDistance(other Vec3) int {
	return utils.IntAbs(v.X-other.X) + utils.IntAbs(v.Y-other.Y) + utils.IntAbs(v.Z-other.Z)
}

func (v Vec3) Distance(other Vec3) float64 {
	dx := float64(v.X - other.X)
	dy := float64(v.Y - other.Y)
	dz := float64(v.Z - other.Z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func (v Vec3) FastDistance(o Vec3) float64 {
	dx := float64(v.X - o.X)
	dy := float64(v.Y - o.Y)
	dz := float64(v.Z - o.Z)
	return dx*dx + dy*dy + dz*dz
}

func (v Vec3) Copy() Vec3 {
	return Vec3{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
}
