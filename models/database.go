package models

type allItems []Item

var Produce = allItems{
	{
		Name:        "Lettuce",
		ProduceCode: "A12T-4GH7-QPL9-3N4M",
		UnitPrice:   3.46,
	},
	{
		Name:        "Peach",
		ProduceCode: "E5T6-9UI3-TH15-QR88",
		UnitPrice:   2.99,
	},
	{
		Name:        "Green Pepper",
		ProduceCode: "YRT6-72AS-K736-L4AR",
		UnitPrice:   0.79,
	},
	{
		Name:        "Gala Apple",
		ProduceCode: "TQ4C-VV6T-75ZX-1RMR",
		UnitPrice:   3.59,
	},
}

func GetAllProduce() []Item {
	return Produce
}

func GetOneItem(pc string) Item {
	empty := Item{}
	for _, item := range Produce {
		if item.ProduceCode == pc {
			return item
		}
	}
	return empty
}

func AddProduce(item Item) {
	Produce = append(Produce, item)
}

func DeleteProduce(pc string) Item {
	empty := Item{}
	for i, item := range Produce {
		if item.ProduceCode == pc {
			Produce = append(Produce[:i], Produce[i+1:]...)
			return item
		}
	}
	return empty
}

func Equal(a, b []Item) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
