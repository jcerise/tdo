package todo

import (
	"encoding/json"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	Position int
	Done     bool
}

func (i *Item) SetPriority(priority int) {
	switch priority {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyPrint() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}

	return " "
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.Position) + "."
}

// ByPriority implements sort.Interface, allowing us to sort lists of Items
type ByPriority []Item

func (s ByPriority) Len() int      { return len(s) }
func (s ByPriority) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPriority) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return !s[i].Done
	}
	if s[i].Priority == s[j].Priority {
		return s[i].Position < s[j].Position
	}

	return s[i].Priority < s[j].Priority
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	// Check if the file exists first

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}

	for i, _ := range items {
		items[i].Position = i + 1
	}

	return items, nil
}

func OpenOrCreateDataFile() {
	// Check if the file exists. If it does not, create it
	if _, err := os.Stat(viper.GetString("datafile")); err != nil {
		_, createErr := os.Create(viper.GetString("datafile"))
		if createErr != nil {
			log.Printf("%v", createErr)
		} else {
			log.Printf("datafile %v does not exist...creating", viper.GetString("datafile"))
		}
	}
}
