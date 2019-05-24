// Code generated by go-bindata.
// sources:
// jsonnet/manifests/ceph-prometheus-rules.yaml
// jsonnet/manifests/gluster-prometheus-rules.yaml
// jsonnet/manifests/noobaa-prometheus-rules.yaml
// DO NOT EDIT!

package manifests

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _jsonnetManifestsCephPrometheusRulesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x58\xdd\x6f\x13\x47\x10\x7f\xcf\x5f\x31\xb2\x40\xc4\xd4\x31\x71\xaa\x54\x60\x11\x24\x94\xd0\x16\x89\x24\x34\xa6\xed\x43\x55\x9d\xd6\xbb\x73\x77\x8b\xf7\x76\x8e\xfd\x70\x30\xd4\xff\x7b\xb5\x7b\xe7\xaf\xe0\x38\xb6\xe3\xf0\xd0\xf2\x84\x6f\x77\x67\x26\xbf\xdf\x7c\xb3\x52\xfe\x81\xc6\x4a\xd2\x5d\x28\x48\x4b\x47\x46\xea\xac\xcd\xc9\x20\xd9\x36\xa7\xe2\xd9\xb0\xb3\x37\x90\x5a\x74\xe1\xbd\xa1\x02\x5d\x8e\xde\x5e\x79\x85\x7b\x05\x3a\x26\x98\x63\xdd\x3d\x00\xc5\xfa\xa8\x6c\xf8\x1f\x40\x39\xbd\xd6\x85\xc1\x73\x1b\xbf\x19\x52\xd8\x05\xa6\xd0\xb8\x03\xe3\x15\x86\xaf\x9a\x15\xd8\x9d\xbb\x7d\xc0\xb1\xcc\x17\x4e\x6d\xc9\x38\x76\x41\x60\xca\xbc\x72\x7b\xb6\x44\x1e\x54\x64\x86\x7c\x19\x95\x1d\xd4\x42\xe2\xcb\x22\x33\x07\xd6\x31\xe7\x6b\x95\x41\x50\x65\xd1\x41\xa5\xb9\x0b\xa7\x58\xe6\xe7\x99\x79\x6b\x5f\xf7\x2d\x6a\x17\x0f\x01\x98\xd6\xe4\x98\x93\xa4\xeb\xfb\xe1\x9f\x40\xcb\x8d\x2c\x5d\x04\x26\xbc\x83\x73\xa6\x59\x86\x06\x72\x66\x41\x48\xcb\xca\x12\x99\x41\x01\xa9\xa1\x62\x0e\x1b\x70\xcc\x64\xe8\xc2\x15\x4e\x43\x34\xa3\xf6\x54\x66\x81\xd6\xb2\x0c\xbb\xd0\x73\x64\x58\x86\x50\xa0\x33\x92\x5b\xe0\xa4\x14\x72\x47\x06\x2c\x9a\xa1\xe4\x08\x9a\x1c\xb0\x21\x93\x8a\xf5\x15\x02\xd3\xa3\x82\x0c\xce\x24\x59\x1c\xa2\x91\x6e\x94\x28\x1c\xa2\xea\xc2\x35\x33\x5a\xea\x6c\x76\x5e\x29\x48\xdc\xa8\xac\xe1\xa9\x8f\xf0\x73\x69\xba\xf0\xcf\xf4\x22\x8b\x40\xec\xfb\xf2\xeb\x47\xea\x9f\x34\x0c\xd1\xe0\x60\x82\x66\x63\x0c\x27\x27\xd0\x69\xd6\x97\x53\x32\x5d\x38\x2e\xea\x5f\xf3\x8c\xcf\x5b\xb4\x68\xcb\xb7\xc8\x9f\x4b\x6b\xa5\xce\xae\xb0\x54\x92\x33\xbb\x1d\x05\xd2\x42\x51\xc9\x01\x53\x0b\xda\x0a\x65\x41\x68\xf5\x13\x07\x39\x1b\x22\x18\xfc\xe4\x65\x20\x54\x13\x50\xba\x44\xf0\xce\x40\xb7\xbe\xb8\x0d\xf1\x26\xbc\x84\xce\xd6\x80\x2f\xc6\x83\xb0\x6b\xc5\x83\xd8\x9a\x93\x73\xa9\x65\xe1\x8b\x19\x72\x13\xcc\x82\xe9\x13\x38\x60\x92\x27\x96\xfb\xf6\x0c\x5e\x80\x73\x99\xe5\x0e\x58\x9a\x22\x77\xe0\x72\x84\x6b\x32\x83\x40\x32\xa5\x53\x69\x5c\x79\xeb\xd0\x2c\xa1\xfb\xad\xb6\x3e\x4d\x25\x97\xa8\xdd\x7a\x96\x3c\x10\xb5\xe1\x38\x29\x84\x4d\x26\xfa\x56\xc4\x16\xbc\x84\xa3\x7b\xd3\xfd\xc9\x93\xf1\xc5\x41\x24\xb5\x3d\x49\xa0\x2b\x18\x27\xfd\x5b\x7c\xf1\xda\x5d\x49\x3b\x58\x9b\xee\xde\x22\x05\xb5\xda\x10\x8c\x8a\xae\xdb\x70\x4a\xda\x31\xee\xa0\xe7\xcb\x92\x8c\x5b\x11\x90\xf5\x43\xe6\xc0\xcc\xf4\x7f\x4b\x02\x1a\x43\x66\x73\x0a\x38\x79\xed\x6a\x12\x48\x27\x95\xb2\xa4\x8a\x84\x95\x4c\x9c\xc0\xfe\xfe\x8d\xc7\x2b\x19\x6c\xc2\x63\x38\x6a\xc2\x0f\x37\x52\x64\x67\x0d\x0e\xb9\x91\x4e\x72\xa6\x96\x92\xf3\xab\xcc\xf2\x0b\x5f\xf4\xd1\x5c\xa6\xef\x90\x09\x34\xa7\x39\xd3\x19\xae\x1f\x98\x4f\xaa\x6c\x59\x95\x73\x68\x7c\xfd\x0a\x8f\x2a\x5b\xda\x1f\xa9\x0f\xe3\x71\xa3\x0b\x52\x5b\xc7\x34\x47\x98\x3b\x9c\x7c\x9b\x0b\xc9\xf1\x38\xd6\x3b\x8b\xa8\xe3\xcd\x21\x53\x1e\xc3\x57\x15\x0d\x03\x5e\x59\x06\x25\x1a\x28\xa4\xf6\x2e\xe4\x51\x8e\xda\xa9\x51\xfb\xc9\xed\x1e\x70\x5a\xbb\xd0\x54\x76\xc1\xf4\xe8\xa6\xcc\xa9\xa0\xdd\x87\xa9\x61\x0e\x67\x2c\x6b\x5f\x24\x18\x4a\x43\x80\x74\x29\xd5\x7f\x1d\x17\x7f\x37\xe1\x29\xfc\x74\x08\xaf\xe0\xb0\xfd\xe2\x78\x37\x29\x5a\x93\xc0\xb5\xc3\xf6\x82\x04\x9e\xd1\xb5\xde\x38\x5e\x83\x96\x79\x9a\xe3\xef\xf1\x18\xae\x43\xa2\x14\x74\xad\xdb\xf0\x5e\x21\xb3\x08\x3c\x47\x3e\x88\x99\x37\xdc\x99\x73\x03\x59\x14\x28\x24\x73\xa8\x56\xb5\x33\x77\x2a\xda\x79\xa8\x57\x6e\xd4\x8d\x4c\x06\x6d\x49\xd0\xd2\xfd\x48\x52\x27\x03\xdf\xc7\x10\xd9\x87\xf3\x54\xfd\x78\x68\x37\x89\xcd\x09\x59\x64\xc5\xda\x34\x5d\xf6\xce\xce\xa4\x1d\x5c\x90\xbb\x42\x5b\x92\x16\x33\x07\xbd\x9b\xb2\xf0\x12\x04\xc6\x32\x39\x07\x64\xfd\x65\x3c\x8e\xb5\xd3\x4c\xe5\xb6\x80\x34\xe4\x64\xdd\xfc\xe5\xf0\x7b\x21\x82\x97\x30\x16\xf5\x2c\xca\xda\x35\x37\xd1\x9a\x24\xd4\x61\xc6\x71\xbf\x0a\x36\xb2\x22\x91\x3a\xe6\x5b\x60\x5a\xc0\xf4\xa3\x2f\x23\x55\xcd\x56\x43\x48\x3b\x68\xb4\x1a\x8f\x3a\x8d\x56\x23\x1e\x0b\x86\x05\xe9\x46\xab\x41\x56\xb4\xf7\xdb\x4f\x9b\x8d\x90\x72\x49\xef\xcf\x9d\x36\xab\x99\x20\x51\x98\xba\xfd\xf0\xf7\xb7\x6a\x10\x9b\x37\xcc\xa8\xde\x48\x3b\x48\x88\x73\x5f\x46\x2e\x5a\x8d\xf0\x62\xa2\x13\x3f\x87\xe2\x85\xc1\xce\x2a\x19\x36\x5a\x8d\x4a\xeb\x42\x8e\xbf\x4f\x8a\xaf\x3d\xe4\x77\x3d\x6d\x82\x76\xeb\x1f\x8c\x73\xb4\x56\x86\xc1\xe1\xbe\xee\x31\x13\xf5\x7d\xdd\xe3\xf0\x7f\xec\x1e\x67\xcc\xb1\x2b\xac\x66\xc7\x0f\x2c\xf4\xbf\x1f\x88\xde\xd1\x26\x59\x24\xf4\xb8\xa6\x16\x11\x6b\x6c\x3f\xd4\x58\xc6\x9d\x1c\x62\xec\x88\xc3\x34\x09\x2e\x67\x1a\x8e\xf2\x75\x3a\xb7\x45\x89\xd2\x82\x55\x74\xbd\xfb\xb2\x1c\xf1\x2f\xb3\xc4\x6b\x81\xc6\xca\x2f\x28\x42\xb9\x9d\x47\xf6\x28\xdf\x7e\xfe\x7c\xff\xcb\x15\x96\x4c\x9a\xed\x40\xed\xa1\x4a\x21\x47\xa6\x80\x4a\x34\xd5\x55\x70\x51\x14\x38\x22\x50\xa4\xb3\xb5\x9a\xe0\xa9\x9c\xd2\x50\x5f\x61\x61\x41\xa0\x43\xee\x50\x3c\x1c\xa0\x52\x73\xd2\x56\x5a\x17\x2a\xf1\x0d\x48\x3b\x9b\x40\x3a\x6d\x5f\xaa\xca\x1b\x07\xcc\xf5\x5b\x98\xba\xed\x7b\x13\x52\x46\x2f\xbc\xdc\x7a\xf6\x90\x16\xa4\xae\x72\x0f\x44\x1b\x6e\xb8\x75\xe7\xb0\x58\xd1\xa8\xdc\x2a\x65\xe7\x0d\x4a\xc0\x3f\x90\xed\xf2\x95\x13\xc8\xab\xc5\xa1\xbf\x73\x78\x9f\x04\x52\xa3\xfc\x67\x45\xda\x2e\x70\xae\xf9\xbf\x37\xd2\x02\x33\xc3\x04\x8a\x3b\xc0\xbe\x9f\xbb\xaf\x01\x77\x68\x40\x36\xc5\xfb\xd6\xb4\x72\xd9\x3b\xab\x37\xa8\xe7\xd2\x16\xcc\xf1\x7c\x6d\xb4\x3f\xe4\x68\x10\x98\xc1\xc5\xd9\x4a\xc8\x34\x45\x13\x42\x75\x58\x09\xb6\x40\x69\xb5\x01\xbb\xec\x9d\x01\xa7\xa2\x24\x8d\xda\xd9\xb9\x3a\x6e\xbc\x0e\xe6\x2d\xa1\x62\xa6\xa4\xf0\xca\xc9\x52\xe1\x82\xd8\xc9\x72\xa4\xde\x89\xd8\x6f\x25\xed\x8e\x9e\x6a\xb8\x9e\x8d\xd8\xa1\xb6\xdf\x31\x62\xf7\x47\x50\xdd\xad\x6d\x6e\x36\xb7\x09\x96\xdb\x77\x92\xa4\xbf\x1f\x79\xe7\xa4\xff\x43\xe4\xad\xb1\x1f\xd9\x35\x79\x37\xab\x8f\x77\x52\xc9\x2f\x91\xa5\x4d\x6b\xd0\x05\x32\xf3\xb3\x57\x6a\xeb\xcc\x38\xa7\x3b\xb6\x58\xdc\x90\xb5\x28\xe0\xf9\xf1\xe3\xf5\xf2\xa1\x46\x66\x42\x52\x4d\xbd\x52\x6d\x78\xf3\xb9\x64\x3a\x60\x14\x8e\x26\x0b\xd5\x87\x5c\x50\x86\xd8\x0b\x39\x32\xe9\x8f\x1c\xda\xc4\x5b\x14\x4d\x78\x76\xdb\x79\x33\xee\x3d\x9e\x6f\xbf\xf7\x58\xce\xc2\x69\x5d\xc4\xd4\xe8\x41\xb8\x78\xb1\x2e\x17\x7c\x6a\x47\xa4\x23\x4e\x1f\x1a\x51\xd8\xd9\xe6\x23\x20\x59\x31\xb4\xeb\x26\xe1\x5e\x94\x6c\xbc\x8a\x9a\xf6\x0d\xff\x06\x00\x00\xff\xff\x53\xe0\x8d\xe8\xfa\x1b\x00\x00")

func jsonnetManifestsCephPrometheusRulesYamlBytes() ([]byte, error) {
	return bindataRead(
		_jsonnetManifestsCephPrometheusRulesYaml,
		"jsonnet/manifests/ceph-prometheus-rules.yaml",
	)
}

func jsonnetManifestsCephPrometheusRulesYaml() (*asset, error) {
	bytes, err := jsonnetManifestsCephPrometheusRulesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jsonnet/manifests/ceph-prometheus-rules.yaml", size: 7162, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsonnetManifestsGlusterPrometheusRulesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x97\x4d\x6f\xda\x4c\x10\xc7\xef\x7c\x8a\x51\xf4\x3c\x52\x5b\x09\x02\x95\x2a\x91\x95\xe8\xa1\x4a\xd5\x53\xa5\xaa\x2f\xb9\xa2\xc1\x9e\xc0\x36\xeb\xdd\xd5\xee\x98\x84\x52\x7f\xf7\x6a\x6d\x2f\x35\x21\x4e\x30\x49\x1a\x7c\x82\xd5\xfc\x67\x66\x7f\xf3\x02\x46\x2b\x2f\xc8\x79\x69\xb4\x80\xcc\x68\xc9\xc6\x49\x3d\x1f\x24\xc6\x91\xf1\x83\xc4\x64\xa7\xcb\x51\xef\x4a\xea\x54\xc0\x17\x67\x32\xe2\x05\xe5\xfe\x6b\xae\xa8\x97\x11\x63\x8a\x8c\xa2\x07\xa0\x70\x46\xca\x87\x4f\x00\x76\x63\x26\xe0\x6a\xec\xcb\x33\x67\x14\x09\x40\x45\x8e\xfb\x2e\x57\x14\x4e\x35\x66\x24\x1a\xd6\xfd\xb9\xca\x3d\x93\xdb\x32\xf0\x16\x13\x12\x90\xd2\x25\xe6\x8a\x7b\xde\x52\x12\xa2\xcc\x9d\xc9\x6d\x19\xaf\x5f\xfb\xa1\x1b\x6b\x5c\x50\xe3\xcc\x93\xe6\x2a\x6a\x70\x54\x25\xd5\xaf\x82\x0b\xf8\x54\x05\xf9\x58\x9b\x9f\x9b\x6b\x5d\x1a\x00\xa0\xd6\x86\x91\xa5\xd1\xb5\x26\x3c\x19\x79\x8f\x73\xda\xd1\xc1\x02\x3d\xa4\xd2\xa3\xb5\x84\x8e\x52\xb8\x74\x26\x6b\x10\x02\x46\x37\x27\x0e\x26\x89\x59\x92\x5b\x0d\x6a\x97\x74\x63\x9d\x80\xdf\x9b\x00\x55\xba\xaf\x72\xbb\xfe\x69\x66\x93\x93\x9a\x41\xfa\xb6\x9f\x28\x49\x9a\x4f\x8a\xc9\x64\xf4\xba\xb6\xbe\x34\x4e\xc0\xe8\x5d\x56\x7f\x6d\x42\x0f\x8f\xa7\x25\x39\xc9\x2b\x01\x89\x93\x2c\x13\x54\x0d\x3e\x9e\x91\x73\xdf\x2f\x29\x0c\x22\xe1\x7b\x11\x7d\x70\x32\xb9\xfa\x56\xca\xba\x10\x82\x52\x07\xeb\xf5\x7f\x55\x7e\x83\x85\xf1\x1c\x72\x28\x0a\xf1\xf7\x70\x16\x8c\xa6\x16\x79\x51\x14\x20\x3d\xa4\xe6\x5a\xb7\x21\xaa\x99\x4c\x2b\x4d\x3b\x29\x98\x4c\x60\xb8\x85\xaa\x13\xa9\x5d\x02\x17\x46\xe5\x19\x1d\x80\xa0\x12\x36\x18\x2c\xcb\x83\xfd\xaf\x5a\xd9\x3f\xdf\x5d\x63\x57\xc4\x91\xcb\x59\x2a\xf9\xab\xbc\xd7\x43\x6d\x51\xdd\xed\xc7\x2d\xc1\xe3\xc9\x34\x3c\x42\x66\x1c\x01\x2f\x50\xc3\x78\xf8\x7f\x0b\xaa\xd1\x70\x08\x6f\xe2\x05\x44\x0d\x2c\x41\x8b\x89\xe4\xd5\x34\xf7\x94\x4e\x67\x2b\x26\x3f\x65\xc3\xa8\x84\xcf\xb3\x8d\x34\x3c\xa7\xad\xd2\xd2\xbe\xd2\x06\x15\xbc\x87\xf1\x16\xe9\x3d\xe6\xef\x1a\x9d\x96\x7a\x7e\x04\xfc\xce\x8e\x81\xdf\x59\x57\x7e\xf7\x4e\x65\xb9\x5f\x0e\xe5\x77\xd7\x72\x6a\x5f\x4c\xbb\x51\x4a\xa7\x5d\xbb\xb3\xde\x5c\x77\xc0\x6d\x1b\xef\xbb\x51\xdf\xf6\xd3\xa8\x4f\xeb\x9e\x78\xe2\xee\x3d\x22\xf8\x7b\xb6\xf6\xcb\xc2\x7f\x4c\xeb\xc7\x25\xcd\x0b\xa9\xad\x31\xaa\xcb\x96\xfe\x5e\x6b\xce\x91\xf1\xd0\x72\x45\x1f\xb0\x5e\x43\xac\x50\xcc\x65\x1a\x52\x83\xa2\x80\x10\x60\x67\xf7\x34\x08\x3e\x30\x27\x11\xee\xc6\x6f\xf8\x47\xb9\x47\x8d\x1a\x75\xd9\x96\x36\xb6\xcf\x3d\x65\x19\x0e\xc6\x4f\x38\x14\xc7\x05\xbb\x7d\x2e\x5e\x0c\xf6\xd9\x53\xee\xff\x48\xea\x73\xfd\xfe\xf1\xac\xc4\x63\x90\x47\xfc\x14\xec\xa0\x8b\x2f\x4e\x07\x92\xdf\xc8\x5f\xb2\xd5\x8f\x0b\x7e\x87\x96\xff\xd7\xf0\x0f\x6e\xfd\x3f\x01\x00\x00\xff\xff\x63\x35\x6e\x54\x99\x0f\x00\x00")

func jsonnetManifestsGlusterPrometheusRulesYamlBytes() ([]byte, error) {
	return bindataRead(
		_jsonnetManifestsGlusterPrometheusRulesYaml,
		"jsonnet/manifests/gluster-prometheus-rules.yaml",
	)
}

func jsonnetManifestsGlusterPrometheusRulesYaml() (*asset, error) {
	bytes, err := jsonnetManifestsGlusterPrometheusRulesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jsonnet/manifests/gluster-prometheus-rules.yaml", size: 3993, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsonnetManifestsNoobaaPrometheusRulesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x51\xcd\x8a\x14\x31\x10\xbe\xcf\x53\xd4\xc5\x9b\xdd\xba\x0b\x0b\x4b\x0e\xc2\x2e\x78\x98\x8b\xe8\x0a\x7a\x6c\x6a\x32\x35\xdd\x61\xd2\xa9\x58\xa9\x8c\x2e\xea\xbb\x4b\xd2\x3f\xf6\x22\x28\x8b\x7d\x4a\x27\xdf\x5f\xd5\x87\xd1\x7d\x22\x49\x8e\x83\x81\x91\x83\x53\x16\x17\xfa\xd6\xb2\x10\xa7\xd6\xf2\xf8\xea\x72\xb5\x3b\xbb\x70\x34\xf0\x5e\x78\x24\x1d\x28\xa7\x87\xec\x69\x37\x92\xe2\x11\x15\xcd\x0e\xc0\xe3\x81\x7c\x2a\x27\x80\xb8\xc2\x0c\x9c\x6f\x53\xbd\x13\xf6\x64\x00\x3d\x89\x36\x92\x3d\x95\xdb\x80\x23\x99\x0d\xba\x09\xcc\x07\xc4\x27\xef\x29\xa2\x25\x03\x47\x3a\x61\xf6\xba\x4b\x91\x6c\x31\xe9\x85\x73\xac\x76\xcd\x2c\x73\xc8\xf6\x4c\xda\x7c\xc9\xac\xd8\x54\x9f\x76\xd1\x01\xa8\xa7\x29\x5c\x33\x85\x30\xf0\x8e\xf9\x1e\xf1\xbe\xd2\x3e\x14\xd6\x67\x94\xe0\x42\x5f\x51\x00\x18\x02\x2b\xaa\xe3\x30\x13\xcb\x77\xa4\x64\xc5\x45\xad\xcb\xba\x9b\x25\x66\x6b\x70\x09\x30\x46\x61\xb4\x83\x0b\x3d\x38\x4d\x50\xd3\xbc\x84\x9c\xb0\xa7\xf2\x3e\xb2\x10\xe8\x80\x61\x55\x04\xb8\xbd\x79\xb1\xfe\x8d\x94\x0a\x74\xa3\x3d\xe5\x83\x7d\x82\xbb\x8d\xf6\x5e\x13\xd4\xcc\x2b\x33\xd1\x85\xc4\xe9\x63\xe7\xe9\x42\xde\xc0\xd7\x27\xc3\x00\x24\x65\xc1\x9e\x3a\x7d\x8c\xb4\xcc\x3e\x3f\xd2\xb7\x28\x06\x7e\xac\xd0\xa9\x85\x6e\x9a\xaa\xab\x23\x7c\xff\x09\x6f\xe0\x75\x7b\x7b\x33\x83\xb6\x6d\x6f\xdd\xb7\xbe\x4b\x31\x42\x89\xb3\x58\x6a\x92\xa2\xd2\x33\xaa\x79\x98\x89\x6f\x45\x58\x3e\x16\xf2\xf3\xab\x59\xcc\xcb\xf2\x5d\x00\x2a\x52\x50\x83\xc0\x89\xe5\x77\x1f\x70\x3d\xb6\x7f\x69\x61\x89\x52\x7a\xd8\x07\xa8\x89\x60\x1b\xe9\xcf\x02\xaa\xd5\x7f\xac\x7f\x49\xde\x0d\x84\x5e\x87\xae\x84\xce\xa9\x16\x71\x35\x63\x4f\x2c\x06\xae\xc7\x7f\x76\x62\xc5\xa9\xb3\xe8\x77\xbf\x02\x00\x00\xff\xff\xc9\xf5\x0b\xfb\xeb\x03\x00\x00")

func jsonnetManifestsNoobaaPrometheusRulesYamlBytes() ([]byte, error) {
	return bindataRead(
		_jsonnetManifestsNoobaaPrometheusRulesYaml,
		"jsonnet/manifests/noobaa-prometheus-rules.yaml",
	)
}

func jsonnetManifestsNoobaaPrometheusRulesYaml() (*asset, error) {
	bytes, err := jsonnetManifestsNoobaaPrometheusRulesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jsonnet/manifests/noobaa-prometheus-rules.yaml", size: 1003, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"jsonnet/manifests/ceph-prometheus-rules.yaml": jsonnetManifestsCephPrometheusRulesYaml,
	"jsonnet/manifests/gluster-prometheus-rules.yaml": jsonnetManifestsGlusterPrometheusRulesYaml,
	"jsonnet/manifests/noobaa-prometheus-rules.yaml": jsonnetManifestsNoobaaPrometheusRulesYaml,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"jsonnet": &bintree{nil, map[string]*bintree{
		"manifests": &bintree{nil, map[string]*bintree{
			"ceph-prometheus-rules.yaml": &bintree{jsonnetManifestsCephPrometheusRulesYaml, map[string]*bintree{}},
			"gluster-prometheus-rules.yaml": &bintree{jsonnetManifestsGlusterPrometheusRulesYaml, map[string]*bintree{}},
			"noobaa-prometheus-rules.yaml": &bintree{jsonnetManifestsNoobaaPrometheusRulesYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

