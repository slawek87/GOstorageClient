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

}
