package mnemonic

import (
	"context"
	"database/sql"
	"fmt"
	"mnemonics/database"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func getMnemonic(MnemonicID string) (*Mnemonic, error) {
	row := database.DbConn.QueryRow(`Select * from mnems where MnemonicID = ?`, MnemonicID)
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

func getMnemonicList() ([]Mnemonic, error) {
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

func addOrUpdateMnenomic(mnem *Mnemonic) {

}

func dynamoaddOrUpdateMnenomic(svc *dynamodb.Client, Mnem Mnemonic) {
	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("Mnemonic"),
		Item: map[string]types.AttributeValue{
			"MnemonicID":     &types.AttributeValueMemberS{Value: Mnem.MnemonicID},
			"ItemtoRemember": &types.AttributeValueMemberS{Value: Mnem.ItemToRemember},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)
}
