// Code generated by protoc-gen-gogo.
// source: controller.proto
// DO NOT EDIT!

/*
Package controller is a generated protocol buffer package.

It is generated from these files:
	controller.proto

It has these top-level messages:
	BDParms
	SystemParameters
	ExternalEntity
	HostEntity
	CustomInfoType
	SfcEntity
*/
package controller

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type RxModeType int32

const (
	RxModeType_RX_MODE_UNKNOWN   RxModeType = 0
	RxModeType_RX_MODE_POLLING   RxModeType = 1
	RxModeType_RX_MODE_INTERRUPT RxModeType = 2
)

var RxModeType_name = map[int32]string{
	0: "RX_MODE_UNKNOWN",
	1: "RX_MODE_POLLING",
	2: "RX_MODE_INTERRUPT",
}
var RxModeType_value = map[string]int32{
	"RX_MODE_UNKNOWN":   0,
	"RX_MODE_POLLING":   1,
	"RX_MODE_INTERRUPT": 2,
}

func (x RxModeType) String() string {
	return proto.EnumName(RxModeType_name, int32(x))
}

type ExtEntDriverType int32

const (
	ExtEntDriverType_EE_DRIVER_TYPE_UNKNOWN   ExtEntDriverType = 0
	ExtEntDriverType_EE_DRIVER_TYPE_IOSXE_SSH ExtEntDriverType = 1
)

var ExtEntDriverType_name = map[int32]string{
	0: "EE_DRIVER_TYPE_UNKNOWN",
	1: "EE_DRIVER_TYPE_IOSXE_SSH",
}
var ExtEntDriverType_value = map[string]int32{
	"EE_DRIVER_TYPE_UNKNOWN":   0,
	"EE_DRIVER_TYPE_IOSXE_SSH": 1,
}

func (x ExtEntDriverType) String() string {
	return proto.EnumName(ExtEntDriverType_name, int32(x))
}

type SfcType int32

const (
	SfcType_SFC_UNKNOWN_TYPE   SfcType = 0
	SfcType_SFC_NS_VXLAN       SfcType = 1
	SfcType_SFC_NS_NIC_BD      SfcType = 3
	SfcType_SFC_NS_NIC_L2XCONN SfcType = 4
	SfcType_SFC_EW_BD          SfcType = 2
	SfcType_SFC_EW_L2XCONN     SfcType = 5
	SfcType_SFC_EW_BD_L2FIB    SfcType = 8
	SfcType_SFC_EW_MEMIF       SfcType = 6
	SfcType_SFC_EW_VETH        SfcType = 7
)

var SfcType_name = map[int32]string{
	0: "SFC_UNKNOWN_TYPE",
	1: "SFC_NS_VXLAN",
	3: "SFC_NS_NIC_BD",
	4: "SFC_NS_NIC_L2XCONN",
	2: "SFC_EW_BD",
	5: "SFC_EW_L2XCONN",
	8: "SFC_EW_BD_L2FIB",
	6: "SFC_EW_MEMIF",
	7: "SFC_EW_VETH",
}
var SfcType_value = map[string]int32{
	"SFC_UNKNOWN_TYPE":   0,
	"SFC_NS_VXLAN":       1,
	"SFC_NS_NIC_BD":      3,
	"SFC_NS_NIC_L2XCONN": 4,
	"SFC_EW_BD":          2,
	"SFC_EW_L2XCONN":     5,
	"SFC_EW_BD_L2FIB":    8,
	"SFC_EW_MEMIF":       6,
	"SFC_EW_VETH":        7,
}

func (x SfcType) String() string {
	return proto.EnumName(SfcType_name, int32(x))
}

type SfcElementType int32

const (
	SfcElementType_ELEMENT_UNKNOWN         SfcElementType = 0
	SfcElementType_EXTERNAL_ENTITY         SfcElementType = 1
	SfcElementType_VPP_CONTAINER_MEMIF     SfcElementType = 2
	SfcElementType_NON_VPP_CONTAINER_AFP   SfcElementType = 3
	SfcElementType_NON_VPP_CONTAINER_MEMIF SfcElementType = 4
	SfcElementType_HOST_ENTITY             SfcElementType = 5
	SfcElementType_VPP_CONTAINER_AFP       SfcElementType = 6
)

var SfcElementType_name = map[int32]string{
	0: "ELEMENT_UNKNOWN",
	1: "EXTERNAL_ENTITY",
	2: "VPP_CONTAINER_MEMIF",
	3: "NON_VPP_CONTAINER_AFP",
	4: "NON_VPP_CONTAINER_MEMIF",
	5: "HOST_ENTITY",
	6: "VPP_CONTAINER_AFP",
}
var SfcElementType_value = map[string]int32{
	"ELEMENT_UNKNOWN":         0,
	"EXTERNAL_ENTITY":         1,
	"VPP_CONTAINER_MEMIF":     2,
	"NON_VPP_CONTAINER_AFP":   3,
	"NON_VPP_CONTAINER_MEMIF": 4,
	"HOST_ENTITY":             5,
	"VPP_CONTAINER_AFP":       6,
}

func (x SfcElementType) String() string {
	return proto.EnumName(SfcElementType_name, int32(x))
}

type BDParms struct {
	Flood               bool   `protobuf:"varint,1,opt,name=flood,proto3" json:"flood,omitempty"`
	UnknownUnicastFlood bool   `protobuf:"varint,2,opt,name=unknown_unicast_flood,proto3" json:"unknown_unicast_flood,omitempty"`
	Forward             bool   `protobuf:"varint,3,opt,name=forward,proto3" json:"forward,omitempty"`
	Learn               bool   `protobuf:"varint,4,opt,name=learn,proto3" json:"learn,omitempty"`
	ArpTermination      bool   `protobuf:"varint,5,opt,name=arp_termination,proto3" json:"arp_termination,omitempty"`
	MacAge              uint32 `protobuf:"varint,6,opt,name=mac_age,proto3" json:"mac_age,omitempty"`
}

func (m *BDParms) Reset()         { *m = BDParms{} }
func (m *BDParms) String() string { return proto.CompactTextString(m) }
func (*BDParms) ProtoMessage()    {}

type SystemParameters struct {
	Mtu                uint32   `protobuf:"varint,1,opt,name=mtu,proto3" json:"mtu,omitempty"`
	StartingVlanId     uint32   `protobuf:"varint,2,opt,name=starting_vlan_id,proto3" json:"starting_vlan_id,omitempty"`
	DynamicBridgeParms *BDParms `protobuf:"bytes,3,opt,name=dynamic_bridge_parms" json:"dynamic_bridge_parms,omitempty"`
	StaticBridgeParms  *BDParms `protobuf:"bytes,4,opt,name=static_bridge_parms" json:"static_bridge_parms,omitempty"`
}

func (m *SystemParameters) Reset()         { *m = SystemParameters{} }
func (m *SystemParameters) String() string { return proto.CompactTextString(m) }
func (*SystemParameters) ProtoMessage()    {}

func (m *SystemParameters) GetDynamicBridgeParms() *BDParms {
	if m != nil {
		return m.DynamicBridgeParms
	}
	return nil
}

func (m *SystemParameters) GetStaticBridgeParms() *BDParms {
	if m != nil {
		return m.StaticBridgeParms
	}
	return nil
}

type ExternalEntity struct {
	Name            string                        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MgmntIpAddress  string                        `protobuf:"bytes,2,opt,name=mgmnt_ip_address,proto3" json:"mgmnt_ip_address,omitempty"`
	MgmntPort       uint32                        `protobuf:"varint,3,opt,name=mgmnt_port,proto3" json:"mgmnt_port,omitempty"`
	BasicAuthUser   string                        `protobuf:"bytes,4,opt,name=basic_auth_user,proto3" json:"basic_auth_user,omitempty"`
	BasicAuthPasswd string                        `protobuf:"bytes,5,opt,name=basic_auth_passwd,proto3" json:"basic_auth_passwd,omitempty"`
	EeDriverType    ExtEntDriverType              `protobuf:"varint,6,opt,name=ee_driver_type,proto3,enum=controller.ExtEntDriverType" json:"ee_driver_type,omitempty"`
	HostInterface   *ExternalEntity_HostInterface `protobuf:"bytes,7,opt,name=host_interface" json:"host_interface,omitempty"`
	HostVxlan       *ExternalEntity_HostVxlan     `protobuf:"bytes,8,opt,name=host_vxlan" json:"host_vxlan,omitempty"`
	HostBd          *ExternalEntity_HostBD        `protobuf:"bytes,9,opt,name=host_bd" json:"host_bd,omitempty"`
}

func (m *ExternalEntity) Reset()         { *m = ExternalEntity{} }
func (m *ExternalEntity) String() string { return proto.CompactTextString(m) }
func (*ExternalEntity) ProtoMessage()    {}

func (m *ExternalEntity) GetHostInterface() *ExternalEntity_HostInterface {
	if m != nil {
		return m.HostInterface
	}
	return nil
}

func (m *ExternalEntity) GetHostVxlan() *ExternalEntity_HostVxlan {
	if m != nil {
		return m.HostVxlan
	}
	return nil
}

func (m *ExternalEntity) GetHostBd() *ExternalEntity_HostBD {
	if m != nil {
		return m.HostBd
	}
	return nil
}

type ExternalEntity_HostInterface struct {
	IfName   string `protobuf:"bytes,1,opt,name=if_name,proto3" json:"if_name,omitempty"`
	Ipv4Addr string `protobuf:"bytes,2,opt,name=ipv4_addr,proto3" json:"ipv4_addr,omitempty"`
}

func (m *ExternalEntity_HostInterface) Reset()         { *m = ExternalEntity_HostInterface{} }
func (m *ExternalEntity_HostInterface) String() string { return proto.CompactTextString(m) }
func (*ExternalEntity_HostInterface) ProtoMessage()    {}

type ExternalEntity_HostVxlan struct {
	IfName     string `protobuf:"bytes,1,opt,name=if_name,proto3" json:"if_name,omitempty"`
	SourceIpv4 string `protobuf:"bytes,2,opt,name=source_ipv4,proto3" json:"source_ipv4,omitempty"`
}

func (m *ExternalEntity_HostVxlan) Reset()         { *m = ExternalEntity_HostVxlan{} }
func (m *ExternalEntity_HostVxlan) String() string { return proto.CompactTextString(m) }
func (*ExternalEntity_HostVxlan) ProtoMessage()    {}

type ExternalEntity_HostBD struct {
	Id         uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BdiIpv4    string   `protobuf:"bytes,2,opt,name=bdi_ipv4,proto3" json:"bdi_ipv4,omitempty"`
	Interfaces []string `protobuf:"bytes,3,rep,name=interfaces" json:"interfaces,omitempty"`
}

func (m *ExternalEntity_HostBD) Reset()         { *m = ExternalEntity_HostBD{} }
func (m *ExternalEntity_HostBD) String() string { return proto.CompactTextString(m) }
func (*ExternalEntity_HostBD) ProtoMessage()    {}

type HostEntity struct {
	Name            string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	EthIfName       string     `protobuf:"bytes,2,opt,name=eth_if_name,proto3" json:"eth_if_name,omitempty"`
	EthIpv4         string     `protobuf:"bytes,3,opt,name=eth_ipv4,proto3" json:"eth_ipv4,omitempty"`
	EthIpv6         string     `protobuf:"bytes,4,opt,name=eth_ipv6,proto3" json:"eth_ipv6,omitempty"`
	LoopbackMacAddr string     `protobuf:"bytes,5,opt,name=loopback_mac_addr,proto3" json:"loopback_mac_addr,omitempty"`
	LoopbackIpv4    string     `protobuf:"bytes,6,opt,name=loopback_ipv4,proto3" json:"loopback_ipv4,omitempty"`
	LoopbackIpv6    string     `protobuf:"bytes,7,opt,name=loopback_ipv6,proto3" json:"loopback_ipv6,omitempty"`
	Mtu             uint32     `protobuf:"varint,8,opt,name=mtu,proto3" json:"mtu,omitempty"`
	RxMode          RxModeType `protobuf:"varint,9,opt,name=rx_mode,proto3,enum=controller.RxModeType" json:"rx_mode,omitempty"`
}

func (m *HostEntity) Reset()         { *m = HostEntity{} }
func (m *HostEntity) String() string { return proto.CompactTextString(m) }
func (*HostEntity) ProtoMessage()    {}

type CustomInfoType struct {
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (m *CustomInfoType) Reset()         { *m = CustomInfoType{} }
func (m *CustomInfoType) String() string { return proto.CompactTextString(m) }
func (*CustomInfoType) ProtoMessage()    {}

type SfcEntity struct {
	Name           string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type           SfcType                 `protobuf:"varint,3,opt,name=type,proto3,enum=controller.SfcType" json:"type,omitempty"`
	SfcIpv4Prefix  string                  `protobuf:"bytes,4,opt,name=sfc_ipv4_prefix,proto3" json:"sfc_ipv4_prefix,omitempty"`
	VnfRepeatCount uint32                  `protobuf:"varint,5,opt,name=vnf_repeat_count,proto3" json:"vnf_repeat_count,omitempty"`
	BdParms        *BDParms                `protobuf:"bytes,6,opt,name=bd_parms" json:"bd_parms,omitempty"`
	Elements       []*SfcEntity_SfcElement `protobuf:"bytes,7,rep,name=elements" json:"elements,omitempty"`
}

func (m *SfcEntity) Reset()         { *m = SfcEntity{} }
func (m *SfcEntity) String() string { return proto.CompactTextString(m) }
func (*SfcEntity) ProtoMessage()    {}

func (m *SfcEntity) GetBdParms() *BDParms {
	if m != nil {
		return m.BdParms
	}
	return nil
}

func (m *SfcEntity) GetElements() []*SfcEntity_SfcElement {
	if m != nil {
		return m.Elements
	}
	return nil
}

type SfcEntity_SfcElement struct {
	Container        string         `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	PortLabel        string         `protobuf:"bytes,2,opt,name=port_label,proto3" json:"port_label,omitempty"`
	EtcdVppSwitchKey string         `protobuf:"bytes,3,opt,name=etcd_vpp_switch_key,proto3" json:"etcd_vpp_switch_key,omitempty"`
	Ipv4Addr         string         `protobuf:"bytes,4,opt,name=ipv4_addr,proto3" json:"ipv4_addr,omitempty"`
	MacAddr          string         `protobuf:"bytes,5,opt,name=mac_addr,proto3" json:"mac_addr,omitempty"`
	Type             SfcElementType `protobuf:"varint,6,opt,name=type,proto3,enum=controller.SfcElementType" json:"type,omitempty"`
	VlanId           uint32         `protobuf:"varint,7,opt,name=vlan_id,proto3" json:"vlan_id,omitempty"`
	Mtu              uint32         `protobuf:"varint,8,opt,name=mtu,proto3" json:"mtu,omitempty"`
	RxMode           RxModeType     `protobuf:"varint,9,opt,name=rx_mode,proto3,enum=controller.RxModeType" json:"rx_mode,omitempty"`
	L2FibMacs        []string       `protobuf:"bytes,10,rep,name=l2fib_macs" json:"l2fib_macs,omitempty"`
	Ipv6Addr         string         `protobuf:"bytes,11,opt,name=ipv6_addr,proto3" json:"ipv6_addr,omitempty"`
}

func (m *SfcEntity_SfcElement) Reset()         { *m = SfcEntity_SfcElement{} }
func (m *SfcEntity_SfcElement) String() string { return proto.CompactTextString(m) }
func (*SfcEntity_SfcElement) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("controller.RxModeType", RxModeType_name, RxModeType_value)
	proto.RegisterEnum("controller.ExtEntDriverType", ExtEntDriverType_name, ExtEntDriverType_value)
	proto.RegisterEnum("controller.SfcType", SfcType_name, SfcType_value)
	proto.RegisterEnum("controller.SfcElementType", SfcElementType_name, SfcElementType_value)
}
