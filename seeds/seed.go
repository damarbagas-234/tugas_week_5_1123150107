package main

import (
	"github.com/joho/godotenv"
	"github.com/damarbagas-234/tugas_week_5_1123150107/config"
	"github.com/damarbagas-234/tugas_week_5_1123150107/models"
	"log"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{
		Name:        "Hoodie Polos Hitam",
		Price:       180000,
		Category:    "Casual",
		Stock:       25,
		Description: "Jaket hoodie polos bahan fleece nyaman dipakai harian",
		ImageURL:    "https://picsum.photos/400",
	},
	{
		Name:        "Bomber Navy",
		Price:       220000,
		Category:    "Casual",
		Stock:       20,
		Description: "Jaket bomber stylish dengan bahan lembut",
		ImageURL:    "https://picsum.photos/401",
	},
	{
		Name:        "Varsity Grey",
		Price:       250000,
		Category:    "Casual",
		Stock:       18,
		Description: "Jaket varsity modern dengan desain premium",
		ImageURL:    "https://picsum.photos/402",
	},
	{
		Name:        "Denim Jacket Blue",
		Price:       275000,
		Category:    "Casual",
		Stock:       15,
		Description: "Jaket denim klasik warna biru",
		ImageURL:    "https://picsum.photos/403",
	},
	{
		Name:        "Sweater Zip Abu",
		Price:       165000,
		Category:    "Casual",
		Stock:       30,
		Description: "Jaket sweater resleting untuk santai",
		ImageURL:    "https://picsum.photos/404",
	},
	{
		Name:        "Windbreaker Merah",
		Price:       230000,
		Category:    "Outdoor",
		Stock:       22,
		Description: "Jaket ringan tahan angin untuk outdoor",
		ImageURL:    "https://picsum.photos/405",
	},
	{
		Name:        "Parka Army",
		Price:       320000,
		Category:    "Outdoor",
		Stock:       12,
		Description: "Jaket parka tebal cocok cuaca dingin",
		ImageURL:    "https://picsum.photos/406",
	},
	{
		Name:        "Jaket Gunung Waterproof",
		Price:       350000,
		Category:    "Outdoor",
		Stock:       10,
		Description: "Jaket tahan air untuk aktivitas outdoor",
		ImageURL:    "https://picsum.photos/407",
	},
	{
		Name:        "Jaket Hiking Hijau",
		Price:       295000,
		Category:    "Outdoor",
		Stock:       14,
		Description: "Jaket hiking breathable dan nyaman",
		ImageURL:    "https://picsum.photos/408",
	},
	{
		Name:        "Softshell Outdoor",
		Price:       310000,
		Category:    "Outdoor",
		Stock:       16,
		Description: "Jaket softshell fleksibel untuk perjalanan",
		ImageURL:    "https://picsum.photos/409",
	},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
