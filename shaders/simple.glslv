#version 330 core
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec4 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
	fragTexCoord = vertTexCoord;
	gl_Position = projection * camera * model * vert;
}