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
- place the main.go and the agility directory into ~/work/src/github.com/<Github user>/terraform-provider-agility
	+ ~ is your OS's user directory
	+ <Github user> is your github username
- go to the newly created ~/work/src/github.com/<Github user>/terraform-provider-agility directory:
	+ run go get
		+ this will download all the dependent go libraries Terraform needs to build your plugin
    + run go build  
    	+ this will build the Agility provider plugin executable and place it in the same directory

#### Istall plugin into Terraform
- Expand the services folder in the Box API Hook you imported and find Box_API_Hook VS

#### Activate Anonymous Contract
- Expand the contracts folder in the Google Sheets API Hook you imported and find the "Anonymous" contract under the "Provided Contracts" folder
- click on the "Activate Contract" workflow activity in the righ-hand Activities portlet
- ensure that the status changes to "Workflow Is Completed"

#### Configure Security
- Go to Box API Hook -> Policies -> Operational Policies ->    ProcessVariables policy
    - Click "modify" in the XML Policy Tab. An XML Policy Content editor dialog will be displayed.
    - change the value of the 'appkey' element to be the 'client_id' value you saved from the Box Developers App Console, above. 
    - Go to the "PM home Dir"/sm70/scripts directory and run the encryptData script/batch file and enter the 'client_secret' you saved from the Box Developers App Console, above.
    - copy the resultant value (including the two '==' at the end of the string)
    - change the value of the 'appsecret' element to be the 'client_secret' value you just copied.
    - change the value of the 'redirectURI' element to be the 'Redirect URL' value you added in the the Box Developers App Console, above.
    - save the changes
    - click on the "Activate Policy" workflow activity in the righ-hand Activities portlet
    - ensure that the status changes to "State: Active"
- Go to Google Sheets API Hook -> Policies -> Operational Policies ->    AddAuthToken policy
    - click on the "Activate Policy" workflow activity in the righ-hand Activities portlet
    - ensure that the status changes to "State: Active"


#### Verify Connectivity
- Using curl -H "authKey:<the value authKey>" http://"URL of the Listener of your ND"/box_api_hook/helloworld
- The correct response should be a JSON object listing the details of the user owning the credentials being used to make the call:
{
    "address": "",
    "avatar_url": "https://app.box.com/api/avatar/large/229214787",
    "created_at": "2014-12-15T01:12:03-08:00",
    "id": "229214787",
    "job_title": "",
    "language": "en",
    "login": "paul.pogonoski@soa.com",
    "max_upload_size": 2147483648,
    "modified_at": "2015-03-16T21:49:50-07:00",
    "name": "Paul Pogonoski",
    "phone": "+61416101363",
    "space_amount": 10737418240,
    "space_used": 790,
    "status": "active",
    "timezone": "Australia/Sydney",
    "type": "user"
}

*Note: the authKey in the curl request, above, is retrieved by using the process in the [Dropbox 3-legged OAuth Client.pdf] (https://github.com/pogo61/Dropbox-API-Hook/blob/master/src/Dropbox%203-legged%20OAuth%20Client.pdf) file in the /src directory*


### How Hello World Works
#### An Akana Integration Primer
The Google_Sheets_API_Hook API is a "Virtual Service". That is, its interface is not that of a real service implementation. It can be a proxy to a "real" implementation, or it can be an aggregate (a combination) of a number of "real" implementations. In Policy Manager a "real" implementation is called a "Physical Service".
Apart from offering a different interface to the Physical Service, a Virtual Service offers the ability to attach Policies for security, logging, QoS, and a number of other non-functional capabilities.
Virtual Services also have the ability to have Custom Process and Scripts run before the Physical Service is called. Here is where a lot of the magic of Integration occurs.

#### Hello World
To create the helloworld operation the following was added to a base RAML document to create the box API Hook HelloWorld.raml document:  
    /helloworld:  
      &nbsp;get:  
        &nbsp;&nbsp;description: "returns details about authorised user"
        &nbsp;&nbsp;&nbsp;responses:  
          &nbsp;&nbsp;&nbsp;&nbsp;200:  
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;body:  
              &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;application/atom+xml:  

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

