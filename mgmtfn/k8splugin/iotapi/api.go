package iotapi

// cni api definitions shared between the cni executable and netplugin

// PluginPath is the path to the listen socket directory for netplugin
const PluginPath = "/run/contiv"

// ContivCniSocket is the full path to the listen socket for netplugin
//const ContivCniSocket = "/run/contiv/contiv-cni.sock"

// EPAddURL is the rest point for adding an endpoint type device IoT
const EPAddURL = "/ContivIOT.AddIotDev"

// EPDelURL is the rest point for deleting an endpoint
//const EPDelURL = "/ContivCNI.DelPod"

// CNIPodAttr holds attributes of the pod to be attached or detached
type IOTDevAttr struct {
	Name          string `json:"IOT_DEV_NAME,omitempty"`
	Tenant        string `json:"IOT_DEV_TENANT,omitempty"`
	InfraIotDevID string `json:"IOT_DEV_INFRA_CONTAINER_ID,omitempty"`
	IntfName      string `json:"IOT_DEV_IFNAME,omitempty"`
	Network       string `json:"IOT_DEV_NETWORK,omitempty"`
	Group         string `json:"IOT_DEV_GROUP,omitempty"`
}

// RspAddIotDev contains the response to the AddIotDev
type RspAddIotDev struct {
	EndpointID string `json:"endpointid,omitempty"`
	ErrMsg     string `json:"errmsg,omitempty"`
	ErrInfo    string `json:"errinfo,omitempty"`
	Attr       *Attr  `json:"epattr,omitempty"`
}

type Attr struct {
	IPAddress   string `json:"ipaddress,omitempty"`
	PortName    string `json:"portname,omitempty"`
	MacAddess   string `json:"gateway,omitempty"`
	IPv6Address string `json:"ipv6address,omitempty"`
}
