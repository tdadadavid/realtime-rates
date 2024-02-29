package exchangerates

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Secrets map[string]string

func GetSecretFromVault(name string) string {
	return getSecretFromAWSSecretManager(name)
}

func getSecretFromAWSSecretManager(secretName string) string {
	region := os.Getenv("AWS_REGION")

	config, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		return ""
	}

	// secret manager client
	secretManager := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	res, err := secretManager.GetSecretValue(context.Background(), input)
	if err != nil {
		log.Println("Error retrieving secret:", err)
		return ""
	}

	var secrets Secrets
	err = json.Unmarshal([]byte(*res.SecretString), &secrets)
	if err != nil {
		return ""
	}

	var result string

	for _, val := range secrets {
		result = val
	}

	return result
}

func mockGetSecretFromAWSSecretManager(secretName string) string {
	// Simulate fetching secrets for testing
	mockSecrets := map[string]string{
		"mockKey": "mockValue",
	}

	secret, exists := mockSecrets[secretName]
	if !exists {
		return ""
	}
	time.Sleep(5 * time.Millisecond)

	return secret
}
