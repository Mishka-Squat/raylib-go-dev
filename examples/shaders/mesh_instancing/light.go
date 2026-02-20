package main

import (
	"fmt"
	"unsafe"

	"github.com/Mishka-Squat/gamemath/vector3"
	"github.com/Mishka-Squat/goex/image/colorex"
	rl "github.com/Mishka-Squat/raylib-go/raylib"
)

type LightType int32

const (
	LightTypeDirectional LightType = iota
	LightTypePoint
)

type Light struct {
	shader    rl.Shader
	lightType LightType
	position  vector3.Float32
	target    vector3.Float32
	color     colorex.RGBA
	enabled   int32
	// shader locations
	enabledLoc int32
	typeLoc    int32
	posLoc     int32
	targetLoc  int32
	colorLoc   int32
}

const maxLightsCount = 4

var lightCount = 0

func NewLight(
	lightType LightType,
	position, target vector3.Float32,
	color colorex.RGBA,
	shader rl.Shader) Light {
	light := Light{
		shader: shader,
	}
	if lightCount < maxLightsCount {
		light.enabled = 1
		light.lightType = lightType
		light.position = position
		light.target = target
		light.color = color
		light.enabledLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].enabled", lightCount))
		light.typeLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].type", lightCount))
		light.posLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].position", lightCount))
		light.targetLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].target", lightCount))
		light.colorLoc = rl.GetShaderLocation(shader, fmt.Sprintf("lights[%d].color", lightCount))
		light.UpdateValues()
		lightCount++
	}
	return light
}

func (lt *Light) UpdateValues() {
	// Send to shader light enabled state and type
	rl.SetShaderValueA(lt.shader, lt.enabledLoc, unsafe.Slice((*float32)(unsafe.Pointer(&lt.enabled)), 4), rl.ShaderUniformInt)
	rl.SetShaderValueA(lt.shader, lt.typeLoc, unsafe.Slice((*float32)(unsafe.Pointer(&lt.lightType)), 4), rl.ShaderUniformInt)

	// Send to shader light position values
	rl.SetShaderValueA(lt.shader, lt.posLoc, []float32{lt.position.X, lt.position.Y, lt.position.Z}, rl.ShaderUniformVec3)

	// Send to shader light target target values
	rl.SetShaderValueA(lt.shader, lt.targetLoc, []float32{lt.target.X, lt.target.Y, lt.target.Z}, rl.ShaderUniformVec3)

	// Send to shader light color values
	rl.SetShaderValueA(lt.shader, lt.colorLoc,
		[]float32{float32(lt.color.R) / 255, float32(lt.color.G) / 255, float32(lt.color.B) / 255, float32(lt.color.A) / 255},
		rl.ShaderUniformVec4)
}
