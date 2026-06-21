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
cat <<EOF > ~/.aws/credentials 
[default]
aws_access_key_id=YOUR_AWS_ACCESS_KEY_ID
aws_secret_access_key=YOUR_AWS_SECRET_ACCESS_KEY
EOF
```
```sh
cat <<EOF > ~/.aws/config 
[default]
region=us-west-2
output=json
EOF
```