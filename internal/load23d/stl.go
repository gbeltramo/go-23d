package load23d

import (
	"bufio"
	"encoding/binary"
	"io"
	"math"
	"os"

	"github.com/gbeltramo/go-23d/internal/sh23d"
)

func LoadSTL(path_to_file string) (*sh23d.Triangulation3D[float32], error) {
	f, err := os.Open(path_to_file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// skip STL header
	_, err = f.Seek(80, 0)

	r := bufio.NewReader(f)
	numTriangles, err := stlReadNumTriangles(r)

	triang := sh23d.NewTriangulation3D[float32](int64(numTriangles))
	if err != nil {
		return nil, err
	}

	err = stlReadTriangulation(r, &triang)
	if err != nil {
		return nil, err
	}

	return &triang, nil
}

func stlReadNumTriangles(r *bufio.Reader) (uint32, error) {
	numTrianglesBuf := make([]byte, 4)
	_, err := io.ReadAtLeast(r, numTrianglesBuf, 4)
	if err != nil {
		return 0, err
	}
	return uint32(binary.LittleEndian.Uint32(numTrianglesBuf)), nil
}

func stlReadTriangulation(r *bufio.Reader, triang *sh23d.Triangulation3D[float32]) error {
	numTriangles := len(triang.Tri)
	numBytesTri := 12*4 + 2
	triangBuf := make([]byte, int(numBytesTri*numTriangles))
	_, err := io.ReadAtLeast(r, triangBuf, int(numBytesTri*numTriangles))
	if err != nil {
		return err
	}

	for idx := 0; idx < int(numTriangles); idx++ {
		offset := idx * numBytesTri
		triBuf := triangBuf[offset : offset+numBytesTri]
		normal := stl12LittleBytesToFloat32(triBuf[0:12])
		vertex1 := stl12LittleBytesToFloat32(triBuf[12:24])
		vertex2 := stl12LittleBytesToFloat32(triBuf[24:36])
		vertex3 := stl12LittleBytesToFloat32(triBuf[36:48])
		triang.Tri[idx] = sh23d.NewTriangle3D[float32](normal, vertex1, vertex2, vertex3)
	}

	verticesSet := make(map[sh23d.Vector3D[float32]]bool)
	for idx := 0; idx < int(numTriangles); idx++ {
		verticesSet[triang.Tri[idx].V1] = true
		verticesSet[triang.Tri[idx].V2] = true
		verticesSet[triang.Tri[idx].V3] = true
	}
	vertices := make([]sh23d.Vector3D[float32], len(verticesSet), len(verticesSet))
	idx := 0
	for vert, _ := range verticesSet {
		vertices[idx] = vert
		idx++
	}
	triang.Vertices = vertices

	return nil
}

func stl12LittleBytesToFloat32(buf []byte) sh23d.Vector3D[float32] {
	x := math.Float32frombits(binary.LittleEndian.Uint32(buf[0:4]))
	y := math.Float32frombits(binary.LittleEndian.Uint32(buf[4:8]))
	z := math.Float32frombits(binary.LittleEndian.Uint32(buf[8:12]))

	return sh23d.NewVector3D[float32](x, y, z)
}
