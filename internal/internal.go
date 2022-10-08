package internal

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

var _internal_internal_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func internal_internal_go() ([]byte, error) {
	return bindata_read(
		_internal_internal_go,
		"internal/internal.go",
	)
}

var _internal_internal_gs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xcd\x6e\x1c\xc7\x11\xbe\xef\x53\x54\x16\x88\x30\x84\x88\x65\x94\x43\x02\x84\x98\x83\xc4\x48\x24\x04\xfd\x30\x24\x93\x1c\x08\x1e\x9a\x3b\xb5\xbb\x2d\xce\x74\x0f\xbb\x6b\x96\x5c\x2e\x08\x04\xc9\x29\x27\x01\x09\xfc\x00\x06\x7c\xb0\x2f\xbe\xd8\x17\x5f\xfc\x36\x12\xe4\xb7\x30\xaa\xbb\x67\xa6\x67\x76\x08\x91\xf6\xc5\x3a\x08\xdc\xee\xaf\xbe\xfa\xfb\xaa\x7b\x7a\x67\x07\x0c\x52\x65\x14\x08\x63\xc4\x0a\x72\x54\x73\x5a\x8c\xa4\x22\xfe\x33\x11\x6a\x75\x7a\x06\x62\x6b\x7d\x3b\x1a\xb5\xd0\x85\xb0\x8b\xa9\xce\xd0\xc1\xf8\x07\xe3\xc0\x3a\x54\x69\xa4\xa2\xdc\x59\x36\x76\x33\x6d\x0a\x41\x16\xc4\x74\xaa\x4d\x26\xd5\x1c\x48\x83\x00\xbf\x0c\xb6\xc4\xa9\x9c\x49\x34\x20\x54\x06\x57\x46\x12\x5a\x06\x58\x12\x2a\x13\x26\x03\x5d\x51\x59\xd1\xc4\x53\xcf\x12\x4b\x86\x29\xbc\xf5\x36\xb0\xa3\xc9\x64\xf2\x4b\x7c\xf9\x74\x2c\xd0\x02\xc1\xa0\xad\x72\x62\xbc\xe7\x9f\x8c\x82\x1f\x7b\x2f\xb7\x50\xfb\xad\xac\xf3\xb9\x40\xc8\x70\x26\xaa\x9c\x9a\x9d\x99\x36\x20\xc9\x82\x2e\xd1\x08\x95\xd9\x7b\xa5\x9b\x74\xfc\x08\x6b\xd1\xd0\xf3\xcb\x4a\xe4\xbe\xc0\x4f\x7c\x20\xe2\x8f\x75\x18\xa2\x2c\x91\xb9\xc7\xcb\x31\x93\x72\x1c\xa8\x32\xd0\x33\x10\xa1\xc3\x63\x31\x1e\x79\x54\xdd\x5d\xcf\xb1\xec\x35\xf9\xe5\xf1\xdb\x37\xa1\x16\x75\x29\x78\xa9\xd3\x58\x0f\xba\xac\xd0\xac\xe0\x4a\xd2\x02\x4a\x41\x8b\x11\x23\x78\x63\x1f\xa9\xae\xda\x3b\xab\xd5\x76\x60\x73\x20\x47\xc0\x0c\x07\x27\x27\x87\x90\xcb\x73\xfe\xfb\x08\x6d\xa9\x95\x45\x07\x1f\xbd\xf0\x85\x77\x4e\x59\x68\xac\xb8\x0e\x47\xf3\x83\xe1\x4c\xe8\x29\x1c\xc3\x82\x8a\xbc\x66\x38\x38\x79\xfd\xea\x33\x0c\x0c\xaf\x73\xe2\x0d\x48\x0c\xe6\x82\xe4\x12\xdd\x4f\x0b\x85\x58\x81\x2e\x24\x4f\x85\x70\xaa\xb2\xb9\xb0\x8b\xad\xba\x32\x7f\xe3\x12\x1c\x0a\x5a\x24\xfd\x1c\x9b\x82\x7a\x99\xcd\xa4\xb1\x04\x4b\x91\x57\xe8\x04\xc1\x6b\x4a\x14\x98\xc1\x54\x17\xa5\x56\xa8\x88\x9b\xc5\xcb\xae\xac\x8d\x0c\x5f\x68\x53\xfc\x83\xcd\x92\xa1\xf8\x2f\x70\xf5\x2b\xbd\x1d\xbe\x3d\x3e\xa9\x7d\x1d\x6a\x4b\x0f\xf4\x07\x06\x2f\x2b\xb4\x04\x46\x5c\xc1\xb9\xce\x56\x35\xd7\x91\x5f\x7f\xa6\xb3\xd5\x46\x6d\xa6\xb9\xb0\x16\x0e\x88\xca\x3d\xad\x08\xaf\x69\x3d\x02\x80\xb8\x41\xbb\x6e\xa1\x27\x81\xa0\x56\xb7\x15\xe1\x59\x05\x90\x7a\xf0\x72\x6b\xb7\xd9\x8e\x75\xe4\x09\x7c\x2a\x4e\x35\x1e\x76\xeb\xfe\xef\xe9\xc4\xbb\x81\x0d\x3f\xac\x15\x48\x61\xb9\xe1\xc1\xd9\xc7\x1e\x9c\xaa\x62\x0f\x81\xe1\xb2\x51\xcb\x00\x7d\x09\x69\x24\x27\x57\xab\xd6\x53\x28\x76\x59\x93\xc6\xac\xb3\x7e\xc7\x5c\x93\xfa\xf4\x4b\x48\x23\x2d\xf9\x38\x19\xb8\xe1\x63\x39\xe4\xa3\x1c\x52\xc6\x9d\x7e\xba\x3a\xba\x87\x2f\xfe\xaf\xe3\xcf\x44\xea\xd9\xf4\xc1\x3a\x83\xb4\x23\xb1\xe1\x7a\x31\xb0\x4e\xc7\x1d\x13\xcf\x24\x1f\xff\xba\x22\x1c\x2d\x88\xca\x03\xa1\xb2\xbc\xc9\xa7\x40\x5a\xe8\xac\x77\x50\xcc\x2a\x35\x85\x24\xd2\xea\x16\x2c\x9c\x55\x08\x8b\x0f\x8e\x70\xf5\x8d\xd9\x22\x1d\x3f\x8e\x62\x89\xec\x60\x4a\xd7\x90\xc6\x2b\x49\x00\x79\xbe\x64\x4a\xd7\x5b\xbb\x3e\xcc\xa3\x4a\x01\xc7\x07\x16\xcd\x12\xcd\xc4\x05\x7b\x54\xa9\x3a\x52\x91\x65\xa6\x9e\xc3\xe3\x95\x25\x2c\x40\x5b\x10\xa5\x0c\xd3\x77\x7a\x06\xfb\x48\x6f\x8f\x9f\x9a\xb9\x4d\x18\x18\xec\xf6\x74\x51\x08\x95\xd5\x34\x7c\x24\x34\xe9\xf2\x65\x63\xe6\x0c\xfe\x27\xdf\x4d\x2f\x64\x5b\x98\x99\xcc\xf1\x4d\x8c\x75\x27\xcb\x36\xf0\xcc\x94\x68\x0a\x36\x3a\xc2\x42\x2f\x37\x2c\x22\xdf\xfb\x48\x57\x59\xd2\x9c\xff\xb6\x13\xb6\x3f\x11\x7c\x2a\xeb\x51\x5d\xd8\x7d\x24\x07\x30\x73\x3b\x89\xd4\x71\x7a\x06\xf3\x28\xbb\x7e\xd3\xa3\xcc\x07\x66\x70\x7a\x8f\x12\xf4\x19\xeb\xb2\x79\x30\x43\x62\x62\xbe\x7d\x89\x50\x91\xd4\xea\x2f\xae\x1e\x20\x6d\xf8\x12\x98\xca\x42\xe4\x0e\x75\xf5\xf0\xa2\x36\x61\xb4\x0d\x69\x8d\x02\xda\x21\x7b\xd1\x18\xd7\x0a\x1b\x9d\xfa\x6c\x36\xf2\x09\x0d\x77\xa9\x71\x15\xda\xd8\xec\x0c\x54\x70\x1e\x1a\x39\x50\x77\x5e\x6f\xe7\x8d\xdb\xfc\x57\x41\x18\x29\x60\xaf\x32\x06\x15\x9d\xc8\xa2\x09\x82\x6e\x9a\x0a\xe4\x62\xa5\x2b\x72\x12\xe1\x2a\xfc\x5d\xc9\xeb\x16\x15\xdd\x1c\x4c\xca\x14\xeb\x5e\x58\xf7\x63\x1f\x88\x3b\x36\x64\x8b\x00\x8d\x93\xe7\x80\xaa\x5e\x40\x7d\x26\x17\x30\xdd\xf4\x4a\xb0\xa7\x8b\x52\xe6\xfc\xe9\xd9\xcc\x27\x64\x55\x51\x3e\x3d\x3e\x49\x1a\x51\x66\xf1\xa8\xf0\xee\xf1\xaa\x38\xd7\x79\x1f\x10\x0a\xf0\x5c\x91\x59\x79\xf7\x7c\x5b\x5d\xe0\x6a\xdb\xe9\xc1\x3b\x76\xbb\xa0\xf0\x9a\xa2\xdf\xee\xfb\xed\x22\xdc\x6e\xdb\x35\x26\xca\xe1\x02\x57\xe9\x45\x7b\x82\x3a\xbe\x34\xba\xec\x98\x2f\x55\x6d\x6a\x3e\x92\xd7\xa2\x5c\xb7\x4e\x4e\xcf\x80\xc4\x79\x8e\x90\xc2\xe9\x93\x3f\x9d\xad\x6f\x77\x9b\xda\x59\x79\x83\xe9\x1f\x76\xfd\x6c\x97\x95\xff\xc0\xe5\xc0\x7d\x44\xec\x2d\x0a\x66\x67\x07\x48\x67\x1a\xa6\x46\x5b\xab\x97\x68\x5e\x4a\x84\x8f\xff\xfd\xe6\xc3\xb7\x3f\x34\x18\x39\x4b\x38\xe6\x74\x3c\x8e\x0c\xdb\x6e\xb4\x91\xdf\xb6\x26\xe1\xdd\xc2\xe5\x84\xd4\x3f\x61\xba\x17\x13\x23\x64\xd8\x72\xa8\xdf\xbb\xf7\x90\xcb\x6a\x0b\x5a\x9c\xaf\x1f\x93\xb8\xad\x53\x79\x16\x71\xcc\x20\x41\xf8\x5d\x0a\x4a\xe6\xbd\xd0\xbc\x19\x61\x51\xfa\xbf\x52\xc0\xdd\x0e\x80\xbf\xd6\x92\x76\x3f\x90\x40\x97\xa5\x0e\x74\xea\x65\x7b\x20\xec\xa2\xce\xa6\x31\x9d\x74\xf3\x8a\x63\xeb\x98\x45\xa9\x3e\x7a\xc4\x0d\xe1\xa5\x2e\xcb\xa6\x73\xfe\xd7\x62\xbc\x54\x20\x52\x60\xff\x5f\xbf\x23\x9b\x9d\xd9\x60\x85\x38\x8a\x56\xcb\xc3\x96\x75\x0f\x20\x0d\x62\x77\xc2\x0a\x47\x24\xf6\xca\xc0\x4a\x7c\xfc\x38\x28\xd1\x71\x01\xe6\x16\x7b\x35\xde\xd9\x81\x0f\xef\xff\xfd\xe9\x5f\xff\x81\x77\xd9\xc5\x93\xc9\x9f\xe1\xa7\xff\xff\xf8\xe9\xcb\xaf\x3f\x7c\xf5\xfd\xc7\xf7\xff\xfb\xf8\xdd\x17\x9f\xf7\xef\xdd\x73\xfb\x86\xfd\xb7\xa9\xc4\x1f\x5c\x3c\x0d\x73\x6c\xc6\x23\x2a\xfe\x6f\x56\xef\x4e\xb3\xe8\x9a\x34\xac\xfa\x3b\xb4\x8a\x03\x1a\x65\x7d\x06\x11\xba\x6d\x16\xe5\x1d\x7a\x1d\xd0\x65\x38\x87\x71\x32\xa0\xc5\xae\x62\x38\x13\xec\xe9\xea\xb6\x33\xc1\x51\x88\x9d\x29\x89\xc6\x04\x07\xc6\xe3\xae\x10\x6e\x07\x34\x16\xb0\x4a\xe6\xfd\x20\x82\x18\xb8\x70\x73\xa4\x63\x79\x83\x03\x17\x2e\x0b\x29\xba\x6a\xf8\xae\x79\x25\xd5\x05\x66\xaf\xa4\xa5\x70\x3e\xfb\x85\x37\x3a\xc3\xf6\xba\x88\x22\x6b\xb7\xa3\xfb\xa2\x5d\x4c\xfa\x4f\x2e\xff\xa4\xec\x3c\x82\x5c\xdb\xd3\x36\x89\xf6\x6e\x68\x83\x59\xf7\x9d\xb9\x07\x2a\xdf\xb3\x96\xee\xb8\x22\x44\x96\x6d\x78\xe7\xb6\x2c\xb9\xf0\x0f\x18\x81\xc8\x29\xbf\xde\xa2\xdc\xe2\xb7\x22\x33\xfb\x47\x73\x3a\x24\xe0\xb0\x05\x79\x57\x52\x1c\x7e\x77\xd5\x1f\x25\xeb\x0d\xd4\x24\x54\xe9\xf3\x04\x1b\x47\x05\x0c\x9c\x15\xfc\x81\x10\x0b\x22\x4a\xd2\x92\x30\xcc\xe9\x42\xee\xcd\xa8\xdf\x1b\x1c\xd1\xfa\x0d\xe3\x20\x5e\xbb\xfd\x63\x2b\x10\x7b\xc4\xd0\xe4\x3c\x5c\xb4\x3f\x07\x00\x00\xff\xff\x23\xf0\x97\x8c\x86\x14\x00\x00")

func internal_internal_gs() ([]byte, error) {
	return bindata_read(
		_internal_internal_gs,
		"internal/internal.gs",
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
	"internal/internal.go": internal_internal_go,
	"internal/internal.gs": internal_internal_gs,
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
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"internal": &_bintree_t{nil, map[string]*_bintree_t{
		"internal.go": &_bintree_t{internal_internal_go, map[string]*_bintree_t{}},
		"internal.gs": &_bintree_t{internal_internal_gs, map[string]*_bintree_t{}},
	}},
}}
