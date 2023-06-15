package models

type Elm struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Dir       string `json:"dir"`
	Path      string `json:"path"`
	ParentDir string `json:"parent_dir"`
	Checksum  string `json:"checksum"`
	Size      uint64 `json:"size"`
}
