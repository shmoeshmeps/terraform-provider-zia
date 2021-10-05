terraform {
    required_providers {
        zia = {
            version = "1.0.0"
            source = "zscaler.com/zia/zia"
        }
    }
}

provider "zia" {}

resource "zia_location_management" "toronto"{
    name = "SGIO-IPSEC-Toronto"
    description = "Created with Terraform"
    //ip_addresses = [ zia_traffic_forwarding_static_ip.example.static_ip_id ]
    vpn_credentials {
       id = zia_traffic_forwarding_vpn_credentials.example.vpn_credental_id
       type = zia_traffic_forwarding_vpn_credentials.example.type
    }

}

resource "zia_traffic_forwarding_vpn_credentials" "example"{
    type = "UFQDN"
    fqdn = "sjc-1-37@securitygeek.io"
    comments = "created automatically"
    pre_shared_key = "newPassword123!"
}

output "zia_location_management"{
    value = zia_traffic_forwarding_vpn_credentials.example
    // sensitive = true
}

/*
resource "zia_traffic_forwarding_static_ip" "example"{
    ip_address =  "50.98.112.169"
    routable_ip = true
    comment = "Created with Terraform"
    geo_override = false
}




*/


/*
data "zia_activation_status" "example"{
}

output "zia_activation_status"{
    value = data.zia_activation_status.example
}

data "zia_location_management" "vancouver"{
    name = "SGIO-IPSEC-Vancouver"
}

output "zia_location_management"{
    value = data.zia_location_management.vancouver
}






*/


