//go:build exclude

package p //Use collection in menory
import (
	"encoding/json"
	"log"
)

var mnemonicsList []Mnemonic

func initFromString() {
	mnemonicsJSON := `[
		{"MnemonicID": 1, 
	"MemoryTip": "Please Excuse My Dear Aunt Sally",
	"ItemToRemember": "Maths Order of Operations - Parentheses, Exponents ..."
	},
	{"MnemonicID": 2, 
	"MemoryTip": "In fourteen hundred ninety-two, Columbus sailed the ocean blue",
	"ItemToRemember": "Date - That Columbus landed in America in 1492"
	},
	{"MnemonicID": 3, 
	"MemoryTip": "gblfblgldflg",
	"ItemToRemember": "fsdfsdfsfsdsffdfsd"
	}
	]
	`
	err := json.Unmarshal([]byte(mnemonicsJSON), &mnemonicsList)
	if err != nil {
		log.Fatal(err)
	}
}

func findMnemonicByID(id int) (*Mnemonic, int) {
	for i, mnemonic := range mnemonicsList {
		if mnemonic.MnemonicID == id {
			return &mnemonic, i
		}
	}
	return nil, 0
}

func getNextID() int {
	highestID := -1
	for _, mnemonic := range mnemonicsList {
		if highestID < mnemonic.MnemonicID {
			highestID = mnemonic.MnemonicID
		}
	}
	return highestID + 1
}
