package mnemonic

type Mnemonic struct {
	MnemonicID     string
	MemoryTip      string //This is the tip i.e. Please excuse my dear aunt sally
	ItemToRemember string //This is the actual thing to remember - PEMDAS - Parentheses, Exponents...
	ImageID        string
	CreatedByID    string
	CreatedDate    string //Consider changing to date
	//Add these fields later
	//SrcLink        string //Could make a new table later to add multiple
	//Tags         //DB Note - this is a M:M relationship - Need a tags table and a map table
}

//Add Getters and Setters?
