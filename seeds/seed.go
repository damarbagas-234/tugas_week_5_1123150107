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
		ImageURL:    "https://i.ibb.co.com/d04dbDb8/black-hoodie-shown-front-back-260nw-2557275213.webp",
	},
	{
		Name:        "Bomber Navy",
		Price:       220000,
		Category:    "Casual",
		Stock:       20,
		Description: "Jaket bomber stylish dengan bahan lembut",
		ImageURL:    "https://i.ibb.co.com/CKnC2XjW/2083aeaa-c3b3-grande.webp",
	},
	{
		Name:        "Varsity Grey",
		Price:       250000,
		Category:    "Casual",
		Stock:       18,
		Description: "Jaket varsity modern dengan desain premium",
		ImageURL:    "https://i.ibb.co.com/PvPbSjv9/Grey-Varsity-Jacket-500x600.webp",
	},
	{
		Name:        "Denim Jacket Blue",
		Price:       275000,
		Category:    "Casual",
		Stock:       15,
		Description: "Jaket denim klasik warna biru",
		ImageURL:    "https://i.ibb.co.com/BHfxxr6s/c5f225d0-e38e-43cf-9c74-4034bc0561bd1739614460637-LOCOMOTIVE-Men-Blue-Solid-Denim-Jacket-42817396144.jpg",
	},
	{
		Name:        "Sweater Zip Abu",
		Price:       165000,
		Category:    "Casual",
		Stock:       30,
		Description: "Jaket sweater resleting untuk santai",
		ImageURL:    "https://i.ibb.co.com/8DCsBBgD/04729702811-000-e1.webp",
	},
	{
		Name:        "Windbreaker Merah",
		Price:       230000,
		Category:    "Outdoor",
		Stock:       22,
		Description: "Jaket ringan tahan angin untuk outdoor",
		ImageURL:    "https://i.ibb.co.com/LX207vfB/triple-jaket-parasut-windbreaker-merah-4853503-3-68398.webp",
	},
	{
		Name:        "Parka Army",
		Price:       320000,
		Category:    "Outdoor",
		Stock:       12,
		Description: "Jaket parka tebal cocok cuaca dingin",
		ImageURL:    "https://i.ibb.co.com/tcfR2Nr/728e82154514bf9e3fa98f13415c7d2f.jpg",
	},
	{
		Name:        "Jaket Gunung Waterproof",
		Price:       350000,
		Category:    "Outdoor",
		Stock:       10,
		Description: "Jaket tahan air untuk aktivitas outdoor",
		ImageURL:    "https://i.ibb.co.com/hJQ81DvC/id-11134207-7rbkb-m7hdlyd9c71120.webp",
	},
	{
		Name:        "Jaket Hiking Hijau",
		Price:       295000,
		Category:    "Outdoor",
		Stock:       14,
		Description: "Jaket hiking breathable dan nyaman",
		ImageURL:    "https://i.ibb.co.com/ZRLnbk77/photo-1669082072806-84c5d3ae7b15.avif",
	},
	{
		Name:        "Softshell Outdoor",
		Price:       310000,
		Category:    "Outdoor",
		Stock:       16,
		Description: "Jaket softshell fleksibel untuk perjalanan",
		ImageURL:    "https://i.ibb.co.com/XZgfh9sL/premium-photo-1697183202186-8ef20ffa6ab0.avif",
	},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
