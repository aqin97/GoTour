package model

import (
	"GoTour/blog_service/global"
	"GoTour/blog_service/pkg/app"
	"GoTour/blog_service/pkg/setting"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//公共Model，包含了所有model包含的公共字段
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreateOn   uint32 `json:"create_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn   uint32 `json:"delete_on"`
	IsDel      uint8  `json:"is_del"`
}

//标签model
type Tag struct {
	*Model
	Name  string `json:"name"`
	State string `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}

//文章的model
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         string `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}

//文章和标签的关联model
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (at ArticleTag) TableName() string {
	return "blog_article_tag"
}

func NewDBEngine(databaseSettings *setting.DatabaseSettings) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSettings.DBType, fmt.Sprintf(s, databaseSettings.Username,
		databaseSettings.Password,
		databaseSettings.Host,
		databaseSettings.DBName,
		databaseSettings.Charset,
		databaseSettings.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSettings.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSettings.MaxOpenConns)

	return db, nil
}
