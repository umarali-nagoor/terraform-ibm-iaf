package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestAccIBMIaf(t *testing.T) {
	t.Parallel()

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../examples/installation",

		Vars: map[string]interface{}{
			"on_vpc":                       false,
			"cluster_id":                   "***************",
			"region":                       "us-south",
			"resource_group":               "Default",
			"ibmcloud_api_key":             "****************",
			"entitled_registry_user_email": "test@in.ibm.com",
			"entitled_registry_key":        "****************",
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
}
