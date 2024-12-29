package terminalshaders

import "math"

// Vector defines the interface for any vector type in this package.
type Vector[T any] interface {
    ToSlice() []float64
    Add(Vector[T]) T
    Sub(Vector[T]) T
    Mul(Vector[T]) T
    Div(Vector[T]) T
    Fill(float64) T
}

// Vec2 represents a 2D vector with X and Y components.
type Vec2 struct {
    X, Y float64
}

// Vec3 represents a 3D vector with R, G, B components (typically for color).
type Vec3 struct {
    R, G, B float64
}

// ToSlice converts Vec2 to a slice of float64 for uniform handling of vectors.
func (v Vec2) ToSlice() []float64 {
    return []float64{v.X, v.Y}
}

// Add returns a new Vec2 which is the sum of v1 and v2, component-wise.
func (v1 Vec2) Add(v2 Vector[Vec2]) Vec2 {
    v, _ := v2.(Vec2)

    return Vec2{
        X: v1.X + v.X,
        Y: v1.Y + v.Y,
    }
}

// Sub returns a new Vec2 which is the subtraction of v1 and v2, component-wise.
func (v1 Vec2) Sub(v2 Vector[Vec2]) Vec2 {
    v, _ := v2.(Vec2)

    return Vec2{
        X: v1.X - v.X,
        Y: v1.Y - v.Y,
    }
}

// Mul returns a new Vec2 which is the multiplication of v1 and v2, component-wise.
func (v1 Vec2) Mul(v2 Vector[Vec2]) Vec2 {
    v, _ := v2.(Vec2)

    return Vec2{
        X: v1.X * v.X,
        Y: v1.Y * v.Y,
    }
}

// Div returns a new Vec2 which is the division of v1 and v2, component-wise.
func (v1 Vec2) Div(v2 Vector[Vec2]) Vec2 {
    v, _ := v2.(Vec2)

    return Vec2{
        X: v1.X / v.X,
        Y: v1.Y / v.Y,
    }
}

// Fill returns a new Vec2 where each component is set to n.
func (v Vec2) Fill(n float64) Vec2 {
    return Vec2{X: n, Y: n}
}

// ToSlice converts Vec3 to a slice of float64 for uniform handling of vectors.
func (v Vec3) ToSlice() []float64 {
    return []float64{v.R, v.G, v.B}
}

// Add returns a new Vec3 which is the sum of v1 and v2, component-wise.
func (v1 Vec3) Add(v2 Vector[Vec3]) Vec3 {
    v, _ := v2.(Vec3)

    return Vec3{
        R: v1.R + v.R,
        G: v1.G + v.G,
        B: v1.B + v.B,
    }
}

// Sub returns a new Vec3 which is the subtraction of v1 and v2, component-wise.
func (v1 Vec3) Sub(v2 Vector[Vec3]) Vec3 {
    v, _ := v2.(Vec3)

    return Vec3{
        R: v1.R - v.R,
        G: v1.G - v.G,
        B: v1.B - v.B,
    }
}

// Mul returns a new Vec3 which is the multiplication of v1 and v2, component-wise.
func (v1 Vec3) Mul(v2 Vector[Vec3]) Vec3 {
    v, _ := v2.(Vec3)

    return Vec3{
        R: v1.R * v.R,
        G: v1.G * v.G,
        B: v1.B * v.B,
    }
}

// Div returns a new Vec3 which is the division of v1 and v2, component-wise.
func (v1 Vec3) Div(v2 Vector[Vec3]) Vec3 {
    v, _ := v2.(Vec3)

    return Vec3{
        R: v1.R / v.R,
        G: v1.G / v.G,
        B: v1.B / v.B,
    }
}

// Fill returns a new Vec3 where each component is set to n.
func (v Vec3) Fill(n float64) Vec3 {
    return Vec3{R: n, G: n, B: n}
}

// Dot computes the dot product of two vectors.
func Dot[T Vector[T]](v1, v2 T) float64 {
    s1 := v1.ToSlice()
    s2 := v2.ToSlice()

    result := 0.0
    for i := 0; i < len(s1); i++ {
        result += s1[i] * s2[i]
    }

    return result
}

// Fract returns the fractional part of x.
func Fract(x float64) float64 {
    return x - math.Floor(x)
}

// Mix performs linear interpolation between a and b using t as the interpolant.
func Mix(a, b, t float64) float64 {
    return a*(1.0-t) + b*t
}

// Smoothstep performs smooth Hermite interpolation between 0 and 1 when x is within edge0 and edge1.
func Smoothstep(edge0, edge1, x float64) float64 {
    t := Clamp((x-edge0)/(edge1-edge0), 0.0, 1.0)

    return t * t * (3.0 - 2.0*t)
}

// Clamp ensures x is within the range [min, max].
func Clamp(x, min, max float64) float64 {
    if x < min {
        return min
    }

    if x > max {
        return max
    }

    return x
}

// Length computes the Euclidean length of a vector.
func Length[T Vector[T]](v T) float64 {
    components := v.ToSlice()
    sum := 0.0

    for _, c := range components {
        sum += c * c
    }

    return math.Sqrt(sum)
}

// Random generates a pseudo-random number based on the input 2D vector.
// This uses a simple noise function based on sine and dot product.
// Note: This function assumes the vector is of type Vec2.
func Random(v Vec2) float64 {
    return Fract(math.Sin(Dot(v, Vec2{127.1, 311.7})) * 43758.5453123)
}
