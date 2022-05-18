package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	//cfg.Region = "us-east-1"
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.EndpointResolver = dynamodb.EndpointResolverFromURL("http://localhost:8000")
		o.Region = "us-east-1"
	})

	//getTables(svc)

	//itemInfo := getItemInfo{tablename: "Mnemonic", primaryKeyName: "mnemonicID", primaryKeyValue: "test2"}
	//getItem(svc, itemInfo)
	addOrUpdateItem(svc)
}

func getTables(svc *dynamodb.Client) {

	p := dynamodb.NewListTablesPaginator(svc, nil, func(o *dynamodb.ListTablesPaginatorOptions) {
		o.StopOnDuplicateToken = true
	})

	for p.HasMorePages() {
		out, err := p.NextPage(context.TODO())
		if err != nil {
			panic(err)
		}

		for _, tn := range out.TableNames {
			fmt.Println(tn)
		}
	}
}

//The lowercase only works b/c this is in the same package.
//If I needed to import it, then these would be private fields (due to the lowercase) therefore I could not initialize them
type getItemInfo struct {
	tablename       string
	primaryKeyName  string
	primaryKeyValue string
}

func getItem(svc *dynamodb.Client, itemtoGet getItemInfo) {
	out, _ := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(itemtoGet.tablename),
		Key: map[string]types.AttributeValue{
			itemtoGet.primaryKeyName: &types.AttributeValueMemberS{Value: itemtoGet.primaryKeyValue},
		},
	})

	fmt.Println(out.Item)
}

func addOrUpdateItem(svc *dynamodb.Client) {
	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("User"),
		Item: map[string]types.AttributeValue{
			"UserID": &types.AttributeValueMemberS{Value: "12343336zz"},
			"name":   &types.AttributeValueMemberS{Value: "John Doe"},
			"email":  &types.AttributeValueMemberS{Value: "john@doe.io"},
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(out.Attributes)
}
