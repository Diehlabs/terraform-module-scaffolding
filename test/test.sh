#!/bin/bash
# Examples for testing

TERRATEST_TIMEOUT=10m

export TF_CLI_PATH="/usr/local/bin/terraform"
export TERRATEST_WORKING_DIR="../examples/build"
#VAULT_SECRET_PATH
#VAULT_APPROLE_ID
#VAULT_WRAPPED_TOKEN

# export SKIP_set_terraformOptions=true
# export SKIP_deployTerraform=true
# export SKIP_redeployTerraform=true
# export SKIP_runTests=true
# export SKIP_terraformDestroy=true

go test -v -timeout=$TERRATEST_TIMEOUT