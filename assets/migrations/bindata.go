// Code generated by go-bindata.
// sources:
// 000001_init.up.sql
// 000002_base_schema.up.sql
// bindata.go
// DO NOT EDIT!

package migrations

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

var __000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x76\xf5\x71\x75\x0e\x51\x30\xb4\xe6\x02\x04\x00\x00\xff\xff\x7f\xa9\x0e\x50\x0a\x00\x00\x00")

func _000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpSql,
		"000001_init.up.sql",
	)
}

func _000001_initUpSql() (*asset, error) {
	bytes, err := _000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.sql", size: 10, mode: os.FileMode(420), modTime: time.Unix(1587154408, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_base_schemaUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xc1\x8e\xdb\x20\x10\xbd\xfb\x2b\xe6\xb8\x2b\xd9\x97\x5e\x7b\x4a\xab\x1c\xaa\xb6\xd2\x2a\xdd\x4b\x54\x55\x16\x31\x13\x87\x86\x80\x05\x38\x5a\xff\x7d\x65\x13\x30\x10\xe3\x6c\x55\xd5\x87\x44\xe2\xc1\x9b\x37\x6f\x60\xa6\xaa\xe0\xf3\x6e\xbb\x79\xdd\xc2\xeb\xfe\x65\x0b\x67\x1c\xae\x84\xf7\x08\x9b\x1f\x45\x55\xc1\xd3\xf8\x33\x7e\x67\x1c\xc0\x7f\x06\xdf\x4c\xe9\x90\x69\x7b\x6d\x86\x0e\x97\x11\x6d\x14\x13\xed\x12\x72\x90\x92\x03\xc0\xf8\x87\x44\x24\x20\x97\xa2\x1d\x41\xd6\x32\x91\x1e\xa4\xb2\x3f\x70\x04\x8a\x0d\xbb\x10\x9e\xb2\x32\x41\xd4\x00\x87\xc1\x20\x19\x91\xe7\x8f\x45\x55\x15\x49\x9a\x5c\xb6\x69\x86\x46\x4f\x7f\x49\xc0\x23\x43\x4e\xb5\xb7\xe5\xe7\xaf\x2c\xa5\xee\x88\xa8\x15\x1e\x53\x5e\x85\x47\xeb\x4e\xe4\x80\x51\xa4\xc1\x9a\x51\xab\xd3\x2f\x4f\x24\x8c\x3a\x19\xd9\x60\x9d\x92\x0d\x6a\x9d\xc6\xd2\xa8\xae\xac\xc1\x5a\x90\x4b\x1a\x8f\xb4\xda\xd7\x2f\xca\xe6\x16\xa3\x70\xf4\x9b\x4f\xdf\xb6\x56\x9e\x2e\x9e\x8a\x48\xab\xfd\x6e\x8a\x23\xb9\x10\x79\xe7\xa1\x13\xd1\xa7\x7b\xa8\x23\x0a\x85\xf1\xe7\x42\x48\x76\xa8\x88\x61\x52\x84\x29\x4c\x75\xe0\x61\x02\x30\x87\x31\x44\x99\xda\xb0\x0b\xde\x71\xd1\xde\x52\x2d\x88\x8b\xdc\x00\xf8\xad\xa5\x38\x58\x84\xcb\x1c\xa2\xf0\x98\x41\x5c\x31\xee\x91\x97\xdd\x97\xef\x9b\xdd\x1e\xbe\x6e\xf7\xf0\xe4\x6c\x2c\x9d\x6d\xe5\x6c\xd2\x73\x71\x57\x82\xb0\x96\xae\x12\x0b\xf5\xbd\x0b\x13\xee\x59\xa0\x8d\x1d\xd6\xf5\xf5\xc3\x12\x37\x04\xec\x93\xc8\x33\x13\x34\x7c\xfc\xab\xc5\xca\xea\x29\x67\xb2\x32\x39\xbe\xe2\xc0\xbc\x91\x09\x8a\x6f\x8f\xf4\x66\x65\xad\xdc\x95\xfc\x15\x5f\xc9\x25\x0e\x54\x06\xfc\x0f\xca\x99\xcf\x63\x16\x7b\xe8\x9b\x33\x9a\xd5\xeb\x9e\x4d\xe0\x7d\xf2\x6d\x84\x07\xb2\xdd\x23\xfa\x47\xeb\xa3\x6c\x00\xc6\x60\xda\x90\x4b\xf7\xf0\xa5\xfe\xef\x9a\x39\x13\x9c\x84\xd0\x8e\xd2\x87\x58\x30\xc6\x90\xf6\x3d\x65\x1c\xb7\xf9\xd9\x19\x2f\xdb\x39\x1b\x3f\xb5\xbf\x2d\x6f\xd4\x80\xc3\x03\x2b\x16\xdc\x24\x95\xb3\x88\xe5\xa4\x7d\x9f\xb2\xd9\x27\x13\x88\x62\x87\x82\xa2\x68\x86\x74\x08\xd9\xf6\x1e\x24\x7c\x5b\x6f\x4e\x8c\x47\x3d\xc4\xad\x13\xce\xeb\x46\xf6\xc2\xa4\xc3\x57\xcb\x5e\x35\xe8\x4f\x2c\x4e\x2a\xaf\x83\x85\xcd\xcc\xe8\x7a\xbe\x72\xc9\x6d\x33\x61\x1b\x4f\x6f\x62\xc0\x96\x6f\xe4\x8e\xbc\x04\xa3\x27\x73\xfe\x04\x00\x00\xff\xff\x33\xec\x22\xee\x45\x09\x00\x00")

func _000002_base_schemaUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_base_schemaUpSql,
		"000002_base_schema.up.sql",
	)
}

func _000002_base_schemaUpSql() (*asset, error) {
	bytes, err := _000002_base_schemaUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_base_schema.up.sql", size: 2373, mode: os.FileMode(420), modTime: time.Unix(1587245256, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\x5b\x6f\xdb\xc8\x15\xc7\x9f\xc9\x4f\x31\x2b\x60\x17\x52\xa1\xca\xbc\x5f\x04\xf8\x65\x93\x14\xc8\x43\xb3\x40\x37\xfb\xd4\x53\x18\x43\x72\x46\x4b\x54\x12\x15\x92\xda\x1e\x3b\xf0\x77\x2f\xfe\x73\x46\xb6\xb2\x76\x12\xa0\x68\x00\xc5\x9c\x11\xe7\x5c\x7f\xe7\x3f\xba\xb9\x51\x6f\x86\xce\xa8\x9d\x39\x9a\x51\xcf\xa6\x53\xcd\xbd\xda\x0d\x7f\x6d\xfa\x63\xa7\x67\xbd\x09\x6f\x6e\xd4\x34\x9c\xc7\xd6\x4c\x5b\x3c\x47\xf8\x17\xdf\xf5\xc7\x7e\xde\x9c\x4f\x9b\xe9\xd3\xfe\x69\x37\xb9\x6b\xf4\x64\xee\xa6\xf6\x77\x73\xd0\x57\x5f\x5e\x4c\xed\x06\xac\xde\xfe\xa2\x3e\xfc\xf2\x51\xbd\x7b\xfb\xfe\xe3\x0f\x61\x78\xd2\xed\xbf\xf5\xce\xa8\x43\xbf\x1b\xf5\xdc\x0f\xc7\x29\x0c\xfb\xc3\x69\x18\x67\xb5\x0c\x83\x45\x73\x3f\x9b\x69\x11\x06\x8b\x76\x38\x9c\x46\x33\x4d\x37\xbb\x87\xfe\x84\x0d\x7b\x98\xf1\xa7\x1f\xe4\xff\x9b\x7e\x38\xcf\xfd\x1e\x8b\xc1\x1d\x38\xe9\xf9\xf7\x1b\xdb\xef\x0d\x1e\xb0\x31\xcd\x63\x7f\xdc\xb9\xef\xe6\xfe\x60\x16\xe1\x2a\x0c\xed\xf9\xd8\x5e\xc2\xfb\x87\xd1\xdd\x12\x0f\xea\x9f\xff\x82\xdb\xb5\x3a\xea\x83\x51\x72\x6c\xa5\x96\x97\x5d\x33\x8e\xc3\xb8\x52\x9f\xc3\x60\xf7\xe0\x56\x6a\x7b\xab\x10\xd5\xe6\x83\xf9\x0f\x8c\x98\x71\xe9\xc2\xc6\xfa\xe7\xb3\xb5\x66\x74\x66\x57\xab\x30\xe8\xad\x3b\xf0\xc3\xad\x3a\xf6\x7b\x98\x08\x46\x33\x9f\xc7\x23\x96\x6b\x65\x0f\xf3\xe6\x1d\xac\xdb\xe5\x02\x86\xd4\x8f\x9f\xb6\xea\xc7\x3f\x16\x12\x89\xf3\xb5\x0a\x83\xc7\x30\x0c\xfe\xd0\xa3\x6a\xce\x56\x89\x1f\x71\x12\x06\x77\x12\xce\xad\xea\x87\xcd\x9b\xe1\x74\xbf\xfc\xa9\x39\xdb\xb5\xda\x3d\xac\xc2\xa0\xdd\xbf\xbb\x44\xba\x79\xb3\x1f\x26\xb3\x5c\x85\xff\xaf\x78\x60\x46\xec\x7f\xc5\x90\x19\x47\x89\xdb\x6f\x36\x67\xbb\xf9\x19\xa1\x2f\x57\x6b\xbc\x11\x3e\x86\xe1\x7c\x7f\x32\x4a\x4f\x93\x99\x51\xf2\x73\x3b\xc3\x8a\xcb\xcf\xf7\x23\x0c\xfa\xa3\x1d\x94\x1a\xa6\xcd\xdf\xfa\xbd\x79\x7f\xb4\xc3\xd3\x39\xdf\xc2\xcb\xfe\x95\x05\xd7\x43\xa5\x7c\x1b\xc3\x60\xea\x1f\xdc\xba\x3f\xce\x45\x16\x06\x07\xa0\xaf\x9e\x8c\xfe\x7d\xe8\x8c\xdb\xfc\xd8\x1f\x8c\x02\x26\x1b\x3c\xc1\x8f\x43\x65\x69\xfb\x3f\xfb\x5a\xa9\x0f\xfa\x60\x96\x2b\xef\x01\x3e\x7d\x96\xb6\xdf\xc0\x7b\xf8\xf8\x8d\xb3\xbf\xf6\x0f\x38\xeb\xa2\xf9\xf2\x28\x02\xfd\xe6\x51\xc4\xba\x5c\x5d\x47\xfe\xa5\x01\xa4\xf6\x3d\x03\x48\x6e\xb9\x7a\x4e\xf4\x85\x05\x9f\xfd\xd7\x8d\xbc\x9f\xde\xf6\xe3\x72\xa5\x9a\x61\xd8\x5f\x9f\xd6\xfb\xe9\x3b\x99\xdf\x4f\x92\xb8\x19\xad\x6e\xcd\xe7\xc7\xab\xd3\x1e\x09\x50\x7e\x77\x77\x25\x38\xbf\x9d\x7e\xfd\xb4\x57\xb7\x1e\x88\xe5\x82\x38\xb6\xc4\x55\x43\x1c\x55\xc4\x51\xf4\xfa\xc7\x5a\xe2\x48\x13\x97\x05\xb1\xcd\x89\xcb\x98\xb8\xcc\x89\x23\x43\x9c\xc7\xc4\x69\x44\xdc\x64\xc4\xa6\x20\x8e\x12\xe2\x28\xfb\xf2\x2c\x3e\xa5\x25\xd6\xb5\x3f\x13\x89\xbd\x67\x1f\x8b\x8b\x94\xbc\x08\xd6\x53\xfe\x9a\x7a\x5c\x66\xe1\x4a\x7d\xc2\x20\x78\x99\xef\x3a\x0c\x82\xc5\x4b\xd5\x5d\xac\xc3\x60\xf5\x04\xe6\x8b\x53\xf0\xf9\x17\x37\x4e\xd7\x3e\xdd\x3c\x3d\x89\xd6\xd7\xa2\xfd\x9e\x2e\x3c\x8d\xb3\x1b\xc8\xed\xed\x9f\x9b\xfb\x19\xd8\x6f\xd5\xab\x41\x2b\x70\xbd\x55\x71\xb4\x56\x00\x74\x7b\xcd\xef\x32\x4b\xa2\x95\xdb\x07\x76\x5b\xc1\xf2\xb7\x63\xcf\xcb\x38\xaf\xca\x38\xcf\xb2\xa8\x5a\xab\x68\xf5\x18\x06\x1a\x6e\x7f\x72\xf9\x7d\x76\x49\x6d\x95\xcf\x0d\x31\x6d\xdd\xff\x8f\x4f\x35\xd6\xeb\x57\x90\xfa\xe2\xb6\xfa\xdf\xc9\x02\x39\x79\x4e\xdc\xc6\xc4\x95\x21\xee\x1a\xe2\x24\x22\x8e\x41\x55\x47\x6c\xb1\x6e\x84\xae\xa6\x92\xe7\xae\x26\xae\x4b\xe2\xdc\x10\x97\x0d\x71\xa6\x89\x75\x43\x1c\xb7\xc4\x5a\x13\x37\x05\x71\x97\x10\x27\x9a\xb8\xeb\x88\xb3\x86\x38\xf7\x7e\xe2\x82\x38\x8d\x89\xe3\x94\xb8\x2a\x89\xab\x82\xb8\x42\x3c\x39\x71\x5a\x11\xe7\xda\x13\xdb\x11\x17\xb9\xbc\x97\xfa\x78\x4c\x4a\x5c\xb4\x62\xa7\xcb\xe5\x7c\x06\xf2\x13\x89\xbf\x6e\x88\xd3\x92\xb8\xb0\xc4\x45\x44\xac\x0b\x89\xc7\x20\x57\x9c\x35\xc4\x4d\x43\x5c\xd6\x12\x57\x9b\x10\x9b\x86\xd8\x1a\xf1\x15\x35\xc4\x45\xe9\xf3\x30\xc4\x55\x46\x6c\x4b\xa9\x21\x6c\xa3\xa6\x59\xfe\x5c\xaf\x2e\x25\xb6\x15\x71\x9a\x12\x97\xe6\xf9\x6c\x1b\xc9\xc4\x45\xa8\x03\xce\xb4\xc4\x06\x35\x8b\x88\x8b\x5a\x6a\x56\x74\x92\x3b\x26\x12\xb5\x8c\x63\xd9\x8b\x33\xc9\xd9\xe0\x39\x21\x2e\x13\x39\x57\x63\xb2\x53\xb1\x0d\x9f\x97\xdc\x93\x4c\xfa\x05\x1b\x3a\x21\x8e\x51\x6f\xc4\x86\x3a\xe3\x2c\x6a\x6d\x88\x75\x46\xdc\x24\xc4\x29\x6a\x1b\x89\x4a\x54\x50\x81\x4e\x6a\x82\xfa\xd6\xc6\xbf\x93\x78\x1b\x5e\x49\xe0\x0b\xf9\x3a\x2e\x6a\xb1\x69\xd0\xbb\xab\x7a\xc4\x39\x71\x86\xef\x34\x71\xde\x0a\x03\xc8\x15\x39\x66\xf8\x78\x15\xc2\x3b\xc8\x21\x49\x89\x33\x7c\xe0\x33\x27\x6e\x4a\x62\x93\x8b\x5d\x6d\x89\x13\xf4\x00\x6b\xf4\xa1\x22\xd6\x95\xf8\x40\x2e\x79\x2a\xec\x55\xf0\x59\x4a\x0f\x61\xc7\x64\xc2\x12\x54\x11\xf9\x3a\x26\xb4\x5f\x37\xd2\xaf\xc4\x4a\x2c\xa8\x5b\xd5\x12\xa7\x5a\x72\x02\xd3\x60\xa6\xee\x7c\xad\xc1\x9f\x96\x75\xeb\x79\x46\x0c\xe0\x42\xb7\x62\x1b\xac\xa2\x37\xe0\x3b\xd6\x52\x0f\xcc\x53\x07\xee\x4b\xf1\xd5\x6a\x99\x1d\x70\xaf\x53\xa9\xbd\xcd\x84\x27\xc4\x01\x3e\x50\xab\xbc\x90\xfa\x27\x46\xfe\xa2\x37\x59\x25\x71\xdb\x4e\xd8\xc5\x1e\xea\xd6\xd4\xd2\x2f\x70\xec\xea\xe5\xe7\xa9\x43\x9e\xa5\xf0\x85\x58\xf1\x2e\x94\x1e\x75\xb2\xbe\x67\x59\x22\x76\x71\x9b\xe0\x1d\xd4\xb6\x40\x2e\x89\xf0\x9e\xd4\xc2\x2b\x6e\x18\xcc\x0d\xbe\xc3\xad\x92\x7a\x0e\x30\xbb\xe0\xae\xae\x25\xcf\x26\x92\xb9\xe9\x8c\xdc\x4c\x88\x01\xcf\xb0\x97\x74\x62\x1f\x3a\xd4\xb5\x5e\x73\x90\x53\x27\xbd\x75\xf3\x5e\x89\xdf\xf6\xa2\x1f\x09\xb1\x45\x1f\x2a\x61\x10\x2c\xc1\x27\x38\x00\x77\xe0\x1c\xb3\x8b\xbf\xc6\x4a\x1d\xb1\x06\x17\x8d\x91\x3a\x61\x66\x61\x03\x3c\x40\x2f\xc0\x12\xfa\x88\xf9\x02\x63\xd8\x43\x3f\xcb\x54\x62\xce\xc1\x72\x22\xb9\xd4\x85\xf4\x18\xb3\x87\xbc\x70\xe3\x42\x07\xf1\xec\xe6\x12\xb9\xd5\xc4\x1a\xf5\xe8\x64\xde\x3a\x7f\x33\xe3\xd3\xa6\x72\x16\x3a\x84\x19\x33\xf0\x8d\xb9\xf5\x3a\xe1\xf4\x2c\x23\x2e\x5b\x62\xdb\x4a\x7f\xdb\xdc\x73\xa2\xa5\xfe\xd0\x91\xcb\x1c\xa6\xb5\xe4\xe5\x34\x0e\x9a\xdc\x4a\x1d\xa2\x5a\x78\x80\xd6\x81\x3b\x30\x55\x14\xa2\x63\xa8\x2f\xea\x5d\xe7\xe2\x03\x73\x93\x83\xc5\x9a\x38\xb9\xfc\x72\xc8\x44\x9f\xf0\x7d\xe4\x39\x45\x4f\x5b\x68\x67\x2a\xbc\x82\x23\xe3\xb5\x0f\xb1\x63\x96\xa0\xbb\x98\x37\x30\x9a\x75\xc2\x08\xb4\xda\x26\x5e\xd3\x60\xd7\xc7\x8f\xbe\x80\x41\xab\x9f\xe7\x54\x7b\x0d\x07\x0b\x98\x2b\xcc\x5c\x77\x35\x47\x97\xde\x81\x0f\xe8\xb2\xeb\xb1\x96\x3a\xa0\x76\x75\x2b\x3e\xba\xc8\xdf\x57\x89\x70\x09\x8e\x52\x6f\x13\xb6\x2c\xf4\xef\x72\x87\xb4\xa2\x2f\xd6\xcf\x39\xe6\x1d\xb3\x89\xbb\xce\xd9\x6d\x44\x17\x9d\x0e\x59\xe1\x16\x75\xed\xbc\x9e\xa0\xa7\xd0\x5e\xdc\x7d\xa8\x0b\x6a\x0b\xfd\x40\x7d\x9b\x54\x18\x07\x57\xd0\x1c\xd4\xa2\xb6\x3e\x7f\xf8\x29\xfd\x5d\x87\x39\xf3\xb1\x82\x35\x70\x5e\x54\xa2\x8f\x65\x26\x77\x0c\xde\x47\xdc\x60\x39\xf1\x8c\xe4\x8d\xc4\x02\x7d\x83\x5e\xa1\x8e\xc8\xab\xcd\xa4\xcf\xb0\xdd\x1a\xaf\x81\x85\x9c\xc7\xbd\xe6\xf4\x2f\x15\xb6\x31\x5b\xd0\xcb\x34\x97\x7e\x66\x5e\x5f\x61\x13\x77\x34\xf4\xb6\x4a\x45\x27\x5a\xd8\x6e\xe5\xbd\x52\x0b\x7b\xb8\x83\xc0\x0e\xfa\x0b\x0e\xa0\x09\xb1\xd7\x2f\x77\xd7\x7a\x36\x31\x3b\x58\x3b\x1d\x36\xc2\x2a\x66\x00\x3d\x47\x1d\x30\x6f\xe8\xe9\x6b\xbf\x50\x61\x17\xb3\x92\x24\x32\x63\xb8\x57\xc0\xf8\xab\xbf\x50\xaf\x7e\xfb\xfc\x37\x00\x00\xff\xff\x30\x98\xda\x9a\x00\x10\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 8192, mode: os.FileMode(420), modTime: time.Unix(1587247127, 0)}
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
	"000001_init.up.sql": _000001_initUpSql,
	"000002_base_schema.up.sql": _000002_base_schemaUpSql,
	"bindata.go": bindataGo,
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
	"000001_init.up.sql": &bintree{_000001_initUpSql, map[string]*bintree{}},
	"000002_base_schema.up.sql": &bintree{_000002_base_schemaUpSql, map[string]*bintree{}},
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
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

