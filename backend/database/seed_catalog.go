package database

import (
	"fmt"
	"math"
)

// Catalog expansion: each family is multiplied by its category's variants
// (e.g. storage tiers, sizes, colors) to produce ~500 extra seed products.
// Images come from Wikimedia Commons (credits: frontend/public/products/CREDITS.md);
// variants of a family share one photo.

type family struct {
	Name, Desc, Pic string
	Base            float64
	Extra           []string // categories beyond the group's own
}

type variant struct {
	Label string
	Mult  float64
}

type catalogGroup struct {
	Category string
	Variants []variant
	Families []family
}

var catalog = []catalogGroup{
	{"Laptops", []variant{{"8GB/256GB", 1}, {"16GB/512GB", 1.2}, {"32GB/1TB", 1.5}, {"64GB/2TB", 2}}, []family{
		{"MacBook Air 13 M3", "Fanless ultraportable with all-day battery and Liquid Retina display.", "/products/macbook-air-13-m3.jpg", 1099, nil},
		{"MacBook Pro 16 M3 Max", "Pro-grade laptop with Liquid Retina XDR display and 22-hour battery.", "/products/macbook-pro-16-m3-max.jpg", 2499, nil},
		{"ASUS Zenbook 14 OLED", "Slim aluminium laptop with a vivid 2.8K OLED panel.", "/products/asus-zenbook-14-oled.jpg", 899, nil},
		{"ASUS TUF Gaming A15", "Military-grade durable gaming laptop with RTX graphics.", "/products/asus-tuf-gaming-a15.jpg", 999, []string{"Gaming"}},
		{"Acer Swift Go 14", "Lightweight everyday laptop with Intel Core Ultra and OLED option.", "/products/acer-swift-go-14.jpg", 749, nil},
		{"Acer Predator Helios 16", "Desktop-class gaming laptop with liquid-metal cooling.", "/products/acer-predator-helios-16.jpg", 1599, []string{"Gaming"}},
		{"Lenovo Legion Pro 7", "Esports-ready laptop with 240Hz display and RTX 4080.", "/products/lenovo-legion-pro-7.jpg", 2199, []string{"Gaming"}},
		{"Lenovo Yoga Slim 7", "Premium thin-and-light with excellent speakers and battery life.", "/products/lenovo-yoga-slim-7.jpg", 949, nil},
		{"Dell XPS 13", "Compact flagship with edge-to-edge InfinityEdge display.", "/products/dell-xps-13.jpg", 999, nil},
		{"Dell Alienware m18", "Massive 18-inch gaming powerhouse with Cherry MX keyboard.", "/products/dell-alienware-m18.jpg", 2499, []string{"Gaming"}},
		{"HP Envy x360 15", "Versatile convertible with pen support and OLED touchscreen.", "/products/hp-envy-x360-15.jpg", 899, nil},
		{"Microsoft Surface Laptop 6", "Clean Windows experience with PixelSense touchscreen.", "/products/microsoft-surface-laptop-6.jpg", 999, nil},
	}},
	{"Phones", []variant{{"128GB", 1}, {"256GB", 1.12}, {"512GB", 1.3}, {"1TB", 1.55}}, []family{
		{"iPhone 15 Pro Max", "Titanium flagship with 5x telephoto camera and A17 Pro chip.", "/products/iphone-15-pro-max.jpg", 1199, nil},
		{"iPhone 15 Plus", "Big-screen iPhone with Dynamic Island and 48MP camera.", "/products/iphone-15-plus.jpg", 899, nil},
		{"Samsung Galaxy S24 Ultra", "Ultra flagship with built-in S Pen and 200MP camera.", "/products/samsung-galaxy-s24-ultra.jpg", 1299, nil},
		{"Samsung Galaxy Z Fold 6", "Foldable phone that opens into a 7.6-inch tablet.", "/products/samsung-galaxy-z-fold-6.jpg", 1899, nil},
		{"Samsung Galaxy A55", "Mid-range all-rounder with IP67 rating and 5000mAh battery.", "/products/samsung-galaxy-a55.jpg", 449, nil},
		{"Google Pixel 9 Pro", "Best-in-class computational photography with Gemini AI.", "/products/google-pixel-9-pro.jpg", 999, nil},
		{"Google Pixel 8a", "Affordable Pixel with Tensor G3 and seven years of updates.", "/products/google-pixel-8a.jpg", 499, nil},
		{"OnePlus 12R", "Flagship performance at a mid-range price with 100W charging.", "/products/oneplus-12r.jpg", 499, nil},
		{"Xiaomi 14 Ultra", "Leica-tuned quad camera system with 1-inch sensor.", "/products/xiaomi-14-ultra.jpg", 1099, nil},
		{"Xiaomi Poco F6", "Performance-first phone with Snapdragon 8s Gen 3.", "/products/xiaomi-poco-f6.jpg", 399, []string{"Gaming"}},
		{"Nothing Phone (2a)", "Distinctive Glyph design with a clean, fast interface.", "/products/nothing-phone-2a.jpg", 349, nil},
		{"Oppo Reno 12 Pro", "Slim design with portrait-focused telephoto camera.", "/products/oppo-reno-12-pro.jpg", 549, nil},
	}},
	{"Tablets", []variant{{"Wi-Fi 64GB", 1}, {"Wi-Fi 256GB", 1.25}, {"5G 256GB", 1.45}}, []family{
		{"iPad Pro 13 M4", "Ultra-thin tablet with Tandem OLED and M4 performance.", "/products/ipad-pro-13-m4.jpg", 1299, nil},
		{"iPad Air 11", "Powerful mid-size iPad with M2 chip and Apple Pencil support.", "/products/ipad-air-11.jpg", 599, nil},
		{"iPad mini", "Pocketable 8.3-inch tablet with A17 Pro chip.", "/products/ipad-mini.jpg", 499, nil},
		{"Samsung Galaxy Tab S9", "AMOLED Android tablet with S Pen included.", "/products/samsung-galaxy-tab-s9.jpg", 799, nil},
		{"Samsung Galaxy Tab S9 FE", "Water-resistant tablet for study and streaming.", "/products/samsung-galaxy-tab-s9-fe.jpg", 449, nil},
		{"Microsoft Surface Pro 10", "2-in-1 tablet that replaces your laptop.", "/products/microsoft-surface-pro-10.jpg", 1099, []string{"Laptops"}},
		{"Lenovo Tab P12", "Large 12.7-inch 3K display tablet for entertainment.", "/products/lenovo-tab-p12.jpg", 349, nil},
		{"Xiaomi Pad 6", "144Hz tablet with Snapdragon 870 and quad speakers.", "/products/xiaomi-pad-6.jpg", 399, nil},
		{"OnePlus Pad", "7:5 aspect ratio display with Dolby Vision.", "/products/oneplus-pad.jpg", 479, nil},
		{"Google Pixel Tablet", "Tablet that doubles as a smart home hub with speaker dock.", "/products/google-pixel-tablet.jpg", 499, []string{"Smart Home"}},
		{"Amazon Fire HD 10", "Budget-friendly tablet for reading and video.", "/products/amazon-fire-hd-10.jpg", 149, nil},
		{"Huawei MatePad Pro", "OLED tablet with ultra-slim bezels and stylus support.", "/products/huawei-matepad-pro.jpg", 699, nil},
	}},
	{"Audio", []variant{{"Black", 1}, {"White", 1}, {"Midnight Blue", 1.03}}, []family{
		{"Sony WH-1000XM5", "Industry-leading noise-cancelling over-ear headphones.", "/products/sony-wh-1000xm5.jpg", 399, nil},
		{"Bose QuietComfort Ultra", "Immersive spatial audio with world-class ANC.", "/products/bose-quietcomfort-ultra.jpg", 429, nil},
		{"Apple AirPods Max", "Over-ear headphones with computational audio.", "/products/apple-airpods-max.jpg", 549, nil},
		{"AirPods 4", "Open-fit earbuds with personalized spatial audio.", "/products/airpods-4.jpg", 129, nil},
		{"Sony WF-1000XM5", "Compact true-wireless earbuds with premium ANC.", "/products/sony-wf-1000xm5.jpg", 299, nil},
		{"Samsung Galaxy Buds3 Pro", "Blade-design earbuds with intelligent ANC.", "/products/samsung-galaxy-buds3-pro.jpg", 249, nil},
		{"Huawei FreeBuds Pro 3", "Hi-res wireless earbuds with dual-driver design.", "/products/huawei-freebuds-pro-3.jpg", 199, nil},
		{"Nothing Ear (a)", "Transparent earbuds with LDAC and strong ANC.", "/products/nothing-ear-a.jpg", 99, nil},
		{"JBL Charge 5", "Rugged portable Bluetooth speaker with 20-hour playtime.", "/products/jbl-charge-5.jpg", 179, nil},
		{"JBL Flip 6", "Waterproof compact speaker with punchy bass.", "/products/jbl-flip-6.jpg", 129, nil},
		{"Bose SoundLink Flex", "Durable outdoor speaker with PositionIQ technology.", "/products/bose-soundlink-flex.jpg", 149, nil},
		{"Sonos Era 100", "Smart speaker with stereo sound and Wi-Fi streaming.", "/products/sonos-era-100.jpg", 249, []string{"Smart Home"}},
		{"SteelSeries Arctis Nova 7", "Wireless multi-platform gaming headset.", "/products/steelseries-arctis-nova-7.jpg", 179, []string{"Gaming"}},
		{"HyperX Cloud III", "Comfortable wired gaming headset with DTS spatial audio.", "/products/hyperx-cloud-iii.jpg", 99, []string{"Gaming"}},
	}},
	{"Wearables", []variant{{"40mm", 1}, {"44mm", 1.08}, {"LTE 44mm", 1.25}}, []family{
		{"Apple Watch Ultra 2", "Rugged titanium smartwatch for endurance athletes.", "/products/apple-watch-ultra-2.jpg", 799, nil},
		{"Apple Watch SE", "Essential Apple Watch features at an accessible price.", "/products/apple-watch-se.jpg", 249, nil},
		{"Samsung Galaxy Watch 7", "Wear OS smartwatch with advanced sleep coaching.", "/products/samsung-galaxy-watch-7.jpg", 299, nil},
		{"Samsung Galaxy Watch Ultra", "Titanium outdoor smartwatch with 100-hour battery.", "/products/samsung-galaxy-watch-ultra.jpg", 649, nil},
		{"Google Pixel Watch 3", "Domed AMOLED smartwatch with Fitbit health tracking.", "/products/google-pixel-watch-3.jpg", 349, nil},
		{"Garmin Fenix 8", "Multisport GPS watch with AMOLED and flashlight.", "/products/garmin-fenix-8.jpg", 999, nil},
		{"Garmin Forerunner 265", "Running watch with training readiness and AMOLED display.", "/products/garmin-forerunner-265.jpg", 449, nil},
		{"Garmin Venu 3", "Lifestyle fitness smartwatch with voice calls.", "/products/garmin-venu-3.jpg", 449, nil},
		{"Fitbit Charge 6", "Slim fitness tracker with built-in GPS and ECG.", "/products/fitbit-charge-6.jpg", 159, nil},
		{"Huawei Watch GT 4", "Elegant smartwatch with two-week battery life.", "/products/huawei-watch-gt-4.jpg", 249, nil},
		{"Amazfit GTR 4", "Feature-rich smartwatch with dual-band GPS.", "/products/watch.png", 199, nil},
		{"OnePlus Watch 2", "Dual-engine Wear OS watch with 100-hour battery.", "/products/oneplus-watch-2.jpg", 299, nil},
	}},
	{"Monitors", []variant{{"24-inch", 1}, {"27-inch", 1.3}, {"32-inch", 1.6}}, []family{
		{"Samsung Odyssey OLED G8", "QD-OLED gaming monitor with 240Hz refresh.", "/products/g7.png", 999, []string{"Gaming"}},
		{"Samsung ViewFinity S8", "4K UHD monitor for creators with HDR10.", "/products/samsung-viewfinity-s8.jpg", 399, nil},
		{"LG UltraGear GP850", "Nano IPS 180Hz gaming monitor with 1ms response.", "/products/samsung-viewfinity-s8.jpg", 349, []string{"Gaming"}},
		{"LG UltraFine 4K", "Color-accurate 4K monitor with USB-C power delivery.", "/products/lg-ultrafine-4k.jpg", 549, nil},
		{"Dell UltraSharp U27", "Professional IPS Black monitor with USB-C hub.", "/products/dell-s-series-curved.jpg", 579, nil},
		{"Dell S-Series Curved", "Curved everyday monitor with built-in speakers.", "/products/dell-s-series-curved.jpg", 249, nil},
		{"ASUS ROG Swift PG", "Esports monitor with 360Hz refresh and G-Sync.", "/products/g7.png", 799, []string{"Gaming"}},
		{"ASUS ProArt PA", "Calibrated monitor for photo and video editing.", "/products/benq-pd-designer.jpg", 499, nil},
		{"BenQ Mobiuz EX", "HDRi gaming monitor with treVolo speakers.", "/products/benq-mobiuz-ex.jpg", 399, []string{"Gaming"}},
		{"BenQ PD Designer", "AQCOLOR monitor tuned for designers.", "/products/benq-pd-designer.jpg", 449, nil},
		{"Acer Nitro XV", "Budget-friendly high-refresh gaming monitor.", "/products/acer-nitro-xv.jpg", 199, []string{"Gaming"}},
		{"Apple Studio Display", "5K Retina display with 12MP Center Stage camera.", "/products/apple-studio-display.jpg", 1599, nil},
	}},
	{"Cameras", []variant{{"Body Only", 1}, {"with 18-55mm Kit", 1.2}, {"with 24-105mm Kit", 1.5}}, []family{
		{"Sony Alpha a7 IV", "33MP full-frame hybrid camera with 4K60 video.", "/products/sony-alpha-a7-iv.jpg", 2499, nil},
		{"Sony Alpha a6700", "APS-C mirrorless with AI subject recognition.", "/products/sony-alpha-a6700.jpg", 1399, nil},
		{"Canon EOS R6 Mark II", "Fast full-frame mirrorless for action and events.", "/products/canon-eos-r6-mark-ii.jpg", 2299, nil},
		{"Canon EOS R50", "Compact creator camera with easy vlogging features.", "/products/canon-eos-r50.jpg", 679, nil},
		{"Nikon Z6 III", "Partially stacked sensor full-frame for stills and video.", "/products/nikon-z6-iii.jpg", 2499, nil},
		{"Nikon Zf", "Retro-styled full-frame camera with modern autofocus.", "/products/nikon-zf.jpg", 1999, nil},
		{"Fujifilm X-T5", "40MP APS-C camera with classic dial controls.", "/products/fujifilm-x-t5.jpg", 1699, nil},
		{"Fujifilm X-S20", "Lightweight hybrid with 6.2K video and big battery.", "/products/fujifilm-x-s20.jpg", 1299, nil},
		{"Panasonic Lumix S5 II", "Phase-detect full-frame camera for videographers.", "/products/panasonic-lumix-s5-ii.jpg", 1999, nil},
		{"OM System OM-5", "Weather-sealed micro four-thirds travel camera.", "/products/om-system-om-5.jpg", 1199, nil},
	}},
	{"Gaming", []variant{{"Standard", 1}, {"Bundle", 1.25}, {"Limited Edition", 1.4}}, []family{
		{"PlayStation 5 Slim", "Slimmer PS5 with 1TB SSD and lightning-fast loads.", "/products/playstation-5-slim.jpg", 499, nil},
		{"PlayStation 5 Pro", "Enhanced GPU and PSSR upscaling for 4K gaming.", "/products/playstation-5-pro.jpg", 699, nil},
		{"PlayStation Portal", "Remote player for streaming PS5 games over Wi-Fi.", "/products/playstation-portal.jpg", 199, nil},
		{"Xbox Series S", "Compact all-digital next-gen console.", "/products/xbox-series-s.jpg", 299, nil},
		{"Xbox Elite Controller Series 2", "Pro controller with adjustable-tension thumbsticks.", "/products/xbox-elite-controller-series-2.jpg", 179, []string{"Accessories"}},
		{"Nintendo Switch Lite", "Handheld-only Switch in a compact design.", "/products/nintendo-switch-lite.jpg", 199, nil},
		{"Steam Deck OLED", "Handheld PC with HDR OLED screen and Steam library.", "/products/steam-deck-oled.jpg", 549, nil},
		{"ASUS ROG Ally X", "Windows handheld with Ryzen Z1 Extreme.", "/products/asus-rog-ally-x.jpg", 799, nil},
		{"Lenovo Legion Go", "Handheld with detachable controllers and 8.8-inch screen.", "/products/lenovo-legion-go.jpg", 699, nil},
		{"Meta Quest 3", "Mixed-reality headset with full-color passthrough.", "/products/meta-quest-3.jpg", 499, []string{"Wearables"}},
		{"Razer BlackWidow V4", "Mechanical gaming keyboard with Razer switches.", "/products/razer-blackwidow-v4.jpg", 169, []string{"Accessories"}},
		{"Logitech G Pro X Superlight 2", "Ultra-light wireless esports mouse.", "/products/logitech-g-pro-x-superlight-2.jpg", 159, []string{"Accessories"}},
	}},
	{"Accessories", []variant{{"Black", 1}, {"White", 1}, {"Graphite", 1.03}}, []family{
		{"Logitech MX Keys S", "Low-profile wireless keyboard with smart backlighting.", "/products/logitech-mx-keys-s.jpg", 109, nil},
		{"Logitech MX Anywhere 3S", "Compact performance mouse that tracks on glass.", "/products/logitech-mx-anywhere-3s.jpg", 79, nil},
		{"Apple Magic Keyboard", "Wireless keyboard with Touch ID.", "/products/apple-magic-keyboard.jpg", 149, nil},
		{"Apple Magic Mouse", "Multi-touch wireless mouse with rechargeable battery.", "/products/apple-magic-mouse.jpg", 79, nil},
		{"Anker 737 Power Bank", "24000mAh power bank with 140W output.", "/products/anker-737-power-bank.jpg", 149, nil},
		{"Anker Nano Charger 65W", "Compact GaN charger for laptops and phones.", "/products/anker-nano-charger-65w.jpg", 49, nil},
		{"Belkin 3-in-1 MagSafe Stand", "Charge iPhone, Watch, and AirPods at once.", "/products/belkin-3-in-1-magsafe-stand.jpg", 129, nil},
		{"Samsung 45W Super Fast Charger", "USB-C PD charger for Galaxy devices.", "/products/samsung-45w-super-fast-charger.jpg", 49, nil},
		{"Apple Pencil Pro", "Precision stylus with squeeze and barrel roll.", "/products/apple-pencil-pro.jpg", 129, []string{"Tablets"}},
		{"Satechi USB-C Hub", "7-in-1 aluminium hub with HDMI and card reader.", "/products/satechi-usb-c-hub.jpg", 69, nil},
		{"Spigen Tough Armor Case", "Dual-layer protective phone case with kickstand.", "/products/spigen-tough-armor-case.jpg", 29, []string{"Phones"}},
		{"Peak Design Everyday Sling", "Weatherproof sling bag for cameras and tech.", "/products/peak-design-everyday-sling.jpg", 99, []string{"Cameras"}},
		{"Elgato Stream Deck MK.2", "15 customizable LCD keys for streamers.", "/products/elgato-stream-deck-mk-2.jpg", 149, []string{"Gaming"}},
		{"Razer Kiyo Pro Webcam", "1080p60 webcam with adaptive light sensor.", "/products/razer-kiyo-pro-webcam.jpg", 129, nil},
	}},
	{"Smart Home", []variant{{"Single", 1}, {"2-Pack", 1.8}, {"4-Pack", 3.2}}, []family{
		{"Amazon Echo Dot 5th Gen", "Compact smart speaker with Alexa.", "/products/amazon-echo-show-8.jpg", 49, []string{"Audio"}},
		{"Amazon Echo Show 8", "Smart display for video calls and recipes.", "/products/amazon-echo-show-8.jpg", 149, nil},
		{"Google Nest Hub", "Smart display with sleep sensing.", "/products/google-nest-hub.jpg", 99, nil},
		{"Google Nest Mini", "Small smart speaker with Google Assistant.", "/products/google-nest-mini.jpg", 49, []string{"Audio"}},
		{"Google Nest Thermostat", "Energy-saving smart thermostat.", "/products/google-nest-thermostat.jpg", 129, nil},
		{"Philips Hue White & Color Bulb", "Millions of colors controlled from your phone.", "/products/philips-hue-white-color-bulb.jpg", 49, nil},
		{"Philips Hue Lightstrip Plus", "Flexible LED strip for ambient lighting.", "/products/philips-hue-lightstrip-plus.jpg", 79, nil},
		{"Ring Video Doorbell 4", "1080p doorbell with pre-roll video previews.", "/products/ring-video-doorbell-4.jpg", 199, []string{"Cameras"}},
		{"Arlo Pro 5S Security Camera", "2K wire-free outdoor camera with spotlight.", "/products/arlo-pro-5s-security-camera.jpg", 249, []string{"Cameras"}},
		{"TP-Link Tapo Smart Plug", "Wi-Fi plug with energy monitoring.", "/products/tp-link-tapo-smart-plug.jpg", 15, nil},
		{"August Wi-Fi Smart Lock", "Retrofit smart lock with auto-unlock.", "/products/august-wi-fi-smart-lock.jpg", 229, nil},
		{"iRobot Roomba j7+", "Self-emptying robot vacuum that avoids obstacles.", "/products/irobot-roomba-j7.jpg", 599, nil},
	}},
	{"Storage", []variant{{"500GB", 1}, {"1TB", 1.6}, {"2TB", 2.6}, {"4TB", 4.5}}, []family{
		{"Samsung 990 Pro NVMe SSD", "PCIe 4.0 SSD with up to 7450MB/s reads.", "/products/samsung-990-pro-nvme-ssd.jpg", 79, []string{"Components"}},
		{"Samsung T7 Shield Portable SSD", "Rugged USB-C portable SSD.", "/products/samsung-t7-shield-portable-ssd.jpg", 69, nil},
		{"WD Black SN850X", "Gaming NVMe SSD with heatsink option.", "/products/wd-black-sn850x.jpg", 69, []string{"Components", "Gaming"}},
		{"WD My Passport", "Slim portable hard drive with password protection.", "/products/wd-my-passport.jpg", 49, nil},
		{"Crucial P5 Plus", "Fast PCIe Gen4 SSD for PCs and PS5.", "/products/crucial-p5-plus.jpg", 59, []string{"Components"}},
		{"Crucial X9 Pro Portable SSD", "Pocket-size SSD with 1050MB/s speeds.", "/products/crucial-x9-pro-portable-ssd.jpg", 69, nil},
		{"Seagate FireCuda 530", "High-endurance NVMe SSD for creators.", "/products/wd-black-sn850x.jpg", 79, []string{"Components"}},
		{"SanDisk Extreme Portable SSD", "IP65 drop-resistant portable SSD.", "/products/sandisk-extreme-portable-ssd.jpg", 69, nil},
		{"Kingston Fury Renegade", "Gen4 NVMe SSD tuned for gaming.", "/products/kingston-fury-renegade.jpg", 59, []string{"Components"}},
		{"Seagate Expansion Desktop HDD", "Plug-and-play desktop backup drive.", "/products/seagate-expansion-desktop-hdd.jpg", 39, nil},
	}},
	{"Networking", []variant{{"Standard", 1}, {"Pro", 1.4}, {"Mesh 3-Pack", 2.6}}, []family{
		{"TP-Link Archer AX73 Router", "Wi-Fi 6 router with 6 antennas.", "/products/tp-link-archer-ax73-router.jpg", 129, nil},
		{"TP-Link Deco XE75", "Wi-Fi 6E whole-home mesh system.", "/products/tp-link-deco-xe75.jpg", 199, []string{"Smart Home"}},
		{"ASUS RT-AX86U Pro", "Gaming router with mobile game mode.", "/products/asus-rt-ax86u-pro.jpg", 249, []string{"Gaming"}},
		{"ASUS ZenWiFi XT9", "Tri-band mesh Wi-Fi 6 system.", "/products/asus-zenwifi-xt9.jpg", 229, nil},
		{"Netgear Nighthawk RAXE300", "Tri-band Wi-Fi 6E router.", "/products/netgear-nighthawk-raxe300.jpg", 299, nil},
		{"Netgear Orbi 970", "Wi-Fi 7 mesh system for gigabit homes.", "/products/netgear-orbi-970.jpg", 699, nil},
		{"Eero Pro 6E", "Simple-setup mesh router with Thread support.", "/products/eero-pro-6e.jpg", 299, []string{"Smart Home"}},
		{"Google Nest Wifi Pro", "Wi-Fi 6E mesh with smart home hub.", "/products/google-nest-wifi-pro.jpg", 199, []string{"Smart Home"}},
		{"Ubiquiti UniFi Dream Router", "All-in-one prosumer router and controller.", "/products/ubiquiti-unifi-dream-router.jpg", 199, nil},
		{"Linksys Velop Pro 7", "Wi-Fi 7 mesh with cognitive security.", "/products/tp-link-deco-xe75.jpg", 349, nil},
	}},
	{"Components", []variant{{"Standard", 1}, {"OC Edition", 1.08}, {"White Edition", 1.1}}, []family{
		{"NVIDIA GeForce RTX 4090", "Flagship GPU for 4K ray-traced gaming.", "/products/nvidia-geforce-rtx-4090.jpg", 1599, []string{"Gaming"}},
		{"NVIDIA GeForce RTX 4070 Super", "Efficient 1440p GPU with DLSS 3.", "/products/nvidia-geforce-rtx-4070-super.jpg", 599, []string{"Gaming"}},
		{"NVIDIA GeForce RTX 4060", "Budget 1080p ray-tracing GPU.", "/products/nvidia-geforce-rtx-4060.jpg", 299, []string{"Gaming"}},
		{"AMD Radeon RX 7900 XTX", "High-end GPU with 24GB VRAM.", "/products/amd-radeon-rx-7900-xtx.jpg", 949, []string{"Gaming"}},
		{"AMD Radeon RX 7800 XT", "Great-value 1440p graphics card.", "/products/amd-radeon-rx-7900-xtx.jpg", 499, []string{"Gaming"}},
		{"AMD Ryzen 7 7800X3D", "Gaming CPU with 3D V-Cache.", "/products/amd-ryzen-7-7800x3d.jpg", 449, []string{"Gaming"}},
		{"AMD Ryzen 9 7950X", "16-core CPU for creators.", "/products/amd-ryzen-9-7950x.jpg", 599, nil},
		{"Intel Core i9-14900K", "24-core desktop processor up to 6GHz.", "/products/intel-core-i9-14900k.jpg", 549, nil},
		{"Intel Core i5-14600K", "Unlocked mid-range CPU with great gaming performance.", "/products/intel-core-i5-14600k.jpg", 319, []string{"Gaming"}},
		{"Corsair Vengeance DDR5 32GB", "High-speed DDR5 memory kit.", "/products/corsair-vengeance-ddr5-32gb.jpg", 109, nil},
		{"Corsair RM850x PSU", "850W fully modular 80+ Gold power supply.", "/products/corsair-rm850x-psu.jpg", 139, nil},
		{"NZXT Kraken 360 AIO", "360mm liquid cooler with LCD display.", "/products/nzxt-kraken-360-aio.jpg", 279, nil},
	}},
}

// generatedProducts expands catalog into seedProduct entries.
func generatedProducts() []seedProduct {
	var out []seedProduct
	for _, g := range catalog {
		for _, f := range g.Families {
			for _, v := range g.Variants {
				out = append(out, seedProduct{
					Name:        fmt.Sprintf("%s (%s)", f.Name, v.Label),
					Description: fmt.Sprintf("%s Configuration: %s.", f.Desc, v.Label),
					Price:       math.Max(9, math.Round(f.Base*v.Mult/10)*10-1), // $X9 pricing
					Picture:     f.Pic,
					Categories:  append([]string{g.Category}, f.Extra...),
				})
			}
		}
	}
	return out
}
