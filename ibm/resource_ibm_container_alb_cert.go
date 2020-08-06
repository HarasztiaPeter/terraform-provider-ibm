package ibm

import (
	"fmt"
	"strings"
	"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIBMContainerALBCert() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerALBCertCreate,
		Read:     resourceIBMContainerALBCertRead,
		Update:   resourceIBMContainerALBCertUpdate,
		Delete:   resourceIBMContainerALBCertDelete,
		Exists:   resourceIBMContainerALBCertExists,
		Importer: &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"cert_crn": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    false,
				Description: "Certificate CRN id",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster ID",
			},
			"secret_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Secret name",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Domain name",
			},
			"expires_on": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate expaire on date",
			},
			"issuer_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "certificate issuer name",
			},
			"cluster_crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cluster CRN",
			},
			"cloud_cert_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "cloud cert instance ID",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Deprecated:  "This field is deprecated",
				Description: "region name",
			},
		},
	}
}

func resourceIBMContainerALBCertCreate(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	certCRN := d.Get("cert_crn").(string)
	cluster := d.Get("cluster_id").(string)
	secretName := d.Get("secret_name").(string)

	params := v1.ALBSecretConfig{
		CertCrn:    certCRN,
		ClusterID:  cluster,
		SecretName: secretName,
	}
	params.State = "update_false"

	targetEnv, err := getAlbTargetHeader(d, meta)
	if err != nil {
		return err
	}

	albAPI := albClient.Albs()
	err = albAPI.DeployALBCert(params, targetEnv)

	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s", cluster, secretName))
	_, err = waitForContainerALBCert(d, meta, schema.TimeoutCreate)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for create resource alb cert (%s) : %s", d.Id(), err)
	}

	return resourceIBMContainerALBCertRead(d, meta)
}

func resourceIBMContainerALBCertRead(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterID := parts[0]
	secretName := parts[1]

	targetEnv, err := getAlbTargetHeader(d, meta)
	if err != nil {
		return err
	}

	albAPI := albClient.Albs()
	albSecretConfig, err := albAPI.GetClusterALBCertBySecretName(clusterID, secretName, targetEnv)
	if err != nil {
		return err
	}

	d.Set("cluster_id", albSecretConfig.ClusterID)
	d.Set("secret_name", albSecretConfig.SecretName)
	d.Set("cert_crn", albSecretConfig.CertCrn)
	d.Set("cloud_cert_instance_id", albSecretConfig.CloudCertInstanceID)
	d.Set("domain_name", albSecretConfig.DomainName)
	d.Set("expires_on", albSecretConfig.ExpiresOn)
	d.Set("issuer_name", albSecretConfig.IssuerName)

	return nil
}

func resourceIBMContainerALBCertDelete(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}

	albAPI := albClient.Albs()

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterID := parts[0]
	secretName := parts[1]
	targetEnv, err := getAlbTargetHeader(d, meta)
	if err != nil {
		return err
	}

	err = albAPI.RemoveALBCertBySecretName(clusterID, secretName, targetEnv)
	if err != nil {
		return err
	}
	_, albCertDeletionError := waitForALBCertDelete(d, meta)
	if albCertDeletionError != nil {
		return albCertDeletionError
	}
	d.SetId("")
	return nil
}

func waitForALBCertDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{"exists"},
		Target:  []string{"deleted"},
		Refresh: func() (interface{}, string, error) {
			resp, err := resourceIBMContainerALBCertExists(d, meta)
			if resp {
				return resp, "exists", nil
			}
			return resp, "deleted", err
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIBMContainerALBCertUpdate(d *schema.ResourceData, meta interface{}) error {
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	cluster := parts[0]
	secretName := parts[1]

	if d.HasChange("cert_crn") {
		crn := d.Get("cert_crn").(string)
		params := v1.ALBSecretConfig{
			CertCrn:    crn,
			ClusterID:  cluster,
			SecretName: secretName,
		}
		params.State = "update_true"
		targetEnv, err := getAlbTargetHeader(d, meta)
		if err != nil {
			return err
		}

		albAPI := albClient.Albs()
		err = albAPI.UpdateALBCert(params, targetEnv)
		if err != nil {
			return err
		}

		_, err = waitForContainerALBCert(d, meta, schema.TimeoutUpdate)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for updating resource alb cert (%s) : %s", d.Id(), err)
		}
	}
	return resourceIBMContainerALBCertRead(d, meta)
}

func resourceIBMContainerALBCertExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	clusterID := parts[0]
	secretName := parts[1]

	targetEnv, err := getAlbTargetHeader(d, meta)
	if err != nil {
		return false, err
	}

	albAPI := albClient.Albs()
	albSecretConfig, err := albAPI.GetClusterALBCertBySecretName(clusterID, secretName, targetEnv)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return albSecretConfig.ClusterID == clusterID && albSecretConfig.SecretName == secretName, nil
}

func getAlbTargetHeader(d *schema.ResourceData, meta interface{}) (v1.ClusterTargetHeader, error) {
	var region string
	if v, ok := d.GetOk("region"); ok {
		region = v.(string)
	}

	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return v1.ClusterTargetHeader{}, err
	}

	if region == "" {
		region = sess.Config.Region
	}

	targetEnv := v1.ClusterTargetHeader{
		Region: region,
	}

	return targetEnv, nil
}

func waitForContainerALBCert(d *schema.ResourceData, meta interface{}, timeout string) (interface{}, error) {
	albClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	clusterID := parts[0]
	secretName := parts[1]

	stateConf := &resource.StateChangeConf{
		Pending: []string{"creating"},
		Target:  []string{"done"},
		Refresh: func() (interface{}, string, error) {
			targetEnv, err := getAlbTargetHeader(d, meta)
			if err != nil {
				return nil, "", err
			}
			alb, err := albClient.Albs().GetClusterALBCertBySecretName(clusterID, secretName, targetEnv)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The resource alb cert %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if alb.State != "created" {
				if strings.Contains(alb.State, "failed") {
					return alb, "failed", fmt.Errorf("The resource alb cert %s failed: %v", d.Id(), err)
				}

				if alb.State == "updated" {
					return alb, "done", nil
				}
				return alb, "creating", nil
			}
			return alb, "done", nil
		},
		Timeout:    d.Timeout(timeout),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}
