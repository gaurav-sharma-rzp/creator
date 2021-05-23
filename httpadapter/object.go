package httpadapter

type Object3D struct {
	Mesh Mesh
}

type Vector3 struct {
	X, Y, Z int64
}
type Mesh struct {
	Vertices  []Vector3
	Triangles []Triangle
}

type Triangle struct {
	A, B, C Vector3
}
