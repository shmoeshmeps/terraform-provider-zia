package zia

import (
	"log"

	"github.com/willguibr/terraform-provider-zia/gozscaler/activation"
	"github.com/willguibr/terraform-provider-zia/gozscaler/adminauditlogs"
	"github.com/willguibr/terraform-provider-zia/gozscaler/adminrolemgmt"
	"github.com/willguibr/terraform-provider-zia/gozscaler/client"
	"github.com/willguibr/terraform-provider-zia/gozscaler/common"
	"github.com/willguibr/terraform-provider-zia/gozscaler/dlpdictionary"
	"github.com/willguibr/terraform-provider-zia/gozscaler/trafficforwarding/gretunnels"
	"github.com/willguibr/terraform-provider-zia/gozscaler/trafficforwarding/grevirtualiplist"
	"github.com/willguibr/terraform-provider-zia/gozscaler/trafficforwarding/publicnodevips"
	"github.com/willguibr/terraform-provider-zia/gozscaler/trafficforwarding/staticips"
	"github.com/willguibr/terraform-provider-zia/gozscaler/trafficforwarding/vpncredentials"
	"github.com/willguibr/terraform-provider-zia/gozscaler/usermanagement"
)

func init() {
	// remove timestamp from Zscaler provider logger, use the timestamp from the default terraform logger
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
}

type Client struct {
	adminauditlogs   adminauditlogs.Service
	adminrolemgmt    adminrolemgmt.Service
	dlpdictionary    dlpdictionary.Service
	usermanagement   usermanagement.Service
	gretunnels       gretunnels.Service
	staticips        staticips.Service
	publicnodevips   publicnodevips.Service
	grevirtualiplist grevirtualiplist.Service
	vpncredentials   vpncredentials.Service
	common           common.Service
	activation       activation.Service
}

type Config struct {
	Username   string
	Password   string
	APIKey     string
	ZIABaseURL string
}

func (c *Config) Client() (*Client, error) {
	config, err := client.NewClientZIA(c.Username, c.Password, c.APIKey, c.ZIABaseURL)
	if err != nil {
		return nil, err
	}

	client := &Client{
		adminauditlogs:   *adminauditlogs.New(config),
		adminrolemgmt:    *adminrolemgmt.New(config),
		dlpdictionary:    *dlpdictionary.New(config),
		usermanagement:   *usermanagement.New(config),
		grevirtualiplist: *grevirtualiplist.New(config),
		publicnodevips:   *publicnodevips.New(config),
		vpncredentials:   *vpncredentials.New(config),
		gretunnels:       *gretunnels.New(config),
		staticips:        *staticips.New(config),
		common:           *common.New(config),
		activation:       *activation.New(config),
	}

	log.Println("[INFO] initialized ZIA client")
	return client, nil
}
