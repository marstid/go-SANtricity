package santricity

import (
	"fmt"
	"os"
	"testing"
)

func TestCluster(t *testing.T) {
	fmt.Println("Test: Get Cluster Info")

	c, err := NewClient(os.Getenv("APIUSER"), os.Getenv("APIPW"), "127.0.0.1:8443", true, true)
	if err != nil {
		t.Error(err)
	}

	data, err := c.GetSystems()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data[0].ID)

	data2, err := c.GetSystem(data[0].ID)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data2.ID)

}
