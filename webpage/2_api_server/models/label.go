package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	logger "github.com/paulxiong/cervical/webpage/2_api_server/log"
)

// Label 标注的信息
type Label struct {
	ID        int64     `json:"id"            gorm:"column:ID; primary_key" example:"1"`    //标注信息ID
	Imgid     int64     `json:"imgid"         gorm:"column:IMGID"           example:"2"`    //所属图片的ID
	Type      int       `json:"type"          gorm:"column:TYPE"            example:"1"`    //标注的细胞的类型
	TypeOut   string    `json:"typeout"       gorm:"-"                      example:"Norm"` //类型字符串，前端使用数据库没有
	X         int       `json:"x"             gorm:"column:X"               example:"320"`  //X中心
	Y         int       `json:"y"             gorm:"column:Y"               example:"480"`  //Y中心
	W         int       `json:"w"             gorm:"column:W"               example:"100"`  //宽
	H         int       `json:"h"             gorm:"column:H"               example:"100"`  //高
	Status    int       `json:"status"        gorm:"column:status"          example:"1"`    //状态, 0 未审核 1 已审核 2 移除
	CreatedBy int64     `json:"created_by"    gorm:"column:created_by"`                     //创建者
	CreatedAt time.Time `json:"created_time"  gorm:"column:CREATED_TIME"`
	UpdatedAt time.Time `json:"updated_time"  gorm:"column:UPDATED_TIME"`
}

/*
alter table c_label add created_by BIGINT   DEFAULT 0 COMMENT '创建者' after H;
alter table c_label add status INT   DEFAULT 0 COMMENT '状态 0 未审核 1 已审核 2 移除' after H;
*/

// BeforeCreate insert之前的hook
func (l *Label) BeforeCreate(scope *gorm.Scope) error {
	if l.CreatedAt.IsZero() {
		l.CreatedAt = time.Now()
	}
	if l.UpdatedAt.IsZero() {
		l.UpdatedAt = time.Now()
	}
	logger.Info.Println(l.ID)
	return nil
}

// InsertLabel 新建标注信息
func (l *Label) InsertLabel() (e error) {
	_l := Label{}
	ret := db.Model(l).Where("X=? AND Y=? AND W=? AND H=? AND TYPE=?", l.X, l.Y, l.W, l.H, l.Type).First(&_l)
	if ret.Error == nil && _l.W > 0 {
		return nil
	}

	l.ID = 0
	//ret2 := db.Debug().Model(l).Save(l)
	ret2 := db.Model(l).Save(l)
	if ret2.Error != nil {
		logger.Info.Println(ret2.Error)
	}

	return ret2.Error
}

// RemoveLabel 删除标注信息
func (l *Label) RemoveLabel() (e error) {
	_l := Label{}
	ret := db.Model(l).Where("X=? AND Y=? AND W=? AND H=? AND TYPE=?", l.X, l.Y, l.W, l.H, l.Type).First(&_l)
	if ret.Error != nil {
		return ret.Error
	}

	ret2 := db.Model(l).Where("ID=?", l.ID).Updates(map[string]interface{}{"status": 2})
	if ret2.Error != nil {
		logger.Info.Println(ret2.Error)
	}

	return ret2.Error
}

// UpdateLabel 更新标注信息
func (l *Label) UpdateLabel() (e error) {
	_l := Label{}
	ret := db.Model(l).Where("X=? AND Y=? AND W=? AND H=? AND TYPE=?", l.X, l.Y, l.W, l.H, l.Type).First(&_l)
	if ret.Error != nil {
		return ret.Error
	}

	ret2 := db.Model(l).Where("ID=?", l.ID).Updates(map[string]interface{}{"X": l.X, "Y": l.Y, "W": l.W, "H": l.H, "TYPE": l.Type})
	if ret2.Error != nil {
		logger.Info.Println(ret2.Error)
	}

	return ret2.Error
}

// ListLabel 依次列出标注信息
func ListLabel(limit int, skip int) (totalNum int64, l []Label, e error) {
	var _l []Label
	var total int64 = 0

	db.Model(&Label{}).Count(&total)
	ret := db.Model(&Label{}).Limit(limit).Offset(skip).Find(&_l)
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
	}
	return total, _l, ret.Error
}

// ListLabelByType 通过标注细胞类型列出标注信息
func ListLabelByType(limit int, skip int, t int) (totalNum int64, l []Label, e error) {
	var _l []Label
	var total int64 = 0

	db.Model(&Label{}).Where("TYPE = ?", t).Count(&total)
	ret := db.Model(&Label{}).Where("TYPE = ?", t).Limit(limit).Offset(skip).Find(&_l)
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
	}
	return total, _l, ret.Error
}

// ListLabelByImageID 找出一张图片里面的所有标注信息
func ListLabelByImageID(limit int, skip int, imgid int) (totalNum int64, l []Label, e error) {
	var _l []Label
	var total int64 = 0

	db.Model(&Label{}).Where("IMGID = ?", imgid).Count(&total)
	ret := db.Model(&Label{}).Where("IMGID = ?", imgid).Limit(limit).Offset(skip).Find(&_l)
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
	}
	return total, _l, ret.Error
}

// ListLabelCountByPN 统计P、N的标注总数
func ListLabelCountByPN(pn int) (totalNum int64, e error) {
	type res struct {
		Total int64
	}
	selector1 := "SELECT COUNT(c_label.IMGID) as total FROM c_label,c_category  where c_label.TYPE=c_category.ID AND c_category.P1N0=?;"
	ress := res{}
	ret := db.Raw(selector1, pn).Scan(&ress)
	if ret.Error != nil {
		log.Println(ret.Error)
	}

	return ress.Total, ret.Error
}
