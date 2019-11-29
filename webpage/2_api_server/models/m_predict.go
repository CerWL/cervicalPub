package models

import (
	"time"

	"github.com/jinzhu/gorm"
	logger "github.com/paulxiong/cervical/webpage/2_api_server/log"
)

// Predict 单个细胞的预测结果
type Predict struct {
	ID           int64     `json:"id"            gorm:"column:id; primary_key" example:"7"`   //预测信息ID
	ImgID        int64     `json:"imgid"         gorm:"column:imgid"           example:"7"`   //所属图片的ID
	PID          int64     `json:"pid"           gorm:"column:pid"             example:"7"`   //所属项目的ID
	X1           int       `json:"x1"            gorm:"column:x1"              example:"100"` //左上角X坐标 单位是像素
	Y1           int       `json:"y1"            gorm:"column:y1"              example:"100"` //左上角Y坐标 单位是像素
	X2           int       `json:"x2"            gorm:"column:x2"              example:"100"` //右下角X坐标 单位是像素
	Y2           int       `json:"y2"            gorm:"column:y2"              example:"100"` //预右下角Y坐标 单位是像素
	CellPath     string    `json:"cellpath"      gorm:"column:cellpath"        example:"100"` //上述坐标切割出来的细胞
	PredictScore int       `json:"predict_score" gorm:"column:predict_score"   example:"100"` //预测得分 50表示50%
	PredictType  int       `json:"predict_type"  gorm:"column:predict_type"    example:"100"` //预测的细胞类型,1到15是细胞类型, 50阴性 51阳性 100 未知, 200 不是细胞
	PredictP1n0  int       `json:"predict_p1n0"  gorm:"column:predict_p1n0"    example:"100"` //预测阴/阳性
	TrueType     int       `json:"true_type"     gorm:"column:true_type"       example:"100"` //医生标注的细胞类型 默认等于predict_type
	TrueP1n0     int       `json:"true_p1n0"     gorm:"column:true_p1n0"       example:"100"` //医生标注的阴/阳性 默认等于predict_p1n0
	VID          int64     `json:"vid"           gorm:"column:vid"             example:"7"`   //谁去做审核(verify)的用户ID
	Status       int       `json:"status"        gorm:"column:status"          example:"100"` //状态 0 未审核 1 已审核 2 移除 3 管理员确认
	CreatedBy    int64     `json:"created_by"    gorm:"column:created_by"`                    //创建者
	CreatedAt    time.Time `json:"created_at"    gorm:"column:created_at"`                    //创建时间
	UpdatedAt    time.Time `json:"updated_at"    gorm:"column:updated_at"`                    //更新时间
}

// BeforeCreate insert 之前的hook
func (p *Predict) BeforeCreate(scope *gorm.Scope) error {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = time.Now()
	}
	return nil
}

// FindPredictbyID 按照ID查询审核
func FindPredictbyID(id int64) (*Predict, error) {
	_p := &Predict{}
	ret := db.Model(_p).Where("id=?", id).First(_p)
	if ret.Error != nil {
		return _p, ret.Error
	}
	return _p, nil
}

// CreatePredicts 新建细胞预测结果，一个很长的数组
func CreatePredicts(predicts []*Predict, pid int64) (e error) {
	// 删除当前项目的所有预测结果
	db.Where("pid=?", pid).Delete(Predict{})

	_db := db.Begin()
	// 新增预测结果
	for _, v := range predicts {
		ret2 := _db.Create(&v)
		if ret2.Error != nil {
			logger.Info.Println(ret2.Error)
		}
	}
	_db.Commit()

	return nil
}

// ListPredict 通过图片ID查找预测
func ListPredict(limit int, skip int, pid int, imgid int, status int) (totalNum int64, p []Predict, e error) {
	var _p []Predict
	var total int64 = 0

	//0 未审核 1 已审核 2 移除 3 管理员已确认 4 未审核+已审核
	if status == 4 {
		var _p2 []Predict
		var total2 int64 = 0
		var _p3 []Predict
		var total3 int64 = 0
		db.Model(&Predict{}).Where("pid=? AND imgid=? AND status=?", pid, imgid, 0).Count(&total2)
		ret := db.Model(&Predict{}).Where("pid=? AND imgid=? AND status=?", pid, imgid, 0).Limit(limit).Offset(skip).Find(&_p2)
		if ret.Error != nil {
			logger.Info.Println(ret.Error)
		}
		db.Model(&Predict{}).Where("pid=? AND imgid=? AND status=?", pid, imgid, 1).Count(&total3)
		ret = db.Model(&Predict{}).Where("pid=? AND imgid=? AND status=?", pid, imgid, 1).Limit(limit).Offset(skip).Find(&_p3)
		if ret.Error != nil {
			logger.Info.Println(ret.Error)
		}
		_p = append(append(_p, _p2...), _p3...)
		total = total + total2 + total3

		return total, _p, ret.Error
	}

	db.Model(&Predict{}).Where("pid=? AND imgid=? AND status=?", pid, imgid, status).Count(&total)
	ret := db.Model(&Predict{}).Where("pid=? AND imgid=? AND status=?", pid, imgid, status).Limit(limit).Offset(skip).Find(&_p)
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
	}
	return total, _p, ret.Error
}

// UpdatePredict 更新审核信息
func UpdatePredict(id int64, trueType int) (e error) {
	_, err := FindPredictbyID(id)
	if err != nil {
		return err
	}

	// 审核细胞类型,1到15是细胞类型, 50 阴性 51 阳性 100 未知, 200 不是细胞
	// TODO: 类型计算
	var trueP1n0 int = 0

	// 状态 0 未审核 1 已审核 2 移除 3 管理员确认
	var status int = 1

	ret := db.Model(&Predict{}).Where("id=?", id).Updates(map[string]interface{}{"true_type": trueType, "true_p1n0": trueP1n0, "status": status})
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
	}
	return ret.Error
}

// GetPredictPercentByImgID 根据图片ID返回当前图片以及当前项目的审核进度
func GetPredictPercentByImgID(imgid int64, pid int64, status int) (int, int, int, int) {
	// 项目所有细胞，　项目已经审核的细胞，　当前图片的所有细胞，当前图片已经审核的细胞，
	cntCellsAll, cntCellsVerified, cntImgCellsAll, cntImgCellsVerified := 0, 0, 0, 0
	db.Model(&Predict{}).Where("pid=?", pid).Count(&cntCellsAll)
	db.Model(&Predict{}).Where("pid=? AND status=?", pid, status).Count(&cntCellsVerified)
	db.Model(&Predict{}).Where("pid=? AND imgid=?", pid, imgid).Count(&cntImgCellsAll)
	db.Model(&Predict{}).Where("pid=? AND imgid=? AND status=?", pid, imgid, status).Count(&cntImgCellsVerified)
	return cntCellsAll, cntCellsVerified, cntImgCellsAll, cntImgCellsVerified
}

// GetPredictPercent 更新审核信息时候顺带返回当前图片以及当前项目的审核进度
func GetPredictPercent(id int64, status int) (int, int, int, int) {
	var _p Predict
	// 项目所有细胞，　项目已经审核的细胞，　当前图片的所有细胞，当前图片已经审核的细胞，
	ret := db.Model(&Predict{}).Where("id=?", id).First(&_p)
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
		return 0, 0, 0, 0
	}
	return GetPredictPercentByImgID(_p.ImgID, _p.PID, status)
}

// ReviewPredict 管理员检查医生审核过后的信息
func ReviewPredict(id int64, status int) (e error) {
	_, err := FindPredictbyID(id)
	if err != nil {
		return err
	}

	ret := db.Model(&Predict{}).Where("id=?", id).Updates(map[string]interface{}{"status": status})
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
	}
	return ret.Error
}