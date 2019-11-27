package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	logger "github.com/paulxiong/cervical/webpage/2_api_server/log"
)

//Syscfg 系统配置的信息
type Syscfg struct {
	ID                   int       `json:"id"                     gorm:"column:id; primary_key"`        //记录ID
	Host                 string    `json:"host"                   gorm:"column:host"`                   //本服务域名+端口
	EmailRegisterContent string    `json:"email_register_content" gorm:"column:email_register_content"` //发送注册邮件的内容
	EmailForgotContent   string    `json:"email_forgot_content"   gorm:"column:email_forgot_content"`   //发送忘记密码邮件的格式
	RefererEn            int64     `json:"referer_en"             gorm:"column:referer_en"`             //开启图片防盗链,0关闭1开启
	Referer404URL        string    `json:"referer_404_url"        gorm:"column:referer_404_url"`        //图片不存在,默认图片
	Referer401URL        string    `json:"referer_401_url"        gorm:"column:referer_401_url"`        //非法请求,默认图片
	CreatedBy            int64     `json:"created_by"             gorm:"column:created_by"`             //创建者
	CreatedAt            time.Time `json:"created_at"             gorm:"column:created_at"`             //创建时间
	UpdatedAt            time.Time `json:"updated_at"             gorm:"column:updated_at"`             //更新时间
}

// BeforeCreate insert 之前的hook
func (s *Syscfg) BeforeCreate(scope *gorm.Scope) error {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}
	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = time.Now()
	}
	return nil
}

// UpdateSysCfg 更新系统配置的信息
func (s *Syscfg) UpdateSysCfg() (err error) {
	var updates map[string]interface{}
	updates = make(map[string]interface{})

	if s.Host != "" {
		updates["host"] = s.Host
	}
	if s.EmailRegisterContent != "" {
		updates["email_register_content"] = s.EmailRegisterContent
	}
	if s.EmailForgotContent != "" {
		updates["email_forgot_content"] = s.EmailForgotContent
	}
	if len(updates) < 1 {
		return nil
	}

	ret := db.Model(s).Where("id=?", s.ID).Updates(updates)
	if ret.Error != nil {
		logger.Info.Println(ret.Error)
	}
	return ret.Error
}

// FindSysCfg 查找系统配置信息
func FindSysCfg() (*Syscfg, error) {
	_s := Syscfg{}
	ret := db.Model(&_s).First(&_s)
	return &_s, ret.Error
}

// NewOrUpdateSysCfg 新建系统配置信息, 如果已经存在就更新
func (s *Syscfg) NewOrUpdateSysCfg() error {
	_s := Syscfg{}
	ret := db.Model(&_s).First(&_s)
	if ret.Error != nil || _s.ID < 1 {
		ret2 := db.Create(s)
		if ret2.Error != nil {
			return ret2.Error
		}
	}

	s.ID = _s.ID
	err := s.UpdateSysCfg()
	return err
}

// GetEmailBody 根据注册码和邮箱地址生成邮件内容
func GetEmailBody(toaddr string, code string, _type int) string {
	_s, err := FindSysCfg()
	if err != nil || len(_s.EmailRegisterContent) < 1 {
		return ""
	}

	emailbody := strings.Replace(_s.EmailRegisterContent, "email@gmail.com", toaddr, -1)
	if _type == 2 {
		emailbody = strings.Replace(_s.EmailForgotContent, "email@gmail.com", toaddr, -1)
	}
	emailbody = strings.Replace(emailbody, "000000", code, -1)
	return emailbody
}
