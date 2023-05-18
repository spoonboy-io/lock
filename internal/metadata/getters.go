package metadata

import "errors"

var (
	ERR_ID_NOT_FOUND   = errors.New("no template with that id found")
	ERR_NAME_NOT_FOUND = errors.New("no template with that name found")
)

// GetByName will iterate the metadata to retrieve by name key
// we also return the index as useful
func (md *Metadata) GetByName(key string) (Plugin, int, error) {
	id := 0
	for _, p := range *md {
		id++
		if p.Name == key {
			return p.Plugin, id, nil
		}
	}
	return Plugin{}, id, ERR_NAME_NOT_FOUND
}

// GetByIndex will iterate the metadata to retrieve by index
func (md *Metadata) GetByIndex(id int) (Plugin, error) {
	for i, p := range *md {
		if id == i+1 {
			return p.Plugin, nil
		}
	}
	return Plugin{}, ERR_ID_NOT_FOUND
}
