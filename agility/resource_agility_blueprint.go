package agility

import (
	"fmt"
	"log"
	"strings"
	"encoding/xml"
	"errors"
	"os"
	//"encoding/json"

	"github.com/csc/csc-agility-terraform-provider-plug-in/agility/api"
	"github.com/hashicorp/terraform/helper/schema"
)

type DeployTask struct {
	XMLName struct{}    `xml:Task"`
	Id  	string 		`xml:"id"`
	Name  	string 		`xml:"name"`
	Type  	string 		`xml:"type"`
	Status  string 		`xml:"status"`
	StartedTime  	string 		`xml:"startedTime"`
	User  	string 		`xml:"user"`
	Completed  	string 		`xml:"completed"`
	Cancelled  	string 		`xml:"cancelled"`
	Visible  	string 		`xml:"visible"`
	Progress  	string 		`xml:"progress"`
	HasDependency  	string 		`xml:"hasDependency"`
	Cancellable  	string 		`xml:"cancellable"`
}

type Result struct {
	XMLName struct{}    `xml:result"`
	ResultName  	string 		`xml:"name"`
	Href			string   	`xml:"href"`
	Id 				string   	`xml:"id"`
	Type 			string   	`xml:"type"`
}

type ReadResult struct {
	XMLName struct{}    `xml:stats"`
	Id  		string 		`xml:"topologyId"`
	Unknown		string   	`xml:"numUnknown"`
	Starting 	string   	`xml:"numStarting"`
	Running 	string   	`xml:"numRunning"`
	Paused 		string   	`xml:"numPaused"`
	Stopping 	string   	`xml:"numStopping"`
	Stopped 	string   	`xml:"numStopped"`
	Destroyed 	string   	`xml:"numDestroyed"`
	Failed 		string   	`xml:"numFailed"`
	Degraded 	string   	`xml:"numDegraded"`
	Instances 	string   	`xml:"numInstances"`
	Templates 	string   	`xml:"numTemplates"`
	Min 		string   	`xml:"minCount"`
	Max 		string   	`xml:"maxCount"`
	Status 		string   	`xml:"status"`
}

type DeploymentPlan struct {
	XMLName struct{}    `xml:"DeploymentPlan"`
	XMLNS 				string `xml:"xmlns,attr,omitempty"`
	Item 				DPOptionItem	`xml:"item"`
	Rank 				string   	`xml:"rank"`
	Error 				string   	`xml:"error"`
	ResourceAffinity 	string   	`xml:"resourceAffinity,omitempty"`
	ResourceList 		[]DPOptionResource	`xml:"resource,omitempty"`
	ChildList			[]DPOptionChild	`xml:"child,omitempty"`
}

type DPOption struct {
	XMLName struct{}    `xml:"option"`
	//XMLNS 				string 		`xml:"xmlns:ns1=http://servicemesh.com/agility/api,attr"`
	Name  				string 		`xml:"name"`
	Item 				DPOptionItem	`xml:"item"`	
	Rank 				string   	`xml:"rank"`
	Error 				string   	`xml:"error"`
	ResourceAffinity 	string   	`xml:"resourceAffinity,omitempty"`
	ResourceList 		[]DPOptionResource	`xml:"resource,omitempty"`
	ChildList			[]DPOptionChild	`xml:"child,omitempty"`
}

type DPChildOption struct {
	XMLName struct{}    `xml:"option"`
	//XMLNS 				string 		`xml:"xmlns:ns1=http://servicemesh.com/agility/api,attr"`
	Name  				string 		`xml:"name"`
	Item 				DPOptionItem	`xml:"item"`	
	Rank 				string   	`xml:"rank"`
	Error 				string   	`xml:"error"`
	ResourceAffinity 	string   	`xml:"resourceAffinity,omitempty"`
	ResourceList 		[]DPOptionResource	`xml:"resource,omitempty"`
	ChildList			[]DPOptionChild	`xml:"child,omitempty"`
}

type DPOptionItem struct {
	XMLName struct{}    `xml:ns1:item"`
	//XMLNS 		string 		`xml:"xmlns:ns1=http://servicemesh.com/agility/api,attr"`
	Name  		string 		`xml:"name"`		
	HREF 		string   	`xml:"href"`
	Id 			string   	`xml:"id"`
	Rel 		string   	`xml:"rel,omitempty"`
	Type 		string   	`xml:"type"`
}

type DPOptionResource struct {
	XMLName struct{}    `xml:resource"`
	//XMLNS 		string 		`xml:"xmlns:ns1=http://servicemesh.com/agility/api,attr"`
	Name  			string 		`xml:"name"`		
	HREF 			string   	`xml:"href"`
	Id 				string   	`xml:"id"`
	Type 			string   	`xml:"type"`
	Description		string   	`xml:"description,omitempty"`
	Rank 			string   	`xml:"rank,omitempty"`
	LinkProperty	[]DPlinkProperty	`xml:"linkProperty,omitempty"`
	MetricList		[]DPOptionMetric	`xml:"metric,omitempty"`
}

type DPOptionMetric struct {
	XMLName struct{}    `xml:metric"`
	//XMLNS 		string 		`xml:"xmlns:ns1=http://servicemesh.com/agility/api,attr"`
	Type  		string 		`xml:"type"`		
	Quantity 	string   	`xml:"quantity"`
	Capacity 	string   	`xml:"capacity"`
}

type DPlinkProperty struct {
	XMLName struct{}    `xml:linkProperty"`
	//XMLNS 		string 		`xml:"xmlns:ns1=http://servicemesh.com/agility/api,attr"`
	Name  	string 		`xml:"name"`		
	Value 	string   	`xml:"value"`
}

type DPOptionChild struct {
	XMLName struct{}    `xml:child"`
	//XMLNS 				string 		`xml:"xmlns:ns1=http://servicemesh.com/agility/api,attr"`
	Name  				string 		`xml:"name,omitempty"`
	Item 				DPOptionItem	`xml:"item"`
	Rank 				string   	`xml:"rank"`
	Error 				string   	`xml:"error"`
	ResourceAffinity 	string   	`xml:"resourceAffinity"`
	ResourceList 		[]DPOptionResource	`xml:"resource,omitempty"`
	OptionList			[]DPChildOption	`xml:"option,omitempty"`
}

type Id struct {
	Value 	[]byte `xml:",chardata"`
}

type Status struct {
	Value 	[]byte `xml:",chardata"`
}

/*func resourceAgilityBlueprint() *schema.Resource {
	// Our schema is shared also with aws_ami_copy and aws_ami_from_instance
	//resourceSchema := resourceAwsAmiCommonSchema(false)

	return &schema.Resource{
		Create: resourceAgilityBlueprintDeploy,
		Read:   resourceAgilityBlueprintRead,
		Delete: resourceAgilityBlueprintDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: 	true,
				ForceNew:	true,
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
			"EnvironmentId": &schema.Schema{
				Type:     	schema.TypeString,
				Required:	true,
				ForceNew:	true,
			},
			"ProjectId": &schema.Schema{
				Type:     	schema.TypeString,
				Required:	true,
				ForceNew:	true,
			},
			"TopologyId": &schema.Schema{
				Type:     schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
			"BlueprintId": &schema.Schema{
				Type:     schema.TypeString,
				Computed: 	true,
				ForceNew:	true,
			},
		},
	}
}*/

type Config struct {
    MaxRetries 	string
    APIURL     	string
    AWSXS  		string
    AWSS  		string
    AWSM  		string
    AWSL  		string
    AWSXL  		string
    AWSXXL 		string
    BCXS    	string
    BCS    		string
    BCM    		string
    BCL    		string
    BCXL    	string
    BCXXL    	string
    AWSCloud 	string
    BizCloud	string
}

var configuration Config

//Deleted as no longer used as a Terraform Resource
/*func init(){
    //get the confing file and load the config variables
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

func resourceAgilityBlueprintDeploy(d *schema.ResourceData, meta interface{}) error {
	//Set up logging
    f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

	// get the value of the Project ID created by the creation of the Project Resource.
	// if that is OK as well as the Name, the get the right bluerint ID, depending on whether 
	// a Version number was supplied of not
	projectId, ok_projectId := d.GetOk("ProjectId")
	log.Println("Project ID is : ", projectId)
	var blueprintId string
	if ok_projectId {
		blueprintName, ok_blueprintName := d.GetOk("name")
		log.Println("Blueprint name is : ", blueprintName.(string))
		if ok_blueprintName {
			version, ok_version := d.GetOk("version")
			if ok_version {
				response, err := api.GetBlueprintIdForVersion(blueprintName.(string), projectId.(string), version.(string))
				if err != nil {
					return err
				}
				blueprintId = response
			} else {
				response, err := api.GetBlueprintId(blueprintName.(string), projectId.(string))
				if err != nil {
					return err
				}
				blueprintId = response
			}

			log.Println("BlueprintId is : ", blueprintId)
			if blueprintId != "" {
				d.Set("BlueprintId",blueprintId)
				d.SetId(blueprintId)
			} else {
				return fmt.Errorf("The blueprint does not have that version")
			}
			
		} else {
			log.Println("No blueprintName was provided")
			return fmt.Errorf("No blueprintName was provided")
		}
	} else {
		log.Println("No ProjectId was provided")
		return fmt.Errorf("No ProjectId was provided")
	}

	environmentId, ok_environmentId := d.GetOk("EnvironmentId")

	//if the environment ID was provided then get the right deployment plan for the blueprint ID
	// then deploy the blueprint.
	// this will take some time so call GetTaskStatus for it to check when this is done
	if ok_environmentId {
		blueprintId, ok_blueprintId := d.GetOk("BlueprintId")
		log.Println("BlueprintId is : ", blueprintId.(string))
		if ok_blueprintId {
			log.Println("BlueprintId is : ", blueprintId.(string))
			compType, ok_type:= d.GetOk("type")
			if ok_type {
				log.Println("type is : ", compType.(string))
				var deploymentPlan = "No Plan"
				var err error
				// based on the blueprint, the environment, and the size (if supplied), ge the right 
				// deployment plan
				deploymentPlan, err = GetDeploymentPlan(d, blueprintId.(string), environmentId.(string))
				if err != nil {
					return err
				}
				f1, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
			    if errf != nil {
			        log.Println("error opening file: ", errf)
			    }
			    defer f1.Close()

			    log.SetOutput(f1)

				log.Println("Determined the Deployment Plan ")
				log.Println("deploymentPlan is : ", deploymentPlan)
				// when the deployment plan is supplied then start the bluerint using that plan
				if deploymentPlan != "No Plan" {
					deployResponse := api.DeploymentPlanBlueprintDeploy(blueprintId.(string), environmentId.(string), deploymentPlan)
					r := strings.NewReader(string(deployResponse))
					log.Println("Deploy response is : ", r)
					decoder := xml.NewDecoder(r)
					finish := false
					for {
						if finish {
	            			break
	        			}
						// Read tokens from the XML document in a stream.
						t, _ := decoder.Token()
						if t == nil {
							break
						}
						
						// Inspect the type of the token just read.
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
										topologyId := d.Get("TopologyId").(string)
										d.SetId(topologyId)
										finish = true
									}
								}
							default:
						}
					} 
				} else {
					return errors.New("There is no Deployment Plan for the Blueprint in that environment")
				}
			} else {
				// if the type (size of machine) is not supplied then just deploy th eblueprint using the default deployment plan
				// that Agility creates
				log.Println("BlueprintId is : ", blueprintId.(string))
				deployResponse := api.SimpleBlueprintDeploy(blueprintId.(string), environmentId.(string))
				r := strings.NewReader(string(deployResponse))
				log.Println("Deploy response is : ", r)
				decoder := xml.NewDecoder(r)
				finish := false
				for {
					if finish {
            			break
        			}
					// Read tokens from the XML document in a stream.
					t, _ := decoder.Token()
					if t == nil {
						break
					}
					
					// Inspect the type of the token just read.
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
									topologyId := d.Get("TopologyId").(string)
									d.SetId(topologyId)
									finish = true
								}
							}
						default:
					}
				}	
			}
		} else {
			log.Println("No blueprintId was provided")
			return errors.New("No blueprintId was provided")
		}
	} else {
		log.Println("No EnvironmentId was provided")
		return errors.New("No EnvironmentId was provided")
	}
	
	return nil
}*/

func deployBlueprint(d *schema.ResourceData) error {
	//Set up logging
    f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

	// get the value of the Project ID created by the creation of the Project Resource.
	// if that is OK as well as the Name, the get the right bluerint ID, depending on whether 
	// a Version number was supplied of not
	projectId, ok_projectId := d.GetOk("ProjectId")
	log.Println("Project ID is : ", projectId)
	var blueprintId string
	if ok_projectId {
		blueprintName, ok_blueprintName := d.GetOk("blueprint")
		log.Println("Blueprint name is : ", blueprintName.(string))
		log.Println("username for blueprint creation is : ", credentials.UserName)
		//log.Println("password for blueprint creation is : ", credentials.Password)
		log.Println("Blueprint name is : ", blueprintName.(string))
		if ok_blueprintName {
			version, ok_version := d.GetOk("version")
			if ok_version {
				response, err := api.GetBlueprintIdForVersion(blueprintName.(string), projectId.(string), version.(string), credentials.UserName, credentials.Password)
				if err != nil {
					return err
				}
				blueprintId = response
			} else {
				response, err := api.GetBlueprintId(blueprintName.(string), projectId.(string), credentials.UserName, credentials.Password)
				if err != nil {
					return err
				}
				blueprintId = response
			}

			log.Println("BlueprintId is : ", blueprintId)
			if blueprintId != "" {
				d.Set("BlueprintId",blueprintId)
				//d.SetId(blueprintId)
			} else {
				return fmt.Errorf("The blueprint does not have that version")
			}
			
		} else {
			log.Println("No blueprintName was provided")
			return fmt.Errorf("No blueprintName was provided")
		}
	} else {
		log.Println("No ProjectId was provided")
		return fmt.Errorf("No ProjectId was provided")
	}

	environmentId, ok_environmentId := d.GetOk("EnvironmentId")

	//if the environment ID was provided then get the right deployment plan for the blueprint ID
	// then deploy the blueprint.
	// this will take some time so call GetTaskStatus for it to check when this is done
	if ok_environmentId {
		log.Println("EnvironmentId is : ", environmentId.(string))
		blueprintId, ok_blueprintId := d.GetOk("BlueprintId")
		if ok_blueprintId {
			log.Println("BlueprintId is : ", blueprintId.(string))
			compType, ok_type:= d.GetOk("type")
			if ok_type {
				log.Println("type is : ", compType.(string))
				var deploymentPlan = "No Plan"
				var err error
				// based on the blueprint, the environment, and the size (if supplied), ge the right 
				// deployment plan
				deploymentPlan, err = GetDeploymentPlan(d, blueprintId.(string), environmentId.(string))
				if err != nil {
					return err
				}
				f1, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
			    if errf != nil {
			        log.Println("error opening file: ", errf)
			    }
			    defer f1.Close()

			    log.SetOutput(f1)

				log.Println("Determined the Deployment Plan ")
				log.Println("deploymentPlan is : ", deploymentPlan)
				// when the deployment plan is supplied then start the bluerint using that plan
				if deploymentPlan != "No Plan" {
					deployResponse := api.DeploymentPlanBlueprintDeploy(blueprintId.(string), environmentId.(string), deploymentPlan, credentials.UserName, credentials.Password)
					r := strings.NewReader(string(deployResponse))
					log.Println("Deploy response is : ", r)
					decoder := xml.NewDecoder(r)
					finish := false
					for {
						if finish {
	            			break
	        			}
						// Read tokens from the XML document in a stream.
						t, _ := decoder.Token()
						if t == nil {
							break
						}
						
						// Inspect the type of the token just read.
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
										err := GetTaskStatus(d,q.Id)
										if err != nil {
											log.Println(err)
										}
										return err
										//topologyId := d.Get("TopologyId").(string)
										//d.SetId(topologyId)
										//finish = true
									}
								}
							default:
						}
					} 
				} else {
					return errors.New("There is no Deployment Plan for the Blueprint in that environment")
				}
			} else {
				// if the type (size of machine) is not supplied then just deploy th eblueprint using the default deployment plan
				// that Agility creates
				log.Println("BlueprintId is : ", blueprintId.(string))
				deployResponse := api.SimpleBlueprintDeploy(blueprintId.(string), environmentId.(string), credentials.UserName, credentials.Password)
				r := strings.NewReader(string(deployResponse))
				log.Println("Deploy response is : ", r)
				decoder := xml.NewDecoder(r)
				finish := false
				for {
					if finish {
            			break
        			}
					// Read tokens from the XML document in a stream.
					t, _ := decoder.Token()
					if t == nil {
						break
					}
					
					// Inspect the type of the token just read.
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
									//topologyId := d.Get("TopologyId").(string)
									//d.SetId(topologyId)
									finish = true
								}
							}
						default:
					}
				}	
			}
		} else {
			log.Println("No blueprintId was provided")
			return errors.New("No blueprintId was provided")
		}
	} else {
		log.Println("No EnvironmentId was provided")
		return errors.New("No EnvironmentId was provided")
	}
	
	return nil
}

//Deleted as no longer used as a Terraform Resource
// func resourceAgilityBlueprintRead(d *schema.ResourceData, meta interface{}) error {
// 	// there is no analogy to reding the blueprint so just return Nill

// 	return nil
// }

// func resourceAgilityBlueprintDelete(d *schema.ResourceData, meta interface{}) error {
// 	//Set up logging
//     f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
//     if errf != nil {
//         log.Println("error opening file: ", errf)
//     }
//     defer f.Close()

//     log.SetOutput(f)
// 	// call the API to delete the topology identified by the ID if this resource.
// 	// this will take a while, so call the GeTaksStatus function to loop and wait until it completes
//     log.Println("Starting the Topology Destroy")
//     log.Println("TopologyId is: ", d.Id())
// 	response := api.DestroyTopology(d.Id())

// 	r := strings.NewReader(string(response))
// 	decoder := xml.NewDecoder(r)
// 	finish := false
// 	for {
// 		// Read tokens from the XML document in a stream.
// 		t, _ := decoder.Token()
// 		if t == nil {
// 			break
// 		}
// 		if finish {
//             break
//         }
// 		// parse the result and loop for and of created task
// 		switch Element := t.(type) {
// 		case xml.StartElement:
// 			if Element.Name.Local == "Task" {
// 				log.Println("Element name is : ", Element.Name.Local)

// 				var q DeployTask
// 				err := decoder.DecodeElement(&q, &Element)
// 				if err != nil {
// 					log.Println(err)
// 				}

// 				log.Println("status value is :", q.Status)
// 				if q.Status == "Pending" {
// 					GetTaskStatus(d,q.Id)
// 					d.SetId("")
// 					finish = true
// 				}
// 			}
// 		default:
// 		}
// 	}

// 	return nil
// }

func GetDeploymentPlan(d *schema.ResourceData, blueprintId string, environmentId string) (result string, err error) {
	//Set up logging
    f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

	var finished bool = false
	var dp DeploymentPlan
	var i,j int 
	var size string
	var childDepth int
	var resourceIndex int

	//Call the API to get all the Deployment plans calculated for the blueprint in the selected environment
	// Parse the result for <DeploymentPlan> elements. When one is discovered then Get the plan's machine configuration
	// and compare it with the one selected in the .tf file.
	// The response fromt he original call contains a recursive XML structure (Afinities like Cloud/Location/Tiers cause this variable depth), 
	// so call the cursive ParseDeploymentPlan function to get to the right depth to make the decision
	for {
		if finished == true {
			break
		}
		log.Println("username for deploymentPlan determination is : ", credentials.UserName)
		//log.Println("password for deploymentPlan determination is : ", credentials.Password)
		statusResponse := api.GetDeploymentPlans(blueprintId , environmentId, credentials.UserName, credentials.Password)

		log.Println("\n GetDeploymentPlans response is:",string(statusResponse))
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
			// Inspect the type of the token just read.
			switch Element := st.(type) {
				case xml.StartElement:
					//Look for the <DeploymentPlan> element
					if Element.Name.Local == "DeploymentPlan" {
						log.Println("Element name is : ", Element.Name.Local)

						err := decoder.DecodeElement(&dp, &Element)
						if err != nil {
							log.Println(err)
							finished = true
							break 
						}
						log.Println("dp.Item.Name is :", string(dp.Item.Name))
						for i = 0; i < 10; i++ {
							if finished == true {
								break
							}
							log.Println("dp.ChildList[i].Item.Name is :", dp.ChildList[i].Item.Name)
							for j = 0; j < 10; j++ {
								if finished == true {
									break
								}
								log.Println("dp.ChildList[i].OptionList[j].Name is :", dp.ChildList[i].OptionList[j].Name)
								//if the plan is for an AWS cloud, check to see if the type matches the plan. There shoukd not multiple 
								// clouds avaiable in a plan, if there is nen AWS is chosen befoe the others
								if strings.Contains(dp.ChildList[i].OptionList[j].ResourceList[0].Name, configuration.AWSCloud)  {
									switch d.Get("type") {
										case "XS":
											size = configuration.AWSXS
										case "S":
											size = configuration.AWSS
										case "M":
											size = configuration.AWSM
										case "L":
											size = configuration.AWSL
										case "XL":	
											size = configuration.AWSXL
										case "XXL":
											size = configuration.AWSXXL
									}
									childDepth = i
									ParseDeploymentPlan (d, j, &resourceIndex, &childDepth, dp.ChildList[i], size)
									finished = true
									break
								//if the plan is for an BizCloud cloud, check to see if the type matches the plan. There shoukd not multiple 
								// clouds avaiable in a plan, if there is nen AWS is chosen befoe the others	
								} else if strings.Contains(dp.ChildList[i].OptionList[j].ResourceList[0].Name, configuration.BizCloud)  {
									switch d.Get("type") {
										case "XS":
											size = configuration.BCXS
										case "S":
											size = configuration.BCS
										case "M":
											size = configuration.BCM
										case "L":
											size = configuration.BCL
										case "XL":	
											size = configuration.BCXL
										case "XXL":
											size = configuration.BCXXL
									}
									childDepth = i
									// There is likely to be multiple child/option depths due to afilities, so dig deeper until the right
									// level is found
									ParseDeploymentPlan (d, j, &resourceIndex, &childDepth, dp.ChildList[i], size)
									finished = true
									break
								} else {
									log.Println("No match for the cloud provided")
									return "", errors.New("No match for the cloud provided")
								}
							}
						}
					}
				default:
			}
		}
	}

	index1 := i-1
	index2 := j
	log.Println("about to build the DeploymentPlan")
	var child DPOptionChild
	// as with parsing the Deployment plan XML, because their is levels, we'll recurse in building up
	// the resulting Deployment plan also
	child = *CreateChildList(d, index2, childDepth, index1, resourceIndex, dp.ChildList[index1], size)
	log.Println("just got the right structure. Child is: ", child)
	deploymentPlan := DeploymentPlan {
		XMLNS: "http://servicemesh.com/agility/api",
		Item:	dp.Item,
		Rank:	dp.Rank,
		Error:	dp.Error,
		ResourceAffinity:	dp.ResourceAffinity,
		ChildList: 	[]DPOptionChild { 
			child,
		},
	}

	xmlStr, err := xml.MarshalIndent(deploymentPlan, "", "    ")
	if err != nil {
		log.Printf("error: %v\n", err)													
	}

	result = xml.Header + string(xmlStr)
	return result, nil
}

func CreateChildList(d *schema.ResourceData, index2 int, childDepth int, currentDepth int, resourceIndex int, children DPOptionChild, size string) *DPOptionChild {
	//Set up logging
    f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

    // just keeping track on where we are for debuggin purposes
	log.Println("currentDepth is: ", currentDepth)
	log.Println("childDepth is: ", childDepth)
	var Child DPOptionChild

	// we know how deep to go, based on the parsing experience, so go deeper if this is not the right level
	if childDepth != currentDepth {
		log.Println("Not the right Child, go deeper")
		var child DPOptionChild
		child = *CreateChildList(d, index2, childDepth, currentDepth+1, resourceIndex, children.OptionList[0].ChildList[0], size)
		log.Println("coming back out of the recursion for ", children.OptionList[0].ChildList[currentDepth].Item.Name)
		children.OptionList[0].ChildList[0] = child
		Child = DPOptionChild {
			Name: 	children.Name,
			Item: 	children.Item,
			Rank: 	children.Rank,
			Error: 	children.Error,
			ResourceAffinity: children.ResourceAffinity,
			OptionList:	[]DPChildOption {
				DPChildOption {
					Name: 	children.OptionList[0].Name,
					Item:	children.OptionList[0].Item,
					Rank:	children.OptionList[0].Rank,
					Error:	children.OptionList[0].Error,
					ResourceAffinity:	children.OptionList[0].ResourceAffinity,
					ResourceList:	children.OptionList[0].ResourceList,
					ChildList: 	children.OptionList[0].ChildList,
				},
			},
		}
		log.Println("child created is: ", &child.Item.Name)
		return &Child
	} else {
		// we know how deep to go, based on the parsing experience, 
		// so create the plan for the one we found that matches the requirements
		log.Println("Have gotten to the right Child")
		log.Println("index2 is: ", index2)
		log.Println("resourceIndex is: ",resourceIndex)
		var ResourceList []DPOptionResource
		rlIndex := 0
		length := len(children.OptionList[index2].ResourceList)
		if length != 0 {
			for i := 0; i < length; i++ {
				if i < len(children.OptionList[index2].ResourceList) {
					log.Println("children.OptionList[index2].ResourceList[i].Name is :", children.OptionList[index2].ResourceList[i].Name)
					log.Println("i is :", i, "and len(children.OptionList[index2].ResourceList) is: ", len(children.OptionList[index2].ResourceList))
					// there are some optional elements that we need to plan for, 
					// but it's the particular 'model' value we are after (there could be a number of them)
					if children.OptionList[index2].ResourceList[i].HREF != "" {
						log.Println("children.OptionList[index2].ResourceList[i].HREF is :", children.OptionList[index2].ResourceList[i].HREF)
						if strings.Contains(children.OptionList[index2].ResourceList[i].HREF, "model/") {
							if children.OptionList[index2].ResourceList[i].Name == size {
								log.Println("children.OptionList[index2].ResourceList[i] is: ", children.OptionList[index2].ResourceList[i])
								ResourceList = append(ResourceList, children.OptionList[index2].ResourceList[i])
								log.Println("ResourceList[rlIndex] is: ", ResourceList[rlIndex])
								rlIndex = rlIndex + 1
							}
						} else {
							// we still have to add the other resources
							log.Println("children.OptionList[index2].ResourceList[i] is: ", children.OptionList[index2].ResourceList[i])
							ResourceList = append(ResourceList, children.OptionList[index2].ResourceList[i])
							log.Println("ResourceList[rlIndex] is: ", ResourceList[rlIndex])
							rlIndex = rlIndex + 1
						}
					} else {
						// we still have to add the other resources
						log.Println("children.OptionList[index2].ResourceList[i] is: ", children.OptionList[index2].ResourceList[i])
						ResourceList = append(ResourceList, children.OptionList[index2].ResourceList[i])
						log.Println("ResourceList[rlIndex] is: ", ResourceList[rlIndex])
						rlIndex = rlIndex + 1
					}
				} else {
					log.Println("No more Resources")
					break
				}
				log.Println("Looping Again")
			}
		} else {
			log.Println("there is no match for the size requirement")
			return nil
		}
		

		log.Println("Finished getting Resource List")

		// create the child element for this level. These are built 'bottom up'
		Child = DPOptionChild {
			Item:	children.Item,
			Rank: 	children.Rank,
			Error: 	children.Error,
			ResourceAffinity: children.ResourceAffinity,
			OptionList:	[]DPChildOption {
				DPChildOption {
					Name: 	children.OptionList[index2].Name,
					Item:	children.OptionList[index2].Item,
					Rank:	children.OptionList[index2].Rank,
					Error:	children.OptionList[index2].Error,
					ResourceAffinity:	children.OptionList[index2].ResourceAffinity,
					ResourceList:	ResourceList,
					ChildList: 	children.OptionList[index2].ChildList,
				},
			},
		}
	}
	log.Println("Child is: ",Child)
	return &Child
}

func ParseDeploymentPlan (d *schema.ResourceData, j int, resourceIndex *int, childDepth *int, child DPOptionChild, size string) {
	//Set up logging
    f, errf := os.OpenFile("agility.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
    if errf != nil {
        log.Println("error opening file: ", errf)
    }
    defer f.Close()

    log.SetOutput(f)

    finished := false
	log.Println("childDepth is: ", *childDepth)
	log.Println("Starting the XML traversing function")
	log.Println("child.Item.HREF is :", child.Item.HREF)
	// its teh Workload HREF we are looking for, but only the one were the value is teh right size
	length := len(child.OptionList[j].ResourceList)
	log.Println("length of child.OptionList[j].ResourceList is :", length)
	if strings.Contains(child.Item.HREF, "workload/") {
		for k := 0; k < length; k++{
			if finished == true {
				break
			}
			log.Println("child.OptionList[j].ResourceList[k].Name is :", child.OptionList[j].ResourceList[k].Name)
			if child.OptionList[j].ResourceList[k].Name == size {
				log.Println("we have found the type definition. it is: ", size)
				resourceIndex = &k
				finished = true
				break 
			}
		}
	} else {
		// if this is not the  one we are looking for, and gobe through all teh resources at this level,
		// then go deeper
		log.Println("child.OptionList[j].Name is :", child.OptionList[j].Name)
		for k := 0; k < length; k++{
			if finished == true {
				break
			}
			log.Println("child.OptionList[j].ChildList[k].Item.name is :", child.OptionList[j].ChildList[k].Item.Name)
			for l := 0; l < length; l++{
				if finished == true {
					break
				}
				*childDepth = *childDepth + 1
				log.Println("child.OptionList[j].ChildList[k].Item.name is :", child.OptionList[j].ChildList[k].Item.Name)
				ParseDeploymentPlan (d, k, resourceIndex, childDepth, child.OptionList[j].ChildList[k], size)
				log.Println("coming out of recursive call into: ", child.OptionList[j].ChildList[k].Item.Name)
				finished = true
				break 
			}
		}
	}
}