package mnemonic_test

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"testing"
)

func Test_getMnemonic(t *testing.T) {
	//Arrange
	var MnemonicMap = struct {
		sync.RWMutex
		m map[int]Mnemonic
	}{m: make(map[int]Mnemonic)}

	fmt.Println("loading mnemonics")
	memMap, err := loadMnemonicMap()
	MnemonicMap.m = memMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mnemonics Loaded")
	//Act
	actual := getMnemonic(2)

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
