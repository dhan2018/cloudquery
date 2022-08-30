# Overview

This is a fork of https://github.com/cloudquery/cloudquery

The key capabilities we enable in this fork are

1. Ensure all commands will have a ORG_ID Set
2. All tables will be created with ORG_ID field
3. Whenever fetch is called all the tables will be filled with correct ORG_ID value

# Setup and build Cloudquery

## Clone the cloudquery repo in WSL
```
NOTE: In Ubuntu shell
cd ~/wsl_finopsroot/code
git clone git@github.com:dhan2018/cloudquery.git

```


## Building cloudquery binaries

```
NOTE: These commands are inside the docker container
cd /home/finopsroot/code/
cd cloudquery

# Full build of all binaries
make build

# verify that all the binaries are generated with latest timestamps
 ls -al bin/

 # copy all the latest cloudquery binaries to the finops folder
 mkdir -p ../finops/bin/cq
ls -al /home/finopsroot/code/finops/bin/cq/  # Verify all the binaries are copied correctly


```

## Debugging the provider 

The below instructions will help you debug the provider
https://www.cloudquery.io/docs/developers/debugging

