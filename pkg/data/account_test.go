package data

import (
	"testing"

	awsmocks "github.com/Optum/dce/pkg/awsiface/mocks"
	"github.com/Optum/dce/pkg/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// account - Handles importing and exporting Accounts and non-exported Properties
type account struct {
	ID                  string                 `json:"id" dynamodbav:"Id"`                                   // AWS Account ID
	Status              string                 `json:"Status" dynamodbav:"Status"`                           // Status of the AWS Account
	LastModifiedOn      int64                  `json:"lastModifiedOn" dynamodbav:"LastModifiedOn"`           // Last Modified Epoch Timestamp
	CreatedOn           int64                  `json:"createdOn"  dynamodbav:"CreatedOn"`                    // Account CreatedOn
	AdminRoleArn        string                 `json:"adminRoleArn"  dynamodbav:"AdminRoleArn"`              // Assumed by the master account, to manage this user account
	PrincipalRoleArn    string                 `json:"principalRoleArn"  dynamodbav:"PrincipalRoleArn"`      // Assumed by principal users
	PrincipalPolicyHash string                 `json:"principalPolicyHash" dynamodbav:"PrincipalPolicyHash"` // The the hash of the policy version deployed
	Metadata            map[string]interface{} `json:"metadata"  dynamodbav:"Metadata"`                      // Any org specific metadata pertaining to the account
}

func TestGet(t *testing.T) {

	t.Run("should return an account object", func(t *testing.T) {
		mockDynamo := awsmocks.DynamoDBAPI{}

		accountID := "abc123"
		currentStatus := "Ready"

		mockDynamo.On("GetItem", &dynamodb.GetItemInput{
			ConsistentRead: aws.Bool(false),
			Key: map[string]*dynamodb.AttributeValue{
				"Id": {
					S: aws.String(accountID),
				},
			},
			TableName: aws.String("Accounts"),
		}).Return(
			&dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{
					"Id": {
						S: aws.String(accountID),
					},
					"Status": {
						S: aws.String(string(currentStatus)),
					},
				},
			}, nil,
		)
		accountData := &Account{
			AwsDynamoDB: &mockDynamo,
			TableName:   "Accounts",
		}
		item := &model.Account{}
		err := accountData.GetAccountByID(accountID, item)
		assert.NoError(t, err)
		assert.NotNil(t, item.ID, accountID)
	})

	t.Run("should return nil object when not found", func(t *testing.T) {
		mockDynamo := awsmocks.DynamoDBAPI{}

		accountID := "abc123"

		mockDynamo.On("GetItem", &dynamodb.GetItemInput{
			ConsistentRead: aws.Bool(false),
			Key: map[string]*dynamodb.AttributeValue{
				"Id": {
					S: aws.String(accountID),
				},
			},
			TableName: aws.String("Accounts"),
		}).Return(
			&dynamodb.GetItemOutput{
				Item: map[string]*dynamodb.AttributeValue{},
			}, nil,
		)
		accountData := &Account{
			AwsDynamoDB: &mockDynamo,
			TableName:   "Accounts",
		}
		item := &model.Account{}
		err := accountData.GetAccountByID(accountID, item)
		assert.Error(t, err, "Account not found")
		assert.NotNil(t, item.ID, "")
	})

}

func TestDelete(t *testing.T) {

	t.Run("should delete an account", func(t *testing.T) {
		mockDynamo := awsmocks.DynamoDBAPI{}

		accountID := "abc123"

		mockDynamo.On("DeleteItem", &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"Id": {
					S: aws.String(accountID),
				},
			},
			TableName:    aws.String("Accounts"),
			ReturnValues: aws.String("ALL_NEW"),
		}).Return(
			&dynamodb.DeleteItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"Id": {
						S: aws.String(accountID),
					},
					"AccountStatus": {
						S: aws.String("Ready"),
					},
				},
			}, nil,
		)
		accountData := &Account{
			AwsDynamoDB: &mockDynamo,
			TableName:   "Accounts",
		}
		item := &model.Account{}
		err := accountData.Delete(accountID, item)
		assert.NoError(t, err)
		assert.Equal(t, item.ID, accountID)
		assert.Equal(t, item.Status, model.Ready)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("should Update", func(t *testing.T) {
		mockDynamo := awsmocks.DynamoDBAPI{}

		accountID := "abc123"
		lastModifiedOn := int64(1573592058)

		mockDynamo.On("PutItem", mock.MatchedBy(func(input *dynamodb.PutItemInput) bool {
			return (*input.TableName == "Accounts" &&
				*input.Item["Id"].S == "abc123" &&
				*input.Item["AccountStatus"].S == "Ready" &&
				*input.Item["Metadata"].M["key"].S == "value" &&
				*input.ReturnValues == "ALL_NEW")
		})).Return(
			&dynamodb.PutItemOutput{
				Attributes: map[string]*dynamodb.AttributeValue{
					"Id": {
						S: aws.String(accountID),
					},
					"LastModifiedOn": {
						N: aws.String("1573592058"),
					},
					"Metadata": {
						M: map[string]*dynamodb.AttributeValue{
							"key": {
								S: aws.String("value"),
							},
						},
					},
				},
			}, nil,
		)
		accountData := &Account{
			AwsDynamoDB: &mockDynamo,
			TableName:   "Accounts",
		}
		item := &model.Account{
			ID: accountID,
			Metadata: map[string]interface{}{
				"key": "value",
			},
			Status: "Ready",
		}

		err := accountData.Update(item, &lastModifiedOn)
		assert.NoError(t, err)
		assert.Equal(t, item.LastModifiedOn, int64(1573592058))
	})

}
