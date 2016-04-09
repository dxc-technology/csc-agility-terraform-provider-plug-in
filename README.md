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
    + go to [INSTALL TERRAFORM](https://www.terraform.io/intro/getting-started/install.html) and follow instructions on intalling and configuring GO for your computer 
 
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

`provider "agility" {}`

`# Create a new Linux instance on a small server`

`resource "agility_compute" "myserver" {`

&nbsp;&nbsp;&nbsp;`    depends_on = ["agility_blueprint.myserver"]`

&nbsp;&nbsp;&nbsp;`    name = "&nbsp;`    

&nbsp;&nbsp;&nbsp;'	   TopologyId = "${agility_blueprint.myserver.TopologyId}"`

`}`

`resource "agility_blueprint" "myserver" {`

&nbsp;&nbsp;&nbsp;`    depends_on = ["agility_environment.Dev"]`

&nbsp;&nbsp;&nbsp;`    name = "Demo Server"`

&nbsp;&nbsp;&nbsp;`    version = "1"`

&nbsp;&nbsp;&nbsp;`    type = "XS"`

&nbsp;&nbsp;&nbsp;`    EnvironmentId = "${agility_environment.Dev.id}"`

&nbsp;&nbsp;&nbsp;`    ProjectId = "${agility_project.Demo.id}"`

`}`

`resource "agility_environment" "Dev" {`

&nbsp;&nbsp;&nbsp;`  	depends_on = ["agility_project.Demo"]`

&nbsp;&nbsp;&nbsp;`  	name = "Dev"`

&nbsp;&nbsp;&nbsp;`  	ProjectId = "${agility_project.Demo.id}"`

`}`

`resource "agility_project" "Demo" {`

&nbsp;&nbsp;&nbsp;`	name = "Demo"`

`} `

Then a VS was created by using the RAML as the definition source.
The root contect of the VS was made exactly the same as the root context of the Box API Hook VS. the consequense of this is that this simple HelloWorld VS and the Hook VS become a single VS.

Then the /helloworld Operation in the VS was mapped to the GET /users/me operation in the Box API Hook PS.

Go to the Box_API_Hook_Helloworld VS -> Operations Tab -> GET /hellowworld operation -> Process tab you'll see this image:
![Helloworld process] 
(https://github.com/pogo61/Box-API-Hook/blob/master/Box%20API%20Hook.png)

Double click on the Process activity to see that it call's the Heloworld Process, which call's an invoke on the GET /users/me operation to make the Hello World operation call successful.


### Create Your Own Integration with the Google Sheets API
The Hello World operation is one simple way of integrating or extending your API's.
Take a look at the [Box API Integration](https://github.com/pogo61/Box-API-Integration).
this will give you a deeper inderstanding of the richness of our gateway product in integrating to API's

### Modify and Build
In the event you need to change the API Hook.   Here are the instructions to do so. 

### License
Put a link to an open source license

