// Code generated by go-bindata.
// sources:
// 1085649617_create_users_table.down.sql
// 1085649617_create_users_table.up.sql
// 1185749658_add_city_to_users.down.sql
// 1185749658_add_city_to_users.up.sql
// DO NOT EDIT!

package testdata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
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

var __1085649617_create_users_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x2d\x4e\x2d\x2a\xb6\xe6\x02\x04\x00\x00\xff\xff\x2c\x02\x3d\xa7\x1c\x00\x00\x00")

func _1085649617_create_users_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1085649617_create_users_tableDownSql,
		"1085649617_create_users_table.down.sql",
	)
}

func _1085649617_create_users_tableDownSql() (*asset, error) {
	bytes, err := _1085649617_create_users_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1085649617_create_users_table.down.sql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1485750305, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1085649617_create_users_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\xd0\xe0\x52\x00\xb3\xe2\x33\x53\x14\x32\xf3\x4a\x52\xd3\x53\x8b\x14\x4a\xf3\x32\x0b\x4b\x53\x75\xb8\x14\x14\xf2\x12\x73\x53\x15\x14\x14\x14\xca\x12\x8b\x92\x33\x12\x8b\x34\x4c\x0c\x34\x41\xc2\xa9\xb9\x89\x99\x39\xa8\xc2\x5c\x9a\xd6\x5c\x80\x00\x00\x00\xff\xff\xa3\x57\xbc\x0b\x5f\x00\x00\x00")

func _1085649617_create_users_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1085649617_create_users_tableUpSql,
		"1085649617_create_users_table.up.sql",
	)
}

func _1085649617_create_users_tableUpSql() (*asset, error) {
	bytes, err := _1085649617_create_users_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1085649617_create_users_table.up.sql", size: 95, mode: os.FileMode(420), modTime: time.Unix(1485803085, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1185749658_add_city_to_usersDownSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\xce\x2c\xa9\xb4\xe6\x02\x04\x00\x00\xff\xff\xb7\x52\x88\xd7\x2e\x00\x00\x00")

func _1185749658_add_city_to_usersDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1185749658_add_city_to_usersDownSql,
		"1185749658_add_city_to_users.down.sql",
	)
}

func _1185749658_add_city_to_usersDownSql() (*asset, error) {
	bytes, err := _1185749658_add_city_to_usersDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1185749658_add_city_to_users.down.sql", size: 46, mode: os.FileMode(420), modTime: time.Unix(1485750443, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1185749658_add_city_to_usersUpSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x2d\x4e\x2d\x2a\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xce\x2c\xa9\x54\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x34\x30\xd0\xb4\xe6\xe2\xe2\x02\x04\x00\x00\xff\xff\xa8\x0f\x49\xc6\x32\x00\x00\x00")

func _1185749658_add_city_to_usersUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1185749658_add_city_to_usersUpSql,
		"1185749658_add_city_to_users.up.sql",
	)
}

func _1185749658_add_city_to_usersUpSql() (*asset, error) {
	bytes, err := _1185749658_add_city_to_usersUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1185749658_add_city_to_users.up.sql", size: 50, mode: os.FileMode(420), modTime: time.Unix(1485843733, 0)}
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
	"1085649617_create_users_table.down.sql": _1085649617_create_users_tableDownSql,
	"1085649617_create_users_table.up.sql":   _1085649617_create_users_tableUpSql,
	"1185749658_add_city_to_users.down.sql":  _1185749658_add_city_to_usersDownSql,
	"1185749658_add_city_to_users.up.sql":    _1185749658_add_city_to_usersUpSql,
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
	"1085649617_create_users_table.down.sql": &bintree{_1085649617_create_users_tableDownSql, map[string]*bintree{}},
	"1085649617_create_users_table.up.sql":   &bintree{_1085649617_create_users_tableUpSql, map[string]*bintree{}},
	"1185749658_add_city_to_users.down.sql":  &bintree{_1185749658_add_city_to_usersDownSql, map[string]*bintree{}},
	"1185749658_add_city_to_users.up.sql":    &bintree{_1185749658_add_city_to_usersUpSql, map[string]*bintree{}},
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
	err = os.WriteFile(_filePath(dir, name), data, info.Mode())
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
