# Terminal Shaders

Terminal shaders is a Go pacakage that allow you to write fake shaders directly in Go and then render them in your terminal.

## Requirements

- go 1.22.5

## Installation

Setup a new Go project, then run `go get -u github.com/theokbokki/terminal-shaders`

## Quickstart 

Getting a shader to run in your terminal is really easy, all you have to do is write the shader in a `Compute` function, register your shader and then render it like so:

```go
package main

import shaders "github.com/theokbokki/terminal-shaders"

type MyShader struct{}

func (s MyShader) Compute(uv shaders.Vec2, time float64) shaders.Vec3 {
	return shaders.Vec3{R: uv.X, G: uv.Y, B: 0.0}
}

func main() {
	shaders.SetFramerate(30) // 30 is the default
	shaders.SetAnsiMode(false) // false is the default

	shaders.RegisterShader("MyShader", MyShader{})

	shaders.Render("MyShader")
}
```

## Available options

### Framerate

You can adjust the framerate of your shader by calling the `SetFramerate` function before you render it.
By default, it will try to run the shader at around 30fps.

```go
shaders.SetFramerate(30) // 30 is the default

shaders.Render("MyShader")
```

### Color mode

You can choose between two color modes for your shader:
- **truecolor**: Use the full RGB colors for terminals that support it 
- **ANSI colors**: Use ANSI color codes if your terminal only supports those (like MacOS default terminal)

The default behaviour is to use truecolor.

```go
shaders.SetAnsiMode(true) // Will use ANSI color codes instead of full RGB 

shaders.Render("MyShader")
```

### Shader size
Coming soon

## Available helpers

### Vector Operations

To make it more like GLSL, I've added a `Vector` type. It can either be a `Vec2` or a `Vec3`.

### `Vector[T any]` Interface
- **Methods**:
  - `ToSlice() []float64`: Converts the vector to a slice of float64 for uniform handling.
  - `Add(Vector[T]) T`: Adds another vector to this vector component-wise.
  - `Sub(Vector[T]) T`: Subtracts another vector from this vector component-wise.
  - `Mul(Vector[T]) T`: Multiplies this vector by another vector component-wise.
  - `Div(Vector[T]) T`: Divides this vector by another vector component-wise.
  - `Fill(float64) T`: Fills all components of the vector with the given value.

### `Vec2`
- **Fields**: 
  - `X, Y float64`
- **Methods**: Implements the `Vector[Vec2]` interface.

### `Vec3`
- **Fields**: 
  - `R, G, B float64` (typically used for color)
- **Methods**: Implements the `Vector[Vec3]` interface.

Each of these vectors implement the methods from the parent `Vector` type which allows you to easily combine two vectors of the same type like so:

```go
vec1 := Vec2{}.Fill(1.0) // Same as `Vec2{X: 1.0, Y: 1.0}`
vec2 := Vec2{X: 2.0, Y: 3.0}
vec3 := Vec2{}.Fill(20.0)

vec4 := vec1.Add(vec2).Div(vec3).Sub(vec1).Mul(vec3)
// This is the same as this:
// vec4 := Vec2{
//   X: (((vec1.X + vec2.X) / vec3.X) - vec1.X) * vec3.X
//   Y: (((vec1.Y + vec2.Y) / vec3.Y) - vec1.Y) * vec3.Y
// }
```

### Math functions 

For convenience, I've reimplemented some of GLSL's math functions.
They are obviously not all there but will be added over time and PRs are more than welcome to add more :))

### `Dot[T Vector[T]](v1, v2 T) float64`
- Computes the dot product of two vectors.

### `Fract(x float64) float64`
- Returns the fractional part of `x`.

### `Mix(a, b, t float64) float64`
- Performs linear interpolation between `a` and `b` using `t` as the interpolant.

### `Smoothstep(edge0, edge1, x float64) float64`
- Provides smooth Hermite interpolation between 0 and 1 when `x` is within `edge0` and `edge1`.

### `Clamp(x, min, max float64) float64`
- Ensures `x` is within the range `[min, max]`.

### `Length[T Vector[T]](v T) float64`
- Computes the Euclidean length of a vector.

### `Random(v Vec2) float64`
- Generates a pseudo-random number based on the input 2D vector using a simple noise function based on sine and dot product. 
