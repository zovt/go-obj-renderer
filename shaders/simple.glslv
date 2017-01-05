#version 330 core
layout (location = 0) in vec4 vert;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

void main() {
	gl_Position = projection * camera * model * vert;
}
