package debugcharts

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func static_index_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x91,
		0x31, 0x53, 0xc4, 0x20, 0x14, 0x84, 0xeb, 0xbb, 0x5f, 0xf1, 0xa4, 0x0f,
		0x04, 0xb5, 0x70, 0x22, 0xb9, 0x46, 0xad, 0xb5, 0xb0, 0xb1, 0x44, 0xc0,
		0x83, 0x9b, 0x10, 0x22, 0xbc, 0xa8, 0xf9, 0xf7, 0x12, 0xb0, 0xc8, 0xa8,
		0x63, 0x65, 0xb5, 0x3c, 0x18, 0xbe, 0xdd, 0x05, 0x71, 0x76, 0x7b, 0x7f,
		0xf3, 0xf8, 0xf4, 0x70, 0x07, 0x16, 0xfd, 0x70, 0xd8, 0x8b, 0x2a, 0x3b,
		0x61, 0x8d, 0xd4, 0x59, 0x77, 0x02, 0x1d, 0x0e, 0xe6, 0x20, 0x58, 0xd5,
		0x75, 0xc7, 0x1b, 0x94, 0xa0, 0xac, 0x8c, 0xc9, 0x60, 0x4f, 0x66, 0x7c,
		0x69, 0xae, 0x08, 0xb0, 0x72, 0x94, 0x54, 0x74, 0x13, 0x42, 0x8a, 0xaa,
		0x27, 0x16, 0x71, 0xea, 0x18, 0x53, 0x41, 0x1b, 0x7a, 0x7a, 0x9d, 0x4d,
		0x5c, 0xa8, 0x0a, 0x9e, 0xd5, 0x65, 0xc3, 0x29, 0xe7, 0xb4, 0xa5, 0xde,
		0x8d, 0xf4, 0x94, 0x48, 0xe6, 0xd7, 0xab, 0x7f, 0x52, 0xac, 0x3b, 0xda,
		0xd5, 0x17, 0x53, 0x21, 0x6d, 0xc6, 0x6f, 0x08, 0xc1, 0xbe, 0xe2, 0x8b,
		0xe7, 0xa0, 0x97, 0xc2, 0xd4, 0xee, 0x0d, 0x9c, 0xee, 0x89, 0x0a, 0x23,
		0x4a, 0x37, 0x9a, 0xc8, 0x09, 0x24, 0x5c, 0x06, 0xd3, 0x93, 0x9c, 0xa1,
		0x79, 0x77, 0x1a, 0x6d, 0x07, 0x17, 0xbc, 0x9d, 0x3e, 0xae, 0xc1, 0x9a,
		0x8c, 0xc6, 0x0e, 0x2e, 0xdb, 0x32, 0x7a, 0x19, 0x8f, 0x6e, 0xec, 0xa0,
		0x05, 0x39, 0x63, 0x58, 0x9d, 0x32, 0xed, 0x77, 0xea, 0xf9, 0x7f, 0x50,
		0xb7, 0xfd, 0xbd, 0xfc, 0xf9, 0x42, 0x82, 0xd5, 0x5a, 0xb9, 0x66, 0xf9,
		0xad, 0xfd, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x41, 0xaf, 0x7c,
		0xc6, 0x01, 0x00, 0x00,
	},
		"static/index.html",
	)
}

func static_main_js() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x55,
		0xdf, 0x8f, 0xe2, 0x36, 0x10, 0x7e, 0x86, 0xbf, 0xc2, 0x3a, 0x6d, 0x95,
		0x44, 0x87, 0xb2, 0x80, 0xda, 0x17, 0x52, 0x54, 0x5d, 0x39, 0x69, 0xfb,
		0x52, 0xf5, 0xa4, 0x22, 0xf5, 0x01, 0xf1, 0x60, 0x92, 0x81, 0xb8, 0x6b,
		0x6c, 0xea, 0x1f, 0x5b, 0xe8, 0x8a, 0xff, 0xbd, 0x33, 0x36, 0x49, 0x08,
		0x9c, 0xee, 0x74, 0xba, 0x45, 0xda, 0x05, 0xcf, 0x7c, 0x1e, 0xcf, 0xf7,
		0xcd, 0x78, 0xfc, 0xc2, 0x0d, 0x2b, 0x6b, 0x6e, 0xdc, 0xa4, 0x18, 0xb6,
		0xbf, 0xa7, 0xc5, 0x30, 0x2c, 0x24, 0xb7, 0xee, 0x69, 0xc1, 0xe6, 0x6c,
		0x5c, 0xb4, 0xeb, 0xa5, 0x76, 0x5c, 0x7e, 0x90, 0x52, 0x97, 0xdc, 0x41,
		0x15, 0x7d, 0xc1, 0xc9, 0x0d, 0xf0, 0x85, 0x56, 0xdb, 0x8f, 0xb0, 0xe5,
		0x5e, 0x3a, 0xf4, 0xbc, 0x0e, 0x07, 0x21, 0xdc, 0x8c, 0x7e, 0x0d, 0xdc,
		0xe9, 0x00, 0x33, 0x96, 0xd8, 0x83, 0x14, 0x0a, 0x92, 0x11, 0x5a, 0x0c,
		0xa8, 0x0a, 0xcc, 0x52, 0xa3, 0xb5, 0xd4, 0xca, 0x71, 0xb4, 0x9b, 0x69,
		0xf0, 0xfc, 0xa7, 0xf5, 0x7e, 0x19, 0xf1, 0x47, 0x32, 0x9c, 0xf1, 0x4f,
		0x6a, 0x5e, 0x09, 0xb5, 0x8b, 0xc1, 0x24, 0xdf, 0x80, 0xfc, 0xd3, 0x9d,
		0x24, 0xc4, 0xf5, 0xc0, 0xe9, 0x03, 0x82, 0x7f, 0xfc, 0xe9, 0x87, 0x04,
		0x57, 0xe7, 0xb8, 0xc3, 0x09, 0xd7, 0xf8, 0x1d, 0x1c, 0x31, 0x8d, 0xe4,
		0x77, 0xd8, 0x6b, 0x73, 0x62, 0x42, 0x31, 0x6f, 0x21, 0x89, 0x28, 0xeb,
		0x37, 0x77, 0xc0, 0x4a, 0x97, 0x7e, 0x0f, 0xca, 0xe5, 0x98, 0x96, 0xf6,
		0x65, 0x6d, 0x1d, 0xd2, 0x60, 0xf3, 0xf9, 0x9c, 0x79, 0x4c, 0x79, 0x8b,
		0x89, 0x56, 0xec, 0x17, 0x3a, 0x36, 0x59, 0x48, 0x51, 0x3e, 0x33, 0xae,
		0x2a, 0x56, 0x19, 0xbe, 0xa3, 0xc8, 0xae, 0x06, 0x76, 0x90, 0xda, 0x05,
		0x41, 0x98, 0xd3, 0x8c, 0xc8, 0xa0, 0x23, 0x61, 0xb3, 0xb0, 0xe3, 0x93,
		0x50, 0x65, 0x1d, 0x50, 0x41, 0x9c, 0x6b, 0x44, 0x4c, 0xe8, 0xf4, 0xe1,
		0x28, 0xec, 0x25, 0x9b, 0x2e, 0xb1, 0x86, 0xc2, 0xe6, 0xe4, 0xc0, 0x5e,
		0x91, 0x94, 0xb0, 0x43, 0x19, 0x23, 0x06, 0x14, 0xdf, 0x48, 0xc0, 0xc5,
		0x96, 0x4b, 0x0b, 0xd1, 0x4f, 0xa9, 0xfc, 0x71, 0x70, 0x42, 0xab, 0x4b,
		0x4c, 0x4a, 0xeb, 0x12, 0x72, 0x2b, 0xa4, 0x5c, 0x68, 0xa9, 0xcd, 0x65,
		0x3d, 0xa0, 0xd2, 0x70, 0xf3, 0x64, 0x50, 0x69, 0x64, 0x8f, 0x56, 0x76,
		0x9c, 0xcc, 0xd8, 0x78, 0xc4, 0x4e, 0xf1, 0xeb, 0x38, 0x8d, 0x2b, 0xfc,
		0x9a, 0x50, 0x70, 0xfc, 0x58, 0x54, 0x1e, 0x23, 0xaf, 0xc2, 0x62, 0xb0,
		0x42, 0xef, 0x6f, 0x62, 0x57, 0x07, 0x6a, 0x36, 0xdf, 0x41, 0x73, 0x76,
		0x9a, 0xe5, 0x25, 0x9d, 0x64, 0x57, 0xe3, 0xf5, 0x7a, 0x74, 0x01, 0x4f,
		0x7a, 0xe0, 0x90, 0x49, 0xfa, 0xd5, 0xdd, 0x59, 0x6e, 0xc9, 0xce, 0x4b,
		0xe1, 0x4e, 0xe9, 0x38, 0x23, 0x54, 0x9a, 0x98, 0xdd, 0x86, 0x27, 0xd9,
		0x3a, 0xc4, 0x0d, 0xff, 0x63, 0x76, 0x7b, 0x6e, 0x9e, 0xa1, 0x65, 0x47,
		0xb4, 0x3c, 0xe6, 0x3a, 0xed, 0x00, 0x44, 0xf8, 0x2f, 0x51, 0xb9, 0x1a,
		0x09, 0x05, 0x03, 0x56, 0x1a, 0xf5, 0x6d, 0x76, 0xd4, 0xfa, 0xa5, 0xdb,
		0xde, 0x03, 0x07, 0xcb, 0xb9, 0x0b, 0xe4, 0x6a, 0x03, 0xb6, 0xd6, 0x12,
		0xc5, 0x57, 0x5e, 0xca, 0xb6, 0x3e, 0xd8, 0x5f, 0x60, 0x04, 0x45, 0x5c,
		0x5d, 0x5d, 0x02, 0xaa, 0x41, 0x68, 0x74, 0xc5, 0xf7, 0x70, 0xd7, 0x97,
		0xe4, 0xa8, 0xb8, 0xc3, 0x2a, 0xad, 0x48, 0xaa, 0xf3, 0x7a, 0x78, 0x1e,
		0x0e, 0xb7, 0x5e, 0x95, 0xa4, 0x05, 0xf3, 0x07, 0xf4, 0xc1, 0xd3, 0x22,
		0x25, 0x48, 0x46, 0xb9, 0xd1, 0xed, 0xc3, 0xd3, 0xf1, 0xc6, 0xad, 0xd6,
		0xc5, 0x70, 0xb0, 0xd5, 0x86, 0xa5, 0x64, 0x13, 0xf3, 0x71, 0xc1, 0xc4,
		0xcf, 0x84, 0xcb, 0x9f, 0xca, 0x4f, 0x1c, 0x63, 0xdb, 0x5c, 0x82, 0xda,
		0xb9, 0x1a, 0xed, 0xef, 0xdf, 0x87, 0xcd, 0x78, 0x0b, 0x6d, 0x7e, 0xf0,
		0xb6, 0x4e, 0x57, 0x3d, 0xe0, 0x4a, 0xac, 0xf3, 0xe5, 0x88, 0xdd, 0xd9,
		0x16, 0xeb, 0x6c, 0x48, 0xdc, 0xe2, 0xd0, 0xc8, 0x23, 0x3b, 0x2c, 0x0b,
		0x55, 0xe5, 0x23, 0x82, 0x53, 0x8c, 0x37, 0x62, 0xce, 0x78, 0xc8, 0x8a,
		0xcf, 0xe4, 0x8d, 0x4c, 0xdb, 0xe1, 0xf1, 0x6d, 0x0c, 0x7e, 0xa5, 0xc6,
		0x6f, 0xf7, 0x7e, 0x95, 0x47, 0x1f, 0x7e, 0xcd, 0xe6, 0xde, 0xd3, 0xe3,
		0x34, 0xfd, 0x06, 0x4e, 0x06, 0xfe, 0xf1, 0x60, 0x23, 0x26, 0x64, 0xf1,
		0x90, 0xf3, 0xbf, 0xf9, 0x31, 0xa5, 0x7c, 0xbc, 0x91, 0x33, 0xf6, 0xee,
		0xb1, 0x82, 0x8d, 0xdf, 0x3d, 0xc6, 0x8e, 0x7e, 0xa4, 0xf3, 0xdf, 0x51,
		0x79, 0xad, 0x2f, 0x4b, 0xb0, 0xd8, 0x15, 0x4d, 0xa8, 0x56, 0x0b, 0x6a,
		0xa6, 0x7e, 0x89, 0x8b, 0xce, 0x74, 0xaf, 0x5e, 0x11, 0x36, 0x60, 0x9e,
		0x4b, 0xb1, 0x07, 0xed, 0x5d, 0x7a, 0x95, 0xd2, 0x88, 0x4d, 0xc6, 0xf8,
		0x09, 0x01, 0x42, 0x97, 0x96, 0xbc, 0xac, 0xa1, 0x9b, 0x0e, 0x37, 0x64,
		0x68, 0xc0, 0x2e, 0xb8, 0x94, 0x1b, 0x5e, 0x3e, 0x47, 0x36, 0x5f, 0x08,
		0x1b, 0xf7, 0x3e, 0xa4, 0xcd, 0xa4, 0xcc, 0x72, 0x6c, 0xe9, 0xea, 0x94,
		0xb6, 0x7c, 0xda, 0xba, 0xe2, 0x68, 0xdf, 0x62, 0x61, 0x1f, 0x72, 0x9c,
		0x5f, 0x38, 0xa9, 0x52, 0x92, 0x71, 0xc4, 0x5e, 0xcf, 0xa3, 0xdb, 0x17,
		0x23, 0x70, 0x21, 0x74, 0x1e, 0xd4, 0xca, 0x9b, 0xe7, 0x01, 0x37, 0x77,
		0xef, 0xc3, 0x24, 0x29, 0x7a, 0x20, 0x78, 0xc1, 0xc3, 0xa9, 0x71, 0x5e,
		0x29, 0xfd, 0x59, 0x8f, 0xc4, 0xb9, 0x81, 0x86, 0x31, 0x9a, 0xd3, 0x00,
		0xa5, 0x58, 0xf8, 0xa2, 0x1d, 0x42, 0x3b, 0xb7, 0xa1, 0xc2, 0xc4, 0xbd,
		0x41, 0x29, 0xae, 0xb4, 0x05, 0x74, 0x57, 0x1d, 0xee, 0x48, 0x38, 0x3a,
		0xeb, 0x72, 0x8d, 0xa9, 0x24, 0x0e, 0x05, 0x4a, 0xda, 0x93, 0xba, 0xd6,
		0xa1, 0x9b, 0x7d, 0x7b, 0x5a, 0x73, 0x6b, 0xd0, 0xae, 0xe0, 0xdf, 0xde,
		0xec, 0xa3, 0xaf, 0x94, 0x42, 0x04, 0x15, 0xde, 0x5e, 0xb7, 0x69, 0xf2,
		0x79, 0x31, 0xfa, 0x93, 0xe7, 0x4b, 0x82, 0xc4, 0x87, 0xe7, 0x7b, 0xa4,
		0xb8, 0x3d, 0xab, 0xb9, 0x70, 0x6f, 0x2e, 0xc7, 0x39, 0x1b, 0xfe, 0x1f,
		0x00, 0x00, 0xff, 0xff, 0xf2, 0x87, 0xbc, 0x10, 0xd8, 0x08, 0x00, 0x00,
	},
		"static/main.js",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"static/index.html": static_index_html,
	"static/main.js": static_main_js,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
		"main.js": &_bintree_t{static_main_js, map[string]*_bintree_t{
		}},
	}},
}}