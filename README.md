# terraform-provider-salesforce

> Not a real provider! It can do nothing practically

## Build

To download all dependencies call:
```bash
dep ensure
```

To build the code call:
```bash
go build -o terraform-provider-salesforce && \
    cp ./terraform-provider-salesforce tf/terraform.d/plugins/darwin_amd64
```

## Test terraform

You'll need a developer org, a Connected App with OAuth in the org and few parameters should be added into a `tf/terraform.tfvars` file:
```hcl
client_id = "<Client ID for the app>"
client_secret = "<Client Secret for the app>"
username = "<User name for an admin user>"
password = "<Password for the admin user>"
```

After each change in the provider run `terraform init`

To plan changes:
```bash
tf$ terraform plan -out=terraform.plan -var-file=./terraform.tfvars
```

To apply changes:
```bash
tf$ terraform apply terraform.plan
```
