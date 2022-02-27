package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestMyModule(t *testing.T) {


	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/build",
		Vars: map[string]interface{}{
			// A unique ID we can use to namespace resources so we don't clash with anything already in the AWS account or
			// tests running in parallel
			uniqueID = random.UniqueId(),
		},
	})

	// Destroy the infra after testing is finished
	// export SKIP_terraform_destroy to skip this stage
	defer test_structure.RunTestStage(t, "terraform_destroy", func() {
		terraformDestroy(t, workingDir)
	})

	// Deploy the web app using Terraform
	// export SKIP_deploy_terraform to skip this stage
	test_structure.RunTestStage(t, "deploy_terraform", func() {
		awsRegion := test_structure.LoadString(t, workingDir, "awsRegion")
		deployUsingTerraform(t, awsRegion, workingDir)
	})

	// Perform tests
	// export SKIP_run_tests to skip this stage
	test_structure.RunTestStage(t, "validate", func() {
		runTests(t, workingDir)
	})
}

func deployUsingTerraform(testing t, workingDir) {
	terraform.InitAndApply(t, terraformOptions)
}

func redeployUsingTerraform(testing t, workingDir) {
	terraform.ApplyAndIdempotent(t, terraformOptions)
}

func runTests(testing t, workingDir) {
	// sample test used for this templated repo
	output := terraform.Output(t, terraformOptions, "hello_world")
	assert.Equal(t, "Hello, World!", output)
}

func terraformDestroy(testing t, workingDir) {
	terraform.Destroy(t, terraformOptions)
}
