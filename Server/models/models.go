package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//primitive is basically a package that contains different dataType used in mongodb , like bson, => bson.M and bson.D => M, D for unordered and ordered representaion of key value pairs basically map[string]interface{}(in golang), as key toh string hi hogi we are sure but value can be anything string, slice, etc.... (in MongoDB), primitive.ObjectID, primitive.A, primitive.E, primitive.D, primitive.M, primitive.Binary, primitive.DateTime, primitive.Decimal128, primitive.Regex, primitive.Timestamp etc....
type User struct {
	ID        primitive.ObjectID `json:"_id"  bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Phone     int                `json:"phone"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password,omitempty"`
	URL       string             `json:"url" `
	Addresses []UserAddress      `json:"addresses"`
}

//This we not saving to DB
type UserAddress struct {
	Address  string `json:"address"`
	Landmark string `json:"landmark"`
	PIN      int    `json:"pin"` // 6 digit
}

//omitempty , omits if ID is empty like when we sending or encoding this to bson while sending to mongodb to store it in the database, if ID is empty then it will be omitted, and by default, the user if its created, will have default null values in golang so if we dont write omitempty then yahan se _id:(x0000000) leke jaega mongodb ke paas, toh mongodb dekhega ki pehle se _id hai toh kyun hi assign karna and it will be stored in the database as id = null (in objectID thats (x00000000)), so to avoid that we use omitempty, nahi toh pehli baar toh null objectId ho jaega store lekin dusri baar save karega toh firse jab null objectId store karne gya toh pehle se bhi ek null id hai ab waapas se yahi toh duplicate hai ye, so mongodb will throw error, vishwas nahi hai toh omitempty htake aur post req maarke dekho,

//ID ka type string bhi rakh sakte hain lekin mongodb me objectId me hoga aur wahi store karna chahte ho toh fir string me convert karna padega, so better to use primitive.ObjectID

type Location struct {
	Email    string `json:"email"`
	FLatNo   string `json:"flatno"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
	LandMark string `json:"landmark"`
}

type Food struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type FoodQuantity struct {
	FoodID   int `json:"foodId"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type Cart struct {
	Email      string         `json:"email"`
	Foods      []FoodQuantity `json:"foods"`
	TotalPrice int            `json:"totalPrice"`
}

type Order struct {
	Email           string         `json:"email"`
	Foods           []FoodQuantity `json:"foods"`
	TotalPrice      int            `json:"totalPrice"`
	OrderID         string         `json:"orderid"`
	DeliveryAddress UserAddress    `json:"address"`
	Date            string         `json:"date"`
	Time            string         `json:"time"`
	InvoiceURL      string         `json:"invoiceurl"`
}

type LineItems struct {
	Item     string `json:"item"`
	Price    int    `json:"price"`
	Total    int    `json:"total"`
	Quantity int    `json:"quantity"`
}

//pdfGen
type PDFGenData struct {
	Date          string      `json:"date"`
	Time          string      `json:"time"`
	Email         string      `json:"email"`
	Customer      UserAddress `json:"customer"`
	LineItems     []LineItems `json:"lineItems"`
	InvoiceTotal  int         `json:"invoiceTotal"`
	InvoiceNumber string      `json:"invoiceNumber"`
}
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
