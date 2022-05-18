package mnemonic

type Mnemonic struct {
	MnemonicID     string
	MemoryTip      string
	ItemToRemember string //This is the actual thing to remember
	ImageID        string
	CreatedByID    string
	CreatedDate    string //Consider changing to date, Dynamo DB doesn't have a date type anyway
}
