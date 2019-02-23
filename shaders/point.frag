#version 150 core
in vec4 vertex_color;
in vec3 Idiff;
out vec4 fragment;
void main()
{
    fragment = vec4(Idiff, 1.0);
}
