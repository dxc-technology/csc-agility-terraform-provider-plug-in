package agility

import (
	"log"
	"strings"
	"time"
	"encoding/xml"
	"errors"
	"os"
	"encoding/json"
	"strconv"

	"github.com/csc/csc-agility-terraform-provider-plug-in/agility/api"
	"github.com/hashicorp/terraform/helper/schema"
)

// topology used to update topology name
type Topology struct {
	XMLName struct{}    `xml:"Topology"`
	XMLNS 		string `xml:"xmlns,attr,omitempty"`
	Name 		string `xml:"name"`
	Id 			string `xml:"id"`
	Uuid 		string `xml:"uuid"`
	AType 		AssetType `xml:"assetType"`
	Top 		string `xml:"top"`
	AssetPath 	string `xml:"assetPath"`
	DetailedAssetPath 	string `xml:"detailedAssetPath"`
	LifecycleVersion 	string `xml:"lifecycleVersion"`
	Removable 	string `xml:"removable"`
	Dom 		Domain `xml:"domain"`
	Creatr 		Creator `xml:"creator"`
	Created 	string `xml:"created"`
	LockType 	string `xml:"lockType"`
	Prent		Parent `xml:"parent"`
	Version 	string `xml:"version"`
	Pblisher	Publisher `xml:"publisher"`
	VersionStatus 		string `xml:"versionStatus"`
	CheckoutAllowed 	string `xml:"checkoutAllowed"`
	HeadAllowed 		string `xml:"headAllowed"`
	ParentProject		ParentProject `xml:"parentProject"`
	AnyOrder	AnyOrder `xml:"anyOrder"`
	AccessUriExpanded 	string `xml:"accessUriExpanded"`
	SourceBlueprint		SourceBlueprint `xml:"sourceBlueprint"`
	BlueprintSubcontainer 	string `xml:"blueprintSubcontainer"`
	ResourceAffinity 	string `xml:"resourceAffinity"`
	MandatoryAffinity 	string `xml:"mandatoryAffinity"`
	AntiAffinity 	string `xml:"antiAffinity"`
	Stats		Stats `xml:"stats"`
}

type AssetType struct {
	XMLName struct{}    `xml:"assetType"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
	Latest		string `xml:"latest,omitempty"`
}

type Domain struct {
	XMLName struct{}    `xml:"domain"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
}

type Creator struct {
	XMLName struct{}    `xml:"creator"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
}

type Parent struct {
	XMLName struct{}    `xml:"parent"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
}

type Publisher struct {
	XMLName struct{}    `xml:"publisher"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
}

type ParentProject struct {
	XMLName struct{}    `xml:"parentProject"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
}

type AnyOrder struct {
	XMLName struct{}    `xml:"anyOrder"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
}

type SourceBlueprint struct {
	XMLName struct{}    `xml:"sourceBlueprint"`
	XMLNS 		string `xml:"xmlns:xsi,attr,omitempty"`
	Name 		string `xml:"name"`
	HREF 		string `xml:"href"`
	Id 			string `xml:"id"`
	Rel 		string `xml:"rel"`
	Type 		string `xml:"type"`
	Position	string `xml:"position,omitempty"`
	Version		string `xml:"version,omitempty"`
	Latest		string `xml:"latest,omitempty"`
	VersionStatus	string `xml:"versionStatus,omitempty"`
	LockType	string `xml:"lockType,omitempty"`
}

//result of search for Templates owned by user
type AssetList struct {
	XMLName struct{}    `xml:"Assetlist"`
	XMLNS 	string `xml:"xmlns,attr,omitempty"`
	AList	[]Asset	`xml:"Asset,omitempty"`
}

type Asset struct {
	XMLName struct{}    `xml:"Asset"`
	Name  				string 		`xml:"name"`
	Id 					string		`xml:"id"`	
	Description 		string   	`xml:"description"`
	Uuid 				string   	`xml:"uuid"`
	AssetType 			AssetType	`xml:"assetType,omitempty"`
	Top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	Domain 				Domain		`xml:"domain,omitempty"`
	Creator 			Creator		`xml:"creator,omitempty"`
	Created 			string   	`xml:"created"`
	LockType 			string   	`xml:"lockType"`
	LastModified 		string  	`xml:"lastModified"`
	PolicyAssignment	[]PolicyAssignment	`xml:"policyAssignment,omitempty"`
	Parent 				Parent		`xml:"parent,omitempty"`
	NumInstances 		string  	`xml:"numInstances"`
	Topology 			Topolgy		`xml:"topology,omitempty"`
	Cloud 				Cloud		`xml:"cloud,omitempty"`
	Factory 			string  	`xml:"factory"`
	Stack 				Stack		`xml:"stack,omitempty"`
	Image 				Image		`xml:"image,omitempty"`
	Model 				string  	`xml:"model"`
	Location 			string  	`xml:"location"`
	Packages 			[]Packages	`xml:"packages,omitempty"`
	Instances 			[]Instances	`xml:"instances,omitempty"`
	Credential 			Credential	`xml:"credential,omitempty"`
	Resources 			[]Resources `xml:"resources,omitempty"`
	Onboarded 			string  	`xml:"onboarded"`
	Environment 		Environment	`xml:"environment,omitempty"`
	Project 			Project		`xml:"project,omitempty"`
	ReleaseDisks 		string  	`xml:"releaseDisks"`
	Stats 				StatsTemp		`xml:"stats,omitempty"`
}

type PolicyAssignment struct {
	XMLName struct{}    `xml:"policyAssignment"`
	Id 					string		`xml:"id"`	
	Uuid 				string   	`xml:"uuid"`
	top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	ApplyToSelf 		string   	`xml:"applyToSelf"`
	ApplyChildrenDepth 	string   	`xml:"applyChildrenDepth"`
	AllowChildrenOverride 	string  	`xml:"allowChildrenOverride"`
	Policy				Policy		`xml:"policy,omitempty"`
	PolicyTypeName 		string   	`xml:"policyTypeName"`
	ItemClass 			string   	`xml:"itemClass"`
	ItemId 				string   	`xml:"itemId"`
	ItemName 			string   	`xml:"itemName"`
}

type Policy struct {
	XMLName struct{}    `xml:"policy"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
	Version 	string   	`xml:"version,omitempty"`
	Latest 		string   	`xml:"latest,omitempty"`
	VersionStatus 	string  `xml:"versionStatus,omitempty"`
	LockType 	string   	`xml:"lockType,omitempty"`
}

type Cloud struct {
	XMLName struct{}    `xml:"cloud"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Stack struct {
	XMLName struct{}    `xml:"stack"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Packages struct {
	XMLName struct{}    `xml:"packages"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Topolgy struct {
	XMLName struct{}    `xml:"topology"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Instances struct {
	XMLName struct{}    `xml:"instances"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Credential struct {
	XMLName struct{}    `xml:"credential"`
	Name  				string 		`xml:"name"`
	Id 					string		`xml:"id"`	
	Description 		string   	`xml:"description"`
	Uuid 				string   	`xml:"uuid"`
	AssetType 			AssetType	`xml:"assetType,omitempty"`
	Top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	Domain 				Domain		`xml:"domain,omitempty"`
	Creator 			Creator		`xml:"creator,omitempty"`
	Created 			string   	`xml:"created"`
	LockType 			string   	`xml:"lockType"`
	LastModified 		string  	`xml:"lastModified,omitempty"`
	PolicyAssignment	[]PolicyAssignment	`xml:"policyAssignment,omitempty"`
	Parent 				Parent		`xml:"parent,omitempty"`
	CredentialType 		string  	`xml:"credentialType"`
	PublicKey 			string		`xml:"publicKey,omitempty"`
	PrivateKey 			string		`xml:"privateKey,omitempty"`
	Encrypted 			string		`xml:"encrypted,omitempty"`
	Cloud 				Cloud		`xml:"cloud,omitempty"`
	Factory 			string  	`xml:"factory"`
	Stack 				Stack		`xml:"stack,omitempty"`
	Image 				Image		`xml:"image,omitempty"`
	Model 				string  	`xml:"model"`
	Location 			string  	`xml:"location"`
	Packages 			[]Packages	`xml:"packages,omitempty"`
	Instances 			[]Instances	`xml:"instances,omitempty"`
}

type Resources struct {
	XMLName struct{}    `xml:"resources"`
	Name  				string 		`xml:"name"`
	Id 					string		`xml:"id"`	
	Description 		string   	`xml:"description,omitempty"`
	Uuid 				string   	`xml:"uuid"`
	AssetType 			AssetType	`xml:"assetType,omitempty"`
	Top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	ResourceType 		string   	`xml:"resourceType"`
	Units 				string   	`xml:"units,omitempty"`
	Quantity 			string   	`xml:"quantity,omitempty"`
	HostResource 		string   	`xml:"hostResource,omitempty"`
	Address 			Address		`xml:"address,omitempty"`
	Network 			Network		`xml:"network,omitempty"`
}

type Address struct {
	XMLName struct{}    `xml:"address"`
	Name  				string 		`xml:"name"`
	Id 					string		`xml:"id"`	
	Description 		string   	`xml:"description"`
	Uuid 				string   	`xml:"uuid"`
	AssetType 			AssetType	`xml:"assetType,omitempty"`
	Top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	lockType 			string   	`xml:"lockType"`
	Addr 				string   	`xml:"address,omitempty"`
	InetAddr 			string   	`xml:"inetAddr,omitempty"`
	Inst 				Inst		`xml:"instance,omitempty"`
	Elastic 			string   	`xml:"elastic,omitempty"`
}

type Inst struct {
	XMLName struct{}    `xml:"instance"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Network struct {
	XMLName struct{}    `xml:"network"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Environment struct {
	XMLName struct{}    `xml:"environment"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Project struct {
	XMLName struct{}    `xml:"project"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type StatsTemp struct {
	XMLName struct{}    `xml:"stats"`
	TemplateId  	string 		`xml:"templateId"`		
	NumUnknown 		string   	`xml:"numUnknown"`
	NumStarting 	string   	`xml:"numStarting"`
	NumRunning 		string   	`xml:"numRunning"`
	NumPaused 		string   	`xml:"numPaused"`
	NumStopping 	string   	`xml:"numStopping"`
	NumStopped 		string   	`xml:"numStopped"`
	NumDestroyed 	string   	`xml:"numDestroyed"`
	NumFailed 		string   	`xml:"numFailed"`
	NumDegraded 	string   	`xml:"numDegraded"`
	NumInstances 	string   	`xml:"numInstances"`
	MinCount 		string   	`xml:"minCount"`
	MaxCount 		string   	`xml:"maxCount"`
	Status 			string   	`xml:"status"`
}

type Stats struct {
	XMLName struct{}    `xml:"stats"`
	TemplateId  	string 		`xml:"topologyId"`		
	NumUnknown 		string   	`xml:"numUnknown"`
	NumStarting 	string   	`xml:"numStarting"`
	NumRunning 		string   	`xml:"numRunning"`
	NumPaused 		string   	`xml:"numPaused"`
	NumStopping 	string   	`xml:"numStopping"`
	NumStopped 		string   	`xml:"numStopped"`
	NumDestroyed 	string   	`xml:"numDestroyed"`
	NumFailed 		string   	`xml:"numFailed"`
	NumDegraded 	string   	`xml:"numDegraded"`
	NumInstances 	string   	`xml:"numInstances"`
	MinCount 		string   	`xml:"minCount"`
	MaxCount 		string   	`xml:"maxCount"`
	Status 			string   	`xml:"status"`
}

//Used in uodating and Instance name
type Instance struct {
	XMLName struct{}    `xml:"Instance"`
	XMLNS 				string `xml:"xmlns,attr,omitempty"`
	Name  				string 		`xml:"name"`
	Id 					string		`xml:"id"`	
	Description 		string   	`xml:"description"`
	Uuid 				string   	`xml:"uuid"`
	AssetType 			AssetType	`xml:"assetType,omitempty"`
	Top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	Cloud 				Cloud		`xml:"cloud,omitempty"`
	Stack 				Stack		`xml:"stack,omitempty"`
	Template 			Template	`xml:"template,omitempty"`
	State 				string  	`xml:"state"`
	InstanceId 			string  	`xml:"instanceId"`
	PublicAddress 		string  	`xml:"publicAddress"`
	PrivateAddress 		string  	`xml:"privateAddress"`
	Credentials			[]Credential	`xml:"credential,omitempty"`
	Hostname 			string  	`xml:"hostname"`
	CanonicalName 		string  	`xml:"canonicalName"`
	Environment 		Environment	`xml:"environment,omitempty"`
	Properties 			[]Properties	`xml:"properties,omitempty"`
	Resources 			[]Resources	`xml:"resources,omitempty"`
	Addressess 			[]Addresses	`xml:"addresses,omitempty"`
	ScriptstatusLink 	[]ScriptstatusLink	`xml:"scriptstatusLink,omitempty"`
	onboarded 			string  	`xml:"onboarded"`
	startTime 			string  	`xml:"startTime"`
	createdOn 			string  	`xml:"createdOn"`
	lastUpdate 			string  	`xml:"lastUpdate"`
	lastStartedOrStoppedBy 	string  	`xml:"lastStartedOrStoppedBy"`
	Location 			Location	`xml:"location,omitempty"`
	Model 				Model	`xml:"model,omitempty"`
	Image 				Image	`xml:"image,omitempty"`
	Pending 			string  	`xml:"pending"`
}

type Location struct {
	XMLName struct{}    `xml:"location"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Model struct {
	XMLName struct{}    `xml:"model"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Image struct {
	XMLName struct{}    `xml:"image"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Template struct {
	XMLName struct{}    `xml:"template"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Credentials struct {
	XMLName struct{}    `xml:"credential"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type Properties struct {
	XMLName struct{}    `xml:"properties"`
	Name  				string 		`xml:"name"`
	Id 					string		`xml:"id"`	
	Description 		string   	`xml:"description"`
	Uuid 				string   	`xml:"uuid"`
	AssetType 			AssetType	`xml:"assetType,omitempty"`
	Top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	Value 				string   	`xml:"value"`
	Overridable 		string   	`xml:"overridable"`
	Encrypted 			string   	`xml:"encrypted"`
	DataEncrypted 		string   	`xml:"dataEncrypted"`
}

type Addresses struct {
	XMLName struct{}    `xml:"addresses"`
	Name  				string 		`xml:"name"`
	Id 					string		`xml:"id"`	
	Description 		string   	`xml:"description,omitempty"`
	Uuid 				string   	`xml:"uuid"`
	AssetType 			AssetType	`xml:"assetType,omitempty"`
	Top 				string   	`xml:"top"`
	AssetPath 			string   	`xml:"assetPath"`
	DetailedAssetPath 	string   	`xml:"detailedAssetPath"`
	LifecycleVersion 	string   	`xml:"lifecycleVersion"`
	Removable 			string   	`xml:"removable"`
	lockType 			string   	`xml:"lockType"`
	Addr 				string   	`xml:"address,omitempty"`
	InetAddr 			string   	`xml:"inetAddr,omitempty"`
	Inst 				Inst		`xml:"instance,omitempty"`
	Elastic 			string   	`xml:"elastic,omitempty"`
}

type ScriptstatusLink struct {
	XMLName struct{}    `xml:"scriptstatusLink"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type,omitempty"`
	Position 	string   	`xml:"position,omitempty"`
}

type ProvCredentials struct {
	UserName  	string	
	Password 	string
}

var credentials ProvCredentials

func resourceAgilityCompute() *schema.Resource {

	return &schema.Resource{
		Create: resourceAgilityComputeCreate,
		Read:   resourceAgilityComputeRead,
		Update: resourceAgilityComputeUpdate,
		Delete: resourceAgilityComputeDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			"active": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				Computed: 	false,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional:	true,
				ForceNew:	true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional:	true,
				ForceNew:	true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"environment": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"blueprint": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
			},
			"EnvironmentId": &schema.Schema{
				Type:     	schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"ProjectId": &schema.Schema{
				Type:     	schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"BlueprintId": &schema.Schema{
				Type:     schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"TopologyId": &schema.Schema{
				Type:     schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"CreatedStopped": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: 	true,
				Computed:	true,
			},
		},
	}
}

func init(){
    file, err1 := os.Open("conf.json")
    if err1 != nil {
        log.Println("error:", err1)
    }
    decoder := json.NewDecoder(file)
    configuration = Config{}
    err := decoder.Decode(&configuration)
    if err != nil {
        log.Println("error:", err)
    }

    err2 := file.Close()
    log.Printf("err2: %v\n", err2)
}

func resourceAgilityComputeCreate(ResourceData *schema.ResourceData, meta interface{}) error {
	//set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

    // set the credentials from the Provider initilisation
    credentials = meta.(ProvCredentials)


    errProj := checkProject(ResourceData)
    log.Println("errProj is: ", errProj)
    if errProj == nil {
    	log.Println("ProjectId is: ", ResourceData.Get("ProjectId"))
    	errEnv := checkEnvironment(ResourceData)
    	if errEnv == nil {
    		log.Println("EnvironmentId is: ", ResourceData.Get("EnvironmentId"))
	    	errBP := deployBlueprint(ResourceData)
	    	log.Println("BlueprintId is: ", ResourceData.Get("BlueprintId"))
	    	if errBP == nil {
	    		//if the active resource variable is set in the .tf file then rename the topology and
			    // start it. Otherwise just rename the tolpology. If this topology is created but not started also 
			    // set the CreatedStopped resource variable for correct update processing
				if ResourceData.Get("active").(string) == "true" {
					topologyId := ResourceData.Get("TopologyId").(string)
					UpdateTopologyName(ResourceData,topologyId)
					StartTopology(ResourceData,topologyId)
					ResourceData.SetId(topologyId)
					errUIN := UpdateInstanceName(ResourceData,topologyId)
					if errUIN != nil {
						return errUIN
					}
				} else {
					topologyId := ResourceData.Get("TopologyId").(string)
					UpdateTopologyName(ResourceData,topologyId)
					ResourceData.Set("CreatedStopped", true)
					ResourceData.SetId(topologyId)
				}
	    	} else {
	    		return errBP
	    	}	
    	} else {
    		return errEnv
    	}
    } else {
    	log.Println("ERROR! errProj is: ", errProj)
    	return errProj
    }

	return nil
}

func resourceAgilityComputeRead(d *schema.ResourceData, meta interface{}) error {
	// no need to do anything for read state

	return nil
}

func resourceAgilityComputeUpdate(d *schema.ResourceData, meta interface{}) error {
	//set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

    // set the credentials from the Provider initilisation
    credentials = meta.(ProvCredentials)

    // if the active resource variable is changed to 'true' then start the topology
    // if this is the first time this has happend after the creation of the topology (createdStopped is true)
    // then also change the name(s) of the instances started
    // Otherwise stop the topology and all its instances
    if d.HasChange("active") {
        if d.Get("active").(string) == "true" {
			topologyId := d.Get("TopologyId").(string)
			log.Println("Starting the Topology")
			StartTopology(d,topologyId)
			d.SetId(topologyId)
			createdStopped := d.Get("CreatedStopped")
			if createdStopped.(bool) {
				errUIN := UpdateInstanceName(d,topologyId)
				if errUIN != nil {
					return errUIN
				}
				d.Set("CreatedStopped", false)
			}
		} else if d.Get("active").(string) == "false" {
			topologyId := d.Get("TopologyId").(string)
			log.Println("Stoping the Topology")
			StopTopology(d,topologyId)
			d.SetId(topologyId)
		}
    }

    // if the name resource variable has changed then change the topology name
    // and the name of the instance(s)
    if d.HasChange("name") {
	topologyId := d.Get("TopologyId").(string)
	UpdateTopologyName(d,topologyId)
	d.SetId(topologyId)
	errUIN := UpdateInstanceName(d,topologyId)
	if errUIN != nil {
		return errUIN
	}
    }
	
	return nil
}

func resourceAgilityComputeDelete(d *schema.ResourceData, meta interface{}) error {
	//Set up logging
    f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

    // set the credentials from the Provider initilisation
    credentials = meta.(ProvCredentials)

	// call the API to delete the topology identified by the ID if this resource.
	// this will take a while, so call the GeTaksStatus function to loop and wait until it completes
    log.Println("Starting the Topology Destroy")
    log.Println("TopologyId is: ", d.Id())
	response := api.DestroyTopology(d.Id(), credentials.UserName, credentials.Password)

	if response != nil {
		r := strings.NewReader(string(response))
		decoder := xml.NewDecoder(r)
		finish := false
		for {
			// Read tokens from the XML document in a stream.
			t, _ := decoder.Token()
			if t == nil {
				break
			}
			if finish {
	            break
	        }
			// parse the result and loop for and of created task
			switch Element := t.(type) {
			case xml.StartElement:
				if Element.Name.Local == "Task" {
					log.Println("Element name is : ", Element.Name.Local)

					var q DeployTask
					err := decoder.DecodeElement(&q, &Element)
					if err != nil {
						log.Println(err)
					}

					log.Println("status value is :", q.Status)
					if q.Status == "Pending" {
						GetTaskStatus(d,q.Id)
						d.SetId("")
						finish = true
					}
				}
			default:
			}
		}
		return nil
	} else {
		return errors.New("error is deleting the compute resource")
	}
	
}

func StartTopology(d *schema.ResourceData,id string) {
	//set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

    log.Println("Starting Topology "+id)

    // call the Agility API to start the topology with the ID passed in
	response := api.StartTopology(id, credentials.UserName, credentials.Password)

	log.Println("\n response is:",string(response))

	//if the result was successful then then get the Id of the task and 
	// call the GetTaskStatus that loops until the start task completes
	r := strings.NewReader(string(response))
	decoder := xml.NewDecoder(r)
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "id" {
				log.Println("Element name is : ", Element.Name.Local)

				q := new(Id)
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				log.Println("Element value is :", string(q.Value))
				if q.Value != nil {
					GetTaskStatus(d,string(q.Value))
				}
			}
		default:
		}
	}
}

func StopTopology(d *schema.ResourceData,id string) {
	//set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

	response := api.StopTopology(id, credentials.UserName, credentials.Password)

	log.Println("\n response is:",string(response))

	//if the result was successful then then get the Id of the task and 
	// call the GetTaskStatus that loops until the start task completes
	r := strings.NewReader(string(response))
	decoder := xml.NewDecoder(r)
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch Element := t.(type) {
		case xml.StartElement:
			if Element.Name.Local == "id" {
				log.Println("Element name is : ", Element.Name.Local)

				q := new(Id)
				err := decoder.DecodeElement(&q, &Element)
				if err != nil {
					log.Println(err)
				}

				log.Println("Element value is :", string(q.Value))
				if q.Value != nil {
					GetTaskStatus(d,string(q.Value))
				}
			}
		default:
		}
	}
}

func UpdateTopologyName(d *schema.ResourceData, topologyId string) {
	//set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

    var tp Topology
    var finished = false
	// call the Agility API to get the topology details with the ID passed in
	statusResponse := api.GetTopologyDetail(topologyId, credentials.UserName, credentials.Password)

	//if the result was successful then then parse the resultant XML
	// looking for the Topology element. When found UnMarshal it into an
	//  temporary instance of a Topology structure
	sr := strings.NewReader(string(statusResponse))
	decoder := xml.NewDecoder(sr)
	for {
		if finished == true {
			break
		}
		// Read tokens from the XML document in a stream.
		st, _ := decoder.Token()
		if st == nil {
			break
		}

		switch Element := st.(type) {
			case xml.StartElement:
				if Element.Name.Local == "Topology" {
					log.Println("Element name is : ", Element.Name.Local)
					err := decoder.DecodeElement(&tp, &Element)
					if err != nil {
						log.Println(err)
						finished = true
						break 
					}
				}
			default:
		}
	}

	// change the name on the Topology and marshall into XML
	tp.XMLNS = "http://servicemesh.com/agility/api"
	tp.Name = d.Get("name").(string)
	// Marshall this struct into XML
	xmlStr, err := xml.MarshalIndent(tp, "", "    ")
	if err != nil {
		log.Printf("error: %v\n", err)													
	}

	update := xml.Header + string(xmlStr)

	// update the topology changing it's name
	api.UpdateTopology(topologyId, update, credentials.UserName, credentials.Password)

	//log.Println("\n response is:",string(response))

}

func UpdateInstanceName(d *schema.ResourceData, topologyId string) error {
	//set up logging
	f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

	var a AssetList
	var inst Instance
    var finished = false
    var instanceUpdated = false
    var instanceId string

    log.Println("Updating Instance name ")

    //Call the API to search for templates owned by the user defined in the conf.json file
	statusResponse := api.SearchTemplates(credentials.UserName, credentials.UserName, credentials.Password)
	
	err := xml.Unmarshal(statusResponse, &a)
	if err != nil {
		log.Println(err)
		return err
	}
	// if the result had templates, then get the number of templates 
	// and loop through them. When the Topology ID in the template matches the 
	// topology that been created as the Compute resource, then get the detail 
	// of that topology and look for it's parent 
	// (The template is usually linked to the 'Cloud Afility', not the root topology)
	log.Println("length of a.AList[i] is: ", len(a.AList))
	log.Println("capacity of a.AList[i] is: ", cap(a.AList))
	if len(a.AList) > 0 {
		var parent Parent
		for i := 0; i < len(a.AList); i++ {
			if finished == true {
				break
			}
			instanceUpdated = false
			log.Println("a.AList[i].Name is: ", a.AList[i].Name)
			log.Println("topologyId is: ", topologyId)
			log.Println("a.AList[i].Topology.Id is: ", a.AList[i].Topology.Id)
			if topologyId != a.AList[i].Topology.Id {
				statusResponse := api.GetTopologyDetail(a.AList[i].Topology.Id, credentials.UserName, credentials.Password)
				sr := strings.NewReader(string(statusResponse))
				decoder := xml.NewDecoder(sr)
				for {
					if finished == true {
						break
					}
					// Read tokens from the XML document in a stream.
					st, _ := decoder.Token()
					if st == nil {
						break
					}

					switch Element := st.(type) {
						case xml.StartElement:
							if Element.Name.Local == "parent" {
								log.Println("Element name is : ", Element.Name.Local)
								err := decoder.DecodeElement(&parent, &Element)
								if err != nil {
									log.Println(err)
									finished = true
									break 
								}
							}
						default:
					}
				}
				// if there was a parent (a 'Cloud Afinity' was defined) then get its instance
				// and create instance struct, change it's name and call the API to update it
				log.Println("Topology Parent ID is : ", parent.Id)
				if parent.Id == topologyId {
					log.Println("len(a.AList[i].Instances) is : ", len(a.AList[i].Instances))
					for j := 0; j < len(a.AList[i].Instances); j++ {
						if instanceUpdated == true {
							break
						}
						instanceId = a.AList[i].Instances[j].Id
						log.Println("instanceId is : ", instanceId)
						log.Println("Getting the Instance. Instance Id is: ",instanceId)
						statusResponse = api.GetInstanceDetail(instanceId, credentials.UserName, credentials.Password)

						//unmarshall the XML result into a temp struct
						err := xml.Unmarshal(statusResponse, &inst)
						if err != nil {
							log.Println(err)
							return err
						}

						//update struct with name
						inst.Name = d.Get("name").(string)+"-"+strconv.Itoa(j+1)
						inst.XMLNS =  "http://servicemesh.com/agility/api"
						//marshall into XML
						xmlStr, err := xml.MarshalIndent(inst, "", "    ")
						if err != nil {
							log.Printf("error: %v\n", err)													
						}

						update := xml.Header + string(xmlStr)
						log.Println("updated instance is:",update)

						// call the API to update to change the name
						response := api.UpdateInstance(instanceId, update, credentials.UserName, credentials.Password)

						log.Println("\n response is:",string(response))
						instanceUpdated = true
					}
				}
			} else {
				// if there ins't a parent (No 'Cloud Afinity' was defined) then get the original topology's instance
				// and create instance struct, change it's name and call the API to update it
				if a.AList[i].Stats.NumInstances !=  "0" {
					if a.AList[i].Instances != nil && len(a.AList[i].Instances) > 0 {
						instanceId = a.AList[i].Instances[0].Id
						log.Println("instanceId is : ", instanceId)
						log.Println("len(a.AList[i].Instances) is : ", len(a.AList[i].Instances))
						for j := 0; j < len(a.AList[i].Instances); j++ {
							if finished == true {
								break
							}
							instanceId = a.AList[i].Instances[j].Id
							log.Println("instanceId is : ", instanceId)
							log.Println("Getting the Instance. Instance Id is: ",instanceId)
							statusResponse = api.GetInstanceDetail(instanceId, credentials.UserName, credentials.Password)

							//unmarshall the XML result into a temp struct
							err := xml.Unmarshal(statusResponse, &inst)
							if err != nil {
								log.Println(err)
								return err
							}

							//update struct with name
							inst.Name = d.Get("name").(string)+"-"+strconv.Itoa(j+1)
							inst.XMLNS =  "http://servicemesh.com/agility/api"
							//marshall into XML
							xmlStr, err := xml.MarshalIndent(inst, "", "    ")
							if err != nil {
								log.Printf("error: %v\n", err)
							}

							update := xml.Header + string(xmlStr)
							log.Println("updated instance is:",update)

							// call the API to update to change the name
							response := api.UpdateInstance(instanceId, update, credentials.UserName, credentials.Password)

							log.Println("\n response is:",string(response))
							instanceUpdated = true
						}
					} else {
						return errors.New("UpdateInstanceName: template contained no instances")
					}
				} else {
					return errors.New("UpdateInstanceName: template in Error State, please check while it contained no instances")
				}
			}
			
		}
	} else {
		return errors.New("There was valid match for Search for instances")
	}

	return nil
}

func GetTaskStatus(d *schema.ResourceData,id string) error {

	log.Println("\n id is:",id)
	var finished bool = false
	//it might take a while for the Pending status to appear, so keep trying
	for {
		if finished == true {
			break
		}
		//get the ststus of the task
		statusResponse := api.GetTaskStatus(id, credentials.UserName, credentials.Password)

		log.Println("\n response is:",string(statusResponse))
		sr := strings.NewReader(string(statusResponse))
		decoder := xml.NewDecoder(sr)
		for {
			// Read tokens from the XML document in a stream.
			st, _ := decoder.Token()
			if st == nil {
				break
			}
			// Inspect the type of the token just read.
			switch Element := st.(type) {
			case xml.StartElement:
				//if the element of <result> exists then
				if Element.Name.Local == "result" {
					log.Println("Element name is : ", Element.Name.Local)

					var q Result
					err := decoder.DecodeElement(&q, &Element)
					if err != nil {
						log.Println(err)
					}

					log.Println("Topology name is :", q.ResultName)
					log.Println("Topology id is :", q.Id)
					//the task is finsihed, so save the Topology id
					d.Set("TopologyId",q.Id)
					d.Set("TopologyName",q.ResultName)
					d.Set("Status","started")
					d.SetId(q.Id)
					finished = true
				}
				// topology tasks have differnt results, so look for the <status> element
				if Element.Name.Local == "status" {
					log.Println("Element name is : ", Element.Name.Local)

					var q Status
					err := decoder.DecodeElement(&q, &Element)
					if err != nil {
						log.Println(err)
					}

					//if any of these values then we are finished, otherwise sleep for 20 secs and look again
					log.Println("Status value is :", string(q.Value))
					if string(q.Value) == "Completed Topology Start" {
						finished = true
					} else if string(q.Value) == "Completed Topology Stop" {
						finished = true
					} else if string(q.Value) == "Completed" {
						//time.Sleep(20000 * time.Millisecond)
						finished = true
					} else if string(q.Value)[:6] == "Delete" {
						finished = true
					} else if len(string(q.Value)) > 26 {
						if string(q.Value)[:26] == "Unable to deploy blueprint" {
							return errors.New(string(q.Value))
						} else {
							log.Println("Status value is :", string(q.Value))
							time.Sleep(20000 * time.Millisecond)
						}
					} else {
						log.Println("Status value is :", string(q.Value))
						time.Sleep(20000 * time.Millisecond)
					}
				}
			default:
			} 
		}
	}
	return nil
}
