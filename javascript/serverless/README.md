# Serverless
## overview
- [doc](#doc)
- [install](#install)
- [notes](#notes)
## doc
- https://www.serverless.com
## install
```sh
make help
```
## notes
create project
- serverless
```sh
make project
```
- serverless dynamodb
```sh
make project.dynamodb
```
AWS
```sh
export AWS_ACCESS_KEY_ID=YOUR_AWS_ACCESS_KEY_ID
export AWS_SECRET_ACCESS_KEY=YOUR_AWS_SECRET_ACCESS_KEY
export AWS_DEFAULT_REGION=eu-west-1
```
```sh
cat <<EOF > $HOME/.aws/credentials
[default]
aws_access_key_id = $AWS_ACCESS_KEY_ID
aws_secret_access_key = $AWS_SECRET_ACCESS_KEY
region = $AWS_DEFAULT_REGION
EOF
```
```sh
cat <<EOF > $HOME/.aws/config
[default]
region=us-west-2
output=json
EOF
```