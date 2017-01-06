#version 330

in vec3 Normal;

out vec4 outputColor;

uniform vec3 objectColor;
uniform vec3 lightColor;

void main() {
	float ambientStrength = 0.7f;
	vec3 ambient = ambientStrength * lightColor;
	
	vec3 result = ambient * objectColor;
	outputColor = vec4(result, 1.0f);
}
