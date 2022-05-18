package mnemonic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func getMnemonic(MnemonicID string) *Mnemonic {
}

func getMnemonicList() []Mnemonic {

}

func addOrUpdateMnenomic(svc *dynamodb.Client, Mnem Mnemonic) {
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
