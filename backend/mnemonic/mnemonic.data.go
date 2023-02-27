package mnemonic

import (
	"database/sql"
	"fmt"
	"mnemonics/database"
	"time"

	"github.com/google/uuid"
)

func GetMnemonic(MnemonicID string) (*Mnemonic, error) {
	row := database.DbConn.QueryRow(`Select 
	MnemonicID,
	MemoryTip,
	ItemToRemmeber,
	ImageID,
	CreatedByID,
	CreateDate
	from mnems where MnemonicID = ?`, MnemonicID)
	mnemonic := &Mnemonic{}
	err := row.Scan(
		&mnemonic.MnemonicID,
		&mnemonic.MemoryTip,
		&mnemonic.ItemToRemember,
		&mnemonic.ImageID,
		&mnemonic.CreatedByID,
		&mnemonic.CreatedDate)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return mnemonic, nil
}

func GetMnemonicList() ([]Mnemonic, error) {
	results, err := database.DbConn.Query(`select * from mnems`)
	if err != nil {
		fmt.Println("Err getting rows")
		return nil, err

	}
	defer results.Close()
	mnemonics := make([]Mnemonic, 0)
	for results.Next() {
		var mnemonic Mnemonic
		results.Scan(
			&mnemonic.MnemonicID,
			&mnemonic.MemoryTip,
			&mnemonic.ItemToRemember,
			&mnemonic.ImageID,
			&mnemonic.CreatedByID,
			&mnemonic.CreatedDate)
		mnemonics = append(mnemonics, mnemonic)
	}
	fmt.Println(mnemonics[0].MnemonicID)
	return mnemonics, nil
}

func AddMnenomic(mnem Mnemonic) (int, error) {
	mnemid := uuid.NewString()

	//Not null/empty check (MySql is writing empty strings)
	//if  strings.TrimSpace(mnem.MemoryTip)

	mnem = GetForeignKeys(mnem)
	result, err := database.DbConn.Exec(`INSERT INTO mnems 
	(MnemonicID,
	MemoryTip,
	ItemToRemember,
	ImageID,
	CreatedByID,
	CreatedDate)
	 VALUES (?,?,?,?,?,?)`,
		mnemid,
		&mnem.MemoryTip,
		&mnem.ItemToRemember,
		&mnem.ImageID,
		&mnem.CreatedByID,
		time.Now())
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(insertID), nil
}

func GetForeignKeys(mnem Mnemonic) Mnemonic {
	row := database.DbConn.QueryRow(`select UserID FROM user WHERE UserID = ?`, mnem.CreatedByID)
	var userID string
	err := row.Scan(&userID)
	if err == sql.ErrNoRows {
		userID = "Anonymous"
	} else {
		mnem.CreatedByID = userID
	}

	mnem.ImageID = "Figure out later"

	return mnem
}

func UpdateMnenomic(mnem Mnemonic) error {
	_, err := database.DbConn.Exec(`UPDATE mnems SET 
	MnemonicID = ?,
	MemoryTip=?,
	ItemToRemember=?,
	ImageID=?,
	CreatedByID=?,
	CreatedDate=?
	WHERE MnemonicID=?`,
		mnem.MnemonicID,
		mnem.MemoryTip,
		mnem.ItemToRemember,
		mnem.ImageID,
		mnem.CreatedByID,
		mnem.CreatedDate,
		mnem.MnemonicID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteMnemonic(MnemonicId string) error {
	_, err := database.DbConn.Query(`DELETE FROM mnems WHERE MnemonicID = ?`, MnemonicId)
	if err != nil {
		return err
	}
	return nil
}

//When switching to DynamoDB, this is the format to use
// func dynamoaddOrUpdateMnenomic(svc *dynamodb.Client, Mnem Mnemonic) {
// 	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
// 		TableName: aws.String("Mnemonic"),
// 		Item: map[string]types.AttributeValue{
// 			"MnemonicID":     &types.AttributeValueMemberS{Value: Mnem.MnemonicID},
// 			"ItemtoRemember": &types.AttributeValueMemberS{Value: Mnem.ItemToRemember},
// 		},
// 	})

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(out.Attributes)
// }
