package santricity

import (
	"fmt"
	"os"
	"testing"
)

func TestVolume(t *testing.T) {
	fmt.Println("Test: Volume")

	c, err := NewClient(os.Getenv("APIUSER"), os.Getenv("APIPW"), "127.0.0.1:8443", true, true)
	if err != nil {
		t.Error(err)
	}

	data, err := c.GetVolumes("1")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data[0].Name)
	fmt.Println(data[0].ID)
	fmt.Println(data[0].Capacity)
	fmt.Println(data[0].TotalSizeInBytes)

	data2, err := c.GetVolume("1", data[0].ID)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data2.Name)

}

func TestVolumeStats(t *testing.T) {
	fmt.Println("Test: Volume")

	c, err := NewClient(os.Getenv("APIUSER"), os.Getenv("APIPW"), "127.0.0.1:8443", true, true)
	if err != nil {
		t.Error(err)
	}

	data, err := c.GetVolumeStats("1")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(data[0].ReadOps)

}
