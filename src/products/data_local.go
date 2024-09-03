package products

var LocalData = make([]ProductEntity, 0)

var listInit = []ProductEntity{
	{
		ID:                    "1",
		ProductName:           "Tank",
		ProductCategory:       "Alat tempur",
		ProductFreshness:      "Bekas",
		ImageOfProduct:        "./tank.jpg",
		AdditionalDescription: "Alat bantu pertempuran",
		ProductPrice:          50_000_000_000,
	},
	{
		ID:                    "2",
		ProductName:           "Gas Air Mata",
		ProductCategory:       "Alat bantu demo",
		ProductFreshness:      "Bekas",
		ImageOfProduct:        "./gas-air-mata.jpg",
		AdditionalDescription: "Alat penyiksa mata",
		ProductPrice:          500_000,
	},
	{
		ID:                    "3",
		ProductName:           "Rubicorn",
		ProductCategory:       "Kendaraan",
		ProductFreshness:      "Baru",
		ImageOfProduct:        "./mobil-orang.jpg",
		AdditionalDescription: "kendaraan roda 4",
		ProductPrice:          1_800_000_000,
	},
	{
		ID:                    "4",
		ProductName:           "Smartphone",
		ProductCategory:       "Elektronik",
		ProductFreshness:      "Baru",
		ImageOfProduct:        "./hp.jpg",
		AdditionalDescription: "Android Phone",
		ProductPrice:          10_000_000,
	},
	{
		ID:                    "5",
		ProductName:           "Jaket",
		ProductCategory:       "Pakaian",
		ProductFreshness:      "Baru",
		ImageOfProduct:        "./jaket.jpg",
		AdditionalDescription: "Jaket kulit",
		ProductPrice:          1_000_000,
	},
}

func init() {
	LocalData = append(LocalData, listInit...)
}
