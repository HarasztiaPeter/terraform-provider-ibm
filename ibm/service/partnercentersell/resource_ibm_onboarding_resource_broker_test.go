// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package partnercentersell_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/partnercentersell"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmOnboardingResourceBrokerBasic(t *testing.T) {
	var conf partnercentersellv1.Broker
	authScheme := "bearer"
	brokerURL := fmt.Sprintf("tf_broker_url_%d", acctest.RandIntRange(10, 100))
	typeVar := "provision_through"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	authSchemeUpdate := "bearer-crn"
	brokerURLUpdate := fmt.Sprintf("tf_broker_url_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "provision_behind"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingResourceBrokerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingResourceBrokerConfigBasic(authScheme, brokerURL, typeVar, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingResourceBrokerExists("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_scheme", authScheme),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "broker_url", brokerURL),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingResourceBrokerConfigBasic(authSchemeUpdate, brokerURLUpdate, typeVarUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_scheme", authSchemeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "broker_url", brokerURLUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmOnboardingResourceBrokerAllArgs(t *testing.T) {
	var conf partnercentersellv1.Broker
	env := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	authUsername := "apikey"
	authPassword := fmt.Sprintf("tf_auth_password_%d", acctest.RandIntRange(10, 100))
	authScheme := "bearer"
	resourceGroupCrn := fmt.Sprintf("tf_resource_group_crn_%d", acctest.RandIntRange(10, 100))
	state := "removed"
	brokerURL := fmt.Sprintf("tf_broker_url_%d", acctest.RandIntRange(10, 100))
	allowContextUpdates := "false"
	catalogType := fmt.Sprintf("tf_catalog_type_%d", acctest.RandIntRange(10, 100))
	typeVar := "provision_through"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	region := fmt.Sprintf("tf_region_%d", acctest.RandIntRange(10, 100))
	envUpdate := fmt.Sprintf("tf_env_%d", acctest.RandIntRange(10, 100))
	authUsernameUpdate := "apikey"
	authPasswordUpdate := fmt.Sprintf("tf_auth_password_%d", acctest.RandIntRange(10, 100))
	authSchemeUpdate := "bearer-crn"
	resourceGroupCrnUpdate := fmt.Sprintf("tf_resource_group_crn_%d", acctest.RandIntRange(10, 100))
	stateUpdate := "active"
	brokerURLUpdate := fmt.Sprintf("tf_broker_url_%d", acctest.RandIntRange(10, 100))
	allowContextUpdatesUpdate := "true"
	catalogTypeUpdate := fmt.Sprintf("tf_catalog_type_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "provision_behind"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	regionUpdate := fmt.Sprintf("tf_region_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmOnboardingResourceBrokerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmOnboardingResourceBrokerConfig(env, authUsername, authPassword, authScheme, resourceGroupCrn, state, brokerURL, allowContextUpdates, catalogType, typeVar, name, region),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmOnboardingResourceBrokerExists("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", conf),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "env", env),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_username", authUsername),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_password", authPassword),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_scheme", authScheme),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "resource_group_crn", resourceGroupCrn),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "state", state),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "broker_url", brokerURL),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "allow_context_updates", allowContextUpdates),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "catalog_type", catalogType),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "region", region),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmOnboardingResourceBrokerConfig(envUpdate, authUsernameUpdate, authPasswordUpdate, authSchemeUpdate, resourceGroupCrnUpdate, stateUpdate, brokerURLUpdate, allowContextUpdatesUpdate, catalogTypeUpdate, typeVarUpdate, nameUpdate, regionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "env", envUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_username", authUsernameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_password", authPasswordUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "auth_scheme", authSchemeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "resource_group_crn", resourceGroupCrnUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "state", stateUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "broker_url", brokerURLUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "allow_context_updates", allowContextUpdatesUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "catalog_type", catalogTypeUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_onboarding_resource_broker.onboarding_resource_broker_instance", "region", regionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_onboarding_resource_broker.onboarding_resource_broker",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmOnboardingResourceBrokerConfigBasic(authScheme string, brokerURL string, typeVar string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_onboarding_resource_broker" "onboarding_resource_broker_instance" {
			auth_scheme = "%s"
			broker_url = "%s"
			type = "%s"
			name = "%s"
		}
	`, authScheme, brokerURL, typeVar, name)
}

func testAccCheckIbmOnboardingResourceBrokerConfig(env string, authUsername string, authPassword string, authScheme string, resourceGroupCrn string, state string, brokerURL string, allowContextUpdates string, catalogType string, typeVar string, name string, region string) string {
	return fmt.Sprintf(`

		resource "ibm_onboarding_resource_broker" "onboarding_resource_broker_instance" {
			env = "%s"
			auth_username = "%s"
			auth_password = "%s"
			auth_scheme = "%s"
			resource_group_crn = "%s"
			state = "%s"
			broker_url = "%s"
			allow_context_updates = %s
			catalog_type = "%s"
			type = "%s"
			name = "%s"
			region = "%s"
		}
	`, env, authUsername, authPassword, authScheme, resourceGroupCrn, state, brokerURL, allowContextUpdates, catalogType, typeVar, name, region)
}

func testAccCheckIbmOnboardingResourceBrokerExists(n string, obj partnercentersellv1.Broker) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
		if err != nil {
			return err
		}

		getResourceBrokerOptions := &partnercentersellv1.GetResourceBrokerOptions{}

		getResourceBrokerOptions.SetBrokerID(rs.Primary.ID)

		broker, _, err := partnerCenterSellClient.GetResourceBroker(getResourceBrokerOptions)
		if err != nil {
			return err
		}

		obj = *broker
		return nil
	}
}

func testAccCheckIbmOnboardingResourceBrokerDestroy(s *terraform.State) error {
	partnerCenterSellClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PartnerCenterSellV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_onboarding_resource_broker" {
			continue
		}

		getResourceBrokerOptions := &partnercentersellv1.GetResourceBrokerOptions{}

		getResourceBrokerOptions.SetBrokerID(rs.Primary.ID)

		// Try to find the key
		_, response, err := partnerCenterSellClient.GetResourceBroker(getResourceBrokerOptions)

		if err == nil {
			return fmt.Errorf("onboarding_resource_broker still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for onboarding_resource_broker (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmOnboardingResourceBrokerBrokerEventCreatedByUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["user_id"] = "testString"
		model["user_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.BrokerEventCreatedByUser)
	model.UserID = core.StringPtr("testString")
	model.UserName = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingResourceBrokerBrokerEventCreatedByUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingResourceBrokerBrokerEventUpdatedByUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["user_id"] = "testString"
		model["user_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.BrokerEventUpdatedByUser)
	model.UserID = core.StringPtr("testString")
	model.UserName = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingResourceBrokerBrokerEventUpdatedByUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmOnboardingResourceBrokerBrokerEventDeletedByUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["user_id"] = "testString"
		model["user_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(partnercentersellv1.BrokerEventDeletedByUser)
	model.UserID = core.StringPtr("testString")
	model.UserName = core.StringPtr("testString")

	result, err := partnercentersell.ResourceIbmOnboardingResourceBrokerBrokerEventDeletedByUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
