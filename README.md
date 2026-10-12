# terraform-<provider>-<module name>

This repo will be used as a template for new Terraform module Github repos.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0.0 |

## Providers

No providers.

## Modules

No modules.

## Resources

No resources.

## Inputs

No inputs.

## Outputs

No outputs.
<!-- END_TF_DOCS -->

## short cut for using with existing projects.
  
1. Add the bellow snippet to an a `$HOME/.bashrc` or `$HOME/.zshrc` file to quickly use this repo to create sub-modulkes with in existing projects.
```shell
tf_new_module() {
 local usage="tf_new_module module_name"
 [[ -z $1 ]] && echo -e "\033[31;1m[ERROR]\033[0m Needs a module name. \n $usage" $$ return 1
 git subtree add --prefix $1 https://github.com/Diehlabs/terraform-module-scaffolding.git main --squash
}
```
2. Restart reload/restart the profile (or restart your terminal) and run the bellow to create a new modules.
3. In the terraform code repo. ake sure all chages are commited.
4. Run `tf_new_module MODULE_NAME` where module_name will be the path to the new module.
