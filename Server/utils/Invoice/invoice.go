package invoice

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	//cloudinary
	"github.com/abhijitsh/go_restapi/models"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

const url = "https://pdfgen.app/api/generate?templateId=b7fc783"

//IMP Using PDFGen API to generate Invoice PDF and then uploading to Cloudinary then storing public url in DB

//TODO : Delete invoice.pdf after saving url to DB
func GetInvoice(payload models.PDFGenData) string {

	apiKey := os.Getenv("PDFGEN_API_KEY")
	JSONpayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling lineItems:", err)
		 log.Fatal("Error marshalling lineItems:", err)
	} 
	pdfPayload := fmt.Sprintf(`{"data" : %s}`, string(JSONpayload))
	//fmt.Println(pdfPayload)
	req, err := http.NewRequest("POST", url, strings.NewReader(pdfPayload))
	if err != nil {
		log.Fatal("error creating request: ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error executing request: ", err)
	}
	defer resp.Body.Close()
	file, err := os.Create("invoice.pdf")
	if err != nil {
		log.Fatal("error creating file: ", err)
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal("error saving file: ", err)
	}

	fmt.Println("PDF downloaded successfully as invoice.pdf")
	return cloudinaryUpload("invoice.pdf")
}

// Cloudinary Upload => Integration Code
func cloudinaryUpload(filePath string) string {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================
	cld, _ := cloudinary.NewFromParams("dcbrlaot1", "731382434158625", "c00bwaJiOBESlEV91yYUWIaDMCk")
	cld.Config.URL.Secure = true
	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(context.Background(), filePath, uploader.UploadParams{})
	if err != nil {
		fmt.Println("error")
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
	return resp.SecureURL
}
