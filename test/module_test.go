package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type RunSettings struct {
	t                         *testing.T
	workingDir                string `default:"../examples/build"`
	tfCliPath                 string `default:"/usr/local/bin/terraform"`
	approleID                 string
	// secretID                  *auth.SecretID
	vaultSecretPath           string
	uniqueID                  string
}

func (r *RunSettings) setDefaults() {
	if r.t == nil {
		panic("No Terratest module provided")
	}

	r.workingDir = "../examples/build"
	if ttwd := os.Getenv("TERRATEST_WORKING_DIR"); ttwd != "" {
		r.workingDir = ttwd
	}

	r.tfCliPath = "/usr/local/bin/terraform"
	if tfcp := os.Getenv("TF_CLI_PATH"); tfcp != "" {
		r.tfCliPath = tfcp
	}

	// VAULT items are used for interaction with HashiCorp Vault
	if vsecp := os.Getenv("VAULT_SECRET_PATH"); vsecp != "" {
		r.vaultSecretPath = vsecp
	}

	if role_id := os.Getenv("VAULT_APPROLE_ID"); role_id != "" {
		r.approleID = role_id
	}

	// if wrapped_token := os.Getenv("VAULT_WRAPPED_TOKEN"); wrapped_token != "" {
	// 	r.secretID = &auth.SecretID{FromEnv: "VAULT_WRAPPED_TOKEN"}
	// }

	if localId := os.Getenv("GITHUB_RUN_ID"); localId != "" {
		r.uniqueID = localId
	} else {
		r.uniqueID = random.UniqueId()
	}

}

func (r *RunSettings) setTerraformOptions() {
	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(r.t, &terraform.Options{
		TerraformDir:    r.workingDir,
		TerraformBinary: r.tfCliPath,
		Vars: map[string]interface{}{
			"unique_id": r.uniqueID,
		},
	})

	test_structure.SaveTerraformOptions(r.t, r.workingDir, terraformOptions)
}

func TestMyModule(t *testing.T) {
	r := RunSettings{t: t}
	r.setDefaults()
	
	t.Parallel()

	// to set terraform options
	// to skip execution "export SKIP_set_terraformOptions=true" in terminal
	test_structure.RunTestStage(t, "setTerraformOptions", r.setTerraformOptions)

	// Need to conditionally get secrets from Vault (if approle id, wrapped token and secrets path are all present)
	// getSpn, err := terratest_helper.GetSecretWithAppRole(r.approleID, r.secretID, r.vaultSecretPath)
	// if err != nil {
	// 	fmt.Print(err)
	// 	panic(err)
	// }
	// fmt.Print(getSpn)

	// Destroy the infra after testing is finished
	// export SKIP_terraformDestroy=true to skip this stage
	defer test_structure.RunTestStage(t, "terraformDestroy", r.terraformDestroy)

	// Deploy using Terraform
	// export SKIP_deployTerraform=true to skip this stage
	test_structure.RunTestStage(t, "deployTerraform", r.deployUsingTerraform)

		// Redeploy using Terraform and ensure idempotency
	// export SKIP_redeployTerraform=true to skip this stage
	test_structure.RunTestStage(t, "redeployTerraform", r.redeployUsingTerraform)

	// Perform tests
	// export SKIP_runTests=true to skip this stage
	test_structure.RunTestStage(t, "runTests", r.runTests)

}

func (r *RunSettings) deployUsingTerraform() {
	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)
	terraform.InitAndApply(r.t, terraformOptions)
}

func (r *RunSettings) redeployUsingTerraform() {
	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)
	terraform.ApplyAndIdempotent(r.t, terraformOptions)
}

func (r *RunSettings) runTests() {
	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)
	// sample test used for this templated repo
	expectedOutput := "Hello "+r.uniqueID
	output := terraform.Output(r.t, terraformOptions, "hello_world")
	assert.Equal(r.t, expectedOutput, output)
}

func (r *RunSettings) terraformDestroy() {
	terraformOptions := test_structure.LoadTerraformOptions(r.t, r.workingDir)
	terraform.Destroy(r.t, terraformOptions)
}
