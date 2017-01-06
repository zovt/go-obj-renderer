#version 330 core
layout (location = 0) in vec4 position;
layout (location = 1) in vec3 normal;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

out vec3 Normal;

void main() {
	gl_Position = projection * camera * model * position;
	Normal = normal;
}
