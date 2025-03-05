package entity

import "time"

// Author は "Code" を主キーとする
type Author struct {
	// 例: "AU12345" のような一意コード
	Code        string     `gorm:"primaryKey;size:20"`
	Name        string     `gorm:"size:255;not null"`
	Profile     string     `gorm:"type:text"` // 長文可
	Nationality string     `gorm:"size:50"`
	BornDate    *time.Time // 生年月日(無い場合はNULL)
	Email       string     `gorm:"size:255"`
	Website     string     `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Book は "ISBN" を主キーとする
type Book struct {
	ISBN        string  `gorm:"primaryKey;size:20"`
	Title       string  `gorm:"size:255;not null"`
	Genre       string  `gorm:"size:100"`
	Price       float64 // 金額
	PageCount   int
	PublishedAt *time.Time // 出版日 (NULL可)

	// 著者を参照
	AuthorCode string `gorm:"size:20;not null"`
	Author     Author `gorm:"foreignKey:AuthorCode;references:Code"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
