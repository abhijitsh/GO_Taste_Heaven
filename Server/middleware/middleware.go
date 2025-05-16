package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	db "github.com/PSS2134/go_restapi/DB"
	"github.com/PSS2134/go_restapi/models"
	invoice "github.com/PSS2134/go_restapi/utils/Invoice"
	"github.com/PSS2134/go_restapi/utils/whatsapp"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var cartCollection *mongo.Collection
var orderCollection *mongo.Collection

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("The following error has occured in loading env file", err) //more better than panic, jyaada acche se error deta
	}

}

func createDBInstance() {
	connectionURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	userCollName := os.Getenv("USER_COLLNAME")
	cartCollName := os.Getenv("CART_COLLNAME")
	orderCollName := os.Getenv("ORDER_COLLNAME")
	//each time we talk to database we need to provide context, agar konsa context dena hai ye nahi pta toh context.TODO() dedo, else api me read write ke waqt context.Background() dena accha hai

	//options me we are adding our connection string :)
	clientOptions := options.Client().ApplyURI(connectionURI)

	//lets connect to DB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error initiating connection to DB", err)
	}

	//re-assigning to above error variable
	err = client.Ping(context.TODO(), nil) //Ping karke we are checking ki sab sahi toh chal rha hai na?

	if err != nil {
		log.Fatal("Error connecting to DB", err)
	}

	fmt.Println("Connected to MongoDB Successfully :) ")

	userCollection = client.Database(dbName).Collection(userCollName)
	cartCollection = client.Database(dbName).Collection(cartCollName)
	orderCollection = client.Database(dbName).Collection(orderCollName)
	fmt.Println("Collection instance created")
}

func init() {
	loadEnv()
	createDBInstance()
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "GET")

	payload := getallusers()
	json.NewEncoder(w).Encode(&payload)

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "POST")
	var user models.User
	//decode json into struct user
	json.NewDecoder(r.Body).Decode(&user)
	//lets save
	insertOneUser(&user) // normal bhi pass kar sakte ho
	//again send the user by encoding into json
	json.NewEncoder(w).Encode(&user)

}

// API SIGNUP
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside SignUp")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "POST")
	defer r.Body.Close()
	// check if user exists already
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	//id := user.ID
	if user.Email == "" || user.Password == "" || user.Name == "" || user.Phone == 0 {
		res := models.Response{Status: 400, Message: "All fields are required", Data: nil}
		//encode me & ki jarurat nahi lekin decode me & ki jarurat hai as sahi cheez decode honi chahiye copy nahi.....:)
		json.NewEncoder(w).Encode(&res)
		return
	}
	var userExists models.User
	var sendUser models.User
	fmt.Printf("user : %+v", user)
	err := userCollection.FindOne(context.Background(), primitive.M{"email": user.Email}).Decode(&userExists)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No user found")
			insertRes, err := userCollection.InsertOne(context.Background(), user)
			if err != nil {
				log.Fatal("Some error inserting user", err)
			}

			user.ID = insertRes.InsertedID.(primitive.ObjectID)
			sendUser = user
			sendUser.Password = ""
			sendUser.Addresses = nil
			json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "User Inserted Successfully", Data: sendUser})
			//IMP : Whatsapp
			whatsapp.SendInvitation(user.Phone)
			r.Body.Close()
			return
		} else {
			log.Fatal("Some error quering user", err)
		}
	}
	fmt.Println("userExists: ", userExists)
	sendUser = userExists

	//IMP DONT SEND PASSWORD TO FRONTEND
	sendUser.Password = ""
	sendUser.Addresses = nil
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "User Already Exists", Data: userExists})
	
}

// API Login Controller
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "POST")

	defer r.Body.Close()
	var user, sendUser models.User

	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if user.Email == "" || user.Password == "" {
		fmt.Println("Email or Password can not be empty")
		json.NewEncoder(w).Encode(&models.Response{
			Status:  400,
			Message: "Email or Password can not be empty",
			Data:    nil,
		})
		return
	}

	var userExists models.User

	err := userCollection.FindOne(context.Background(), primitive.M{"email": user.Email}).Decode(&userExists)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("User not found, Please SignUp...")
			json.NewEncoder(w).Encode(&models.Response{
				Status:  400,
				Message: "User not found",
				Data:    nil,
			})
			return
		} else {
			log.Fatal("Some error quering user", err)
		}
	}

	if userExists.Password != user.Password {
		fmt.Println("Invalid Password")
		json.NewEncoder(w).Encode(&models.Response{
			Status:  400,
			Message: "Invalid Password",
			Data:    nil,
		})
		return
	}

	//IMP DONT SEND PASSWORD TO FRONTEND
	sendUser = userExists
	sendUser.Password = ""
	sendUser.Addresses = nil
	json.NewEncoder(w).Encode(&models.Response{
		Status:  200,
		Message: "Login Successful",
		Data:    sendUser,
	})

}

// API - POST Address
func PostAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	defer r.Body.Close()

	qparams := r.URL.Query()
	email := qparams.Get("email")
	fmt.Println("email", email)
	var address models.UserAddress
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error decoding user address", Data: err})
		log.Fatal("Error decoding user address", err)
		return
	}
	fmt.Printf("Incoming Address: %+v\n", address)

	var user models.User

	err = userCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&user)
	if err != nil && err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error finding user", Data: err})
		log.Fatal("Error finding user", err)
		return
	}

	if len(user.Addresses) > 0 {

		_, err = userCollection.UpdateOne(context.Background(), primitive.M{"email": email}, primitive.M{
			"$push": primitive.M{"addresses": address},
		})
	} else {
		fmt.Print("HI")
		user.Addresses = []models.UserAddress{address}
		fmt.Printf("User : %+v", user)
		_, err = userCollection.ReplaceOne(context.Background(), primitive.M{"email": email}, user)
	}

	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error updating address", Data: err})
		log.Fatal("Error updating address", err)
		return
	}

	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Address Saved Successfully", Data: address})
}

// //PI - POST Address
// func PostAddress(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Acess-Control-Allow-Origin", "*")
// 	w.Header().Set("Acess-Control-Allow-Methods", "POST")
// 	defer r.Body.Close()
// 	qparams := r.URL.Query()
// 	email := qparams.Get("email")

// 	_ , err := addressCollection.DeleteMany(context.Background(), primitive.M{"email": email})

// 	if err != nil {
// 		json.NewEncoder(w).Encode(&models.Response{ Status: 400, Message: "Error deleting address", Data: err})
// 		log.Fatal("Error deleting address", err)
// 		return
// 	}

//     var userAddress models.Location
// 	err = json.NewDecoder(r.Body).Decode(&userAddress)
// 	if err != nil {
// 		json.NewEncoder(w).Encode(&models.Response{ Status: 400, Message: "Error decoding user address", Data: err})
// 		log.Fatal("Error decoding user address", err)

// 	}
// 	fmt.Println("User Address: ", userAddress)
// 	userAddress.Email = email
// 	_ , err = addressCollection.InsertOne(context.Background(), userAddress)

// 	if err != nil {
// 		json.NewEncoder(w).Encode(&models.Response{ Status: 400, Message: "Error inserting address", Data: err})
// 		log.Fatal("Error inserting address", err)
// 	}
// 	json.NewEncoder(w).Encode(&models.Response{ Status: 200, Message: "Address Inserted Successfully", Data: userAddress})
// 	//fmt.Print("All success at POSTAddress")
// }

// API - GET ALL ADDRESSES
func GetAllAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "GET")
	defer r.Body.Close()

	email := r.URL.Query().Get("email")
	var user models.User
	err := userCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching user", Data: err})
		log.Fatal("Error fetching user", err)
		return
	}
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "User Addresses Fetched Successfully", Data: user.Addresses})

}

// API - GET Confirmation
func GetConfirmation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "GET")
	defer r.Body.Close()
	email := r.URL.Query().Get("email")
	//TODO: Check if email is empty

	//IMP user, address, cart?

	var user models.User
	var cart models.Cart

	err := userCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&user)

	if err != nil {

		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching user", Data: err})
		log.Fatal("Error fetching user", err)
		return
	}

	err = cartCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&cart)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching user cart", Data: err})
			log.Fatal("Error fetching user cart", err)
			return
		}

	}

	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Confirmation Data", Data: map[string]interface{}{"name": user.Name, "food": cart.Foods, "totalPrice": cart.TotalPrice}})
	fmt.Println("All success at GETConfirmation")
}

// API - GET ORDER
func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "GET")
	defer r.Body.Close()
	email := r.URL.Query().Get("email")
	orderId := r.URL.Query().Get("orderId")
	var order models.Order
	err := orderCollection.FindOne(context.Background(), primitive.M{"email": email, "orderid": orderId}).Decode(&order)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching order", Data: err})
			log.Fatal("Error fetching order", err)
			return
		}
	}
	fmt.Println("Order Fetched Successfully")
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Order Data", Data: order})
}

//API POST Order

func PostOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	defer r.Body.Close()

	email := r.URL.Query().Get("email")
	orderId := r.URL.Query().Get("orderId")

	var address models.UserAddress
	json.NewDecoder(r.Body).Decode(&address)
	var cart models.Cart
	err := cartCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&cart)
	if err != nil && err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching cart", Data: err})
		log.Fatal("Error fetching cart", err)
		return
	}

	// Check if an order with the same orderId, email, foods, total price, and delivery address exists
	var existingOrder models.Order
	err = orderCollection.FindOne(context.Background(), primitive.M{
		"orderid":         orderId,
		"email":           email,
		"foods":           cart.Foods,
		"totalprice":      cart.TotalPrice,
		"deliveryaddress": address,
	}).Decode(&existingOrder)
	if err == nil {
		// Order already exists
		json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Order already exists", Data: existingOrder})
		return
	} else if err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error checking existing order", Data: err})
		log.Fatal("Error checking existing order", err)
		return
	}

	date := time.Now().Format("02-01-2006")
	time := time.Now().Format("15:04:05")
	//fmt.Printf("Today's Date: %s and Time: %s\n", date, time)

	var order models.Order
	order.Email = email
	order.Foods = cart.Foods
	order.TotalPrice = cart.TotalPrice
	order.OrderID = orderId
	order.DeliveryAddress = address
	order.Date = date
	order.Time = time


	var argItems []models.LineItems
	for _, val := range order.Foods {
		argItems = append(argItems, models.LineItems{Item: db.MenuItems[val.FoodID-1].Title, Price: val.Price, Total: val.Price * val.Quantity, Quantity: val.Quantity})
	}
	var payload models.PDFGenData
	payload.Date = date
	payload.Time = time
	payload.Customer = address
	payload.Email = email
	payload.LineItems = argItems
	payload.InvoiceTotal =  order.TotalPrice
	payload.InvoiceNumber = orderId
	invoiceUrl := invoice.GetInvoice(payload)
	//Now update invoice url
	order.InvoiceURL = invoiceUrl
	_, err = orderCollection.InsertOne(context.Background(), order)
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error inserting order", Data: err})
		log.Fatal("Error inserting order", err)
		return
	}

	fmt.Println("Order Inserted Successfully")
	_, err = cartCollection.DeleteMany(context.Background(), primitive.M{"email": email})
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error deleting cart", Data: err})
		log.Fatal("Error deleting cart", err)
	}
	
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Order Inserted successfully", Data: nil})
	//IMP : Whatsapp
	var userWA models.User
	err = userCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&userWA)
	if(err != nil) {
		fmt.Println("Error fetching user for whatsapp", err)
		return
	}
	whatsapp.SendOrderPlaced(userWA.Phone, invoiceUrl, userWA.Name, orderId)
}

//API GET PROFILE

func GetProfile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "GET")

	email := r.URL.Query().Get("email")
	cur, err := orderCollection.Find(context.Background(), primitive.M{"email": email})
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching orders for user", Data: err})
		log.Fatal(err)
	}
	var orderSlice []models.Order
	for cur.Next(context.Background()) {

		var single models.Order
		err := cur.Decode(&single)
		if err != nil {
			json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching orders for user", Data: err})
			log.Fatal(err)
		}
		orderSlice = append(orderSlice, single)

	}

	message := fmt.Sprintf("All Orders Fetched successfully for User: %s", email)
	json.NewEncoder(w).Encode(&models.Response{
		Status:  200,
		Message: message,
		Data:    orderSlice,
	})
	fmt.Println(message)

}

//API POST PROFILE

func PostProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Origin", "*")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	defer r.Body.Close()

	email := r.URL.Query().Get("email")
	var userExists models.User
	err := userCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&userExists)
	if err != nil && err != mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching user", Data: err})
		log.Fatal("Error fetching user", err)
		return
	}
	var userReq models.User
	json.NewDecoder(r.Body).Decode(&userReq)
	fmt.Println(userReq.Phone)
	userExists.Name = userReq.Name
	userExists.Phone = userReq.Phone
	_, err = userCollection.ReplaceOne(context.Background(), primitive.M{"email": email}, userExists)
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error updating user", Data: err})
		log.Fatal("Error updating user", err)
		return
	}

	fmt.Println("User Profile Updated Successfully")
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "User Profile Updated Successfully", Data: userExists})
}

// API - PUT FOR IMAGE CHANGE
func UpdateIMG(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IMAGE CHANGE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "POST")
	defer r.Body.Close()

	email := r.URL.Query().Get("email")
	var custom struct {
		URL string `json:"url"`
	}
	json.NewDecoder(r.Body).Decode(&custom)
	fmt.Println("URL: ", custom.URL)
	result, err := userCollection.UpdateOne(context.Background(), primitive.M{"email": email}, primitive.M{"$set": primitive.M{"url": custom.URL}})

	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error updating image", Data: err})
		log.Fatal("Error updating image", err)
		return
	}

	fmt.Println("Image Updated Successfully: ", result)
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Image Updated Successfully", Data: nil})

}

// API CART
func PostCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	defer r.Body.Close()

	email := r.URL.Query().Get("email")
	action := r.URL.Query().Get("action")
	var foodQuantity models.FoodQuantity
	var userCart models.Cart
	json.NewDecoder(r.Body).Decode(&foodQuantity)
	//fmt.Println(foodQuantity)
	if foodQuantity.Quantity == 0 || foodQuantity.Price == 0 {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Invalid food quantity or price", Data: nil})
		return
	}
	err := cartCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&userCart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Insert a new cart if it doesn't exist
			_, err = cartCollection.InsertOne(context.Background(), models.Cart{
				Email:      email,
				Foods:      []models.FoodQuantity{foodQuantity},
				TotalPrice: foodQuantity.Price * foodQuantity.Quantity,
			})
			if err != nil {
				json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error inserting cart", Data: err})
				log.Fatal("Error inserting cart", err)
				return
			}
			json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Cart created successfully", Data: &models.Cart{Email: email, Foods: []models.FoodQuantity{foodQuantity}, TotalPrice: foodQuantity.Price * foodQuantity.Quantity}})
			return
		} else {
			json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching cart", Data: err})
			log.Fatal("Error fetching cart", err)
			return
		}
	}

	// Handle actions
	switch action {
	case "add":
		found := false
		for i, item := range userCart.Foods {
			if item.FoodID == foodQuantity.FoodID {
				userCart.Foods[i].Quantity += foodQuantity.Quantity
				userCart.TotalPrice += foodQuantity.Price * foodQuantity.Quantity
				found = true
				break
			}
		}
		if !found {
			userCart.Foods = append(userCart.Foods, foodQuantity)
			userCart.TotalPrice += foodQuantity.Price * foodQuantity.Quantity
		}
	case "remove":
		found := false
		for i, item := range userCart.Foods {
			if item.FoodID == foodQuantity.FoodID {
				found = true
				newQuantity := item.Quantity - foodQuantity.Quantity
				if newQuantity <= 0 {
					// Remove the item if the new quantity is zero or less
					userCart.Foods = append(userCart.Foods[:i], userCart.Foods[i+1:]...)
				} else {
					// Update the item's quantity
					userCart.Foods[i].Quantity = newQuantity
				}

				// Update the total price
				userCart.TotalPrice -= foodQuantity.Price * foodQuantity.Quantity
				if userCart.TotalPrice < 0 {
					userCart.TotalPrice = 0
				}
				break
			}
		}

		// If the item was not found, respond with an error
		if !found {
			json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Item not found in cart", Data: nil})
			return
		}

	default:
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Invalid action", Data: nil})
		return
	}

	_, err = cartCollection.UpdateOne(context.Background(), primitive.M{"email": email}, primitive.M{"$set": userCart})
	if err != nil {
		json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error updating cart", Data: err})
		log.Fatal("Error updating cart", err)
		return
	}
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Cart Updated successfully", Data: userCart})
}

//API GetCart

func GetCart(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Acess-Control-Allow-Origin", "*")
	w.Header().Set("Acess-Control-Allow-Methods", "GET")
	defer r.Body.Close()
	email := r.URL.Query().Get("email")
	var userCart models.Cart
	err := cartCollection.FindOne(context.Background(), primitive.M{"email": email}).Decode(&userCart)
	if err != nil {

		if err == mongo.ErrNoDocuments {
			json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Cart not found", Data: nil})
			return
		} else {
			json.NewEncoder(w).Encode(&models.Response{Status: 400, Message: "Error fetching cart", Data: err})
			log.Fatal("Error fetching cart", err)
			return
		}

	}
	json.NewEncoder(w).Encode(&models.Response{Status: 200, Message: "Cart fetched Successfully", Data: userCart})

}

// primitive.M matlab bson in MongoDB (key value pairs)
func getallusers() []primitive.M {

	cur, err := userCollection.Find(context.TODO(), primitive.M{})
	if err != nil {
		log.Fatal(err)
	}
	//curr => cursor => like a linked list

	var result []primitive.M
	for cur.Next(context.Background()) {
		var single primitive.M
		err := cur.Decode(&single)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, single)
	}

	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return result
}

func insertOneUser(user *models.User) {
	insertRes, err := userCollection.InsertOne(context.Background(), user)

	if err != nil {
		log.Fatal(err)
	}
	/*
		You're welcome! Here's a concise note summarizing what we discussed:

		Go Type Assertion for MongoDB InsertedID
		MongoDB returns InsertedID as interface{}: Go doesn't know the exact type at runtime, even though it's usually primitive.ObjectID.
		Type Assertion is necessary to convert interface{} to the actual type primitive.ObjectID.
		Example:
		 => objectID, ok := insertRes.InsertedID.(primitive.ObjectID)

		Reason: Go is statically typed and requires explicit conversion from interface{} to a specific type for type safety.
		Purpose: This avoids runtime type errors by ensuring that you're working with the correct type, making Go more robust.
	*/
	//id, ok := insertRes.InsertedID.(string) => here ok is false, as insertRes.InsertedID is primitive.ObjectID and we are assigning it string its like its of type objectid but since it is returned as an interface{} to golang, we need to tell golang the type, so we are telling it string but ist actually primitive.ObjectID, so thats the issue here
	id, ok := insertRes.InsertedID.(primitive.ObjectID)
	if ok {
		user.ID = id
	} else {
		fmt.Println("Error while converting ID")
	}
	fmt.Println(insertRes.InsertedID)
	fmt.Println("User Inserted successfully: ", insertRes)
}

// func Errfunc(err error, msg string, status int) {
// 	if err != nil {
// 		log.Fatal(msg, err)
// 		json.NewEncoder(w).Encode(&models.Response{ Status: 400, Message: "Error fetching user", Data: err})
// 		return

// 	}
// }
