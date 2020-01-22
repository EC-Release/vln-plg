/*
 * Copyright (c) 2016 General Electric Company. All rights reserved.
 *
 * The copyright to the computer software herein is the property of
 * General Electric Company. The software may be used and/or copied only
 * with the written permission of General Electric Company or in accordance
 * with the terms and conditions stipulated in the agreement/contract
 * under which the software has been supplied.
 *
 * author: apolo.yasuda@ge.com
 */

package main

import (
	"os"
	"errors"
	"net"
	"strings"
	util "github.build.ge.com/212359746/wzutil"
	//plugin "github.build.ge.com/212359746/wzplugin"
	"flag"
	"github.com/vishvananda/netlink"
	"gopkg.in/yaml.v2"
	"encoding/base64"

)

var (
	//YML_VLAN_FLAG = "vlan"
	REV string= "beta"
)

const (
	EC_LOGO = `
           ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄
          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░
          ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀
          ▐░▌          ▐░▌   
          ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌
          ▐░░░░░░░░░░░▌▐░▌
          ▐░█▀▀▀▀▀▀▀▀▀ ▐░▌
          ▐░▌          ▐░▌
          ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄ 
          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌
           ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  @Digital Connect 
`
	COPY_RIGHT = "Digital Connect,  @GE Corporate"
	ISSUE_TRACKER = "https://github.com/Enterprise-connect/sdk/issues"


	//agent authorization header
	AUTH_HEADER = "Authorization"
	
	//ec service header in predix
	EC_SUB_HEADER  = "Predix-Zone-Id"

	//app index available in a cf environment
	CF_INS_IDX_EV  = "CF_INSTANCE_INDEX"
	//forwarding header targeting a cf environment
	CF_INS_HEADER  = "X-CF-APP-INSTANCE"

	//app index available in a watcher environment
	EC_INS_IDX_EV  = "EC_INSTANCE_INDEX"
	//forwarding header targeting a watcher environement
	EC_INS_HEADER  = "X-EC-APP-INSTANCE"

	//xcalr url
	CA_URL = "https://xcalr.apps.ge.com/v2beta"

	//watcher url
	WATCHER_URL = "https://raw.githubusercontent.com/Enterprise-connect/sdk/v1.1beta.watcher/watcher.yml"

)


type IPRoute struct {}

func (i *IPRoute)RegisterCidrList(ips []string) error {

	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		}
	}()

	lpb,_er1:=net.Interfaces()
	if _er1!=nil{
		return _er1
	}
		
	for _,v:= range lpb {
		if v.Flags==(net.FlagUp|net.FlagLoopback) {
			lo, err := netlink.LinkByName(v.Name)
			if err!=nil {
				return err
			}
			for _,ip := range ips {
				ip=strings.Trim(ip," ")

				util.InfoLog("[VLAN] ip: "+ip)

				addr, e := netlink.ParseAddr(ip)
				if e!=nil{
					return e
				}
				
				if e2:=netlink.AddrReplace(lo, addr);e2!=nil{
				//if e2:=netlink.AddrAdd(lo, addr);e2!=nil{
					return e2
				}

				util.InfoLog("[VLAN] Cidr address "+ip+" has been added/replaced in loopback interface.")

			}

		}

	}

	return nil
}

func GetVLANSetting()(map[string]interface{}, error){

	plg:=flag.String("plg","","Enable support for EC VLAN Plugin.")
	ver:=flag.Bool("ver", false, "Show current plugin revision.")

	flag.Parse()
	
	if *ver {
		util.InfoLog("Rev:"+REV)
		os.Exit(0)
		return nil,nil
	}

	util.InfoLog(*plg)
	f, err := base64.StdEncoding.DecodeString(*plg)
	if err!=nil{
		util.InfoLog(err)
	}
	t:=make(map[string]interface{})
	err=yaml.Unmarshal(f, &t)
	if err!=nil{
		return nil,err
	}

	if len(t)<1 {
		return nil, errors.New("invalid file format in plugin.yml")
	}

	return t, nil

}

func main(){

	defer func(){
		if r:=recover();r!=nil{
			util.PanicRecovery(r)
		} else {
			util.InfoLog("plugin undeployed.")
		}
	}()

	util.Branding("/.ec","ec-plugin","ec-config","Authorization","EC_HTTP_HEADER","EC","EC_LOGO","COPY_RIGHT","https://ca-not-in-use.com","EC")

		bc:=&util.BrandingConfig{
		CONFIG_MAIN: "/.ec",
		BRAND_CONFIG: "EC", 
		LOGO: EC_LOGO,
		COPY_RIGHT: COPY_RIGHT,
		HEADER_PLUGIN: "ec-plugin",
		HEADER_CONFIG: "ec-config",
		HEADER_AUTH: AUTH_HEADER,
		HEADER_SUB_ID: EC_SUB_HEADER,
		HEADER_CF_INST: CF_INS_HEADER,
		HEADER_INST: EC_INS_HEADER,
		ENV_CF_INST_IDX: CF_INS_IDX_EV,
		ENV_INST_IDX: EC_INS_IDX_EV,
		URL_CA: CA_URL,
		URL_WATCHER_CONF: WATCHER_URL,
		//URL_WATCHER_REPO: WATCHER_CONTR_URL,
		URL_ISSUE_TRACKER: ISSUE_TRACKER,
	}
	
	util.Branding(bc)

	util.Init("vlan","0",true)

	t,err:=GetVLANSetting()
	if err!=nil{
		panic(err)
	}
	util.DbgLog(t)
	
        ipr:=&IPRoute{}
	_ips:=strings.Split(t["ips"].(string),",")
	if _er:=ipr.RegisterCidrList(_ips);_er!=nil{
		panic(_er)
	}

}
