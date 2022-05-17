package mnemonic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

//Use DB
//To implement

//Use File as Datasource --
//Maps are not thread safe, so we wrap a map
//in a mutex to allow locking/unlocking (preventing multiple threads from reading/writing the same object concurrently)
var MnemonicMap = struct {
	sync.RWMutex
	m map[int]Mnemonic
}{m: make(map[int]Mnemonic)}

func init() {
	fmt.Println("loading mnemonics")
	memMap, err := loadMnemonicMap()
	MnemonicMap.m = memMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mnemonics Loaded")
}

func loadMnemonicMap() (map[int]Mnemonic, error) {
	fileName := "mnemonics.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	mnemonicList := make([]Mnemonic, 0)
	err = json.Unmarshal([]byte(file), &mnemonicList)
	if err != nil {
		log.Fatal(err)
	}

	memMap := make(map[int]Mnemonic)
	for i := 0; i < len(mnemonicList); i++ {
		memMap[memMap[i].MnemonicID] = mnemonicList[i]
	}

	return memMap, nil
}

func getMnemonic(MnemonicID int) *Mnemonic {
	// MnemonicMap.RLock()

	// defer MnemonicMap.RUnlock()

	fmt.Println(MnemonicID)
	//item := MnemonicMap.m[MnemonicID]

	if mnemonic, ok := MnemonicMap.m[MnemonicID]; ok {
		fmt.Println("Should return Mem")
		return &mnemonic
	}

	return nil
}

func removeMnemonic(MnemonicID int) {
	MnemonicMap.Lock()
	defer MnemonicMap.Unlock()
	delete(MnemonicMap.m, MnemonicID)
}

func getMnemonicList() []Mnemonic {
	MnemonicMap.RLock()
	mnemonics := make([]Mnemonic, 0, len(MnemonicMap.m))
	for _, value := range MnemonicMap.m {
		mnemonics = append(mnemonics, value)
	}
	MnemonicMap.Unlock()
	return mnemonics
}

func getMnemonicIds() []int {
	MnemonicMap.RLock()
	mnemonicIds := []int{}
	for key := range MnemonicMap.m {
		mnemonicIds = append(mnemonicIds, key)
	}
	MnemonicMap.Unlock()
	sort.Ints(mnemonicIds)
	return mnemonicIds
}

func getNextMnemonicID() int {
	mnemonicIds := getMnemonicIds()
	return mnemonicIds[len(mnemonicIds)-1] + 1
}

func addOrUpdateMnenomic(mnemonic Mnemonic) (int, error) {
	//if id is set, update, otherwise add
	addOrUpdateID := -1

	if mnemonic.MnemonicID > 0 {
		//Update
		oldMnemonic := getMnemonic(mnemonic.MnemonicID)
		if oldMnemonic == nil {
			return 0, fmt.Errorf("mnenomic id [%d] doesn't exist", mnemonic.MnemonicID)
		}
		addOrUpdateID = mnemonic.MnemonicID
	} else {
		//New
		addOrUpdateID = getNextMnemonicID()
		mnemonic.MnemonicID = addOrUpdateID
	}
	MnemonicMap.Lock()
	MnemonicMap.m[addOrUpdateID] = mnemonic
	MnemonicMap.Unlock()
	return addOrUpdateID, nil
}
