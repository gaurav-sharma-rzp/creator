package httpadapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qmuntal/gltf"
	"github.com/qmuntal/gltf/modeler"

	"github.com/private/creator/internal/logger"

	"github.com/private/creator/internal/lsystem"
)

func (s *server) addLSystemRoutes() {
	s.Router.HandleFunc("/lsystem/expand", handleLSystemExpand()).Methods(http.MethodPost)
	s.Router.HandleFunc("/gltf", handlegltf()).Methods(http.MethodGet)
}

func handlegltf() http.HandlerFunc {
	//return personCreate
	return func(w http.ResponseWriter, r *http.Request) {
		//simpleGltf := "{\n  \"scenes\" : [\n    {\n      \"nodes\" : [ 0 ]\n    }\n  ],\n\n  \"nodes\" : [\n    {\n      \"mesh\" : 0\n    }\n  ],\n\n  \"meshes\" : [\n    {\n      \"primitives\" : [ {\n        \"attributes\" : {\n          \"POSITION\" : 1\n        },\n        \"indices\" : 0,\n        \"material\" : 0\n      } ]\n    }\n  ],\n\n  \"buffers\" : [\n    {\n      \"uri\" : \"data:application/octet-stream;base64,AAABAAIAAAAAAAAAAAAAAAAAAAAAAIA/AAAAAAAAAAAAAAAAAACAPwAAAAA=\",\n      \"byteLength\" : 44\n    }\n  ],\n  \"bufferViews\" : [\n    {\n      \"buffer\" : 0,\n      \"byteOffset\" : 0,\n      \"byteLength\" : 6,\n      \"target\" : 34963\n    },\n    {\n      \"buffer\" : 0,\n      \"byteOffset\" : 8,\n      \"byteLength\" : 36,\n      \"target\" : 34962\n    }\n  ],\n  \"accessors\" : [\n    {\n      \"bufferView\" : 0,\n      \"byteOffset\" : 0,\n      \"componentType\" : 5123,\n      \"count\" : 3,\n      \"type\" : \"SCALAR\",\n      \"max\" : [ 2 ],\n      \"min\" : [ 0 ]\n    },\n    {\n      \"bufferView\" : 1,\n      \"byteOffset\" : 0,\n      \"componentType\" : 5126,\n      \"count\" : 3,\n      \"type\" : \"VEC3\",\n      \"max\" : [ 1.0, 1.0, 0.0 ],\n      \"min\" : [ 0.0, 0.0, 0.0 ]\n    }\n  ],\n\n  \"materials\" : [\n    {\n      \"pbrMetallicRoughness\": {\n        \"baseColorFactor\": [ 1.000, 0.766, 0.336, 1.0 ],\n        \"metallicFactor\": 0.5,\n        \"roughnessFactor\": 0.1\n      }\n    }\n  ],\n  \"asset\" : {\n    \"version\" : \"2.0\"\n  }\n}\n"
		//w.Header().Set("Content-Type", "application/json")
		//w.Write([]byte(simpleGltf))
		//logger.Info("response sent")

		doc := gltf.NewDocument()
		doc.Meshes = []*gltf.Mesh{{
			Name: "Pyramid",
			Primitives: []*gltf.Primitive{{
				Indices: gltf.Index(modeler.WriteIndices(doc, []uint16{0, 2, 1, 2, 3, 1, 2, 6, 7, 2, 7, 3, 6, 4, 5, 7, 6, 5, 4, 0, 1, 1, 5, 4, 2, 0, 6, 6, 0, 4, 3, 7, 5, 3, 5, 1})),
				Attributes: map[string]uint32{
					gltf.POSITION: modeler.WritePosition(doc, [][3]float32{{0, 0, 0}, {0, 10, 0}, {0, 0, 10}, {0, 10, 10}, {20, 0, 0}, {20, 10, 0}, {20, 0, 10}, {20, 10, 10}}),
					gltf.COLOR_0:  modeler.WriteColor(doc, [][3]uint8{{255, 0, 0}, {0, 255, 0}, {0, 0, 255}, {0, 255, 255}, {255, 0, 0}, {0, 255, 0}, {0, 0, 255}, {0, 255, 255}}),
					gltf.NORMAL:   modeler.WriteNormal(doc, [][3]float32{{0, 0, 10}, {0, 10, 0}, {0, 0, 10}, {0, 10, 10}, {20, 0, 0}, {20, 10, 0}, {20, 0, 10}, {20, 10, 10}}),
					//gltf.TEXCOORD_0: modeler.WriteTextureCoord(doc),
				},
			}},
		}}
		doc.Nodes = []*gltf.Node{
			{Name: "Root", Children: []uint32{1, 2, 3}, Rotation: [4]float32{0.0, 0.0, 0.0, 1.0}, Scale: [3]float32{1.0, 1.0, 1.0}, Translation: [3]float32{0.0, 0.0, 0.0}},
			{Name: "Mesh", Rotation: [4]float32{0.0, 0.0, 0.0, 1.0}, Scale: [3]float32{1.0, 1.0, 1.0}, Translation: [3]float32{0.0, 0.0, 0.0}},
			{Name: "Cube", Mesh: gltf.Index(0), Rotation: [4]float32{0.0, 0.0, 0.0, 1.0}, Scale: [3]float32{1.0, 1.0, 1.0}, Translation: [3]float32{0.0, 0.0, 0.0}},
			{Name: "Texture Group", Rotation: [4]float32{0.0, 0.0, 0.0, 1.0}, Scale: [3]float32{1.0, 1.0, 1.0}, Translation: [3]float32{0.0, 0.0, 0.0}},
		}
		doc.Textures = []*gltf.Texture{
			{
				Extensions: nil,
				Extras:     nil,
				Name:       "",
				Sampler:    gltf.Index(0),
				Source:     gltf.Index(0),
			},
		}
		doc.Samplers = []*gltf.Sampler{
			{
				Extensions: nil,
				Extras:     nil,
				Name:       "",
				MagFilter:  9729,
				MinFilter:  0,
				WrapS:      0,
				WrapT:      0,
			},
		}
		doc.Materials = []*gltf.Material{
			{
				Name: "Default",
				PBRMetallicRoughness: &gltf.PBRMetallicRoughness{
					BaseColorFactor: &[4]float32{
						0.800000011920929,
						0.800000011920929,
						0.800000011920929,
						1.0,
					},
					BaseColorTexture: nil,
					MetallicFactor:   gltf.Float(0.100000001490116),
					RoughnessFactor:  gltf.Float(0.990000005960464),
				},
			},
		}
		//Samplers: []*gltf.Sampler{{WrapS: gltf.WrapRepeat, WrapT: gltf.WrapRepeat}},
		//Scene:    gltf.Index(0),
		doc.Scenes = []*gltf.Scene{{Name: "Root Scene", Nodes: []uint32{0}}}
		for _, buf := range doc.Buffers {
			buf.EmbeddedResource()
		}

		w.Header().Set("Content-Type", "application/json")
		var buf bytes.Buffer
		enc := gltf.NewEncoder(&buf)
		enc.AsBinary = false

		enc.Encode(doc)
		//http.Post("http://example.com/upload", "model/gltf+json", &buf)
		//json.NewEncoder(w).Encode(buf)
		xx, err := buf.WriteTo(w)
		fmt.Println(xx, err)

		//doc := gltf.NewDocument()
		//attrs, _ := modeler.WriteAttributesInterleaved(doc, modeler.Attributes{
		//	Position: [][3]float32{{0, 0, 0}, {0, 10, 0}, {0, 0, 10}},
		//	Color:    [][3]uint8{{255, 0, 0}, {0, 255, 0}, {0, 0, 255}},
		//})
		//doc.Meshes = []*gltf.Mesh{{
		//	Name: "Pyramid",
		//	Primitives: []*gltf.Primitive{{
		//		Indices:    gltf.Index(modeler.WriteIndices(doc, []uint16{0, 1, 2})),
		//		Attributes: attrs,
		//	}},
		//}}
		//fmt.Printf("%+v\n", doc.Buffers[0])
		//res := base64.StdEncoding.EncodeToString(doc.Buffers[0].Data)
		//fmt.Println(res)
		//uriData := fmt.Sprintf("data:application/octet-stream;base64,%s", res)
		//doc.Buffers[0].URI = uriData
		//w.Header().Set("Content-Type", "application/json")
		//var buf bytes.Buffer
		//enc := gltf.NewEncoder(&buf)
		//enc.AsBinary = false
		//enc.Encode(doc)
		////http.Post("http://example.com/upload", "model/gltf+json", &buf)
		////json.NewEncoder(w).Encode(buf)
		//xx, err := buf.WriteTo(w)
		//fmt.Println(xx, err)
		////gbuf := gltf.Buffer{
		////	Extensions: nil,
		////	Extras:     nil,
		////	Name:       "",
		////	URI:        "",
		////	ByteLength: 44,
		////
		////	Data:       nil,
		////}
		////uriDataRaw := base64.StdEncoding.EncodeToString([]byte{97, 110, 121, 32, 99, 97, 114, 110, 97, 108, 32, 112, 108, 101, 97, 115})
		////res, err := base64.StdEncoding.DecodeString("data:application/octet-stream;base64,AAABAAIAAAAAAAAAAAAAAAAAAAAAAIA/AAAAAAAAAAAAAAAAAACAPwAAAAA=")
		////fmt.Println(res)
		////res, err = base64.RawStdEncoding.DecodeString("AAABAAIAAAAAAAAAAAAAAAAAAAAAAIA/AAAAAAAAAAAAAAAAAACAPwAAAAA=")
		////fmt.Println(res)
		////bits := binary.LittleEndian.Uint64(res)
		////fmt.Println(bits)
		////floatRes := math.Float64frombits(bits)
		////fmt.Println(floatRes)
		////uriData := fmt.Sprintf("data:application/octet-stream;base64,%s", "AAABAAIAAAAAAAAAAAAAAAAAAAAAAIA/AAAAAAAAAAAAAAAAAACAPwAAAAA=")
		////doc := &gltf.Document{
		////	Accessors: []*gltf.Accessor{
		////		{BufferView: gltf.Index(0), ByteOffset: 0, ComponentType: gltf.ComponentUshort, Count: 3, Type: gltf.AccessorScalar, Max: []float32{2}, Min: []float32{0}},
		////		{BufferView: gltf.Index(1), ByteOffset: 0, ComponentType: gltf.ComponentFloat, Count: 3, Max: []float32{1.0, 1.0, 0.0}, Min: []float32{0.0, 0.0, 0.0}, Type: gltf.AccessorVec3},
		////		//{BufferView: gltf.Index(2), ComponentType: gltf.ComponentFloat, Count: 24, Type: gltf.AccessorVec3},
		////		//{BufferView: gltf.Index(3), ComponentType: gltf.ComponentFloat, Count: 24, Type: gltf.AccessorVec4},
		////		//{BufferView: gltf.Index(4), ComponentType: gltf.ComponentFloat, Count: 24, Type: gltf.AccessorVec2},
		////	},
		////	Asset: gltf.Asset{
		////		Version:   "2.0",
		////		Generator: "FBX2glTF",
		////	},
		////	BufferViews: []*gltf.BufferView{
		////		{Buffer: 0, ByteLength: 6, ByteOffset: 0, Target: 34963},
		////		{Buffer: 0, ByteLength: 36, ByteOffset: 8, Target: 34962},
		////	},
		////	Buffers: []*gltf.Buffer{{ByteLength: 44, URI: uriData}},
		////	//Materials: []*gltf.Material{{
		////	//	Name: "Default", AlphaMode: gltf.AlphaOpaque, AlphaCutoff: gltf.Float(0.5),
		////	//	PBRMetallicRoughness: &gltf.PBRMetallicRoughness{BaseColorFactor: &[4]float32{0.8, 0.8, 0.8, 1}, MetallicFactor: gltf.Float(0.1), RoughnessFactor: gltf.Float(0.99)},
		////	//}},
		////	Meshes: []*gltf.Mesh{
		////		{Name: "Cube", Primitives: []*gltf.Primitive{
		////			{
		////				Indices: gltf.Index(0),
		////				//Material:   gltf.Index(0),
		////				//Mode:       gltf.PrimitiveTriangles,
		////				Attributes: map[string]uint32{
		////					gltf.POSITION: 1,
		////					gltf.COLOR_0:  3,
		////					//gltf.NORMAL:     2,
		////					//gltf.TEXCOORD_0: 4,
		////				},
		////			},
		////		},
		////		},
		////	},
		////	Nodes: []*gltf.Node{
		////		{Name: "Cube", Mesh: gltf.Index(0)},
		////	},
		////	//Samplers: []*gltf.Sampler{{WrapS: gltf.WrapRepeat, WrapT: gltf.WrapRepeat}},
		////	//Scene:    gltf.Index(0),
		////	Scenes: []*gltf.Scene{{Name: "Root Scene", Nodes: []uint32{0}}},
		////}
		////
		////w.Header().Set("Content-Type", "application/json")
		////var buf bytes.Buffer
		////enc := gltf.NewEncoder(&buf)
		////enc.AsBinary = false
		////enc.Encode(doc)
		//////http.Post("http://example.com/upload", "model/gltf+json", &buf)
		//////json.NewEncoder(w).Encode(buf)
		////xx, err := buf.WriteTo(w)
		////fmt.Println(xx, err)
		////logger.Info("response sent")

	}
}

func handleLSystemExpand() http.HandlerFunc {
	//return personCreate
	return func(w http.ResponseWriter, r *http.Request) {
		var request ExpandGrammarRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		logger.Info("request", fmt.Sprintf("%v", request))
		result := lsystem.Expand(request.Grammar, request.Steps)
		response := ExpandGrammarResponse{Result: result}
		// todo there should be a better way for decoding and encoding request response and should be abstracted out to helper functions
		// todo error handling logic should be common for all handlers
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		logger.Info("response sent")

	}
}

type ExpandGrammarRequest struct {
	Grammar lsystem.Grammar `json:"grammar"`
	Steps   int             `json:"steps"`
}
type ExpandGrammarResponse struct {
	Result []string `json:"result"`
}
