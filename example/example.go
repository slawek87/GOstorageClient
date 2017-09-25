package example

import (
	"github.com/slawek87/GOstorageClient/client"
	"os"
)

func Example() {
	storageClient := client.ClientInterface(client.Client{})

	path := "./example/data/mazda.pdf"
	file, _ := os.Open(path)
	storageClient.UploadFile(file)

	storageClient.DeleteFile("5_4_3_2_1_mazda.pdf")

	path = "./example/data/volvo.pdf"
	file, _ = os.Open(path)

	storageClient.OverwriteFile(file, "8_7_6_5_4_3_2_1_1024x768-3.jpeg")
}
