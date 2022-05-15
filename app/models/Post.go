package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/joshua-ather/sv_be/app/helpers"
	"github.com/joshua-ather/sv_be/config"
	"github.com/labstack/echo/v4"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title        string    `gorm:"size:200;not null" json:"title"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	Category     string    `gorm:"size:100;not null;" json:"category"`
	Created_date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
	Updated_date time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_date"`
	Status       string    `gorm:"size:100;not null" json:"status"`
}

func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Category = html.EscapeString(strings.TrimSpace(p.Category))
	p.Created_date = time.Now()
	p.Updated_date = time.Now()
	p.Status = html.EscapeString(strings.TrimSpace(p.Status))
}

func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("Required Title")
	}
	if p.Content == "" {
		return errors.New("Required Content")
	}
	if p.Category == "" {
		return errors.New("Required Category")
	}
	if p.Status == "" {
		return errors.New("Required Status")
	}

	s := []string{"Publish", "Draft", "Thrash"}
	if !helpers.Contains(s, p.Status) {
		return errors.New("Status is invalid")
	}

	return nil
}

func (p *Post) SavePost() (*Post, error) {
	db := config.DB()

	var err error
	err = db.Debug().Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	return p, nil
}

func (p *Post) FindAllPost() (*[]Post, error) {
	db := config.DB()

	var err error
	var posts []Post
	err = db.Debug().Model(&Post{}).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	return &posts, err
}

func (p *Post) FindPostWLimitOffset(c echo.Context) (*[]Post, error) {
	db := config.DB()
	limit := c.Param("limit")
	offset := c.Param("offset")

	var err error
	var posts []Post
	err = db.Debug().Model(&Post{}).Limit(limit).Offset(offset).Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	return &posts, err
}

func (p *Post) FindPostByID(c echo.Context) (*Post, error) {
	db := config.DB()
	id := c.Param("id")

	var err error
	err = db.Debug().Model(Post{}).Where("id = ?", id).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Post{}, errors.New("Post Not Found")
	}
	return p, err
}

func (p *Post) UpdatePost(c echo.Context) (*Post, error) {
	db := config.DB()
	id := c.Param("id")

	db = db.Debug().Model(&Post{}).Where("id = ?", id).Take(&Post{}).UpdateColumns(
		map[string]interface{}{
			"title":        p.Title,
			"content":      p.Content,
			"category":     p.Category,
			"status":       p.Status,
			"updated_date": time.Now(),
		},
	)
	if db.Error != nil {
		return &Post{}, db.Error
	}
	// This is the display the updated article
	err := db.Debug().Model(&Post{}).Where("id = ?", id).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	return p, nil
}

func (p *Post) DeletePost(c echo.Context) (int64, error) {
	db := config.DB()
	id := c.Param("id")

	db = db.Debug().Model(&Post{}).Where("id = ?", id).Take(&Post{}).UpdateColumns(
		map[string]interface{}{
			"status":       "Thrash",
			"updated_date": time.Now(),
		},
	)

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
