package db

import "github.com/PSS2134/go_restapi/models"

// Define a mock database of menu items
var MenuItems = []models.Food{
	{
		ID:          1,
		Title:       "Masala Chai",
		Description: "A fragrant blend of black tea, milk, and a medley of spices like cardamom, cinnamon, and ginger. Warm, comforting, and perfect for any time of the day.",
		Price:       50,
	},
	{
		ID:          2,
		Title:       "Butter Chicken",
		Description: "Tender chicken cooked in a rich and creamy tomato-based sauce, served with fragrant basmati rice. A beloved classic that will satisfy your taste buds.",
		Price:       259,
	},
	{
		ID:          3,
		Title:       "Tandoori Chicken",
		Description: "Succulent chicken marinated in a blend of yogurt and spices, roasted in a traditional clay oven. Served with mint chutney and naan bread for an authentic experience.",
		Price:       280,
	},
	{
		ID:          4,
		Title:       "Chicken Tikka",
		Description: "Grilled chicken cooked in a luscious tomato and cream sauce, infused with aromatic spices. Pair it with fluffy naan bread for a delightful meal.",
		Price:       269,
	},
	{
		ID:          5,
		Title:       "Manchurian",
		Description: "Crispy cauliflower tossed in a savory sauce with garlic, ginger, and a hint of soy, served with steamed rice. A fusion dish that brings together Chinese and Indian flavors.",
		Price:       179,
	},
	{
		ID:          6,
		Title:       "Nutella Milkshake",
		Description: "A heavenly blend of creamy milk, indulgent Nutella, and a touch of sweetness. A chocolate lover's dream in a glass.",
		Price:       120,
	},
	{
		ID:          7,
		Title:       "Mango Lassi",
		Description: "A refreshing yogurt-based drink infused with sweet, ripe mangoes. A delightful and cooling beverage perfect for hot summer days.",
		Price:       80,
	},
	{
		ID:          8,
		Title:       "Chicken Biryani",
		Description: "Fragrant basmati rice cooked with tender chicken, aromatic spices, and caramelized onions. Served with raita, this traditional dish is a flavor explosion.",
		Price:       369,
	},
	{
		ID:          9,
		Title:       "Hakka Noodles",
		Description: "Stir-fried noodles with a colorful medley of vegetables, flavored with soy sauce and spices. A popular Indo-Chinese dish that will satisfy your noodle cravings.",
		Price:       160,
	},
	{
		ID:          10,
		Title:       "Chole Bhature",
		Description: "Soft and fluffy deep-fried bread paired with spicy chickpea curry. This North Indian specialty is a must-try for a hearty and fulfilling meal.",
		Price:       140,
	},
	{
		ID:          11,
		Title:       "Pakode",
		Description: "Crispy and flavorful fritters made with a variety of vegetables, chickpea flour, and aromatic spices. Served piping hot, these deep-fried delights are perfect as a snack or appetizer, providing a satisfying crunch with every bite.",
		Price:       65,
	},
	{
		ID:          12,
		Title:       "Vada Pav",
		Description: "Vada Pav features a spiced potato fritter (vada) sandwiched between soft pav buns. Topped with tangy chutneys and served with fried green chili, this iconic dish offers a burst of flavors and textures in every mouthful.",
		Price:       40,
	},
}

