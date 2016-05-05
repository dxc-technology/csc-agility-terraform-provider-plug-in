# csc-agility-terraform-provider-plug-in
![Image of CSC] 
(http://assets1.csc.com/home/images/logo.png) 
[CSC.com](http://CSC.com)

Terraform Provider Plugin for Agility



## Terraform Provider Plug-in
Terraform is built on a plugin-based architecture. All providers and provisioners that are used in Terraform configurations are plugins, even the core types such as AWS and Heroku. Users of Terraform are able to write new plugins in order to support new functionality in Terraform. 

## CSC Agility Platform Provider 
### About the Provider
- The Agility Platform Plugin utilises the Agility Platfporm API to manipulate the Blueprints, Templates, Environments and Projects defined in it
- API Documentation can be found at /agility/api/vX.Y on the Instanc e of Agility you have installed. the X.Y variable indicate the version of the API. At this current time the latest version is 3.2

### Pre-Reqs
- Install Go: 
    + go to [Go Lang Getting Started](https://golang.org/doc/install) and follow instructions on intalling and configuring GO for your computer

- Install Terraform:
    + go to [INSTALL TERRAFORM](https://www.terraform.io/intro/getting-started/install.html) and follow instructions on intalling and configuring Terraform for your computer 
 
### Getting Started Instructions
#### Download and Install
- Download the .zip, or use your desktop tools to save the project locally
- place the main.go and the agility directory into ~/work/src/github.com/"Github user"/terraform-provider-agility
	+ ~ is your OS's user directory
	+ "Github user" is your github username
- go to the newly created ~/work/src/github.com/<Github user>/terraform-provider-agility directory:
	+ run go get
		+ this will download all the dependent go libraries Terraform needs to build your plugin
    + run go build  
    	+ this will build the Agility provider plugin executable and place it in the same directory

#### Install plugin into Terraform
- Follow the instructions at [PLUGIN BASICS] (https://www.terraform.io/docs/plugins/basics.html) in the "Installing a Plugin" section
- when developing its best to point to your ~/work/src/github.com/"Github user"/terraform-provider-agility/terraform-provider-agility executable


### Using the Plugin
#### The agility.tf file
The agility.tf file is where you configure Terraform to create/update/delete a "Topology" based on a Blueprint and Policies in the Agility Platform

`variable "agility_userid" {}`

`variable "agility_password" {}`

`provider "agility" {`

&nbsp;&nbsp;&nbsp;    `userid = "${var.agility_userid}"`
    
&nbsp;&nbsp;&nbsp;    `password = "${var.agility_password}"`

`}`

`# Create a new Linux instance on a small server`

`resource "agility_compute" "myserver" {`

&nbsp;&nbsp;&nbsp;   `name = "myserver"`

&nbsp;&nbsp;&nbsp;    `active = "true"`

&nbsp;&nbsp;&nbsp;    `version = "1"`

&nbsp;&nbsp;&nbsp;    `type = "XS"`

&nbsp;&nbsp;&nbsp;    `blueprint = "Demo Server"`

&nbsp;&nbsp;&nbsp;    `environment = "Dev"`

&nbsp;&nbsp;&nbsp;    `project = "Demo"`

`}`

In the above example, an **"agility_compute"** terraform resource with the name of **'myserver'** will be created in the **"agility" provider**. Which means that a topology called myserver with be created in Agility.

*The Provider called "agility" takes two parameters (userid and password). These are set from Terraform Variables (agility_userid, agility_password). These, in turn, are set from Environment Variables (TF_VAR_agility_userid and TF_VAR_agility_password).

*Changing the name parameter of the agility_compute resource will make the plugin change the name of the VM and Topology in Agility, without affecting the VM* **This is Mandatory**

*Changing the active parameter of the agility_compute resource will make the plugin either stop or Start the VM. 'true' means start, 'false' means stop* **This is Mandatory**

*The `version = "1"` tells the plugin to use version 1 of the blueprint. If this parameter is omited the plugin will use the latest version of the blueprint.*

*The `type = "XS"` tells the plugin to use the equivalent computer definition for the cloud the blueprint will be deployed into.* **This is Mandatory**

*The `blueprint = "Demo Server"` tells the plugin to search agility for a blueprint in the project called 'Demo Server'. **This is Mandatory**

*The `environment = "Dev"` tells the plugin to search agility for an environment in the project called 'Dev'. **This is Mandatory**

*The `project = "Demo"` tells the plugin to search agility for a project called 'Demo'. **This is Mandatory**

***Putting this all together, A small VM called myserver will be created in the Dev Environment, of the Demo Project in Agility, based on the 'Demo Server' blueprint***

### Troubleshooting
The Plugin writes details of its actions to an *agility.log* file in the same directory as the agility.tf file

Use the contents of this file when reporting issues with the plugin, or trying to determine what the plugin is doing.

### License
Put a link to an open source license

