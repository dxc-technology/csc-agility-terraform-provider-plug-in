# csc-agility-terraform-provider-plug-in
[<img src="http://assets1.csc.com/home/images/logo.png">](http://csc.com/)

Terraform Provider Plugin for Agility



## Terraform Provider Plug-in
Terraform is built on a plugin-based architecture. All providers and provisioners that are used in Terraform configurations are plugins, even the core types such as AWS and Heroku. Users of Terraform are able to write new plugins in order to support new functionality in Terraform. 

## CSC Agility Platform Provider 
### About the Provider
This demo Agility plugin for Terraform plugin allows Terraform to provision an Agility Blueprint using the default deployer.  Note that at this time "advanced deployment" options are NOT supported.

- The Agility Platform Plugin utilises the Agility Platform REST API to manipulate Agility Blueprints, Templates, Environments and Projects.
- API Documentation can be found at /agility/api/vX.Y on the Instance of Agility you have installed. The X.Y variables indicate the version of the API. At the time of this writing, the latest version is 3.3.


 
### Getting Started 

PreRequisites:

 - Terraform installed
 - Agility Platform installed
 - One of 3 Cloud providers installed, with any of the following hardcoded strings in the Agility Cloud Provider name:
    - Mock
    - BizCloud
    - AWSCloud

This naming convention is a release limitation, it can be fixed in future versions.

#### Download and Install
Download release package from Github: [Plugin Releases](https://github.com/dxc-technology/csc-agility-terraform-provider-plug-in/releases)

Copy plugin to terraform plugins directory (rename if necessary):
terraform-provider-agility_v1.0.0

- Windows: `%appdata%\terraform.d\plugins\windows_amd64`
- OS X:    `~/.terraform.d/plugins/darwin_amd64/`
- Linux:   `~/.terraform.d/plugins/linux_amd64/`

If you get stuck, consult the official Terraform instructions at [Plugin Basics](https://www.terraform.io/docs/plugins/basics.html) in the "Installing a Plugin" section

Copy conf.json and test.tf to your Terraform working directory.  Modify conf.json to point to your Agility Platform.  Modify test.tf to conform to your environment, project, and Cloud Provider.



### Using the Plugin

Run `terraform init` to test.  If the plugin is installed properly, the command should complete successfully.  

If you want to deploy an Agility Store Product, order and approve the Store Product first so the project will already have a Blueprint ready to deploy.

Run `terraform apply` to deploy the project Blueprint and create a topology and running Virtual Machine. 

#### The conf.json file

The conf.json file provides default configuration options for the plugin, including mappings to Agility model sizes and Cloud Providers.

```
{
    "MaxRetries": "2",
    "APIURL": "https://52.XX.XX.XX/agility/api/v3.2/",

    "AWSXS": "t1.micro",
    "BCXS": "TBD",
    "MXS": "mock.small.x64",
    
    "AWSCloud": "Amazon EC2",
    "BizCloud": "BizCloud",
    "Mock": "Mock"
}
```



#### The agility.tf file
The agility.tf file is where you configure Terraform to create/update/delete a "Topology" based on a Blueprint and Policies in the Agility Platform

```
variable "agility_userid" {}
variable "agility_password" {}
provider "agility" {
    userid = "${var.agility_userid}"
    password = "${var.agility_password}"
}
# Create a new Linux instance on a small server
resource "agility_compute" "myserver" {
    name = "myserver"
    active = "true"
    version = "1"
    type = "XS"
    blueprint = "Demo Server"
    environment = "Dev"
    project = "Demo"
}
```

In the above example, an **"agility_compute"** Terraform resource with the name of **'myserver'** will be created in the **"agility" provider**, which means that a topology called myserver with be created in Agility.

- The Provider called "agility" takes two parameters (userid and password). These are set from Terraform Variables ```agility_userid``` ```agility_password```. These, in turn, are set from Environment Variables ```TF_VAR_agility_userid``` and ```TF_VAR_agility_password```.

- Changing the name parameter of the agility_compute resource will make the plugin change the name of the VM and Topology in Agility, without affecting the VM **This is Mandatory**

- Changing the active parameter of the agility_compute resource will make the plugin either stop or Start the VM. 'true' means start, 'false' means stop **This is Mandatory**

- The version = "1"` tells the plugin to use version 1 of the blueprint. If this parameter is omited the plugin will use the latest version of the blueprint.

- The `type = "XS"` tells the plugin to use the equivalent computer definition for the cloud the blueprint will be deployed into. **This is Mandatory**

- The `blueprint = "Demo Server"` tells the plugin to search agility for a blueprint in the project called 'Demo Server'. **This is Mandatory**

- The `environment = "Dev"` tells the plugin to search agility for an environment in the project called 'Dev'. **This is Mandatory**

- The `project = "Demo"` tells the plugin to search agility for a project called 'Demo'. **This is Mandatory**

***Putting this all together, A small VM called myserver will be created in the Dev Environment, of the Demo Project in Agility, based on the 'Demo Server' blueprint***

### Troubleshooting
More debugging information can be be obtained by setting the TF_LOG environment variable before running `terraform apply`

```export TF_LOG=1```

Agility REST interactions are also logged to agility.log, located in the same folder as your terraform file.

### License
The Services SDK and the reference code provided in this repository are licensed under the Apache License, Version 2.0
See [LICENSE](https://github.com/csc/csc-agility-terraform-provider-plug-in/blob/master/LICENSE) for the full license text.

## Building the Agility Terraform Plugin

Note that this plugin is in draft form, including the build process

### PreRequisites

- Install Go: 
    + Go to [Golang Getting Started](https://golang.org/doc/install) and follow instructions on installing and configuring Go for your computer.

- Install Terraform:
   + Go to [Install Terraform](https://www.terraform.io/intro/getting-started/install.html) and follow instructions on intalling and configuring Terraform for your computer.

- Clone repo to the Go projects folder for your OS.  Typically, Go expects all source files to be located in ~/go/src (ex. /Users/dlitster/go/src/github.com/csc/csc-agility-terraform-provider-plug-in/)   


### Building

- change to the folder that contains main.go:
- run ```go get```  
  + this will download all the dependent go libraries Terraform needs to build your plugin  
- run ```go build```  
    + this will build the Agility provider plugin executable and place it in the same directory
    + optionally, create a "build" folder for your compiled binary
    + optionally, create a symlink from the terraform plugins folder to the binary that go builds to avoid copying the plugin to the terraform folder on each build

    `ln -s build/linux_amd64/terraform-provider-agility_v1.0.0 ~/.terraform.d/plugins/darwin_amd64/terraform-provider-agility_v1.0.0`

Note that each new build requires running `terraform init` in the Terraform working directory to use the new binary.

#### Sample OS Build Targets

```
env GOOS=linux GOARCH=amd64 go build -o build/linux_amd64/terraform-provider-agility_v1.0.0
env GOOS=darwin GOARCH=amd64 go build -o build/darwin_amd64/terraform-provider-agility_v1.0.0
env GOOS=windows GOARCH=amd64 go build -o build/windows_amd64/terraform-provider-agility_v1.0.0.exe
```




## Contributing to the Project 
Please follow the standard Git repository process of "Forking" and "Pull Request" [defined here](https://guides.github.com/activities/forking/)

