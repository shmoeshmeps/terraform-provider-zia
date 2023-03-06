package zia

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceZIADLPEngines_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDataSourceDLPEnginesConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceDLPEnginesCheck("data.zia_dlp_engines.credit_cards"),
					testAccDataSourceDLPEnginesCheck("data.zia_dlp_engines.canada_ssn"),
					testAccDataSourceDLPEnginesCheck("data.zia_dlp_engines.us_ssn"),
				),
			},
		},
	})
}

func testAccDataSourceDLPEnginesCheck(name string) resource.TestCheckFunc {
	return resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrSet(name, "name"),
	)
}

var testAccCheckDataSourceDLPEnginesConfig_basic = `
data "zia_dlp_engines" "credit_cards"{
    name = "Credit Cards"
}

data "zia_dlp_engines" "canada_ssn"{
    name = "Canada-SSN"
}

data "zia_dlp_engines" "us_ssn"{
    name = "Social Security Numbers"
}
`
