package database

import (
	"log"
	"service_components/internal/model"

	"github.com/google/uuid"
)

func Seeder() {
	categories := []model.Category{
		{Name: "UI Kit", Slug: "ui-kit"},
		{Name: "Dashboard", Slug: "dashboard"},
		{Name: "Authentication", Slug: "authentication"},
	}

	for _, category := range categories {
		if err := DB.Where("slug = ?", category.Slug).FirstOrCreate(&category).Error; err != nil {
			log.Printf("Gagal menambahkan kategori: %s", err)
		}
	}

	tags := []model.Tag{
		{Name: "React", Slug: "react"},
		{Name: "Tailwind", Slug: "tailwind"},
		{Name: "Bootstrap", Slug: "bootstrap"},
	}

	for _, tag := range tags {
		if err := DB.Where("slug = ?", tag.Slug).FirstOrCreate(&tag).Error; err != nil {
			log.Printf("Gagal menambahkan tag: %s", err)
		}
	}

	var uiKit model.Category
	DB.Where("slug = ?", "ui-kit").First(&uiKit)

	components := []model.Component{
		{
			Name:        "Button",
			Slug:        "button",
			Description: "Reusable button component",
			CategoryID:  uiKit.ID,
			CodeJSX:     "<button className='btn'>Click me</button>",
			CodeCSS:     ".btn { padding: 8px; background-color: blue; }",
			UserID:      uuid.New(),
		},
		{
			Name:        "Card",
			Slug:        "card",
			Description: "Card component with shadow",
			CategoryID:  uiKit.ID,
			CodeJSX:     "<div className='card'>Card Content</div>",
			CodeCSS:     ".card { background: white; border-radius: 4px; }",
			UserID:      uuid.New(),
		},
	}

	for _, component := range components {
		if err := DB.Where("slug = ?", component.Slug).FirstOrCreate(&component).Error; err != nil {
			log.Printf("Gagal menambahkan komponen: %s", err)
		}
	}

	log.Println("Seeder selesai dijalankan.")
}
