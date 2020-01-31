// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/commands.tmpl
// templates/main.tmpl

package main


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataTemplatesCommandsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\x5b\x73\xdb\xb8\x15\x7e\x26\x7f\x05\x96\x93\xec\x90\x89\x4c\x65\xb7\x3b\x7d\x50\x57\x9d\x49\x9c\xeb\xcc\x26\x71\x6d\x67\xf7\xc1\xf5\xd4\x10\x79\x24\x61\x4c\x02\x34\x08\xf9\x52\x2d\xff\x7b\xe7\x00\x20\x09\x5e\x24\xcb\x69\x5f\xfa\xe2\x48\x3c\x20\xce\xfd\x3b\x17\x65\x3a\x25\xc7\x22\x05\xb2\x02\x0e\x92\x2a\x48\xc9\xe2\x81\x88\x02\x38\x2d\xd8\x51\x92\xb1\x23\x4b\x10\x32\x26\x6f\xbf\x92\x2f\x5f\xcf\xc9\xbb\xb7\x9f\xce\x63\x7f\x3a\x25\x67\x00\x64\xad\x54\x51\xce\xa6\xd3\x15\x53\xeb\xcd\x22\x4e\x44\x3e\x85\x7b\x51\x26\x34\x83\xe9\xe8\x35\xbe\x5f\xd0\xe4\x9a\xae\x80\xe4\x94\x71\xdf\x67\x79\x21\xa4\x22\xa1\xef\x6d\xb7\x84\x2d\x49\xfc\x49\x3f\x28\xe3\xf7\xb9\x22\x55\x15\x2c\x73\x15\x6c\xb7\x04\x78\x4a\xaa\x6a\x70\xe8\x4c\x49\xc6\x57\x25\x1e\x2c\xcd\xc7\x3d\x87\xcf\x59\x0e\x78\x52\xb1\x1c\x9c\x63\xbe\x17\x1c\x2c\xfd\x34\xc9\x58\xd0\x7d\xa1\xb8\x5e\x4d\x41\x4a\x21\xcb\x1e\x41\x96\xd3\x7f\x83\x14\x99\x58\x4d\x33\xb1\xea\x11\xcb\x62\xf9\xd3\x5f\xa6\x89\x58\x48\x3a\x4a\xb9\x65\x05\x48\x4d\x11\xc5\xf5\x2a\x66\x7c\xba\xfe\x99\x0b\x3e\x5d\x01\x57\x19\xe4\x94\xc7\xb7\x3f\x07\x7e\xe4\xfb\xdb\x2d\x49\x61\xc9\x38\x90\xa0\xa0\x92\xe6\x65\x60\x55\x3f\x22\x92\xf2\x15\x90\xf8\x6b\xa1\x98\xe0\x34\x3b\xd1\x64\x4d\xd5\x64\xb6\x24\x70\x43\xe2\xf3\x87\x02\x48\xb0\x10\x22\x03\xca\xcd\xcb\x9e\x97\xe4\x69\xfc\x3e\xa3\xab\x32\x8c\xe2\x37\x42\x64\x21\xda\x2b\x3e\xfe\xed\xd3\x17\x6a\x6c\x38\x21\x4b\x9a\x95\x30\x21\x9a\xf0\x16\xca\x44\x32\xcd\x07\x89\x91\xe5\x00\x59\x09\x5d\x36\x8c\xab\xbf\xfe\x32\xc6\xe4\x13\x12\x46\xb8\xbc\x7a\x2a\x87\x65\x26\xe8\x0e\x1e\xef\x0d\x69\x8c\x4b\x7c\x08\x9f\xe1\x8d\x26\xfc\x46\x2e\x0c\x82\x47\xee\x6b\x22\xf4\xa8\x0d\x43\xc7\x67\x7f\x50\xa6\x40\x5a\x67\x0d\x9d\x71\x47\x99\x3a\xc2\xeb\xcd\xb9\xdd\x8e\xb1\xf4\xb3\x35\xe6\x98\xe1\xdf\x61\x99\x64\x2c\x3e\x03\x75\xbc\x29\x95\xc8\x0d\x8f\x24\x4f\x23\xdf\xf7\xd8\x92\xb8\x7c\x3f\xd2\xd2\x7e\x24\x5b\xdf\xf3\x4c\xa8\xc5\x6f\x18\x4f\x4f\x9a\xd7\xea\xc3\x91\xef\x55\xbe\x93\x5f\xdb\x2d\x79\xc6\x51\xbc\xd9\x9c\xc4\x56\x4e\xfd\x90\x16\x4c\x3f\xfb\x20\x7a\x4f\x4f\x36\x8b\x8c\x25\x9a\x66\x3e\xb6\x27\xfc\x5b\x2a\x49\xfd\x72\x55\x9d\x6d\x16\x89\xc8\x73\xca\x53\x82\x21\xec\xfb\xcb\x0d\x4f\x5c\x3a\xc8\x5b\x90\x28\xf6\xc5\x65\x4e\x8b\x0b\x83\x12\x97\xe6\x1f\x54\x45\x82\xda\x48\x3e\x46\xdd\x6a\x5f\x59\x8f\x3c\x2b\xf5\x45\x5a\x24\x7b\xa7\x8d\x87\xd1\xf7\x3c\x2f\x48\x5b\xcf\x07\x33\xed\x0d\x7b\x47\x3f\x26\x26\xe6\xfc\x46\x66\xbd\x73\xdf\x4e\x7f\x6b\xe8\xd5\xc4\x48\x53\x07\x4e\xe5\x1b\xc3\x5a\xe9\x56\x52\x6c\x0a\x63\x4b\xfc\x64\x64\x43\x4b\x25\x79\xaa\x9f\xe0\xb5\xfa\x50\x6b\x6c\x32\x27\x3f\x6a\x00\x8a\x8f\x8d\x05\x51\xf0\x6f\x25\x58\x21\xcc\xe9\x26\xac\x7c\xcf\xd3\x51\xd4\xa1\x36\x71\x85\xe4\xdf\x04\x5f\xcd\xc8\x55\x4b\xc5\x07\xa4\xaa\xae\x26\xc3\x78\xb0\x62\x8b\x02\x61\x15\xed\x80\xa2\x7f\xad\xbf\x19\xf1\xa7\x53\xd2\x0d\x87\xaa\xc2\x90\x6e\xe5\xc7\x6f\xb5\x04\xbe\xe7\x3a\x7e\xfc\x85\xb0\x61\x1c\x9f\xc2\xcd\x86\x49\x48\x1b\x4c\xec\xde\x6c\x1c\x39\x21\x8d\xd0\x26\xde\xc9\x0b\x0d\xcb\xf1\xef\xf8\xd7\x16\x97\x63\xca\x3f\xd2\x5b\x78\x23\xd2\x07\x52\x55\x13\xb2\xc0\x0f\x36\x10\xea\xb7\x23\x12\xbe\x68\x81\xfb\x14\xca\x42\x70\x4c\x51\x64\x7a\xaa\xe3\x4f\x63\x17\xbe\xae\x2b\x89\xc9\xb1\x35\xe5\x69\x06\xf2\x84\xaa\x35\x9a\x47\xe7\xf3\x47\xf3\xac\xf6\x8a\xef\x61\x9e\x8e\x66\x83\x8e\x42\xf7\x0a\x73\x83\x49\xc4\xaa\x22\x01\x79\x49\x1c\xb2\xef\x79\x88\x07\x5e\x1b\xe6\x46\xd5\x0f\xa0\x6a\x90\x33\x24\x0d\x61\x6c\x49\xec\xc1\xf9\x9c\x04\x81\x61\x56\x3f\x19\xcb\xbe\x8b\xe6\xb6\x4f\x5c\xd5\x57\x1d\x31\x9e\xc2\x7d\x10\x5d\x5e\xe8\xe0\xbf\xac\x65\xd8\xc8\x0c\x05\x30\x87\x5e\x6a\xbd\xb5\x06\xa8\xb0\xc5\x4f\x1b\x3f\x6c\x42\x9e\x69\xd7\xe8\xf8\x19\x38\x15\xa5\x6a\x2b\x9d\x39\x19\x7f\xe2\x58\x29\xd5\xba\xae\x10\x9a\xdd\xdc\xba\xac\x8c\x4f\xa1\xc8\x68\x02\xe1\x46\x66\x1a\x43\xaf\xb6\x57\x3a\x3a\xec\xdb\xd6\xf2\xdb\xed\x55\x75\xa5\xd1\xb6\x25\x35\xe1\x33\x21\x3f\x45\x35\xeb\x3a\x5f\xbb\xb0\xeb\x49\xb8\x41\x91\x11\x7e\x8f\x33\x06\x5c\xc5\xa8\xe5\x67\x50\x6b\x81\x47\xc2\x08\x93\x1f\x65\x88\xfc\x0e\x0a\x1d\xa4\xf0\x50\xdf\x9b\x0d\xc8\x87\x46\x61\x64\x3e\x27\x12\x6e\xe2\xd7\x69\xfa\x0f\x24\x99\x0a\xd6\x96\x90\x11\xa5\xac\x46\x6e\xc5\x75\x18\xac\x81\xa6\x20\xc7\x39\x7c\xd4\xb4\xa7\xb0\x68\x6d\xe6\x98\xec\x91\x8e\xc6\xeb\x66\xf0\x6c\x4e\x6c\x8d\xfa\x00\x0a\x49\x3a\xc3\xfe\x24\x8a\xa9\x4c\x63\x41\xbf\x66\x6b\xe6\x26\x99\x9c\x6b\x7e\xd0\xd1\xac\x5f\xfe\xc2\x32\x0d\x3a\x5a\x41\xa7\x7f\x1a\xda\xf7\x00\x03\x2f\x73\x15\x9f\x15\x92\x71\xb5\x0c\x83\xe7\xb7\xc6\x1e\x8e\x25\xa2\x86\x8b\xdb\xe1\x8c\x58\xfa\x10\x53\x3f\x81\x59\x6d\x79\x6f\x18\xb2\xa3\x78\x67\x8d\xa6\x41\xef\x87\x16\x0a\xc6\xa5\x3a\x16\x5c\x01\x57\x47\x68\xcd\xba\x3f\xfa\x0c\x29\xa3\x16\xfd\x02\x6c\x6f\xd2\x07\x8b\x37\x78\x67\xd4\x8a\xe2\x48\x82\x59\x63\x80\xf0\x0d\x2c\x85\x84\xd0\xc1\xb1\x89\x75\xfb\x04\x99\x47\x26\xd5\xca\x42\xe3\x2a\xc6\x04\x4a\xf4\x56\x84\x16\xc3\xf0\xe1\x0f\x73\xc2\x59\x66\xc4\xb6\xbd\x00\x67\xd9\xc4\xfc\x31\x7d\x7d\xfc\x87\xa4\x45\x08\x52\x4e\x48\x80\x29\x07\xa5\x22\x4b\xca\x32\x48\x75\xd4\x68\x99\xb0\xd6\xa6\x90\x88\x14\xd2\x21\xac\xfb\x86\x1d\x4a\x12\x9f\x29\xaa\x36\xa5\x9e\xb9\x7e\x25\xbf\xbc\x7a\x65\x38\x5b\x61\x2c\x24\x7c\xe3\x39\x95\xe5\x9a\x66\x75\xa9\x08\x8d\x12\x3f\x5a\x0e\xd1\xdf\x06\xa2\x1f\x22\x7b\x73\x6d\x86\xbd\x8f\xb4\x77\xbb\xaa\x68\x5b\x57\x26\xe6\xf6\x5a\xe4\x1d\xfe\xb3\x0c\x83\x8f\xe7\xe7\x27\xe4\x79\x3a\x23\xcf\xcb\x60\xd2\x57\xb0\x79\xa0\xfd\x19\x35\xb6\xa2\x4b\x05\x8d\xae\xc6\x91\xaf\xf1\xd1\x2e\x3f\xa2\xea\xb5\xe6\xc6\x92\xe6\x06\x57\xff\xda\xf6\x73\x43\x33\xc1\xca\xa1\xe3\x08\x9c\x3e\x40\x2e\x69\x02\xdb\x0a\x13\x28\x0e\x07\x9e\x8a\x5c\xf8\xb1\x48\xad\x2d\xd0\x91\x42\xdb\x62\x67\x53\x73\xa7\x7b\x6e\x0d\xd0\x6e\x1b\xff\xa4\xfe\xa4\x69\x85\xfe\x17\x9d\x4a\x64\xbc\xa6\x0d\x45\x95\x82\xbc\x50\x28\xdd\x2b\xdf\xf3\x96\x42\x92\xfa\xd1\xaf\x5a\x3a\x23\x7d\xfc\xda\x3c\x2c\x1b\xc8\xb3\xa7\x5e\xbe\xf4\x4d\x5c\x74\xcc\x61\x63\x77\x4c\xbb\x56\x93\xff\x46\xcf\x81\x82\x91\x93\x35\x63\x79\x30\x8c\xfe\x63\xb1\xc9\x52\xc2\x85\x22\x09\xcd\x32\x62\xbd\xd4\x74\xa0\x75\xfc\xe3\x5f\x4c\x66\x9a\xa8\x0d\xcd\x88\x13\x32\x35\x25\xa7\x2a\x59\x9b\x69\xc3\x73\x6b\xb3\x7e\x6e\x1d\xff\xd9\x7c\xae\x6b\x93\x67\x6e\x33\x86\x32\x71\xff\x01\x94\x3e\xf4\x3b\xcd\x36\x26\xbf\x63\x8d\x8f\xf7\xca\x22\xe3\x19\x64\x90\x98\xe6\xba\x2e\x65\xaf\xb3\xec\x0c\x94\xc2\x3e\x25\x8c\x3a\x39\x31\x6e\x8b\x43\x8c\xb1\x02\x45\x6a\xc9\x6f\x51\x16\x63\x08\x6b\x09\x4f\x93\x5c\xb9\xb5\xd0\xa6\xc2\x9c\x23\x10\x6a\xf9\x2e\x2e\x17\x0f\x0a\x74\x3a\xbd\xbb\x2f\x20\x51\x90\x92\x3f\x89\x29\x39\x24\x78\x7e\x83\xd9\x16\x4d\xac\x4d\xbf\x47\xde\x3f\xac\x84\xc6\xf6\x88\x58\x1b\xd9\x48\xda\xd4\x48\x43\xb5\x77\x35\x7d\x90\x46\x24\x20\x41\xfd\x56\x53\x3a\x7b\xec\x6a\x4c\x33\x49\x4b\x12\xca\xd1\x3e\x12\x68\xb2\x26\x29\x94\x18\x9c\xa4\xd4\x57\x2d\x20\xa1\x9b\x12\xc8\xf3\x92\xb0\xd2\x40\xdf\xc0\x65\xfb\x6d\xd1\x88\xe8\x2c\x16\x3c\xcf\x5b\x48\xa0\xd7\x2d\xad\xa9\xc6\x5e\xd5\x6d\x8d\xf0\x9b\x62\x39\xc4\x67\x19\x40\x11\x9a\x45\x43\x46\xb1\x22\xbf\x30\xcf\x21\x11\x3c\x6d\x10\x17\x21\xd3\x66\xf9\xdf\xe7\x7b\xd3\xbc\x6b\x92\x2f\x70\x17\x06\x9f\xe9\x3d\xcb\x37\x79\x7d\x43\x49\xe0\x3e\x01\x48\xdd\xea\xd7\x96\x89\x1e\x2a\xf6\xa6\xf1\x53\x58\xb1\x12\x91\xbe\xec\x8e\xed\x7a\xe0\x91\x42\xa8\xba\x2a\x9c\x0a\xa1\xcc\x12\xa2\xec\x8e\x34\xfa\xd0\xd8\xe0\xea\x4c\xae\xdc\x99\x59\xdd\xa1\x35\x3e\xb7\xcd\x9f\xa1\x98\x79\xd5\xc4\xb4\xbc\x4e\xc5\x1d\x0f\x47\x57\x36\x13\xbf\xe9\x41\x86\x63\xd6\x9c\x28\xb9\x01\xdf\x2d\x9a\xb5\xfc\x76\x36\x9d\xf7\x78\xbb\x27\xf4\x84\x3c\x3f\x40\x06\xdf\x33\x0b\xa2\xbd\xf3\xbe\x0d\x79\x3b\x7d\x9f\x50\x09\x5c\xb5\x4b\xab\xde\x1a\xc0\x90\x5b\x84\xc5\x46\xcd\x9a\x33\xdc\xb3\x34\x88\xfc\x7e\xd8\xa2\x47\x9e\xf8\x72\xbb\x9e\x75\x76\x5f\xba\xbb\xec\xec\xbd\x30\x46\x51\xc7\x9d\xce\xc6\x03\x5d\x2f\xe3\xfb\x44\x17\x35\x92\x80\x54\x94\x71\x02\xb7\x68\x06\x21\x9b\xdc\xc5\x96\x91\x98\x88\x65\x7c\xe5\xc6\x42\xf0\x26\x13\xc9\x35\x06\x38\x24\x1b\x6d\x7b\x74\xf1\xa6\x84\x92\x14\xc2\x74\x4d\x4a\x90\x02\x24\x13\x29\xc3\x2a\xf2\x40\x92\x35\x24\xd7\xdf\xc1\xb1\xf2\x87\x96\x43\x75\x7a\xf3\xde\x8e\x5e\xc2\x33\xdd\x84\x5d\xc6\xd5\xeb\xb8\x76\x6a\xc7\xbc\x35\x18\x93\xe4\xe9\x0e\x13\x3a\x19\x13\x7f\x2b\xdb\xb4\x68\x86\x9b\xf8\x75\xc6\x28\xea\xde\xc0\x93\x7d\x30\x23\x17\x9d\x55\x97\xd7\x19\xce\x06\x6f\x79\x9e\xe6\xe1\x30\xe8\x0f\x1b\xf5\x7e\x6b\x84\xe0\xa6\x6f\x67\xdb\xb4\x27\x81\xed\xde\xc9\x64\x2e\x8a\x2d\x57\xe5\x8c\x18\x0b\x7c\x66\x1c\xc1\xec\x0b\x3e\x43\xdc\xcc\x80\xef\xed\x42\xea\x3b\x4e\x37\x7c\x46\xd0\xe8\x18\xdf\xe4\x45\xc7\x9c\x13\x42\xe5\xaa\x6c\x8c\x52\x3b\xc5\xed\xea\x0f\xec\xfb\x9e\xdd\x77\xe6\xfb\x3d\x72\x21\xc7\x0b\xbc\xf5\x9e\x54\xd5\xe5\xb0\x3d\x1a\x99\x10\x3c\xcf\xcb\xc4\x2a\x7e\x4f\x15\xcd\xc2\x08\xcb\x1d\x16\xd7\x28\xfe\x5c\xae\xc2\x40\x17\x3f\xdd\x14\x61\x84\x46\xb5\x57\x7c\xd7\x39\xe6\x1b\x9e\xe9\xe5\x7b\xe4\xdb\x09\x92\x60\x85\xc8\x74\xb5\xad\x7f\x8c\x68\x95\xa8\x27\xcb\x70\x3f\x0c\x1c\xb2\x1a\xec\x86\xff\x78\xf4\xd7\xed\x1a\xdc\xd3\xbc\xc8\xa0\xb4\xad\xb2\xdf\x6d\xda\xe0\x5e\xdf\xff\xae\x3e\x64\xe3\xae\x79\xe9\xe5\x9c\x04\x44\x2f\xcb\x1a\xd0\xb6\x8a\xe3\x94\x12\x46\xe4\x25\x09\xb4\x77\x1b\x79\x6d\x32\xe9\x87\x80\xde\xf9\x27\x0f\x86\xd5\x7b\x4f\x5e\xee\x48\xcb\x5d\x59\xb9\x33\x29\xf7\xe6\xe4\x20\x25\xfb\x89\x57\x4d\x46\xb6\x02\xfb\xd2\xf1\xc0\x6c\xac\xd5\xf8\xc8\xd2\x14\x78\xc3\xce\x7c\x9d\xe9\xb6\xa9\x21\x8d\x8a\x60\x5d\x35\x6b\x1c\x6b\x4e\x3d\x9a\xe4\xbb\x52\xfb\x7b\x32\xbb\x56\x62\xb8\x0c\xf1\xbc\x85\x48\x1f\x26\xee\x34\xff\x01\x14\x1e\x08\x87\xfb\x0e\x73\xfd\xc5\xb8\x84\x55\x35\xbb\xb4\x79\x38\xda\x2b\xef\x49\xe6\x6f\x9c\x2e\x32\xc0\x5a\x85\xdd\x3d\x0a\x54\x67\x74\x35\x40\x59\xb7\x59\x8e\xbf\x00\xa4\x65\xbd\x74\x20\x55\x85\x63\x49\x5b\xed\xff\xd5\x84\xf0\x61\xa3\xdf\xe3\x00\xf7\x54\x58\xdb\xb3\x73\x6f\x97\xed\xdf\x67\x33\x03\x80\x89\x5d\x8e\xf4\xc6\xc2\x66\x1c\xea\x2d\x6a\xde\x0b\x99\x63\x5b\x2c\xed\xa7\x70\xcf\x82\x66\x1f\x73\x7b\x0f\x72\x76\xb7\x31\x2d\xdb\xb1\xf6\xc8\x4c\x30\x37\x27\x0d\xfa\x8d\x4d\x89\x7e\x93\xde\xa3\x2b\x64\x07\xb6\x46\x97\xc9\x1d\x1e\x17\xc1\x60\x0f\x1e\x5c\x92\x79\x13\xc4\xcf\xb0\x37\xbe\x6c\x19\x76\x22\x6c\xcf\x50\x3d\x1a\x8f\xe3\x3f\x85\x5a\x17\xb4\xab\xdd\x47\x7f\x0f\x6d\xab\xb1\x77\xb7\xbb\x50\x74\x79\xea\x28\x2d\xf5\x30\x27\x4c\xf3\x85\x48\xe2\x76\x34\x4f\x9f\xe7\xdb\xfb\x74\xe2\x37\x46\xed\x8d\xf1\xbb\x43\x77\x6f\x00\x3d\x3e\xcf\xbb\xd1\x84\x0a\xc8\x55\x9d\x66\xa8\xe0\x70\x45\xec\xce\xea\x63\x1e\xea\xfe\x0e\x93\xf6\x0d\x36\x88\xa2\xff\x7f\x93\xd9\xf0\x89\xcf\x40\x69\x2c\x7f\xc6\x52\x0b\xe2\x8f\x9a\x6a\x7f\x23\x68\x83\x76\x17\x5c\x8e\x05\xa2\xe3\x3e\x07\x24\xef\xf6\x36\x7f\x7b\x6c\x61\x57\x1e\x7a\xea\x6f\x95\xaf\xba\x5d\x60\xaf\x0e\x9b\xea\xec\x8c\x9d\x2d\x94\xe8\xf9\xcf\xf9\xef\x12\xd8\xc7\x0c\x7f\x61\xb3\x94\x09\x09\x8e\xba\x0d\xd4\x87\xee\x2f\xcf\xfa\x3f\x53\x98\x5f\xd4\xdc\xd9\xb2\x7f\x7e\xd7\x24\x7b\xd0\xc4\x1a\xf5\x7f\x63\xd8\xd1\xcc\xda\xc5\x70\x18\xb9\x6d\x6b\xe5\xff\x27\x00\x00\xff\xff\x60\x96\x81\xca\x46\x25\x00\x00")

func bindataTemplatesCommandsTmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesCommandsTmpl,
		"templates/commands.tmpl",
	)
}



func bindataTemplatesCommandsTmpl() (*asset, error) {
	bytes, err := bindataTemplatesCommandsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/commands.tmpl",
		size: 9542,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1600784439, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataTemplatesMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\xc1\x4a\xc4\x40\x10\x44\xcf\xd3\x5f\xd1\xe4\x20\x09\xb8\x93\xf5\x9a\xdb\xa2\x39\x78\x71\x45\xc4\xfb\x38\xe9\x64\x1b\x33\xdd\xc3\x64\xb2\x04\x42\xfe\x5d\xb2\x22\x7b\xab\xaa\x47\x15\x15\x9d\xff\x71\x03\x61\x70\x2c\x00\x1c\xa2\xa6\x8c\x25\x98\x62\xe0\x7c\x99\xbf\xad\xd7\x50\xd3\xa2\x93\x77\x23\xd5\x1a\x49\x5c\xe4\x83\x1f\xf9\x30\x90\x50\x72\x59\x53\xed\x47\x2e\xa0\x02\xe8\x67\xf1\xb7\x9d\xb2\xc2\x15\x8c\x1f\xd9\xbe\x0a\xe7\xf2\x61\x57\xcf\x2a\x3d\x0f\x2b\x18\x73\x8a\xf1\xcd\x05\x6a\x10\xb1\x58\x57\xb4\xbb\xc1\x6d\x2b\x1e\xc1\x98\x56\xae\xef\x89\x7a\x5e\x9a\x3b\x6b\xe5\xfa\x8f\xbf\x28\x4d\xac\x72\xab\x3e\xd9\xa3\x3d\xee\xe9\x56\x01\x98\xba\xc6\xcf\xf3\xcb\xb9\xc1\x53\xd7\x61\xa2\x81\xa7\x4c\x09\xbd\x86\xe0\xa4\x9b\xf0\x42\x89\x2c\xfc\x7d\xfa\x50\xcd\xb6\x5d\xc8\xcf\x99\xca\x0a\x36\xf8\x0d\x00\x00\xff\xff\xe5\xd5\x27\x9a\x03\x01\x00\x00")

func bindataTemplatesMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_bindataTemplatesMainTmpl,
		"templates/main.tmpl",
	)
}



func bindataTemplatesMainTmpl() (*asset, error) {
	bytes, err := bindataTemplatesMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "templates/main.tmpl",
		size: 259,
		md5checksum: "",
		mode: os.FileMode(420),
		modTime: time.Unix(1600784439, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"templates/commands.tmpl": bindataTemplatesCommandsTmpl,
	"templates/main.tmpl":     bindataTemplatesMainTmpl,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"templates": {Func: nil, Children: map[string]*bintree{
		"commands.tmpl": {Func: bindataTemplatesCommandsTmpl, Children: map[string]*bintree{}},
		"main.tmpl": {Func: bindataTemplatesMainTmpl, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
