package example


import (
	"github.com/slawek87/GOstorageClient/client"
	"os"
	"fmt"
)

func Example() {
	storageClient := client.ClientInterface(client.Client{})
	getFile, err := os.Open("./example/data/mazda.pdf")

	if err != nil {
		fmt.Println(err)
	}

	storageClient.UploadFile(getFile)
}
