package terminalshaders

import "math"

type Vector interface {
	ToSlice() []float64
}

type Vec2 struct {
	X, Y float64
}

type Vec3 struct {
	R, G, B float64
}

func (v Vec2) ToSlice() []float64 {
	return []float64{v.X, v.Y}
}

func (v Vec3) ToSlice() []float64 {
	return []float64{v.R, v.G, v.B}
}

func Dot(v1, v2 Vector) float64 {
	s1 := v1.ToSlice()
	s2 := v2.ToSlice()

	if len(s1) != len(s2) {
		panic("vectors must have the same length")
	}

	result := 0.0
	for i := 0; i < len(s1); i++ {
		result += s1[i] * s2[i]
	}

	return result
}

func Fract(x float64) float64 {
	return x - math.Floor(x)
}

func Mix(a, b, t float64) float64 {
	return a * (1.0 - t) + b * t
}

func Smoothstep(edge0, edge1, x float64) float64 {
	t := Clamp(x-edge0 / (edge1 - edge0), 0.0, 1.0)
	return t * t * (3.0 - 2.0 * t)
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}

	if x > max {
		return max
	}

	return x
}

func Length(v Vector) float64 {
	components := v.ToSlice()
	sum := 0.0

	for _, c := range components {
		sum += c * c
	}

	return math.Sqrt(sum)
}

func Random(v Vec2) float64 {
	return Fract(math.Sin(Dot(v, Vec2{127.1, 311.7})) * 43758.5453123)
}
