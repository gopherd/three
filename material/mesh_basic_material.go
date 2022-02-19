package material

import "github.com/gopherd/three/driver/renderer/shader"

type MeshBasicMaterialParameters struct {
	Options Options
}

type MeshBasicMaterial struct {
	basicMaterial
	parameters MeshBasicMaterialParameters
	shader     shader.Shader
}

func NewMeshBasicMaterial(paramters MeshBasicMaterialParameters) *MeshBasicMaterial {
	m := &MeshBasicMaterial{
		parameters: paramters,
	}
	return m
}

func (m *MeshBasicMaterial) Parameters() *MeshBasicMaterialParameters {
	return &m.parameters
}

func (m *MeshBasicMaterial) Options() Options {
	return m.parameters.Options
}

func (m *MeshBasicMaterial) Shader() shader.Shader {
	return m.shader
}
