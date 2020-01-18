package santricity

import (
	"fmt"
	"os"
	"testing"
)

func TestGetPools(t *testing.T) {
	fmt.Println("Test: GetPools")

	c, err := NewClient(os.Getenv("APIUSER"), os.Getenv("APIPW"), "127.0.0.1:8443", true, true)
	if err != nil {
		t.Error(err)
	}

	data, err := c.GetPools("1")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data[0].ID)

}
