package sh23d

type Float interface {
	~float32 | ~float64
}

type Vector2D[T Float] struct {
	X T
	Y T
}

func NewVector2D[T Float](x T, y T) Vector2D[T] {
	return Vector2D[T]{X: x, Y: y}
}

type Vector3D[T Float] struct {
	X T
	Y T
	Z T
}

func NewVector3D[T Float](x T, y T, z T) Vector3D[T] {
	return Vector3D[T]{X: x, Y: y, Z: z}
}

type Triangle2D[T Float] struct {
	V1 Vector2D[T]
	V2 Vector2D[T]
	V3 Vector2D[T]
}

func NewTriangle2D[T Float](v1 Vector2D[T], v2 Vector2D[T], v3 Vector2D[T]) Triangle2D[T] {
	return Triangle2D[T]{
		V1: v1,
		V2: v2,
		V3: v3,
	}
}

type Triangle3D[T Float] struct {
	Normal Vector3D[T]
	V1     Vector3D[T]
	V2     Vector3D[T]
	V3     Vector3D[T]
}

func NewTriangle3D[T Float](normal Vector3D[T], v1 Vector3D[T], v2 Vector3D[T], v3 Vector3D[T]) Triangle3D[T] {
	return Triangle3D[T]{
		Normal: normal,
		V1:     v1,
		V2:     v2,
		V3:     v3,
	}
}

type Triangulation2D[T Float] struct {
	Tri []Triangle2D[T]
}

func NewTriangulation2D[T Float](length int64) Triangulation2D[T] {
	triangles := make([]Triangle2D[T], length, length)
	return Triangulation2D[T]{Tri: triangles}
}

type Triangulation3D[T Float] struct {
	Tri      []Triangle3D[T]
	Vertices []Vector3D[T]
}

func NewTriangulation3D[T Float](length int64) Triangulation3D[T] {
	triangles := make([]Triangle3D[T], length, length)
	return Triangulation3D[T]{Tri: triangles}
}
