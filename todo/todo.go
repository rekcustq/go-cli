package todo

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Item struct {
	Text 	 string
	Priority int
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	return items, nil
}

func (i *Item) SetPriority(pri int) {
	i.Priority = 2
	if i.Priority != pri {
		i.Priority = pri
	}
}