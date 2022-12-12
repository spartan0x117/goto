package alias

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spartan0x117/goto/pkg/storage"
)

func getAliasFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := fmt.Sprintf("%s/.config/goto/aliases.json", home)
	return path, nil
}

func getAliasMap() (map[string]string, error) {
	path, err := getAliasFilePath()
	if err != nil {
		return nil, err
	}

	if !storage.FileExists(path) {
		return nil, err
	}

	aliasContents, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("encountered an error trying to read alias file")
		return nil, err
	}
	var aliases map[string]string
	err = json.Unmarshal(aliasContents, &aliases)
	if err != nil {
		fmt.Println("encountered an error trying to unmarshal alias file")
		return nil, err
	}
	return aliases, nil
}

// Searches aliases.json for the input. If nothing is found,
// the input is returned to be used as a label.
func GetLabelOrAlias(in string) string {
	aliases, err := getAliasMap()
	if err != nil {
		return in
	}
	aliasMapping, ok := aliases[in]

	// There was no mapping for this particular input
	if !ok {
		return in
	}
	return aliasMapping
}

func AddAlias(alias string, label string) error {
	label = storage.NormalizeLabel(label)
	aliases, err := getAliasMap()
	if err != nil {
		return err
	}
	aliases[alias] = label
	data, err := json.Marshal(aliases)
	if err != nil {
		return err
	}
	path, err := getAliasFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0666)
}

func RemoveAlias(alias string) error {
	aliases, err := getAliasMap()
	if err != nil {
		return err
	}
	delete(aliases, alias)
	data, err := json.Marshal(aliases)
	if err != nil {
		return err
	}
	path, err := getAliasFilePath()
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0666)
}
