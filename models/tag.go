package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreateBy   string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tag []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tag)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistsTagByName(name string) bool {
	var tag Tag
	db.Where("name = ?", name).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:     name,
		State:    state,
		CreateBy: createBy,
	})
	return true
}
