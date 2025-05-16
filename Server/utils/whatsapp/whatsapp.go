package whatsapp

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const url = "https://graph.facebook.com/v19.0/328629233675098/messages"

// func SendTemplateMessage(phone int) {
// 	fmt.Println("Sending template message to", phone)
// 	token := os.Getenv("WA_TOKEN")
// 	phn := strconv.Itoa(phone) // Convert the integer phone number to a string
//  phn = fmt.Sprintf("91%s", phn) // Add the country code to the phone number
// 	method := "POST"

// 	payload := strings.NewReader(fmt.Sprintf(`{
// 	  "messaging_product": "whatsapp",
// 	  "to": "%s",
// 	  "type": "template",
// 	  "template": {
// 		"name": "hello_world",
// 		"language": {
// 		  "code": "en_US"
// 		}
// 	  }
// 	}`, phn))
	

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}
	
// 	req.Header.Add("Content-Type", "application/json")
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token)) // Replace with actual token
// 	res, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}

// 	fmt.Println("Response:", string(body))
// }

//To send invitation template to a user
func SendInvitation(phone int) {
	fmt.Println("Sending invitation to", phone)
	token := os.Getenv("WA_TOKEN")
	phn := strconv.Itoa(phone) // Convert the integer phone number to a string
	phn = fmt.Sprintf("91%s", phn) // Add the country code to the phone number
	headerImg := os.Getenv("WA_INVITE_HEADER_IMG")
	method := "POST"

	//NewReader accepts string as an argument and returns a reader.....(JSON format, abhi ke liye).
	payload := strings.NewReader(fmt.Sprintf(`{
    "messaging_product": "whatsapp",
    "to": "%s",
    "type": "template",
    "template": {
        "name": "invitation",
        "language": {
            "code": "en"
        },
        "components": [
            {
                "type": "header",
                "parameters": [
                    {
                        "type": "image",
                        "image": {
                            "link": "%s"
                        }
                    }
                ]
            },
            {
                "type": "button",
                "index": "0",
                "sub_type": "url",
                "parameters": [
                    {
                        "type": "text",
                        "text": "/"
                    }
                ]
            }
        ]
    }
}`, phn, headerImg))
	

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}

//To send order_placed template to a user
func SendOrderPlaced(phone int, invoice string, name string, orderId string) {
	fmt.Println("Sending order_placed to", phone)
	token := os.Getenv("WA_TOKEN")
	phn := strconv.Itoa(phone) // Convert the integer phone number to a string
	phn = fmt.Sprintf("91%s", phn) // Add the country code to the phone number
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{
    "messaging_product": "whatsapp",
    "to": "%s",
    "type": "template",
    "template": {
        "name": "order_placed",
        "language": {
            "code": "en"
        },
        "components": [
            {
                "type": "header",
                "parameters": [
                    {
                        "type": "document",
                        "document": {
                            "link": "%s"
                        }
                    }
                ]
            },
            {
                "type": "body",
                "parameters": [
                    {
                        "type": "text",
                        "text": "%s"
                    },
                    {
                        "type": "text",
                        "text": "%s"
                    }
                ]
            }
        ]
    }
}`, phn, invoice, name, orderId))
	

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)

	//IMP:
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
