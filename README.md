# csc-agility-terraform-provider-plug-in
[<img src="http://assets1.csc.com/home/images/logo.png">](http://csc.com/)

Terraform Provider Plugin for Agility



## Terraform Provider Plug-in
Terraform is built on a plugin-based architecture. All providers and provisioners that are used in Terraform configurations are plugins, even the core types such as AWS and Heroku. Users of Terraform are able to write new plugins in order to support new functionality in Terraform. 

## CSC Agility Platform Provider 
### About the Provider
- The Agility Platform Plugin utilises the Agility Platfporm API to manipulate the Blueprints, Templates, Environments and Projects defined in it.
- API Documentation can be found at /agility/api/vX.Y on the Instance of Agility you have installed. The X.Y variables indicate the version of the API. At this current time the latest version is 3.2.

### Pre-Reqs
- Install Go: 
    + Go to [Golang Getting Started](https://golang.org/doc/install) and follow instructions on intalling and configuring Go for your computer.

- Install Terraform:
   + Go to [Install Terraform](https://www.terraform.io/intro/getting-started/install.html) and follow instructions on intalling and configuring Terraform for your computer.
 
### Getting Started Instructions
#### Download and Install
- Download the .zip, or use your desktop tools to save the project locally.
- Place the main.go and the agility directory into ```~/work/src/github.com/csc/csc-agility-terraform-provider-plug-in```
   + ```~``` is your OS's user directory
- go to the newly created ```~/work/src/github.com/csc/csc-agility-terraform-provider-plug-in```:
   + run ```go get```  
      + this will download all the dependent go libraries Terraform needs to build your plugin  
   + run ```go build```  
      + this will build the Agility provider plugin executable and place it in the same directory  

#### Install plugin into Terraform
- Follow the instructions at [Plugin Basics](https://www.terraform.io/docs/plugins/basics.html) in the "Installing a Plugin" section
- when developing it's best to point to your ~/work/src/github.com/csc/csc-agility-terraform-provider-plug-in/terraform-provider-agility executable


### Using the Plugin
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
The Plugin writes details of its actions to an ```agility.log``` file in the same directory as the ```agility.tf``` file

Use the contents of this file when reporting issues with the plugin, or trying to determine what the plugin is doing.

### License
The Services SDK and the reference code provided in this repository are licensed under the Apache License, Version 2.0
See [LICENSE](https://github.com/csc/csc-agility-terraform-provider-plug-in/blob/master/LICENSE) for the full license text.

## Contributing to the Project
Please follow the standard Git repository process of "Forking" and "Pull Request" [defined here](https://guides.github.com/activities/forking/)

