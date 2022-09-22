package Basic

import (
	"encoding/csv"
	"os"
)

type Item struct {
	SKU  string
	Name string
}

func WriteItems(filename string, items []Item) error {
	// create a file, if no errot, then close the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	row := []string{"SKU", "Name"}
	wtr := csv.NewWriter(file)
	defer wtr.Flush()
	if err := wtr.Write(row); err != nil {
		return err
	}

	for _, item := range items {
		row[0] = item.SKU
		row[1] = item.Name
		if err := wtr.Write(row); err != nil {
			return err
		}
	}
	return wtr.Error()
}
