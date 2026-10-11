package database

import (
	"log"
	"strings"

	"gorm.io/gorm"
)

// seedData contains the original 25 products (frontend/lib/seed.js);
// generatedProducts() in seed_catalog.go adds the expanded catalog.
type seedProduct struct {
	Name        string
	Description string
	Price       float64
	Picture     string
	Categories  []string
}

var seedData = []seedProduct{
	{Name: "MacBook Pro", Description: "Powerful laptop for professionals with M3 chip and stunning Retina display.", Price: 1299, Picture: "/products/macbook.png", Categories: []string{"Laptops"}},
	{Name: "ASUS ROG", Description: "High-performance gaming laptop with RTX graphics and 165Hz display.", Price: 1499, Picture: "/products/rog.png", Categories: []string{"Laptops", "Gaming"}},
	{Name: "MSI Gaming", Description: "Premium gaming laptop with mechanical keyboard and RGB lighting.", Price: 1399, Picture: "/products/msi.png", Categories: []string{"Laptops", "Gaming"}},
	{Name: "Dell XPS 15", Description: "Ultra-thin laptop with OLED touchscreen and Intel Core i9 processor.", Price: 1099, Picture: "/products/xps15new.png", Categories: []string{"Laptops"}},
	{Name: "Lenovo ThinkPad X1", Description: "Business-class laptop with legendary keyboard and robust build.", Price: 1249, Picture: "/products/x1.png", Categories: []string{"Laptops"}},
	{Name: "HP Spectre x360", Description: "Convertible 2-in-1 laptop with 4K display and gem-cut design.", Price: 1149, Picture: "/products/x360new.png", Categories: []string{"Laptops"}},
	{Name: "Razer Blade 15", Description: "Sleek gaming laptop with RTX 4070 and 240Hz QHD display.", Price: 1599, Picture: "/products/blade.png", Categories: []string{"Laptops", "Gaming"}},
	{Name: "iPhone 15", Description: "Latest Apple smartphone with A17 Pro chip and titanium design.", Price: 999, Picture: "/products/iphone.png", Categories: []string{"Phones"}},
	{Name: "Samsung Galaxy S24", Description: "Flagship Android phone with AI features and stunning camera.", Price: 899, Picture: "/products/galaxy.png", Categories: []string{"Phones"}},
	{Name: "Xiaomi Redmi Note", Description: "Feature-packed budget smartphone with great battery life.", Price: 399, Picture: "/products/redmi.png", Categories: []string{"Phones"}},
	{Name: "Google Pixel 9", Description: "Pure Android experience with Tensor G4 chip and amazing camera.", Price: 799, Picture: "/products/pixelnew.png", Categories: []string{"Phones"}},
	{Name: "OnePlus 12", Description: "Fast and smooth flagship with 100W charging and Hasselblad cameras.", Price: 699, Picture: "/products/12.png", Categories: []string{"Phones"}},
	{Name: "Nothing Phone (2)", Description: "Unique transparent design with Glyph Interface and clean OS.", Price: 599, Picture: "/products/np2new.png", Categories: []string{"Phones"}},
	{Name: "Oppo Find X7", Description: "Ultra flagship with periscope zoom and MediaTek Dimensity 9300.", Price: 749, Picture: "/products/x7.png", Categories: []string{"Phones"}},
	{Name: "AirPods Pro", Description: "Wireless earbuds with active noise cancellation and spatial audio.", Price: 249, Picture: "/products/airpods.png", Categories: []string{"Accessories"}},
	{Name: "Huawei FreeBuds", Description: "Premium wireless earbuds with long battery life and comfort fit.", Price: 179, Picture: "/products/freebuds.png", Categories: []string{"Accessories"}},
	{Name: "Gaming Headset", Description: "7.1 surround sound headset with noise-cancelling microphone.", Price: 129, Picture: "/products/headset.png", Categories: []string{"Accessories", "Gaming"}},
	{Name: "Apple Watch Series 9", Description: "Advanced health tracking with S9 chip and always-on Retina display.", Price: 399, Picture: "/products/applenew.png", Categories: []string{"Accessories"}},
	{Name: "Samsung Galaxy Watch", Description: "Feature-rich smartwatch with Wear OS and BioActive sensor.", Price: 329, Picture: "/products/watch.png", Categories: []string{"Accessories"}},
	{Name: "Logitech MX Master 3S", Description: "Premium ergonomic mouse with quiet clicks and MagSpeed scroll.", Price: 99, Picture: "/products/mx.png", Categories: []string{"Accessories"}},
	{Name: "Anker Power Bank", Description: "20000mAh high-capacity fast-charging portable battery pack.", Price: 79, Picture: "/products/powerbank.png", Categories: []string{"Accessories"}},
	{Name: "Nintendo Switch OLED", Description: "Hybrid console with vibrant 7-inch OLED screen and enhanced audio.", Price: 349, Picture: "/products/switch.png", Categories: []string{"Gaming"}},
	{Name: "PS5 DualSense", Description: "Wireless controller with haptic feedback and adaptive triggers.", Price: 74, Picture: "/products/ps5.png", Categories: []string{"Gaming"}},
	{Name: "Xbox Series X", Description: "Most powerful Xbox with 12 TFLOPS and 1TB SSD for quick loading.", Price: 499, Picture: "/products/xbox.png", Categories: []string{"Gaming"}},
	{Name: "Samsung Odyssey G7", Description: "32-inch 4K 144Hz curved gaming monitor with 1ms response time.", Price: 799, Picture: "/products/g7.png", Categories: []string{"Gaming"}},
}

// Seed inserts any seed products (original 25 + generated catalog) missing by name.
// Skips entirely once the products table already holds at least that many rows.
func Seed() {
	all := append(append([]seedProduct{}, seedData...), generatedProducts()...)
	syncPictures(all)

	var count int64
	DB.Model(&Product{}).Count(&count)
	if count >= int64(len(all)) {
		log.Printf("Database already seeded (%d products found). Skipping.", count)
		return
	}

	log.Printf("Seeding database (%d products present, %d in catalog)...", count, len(all))

	added := 0
	err := DB.Transaction(func(tx *gorm.DB) error {
		var names []string
		if err := tx.Model(&Product{}).Pluck("name", &names).Error; err != nil {
			return err
		}
		existing := make(map[string]bool, len(names))
		for _, n := range names {
			existing[n] = true
		}

		categoryMap := make(map[string]Category)
		for _, sp := range all {
			if existing[sp.Name] {
				continue
			}
			var cats []Category
			for _, name := range sp.Categories {
				cat, ok := categoryMap[name]
				if !ok {
					cat = Category{Name: name}
					if err := tx.Where("name = ?", name).FirstOrCreate(&cat).Error; err != nil {
						return err
					}
					categoryMap[name] = cat
				}
				cats = append(cats, cat)
			}
			product := Product{
				Name:        sp.Name,
				Description: sp.Description,
				Price:       sp.Price,
				Picture:     sp.Picture,
				Categories:  cats,
			}
			if err := tx.Create(&product).Error; err != nil {
				return err
			}
			added++
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Seed failed: %v", err)
	}

	log.Printf("Seeded %d new products successfully", added)
}

// syncPictures points already-seeded products at the catalog's current image
// in one statement, so image changes reach existing databases.
// ponytail: overwrites admin-edited pictures on seed products; key on a seed flag if that matters.
func syncPictures(all []seedProduct) {
	values := make([]string, len(all))
	args := make([]interface{}, 0, len(all)*2)
	for i, sp := range all {
		values[i] = "(?, ?)"
		args = append(args, sp.Name, sp.Picture)
	}
	res := DB.Exec(`UPDATE products AS p SET picture = v.pic
		FROM (VALUES `+strings.Join(values, ",")+`) AS v(name, pic)
		WHERE p.name = v.name AND p.picture IS DISTINCT FROM v.pic`, args...)
	if res.Error != nil {
		log.Printf("Picture sync failed: %v", res.Error)
		return
	}
	if res.RowsAffected > 0 {
		log.Printf("Updated pictures on %d products", res.RowsAffected)
	}
}
