package terminalshaders

type Shader interface {
	Compute(uv Vec2, time float64) Vec3
}

var shaderRegistry = make(map[string]Shader)

func RegisterShader(name string, shader Shader) {
	shaderRegistry[name] = shader
}

func GetShader(name string) Shader {
	return shaderRegistry[name]
}
