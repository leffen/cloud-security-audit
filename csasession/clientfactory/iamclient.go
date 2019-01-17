package clientfactory

import (
	"github.com/aws/aws-sdk-go/service/iam"
)

type IAMClient interface {
	ListUsers(input *iam.ListUsersInput) (*iam.ListUsersOutput, error)
	ListAccessKeys(input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error)
}

type AWSIAMClient struct {
	api *iam.IAM
}

func (client AWSIAMClient) ListUsers(input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	return client.api.ListUsers(input)
}

func (client AWSIAMClient) ListAccessKeys(input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	return client.api.ListAccessKeys(input)
}