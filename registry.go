package terminalshaders

// The interface that defines what a Shader should be.
type Shader interface {
    // The calculations to make to render the shader.
    Compute(uv Vec2, time float64) Vec3
}

// The registery to keep all shaders in memory to render them later.
var shaderRegistry = make(map[string]Shader)

// Registers a new shader in the registry.
func RegisterShader(name string, shader Shader) {
    shaderRegistry[name] = shader
}

// Retreives a shader from the registry
func GetShader(name string) Shader {
    return shaderRegistry[name]
}
