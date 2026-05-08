package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `grom:"column:created_by" json:"created_by"`
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

func ExistsTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// EditTag 编辑标签
func EditTag(id int, data map[string]interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)
	return true
}

// DeleteTag 删除标签
func DeleteTag(id int) bool {
	db.Where("id = ? ", id).Delete(&Tag{})
	return true
}
