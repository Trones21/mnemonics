package mnemonic_test

import (
	"encoding/json"
	"fmt"
	"mnemonics/mnemonic"
	"testing"
)

func Test_getMnemonic(t *testing.T) {
	//Arrange
	//Connect to DB

	//Act
	actual, err := mnemonic.GetMnemonic("<ID>")

	//Assert
	jsonOBJ, err := json.Marshal(actual)
	if err != nil {
		t.Errorf("Error unmarshalling JSON")
	} else {
		fmt.Print(jsonOBJ)
	}
}

// func Test_loadFile(t *testing.T) {

// 	loadFile()

// }

// func loadFile() (status *bool, err error) {
// 	fileName := "mnemonics.json"
// 	_, err = os.Stat(fileName)
// 	if os.IsNotExist(err) {
// 		return nil, fmt.Errorf("file [%s] does not exist", fileName)
// 	}
// 	return true, nil
// }
