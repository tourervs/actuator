package run

import "flag"
import "wapour/settings"

var initial = []string {

    "SERVER_ADDR",
    "SERVER_PORT",
    "SERVER_PROTO",
    "SERVER_URL",
    "RESTAPI_SERVER_ADDR",
    "RESTAPI_SERVER_PROTO",
    "RESTAPI_SERVER_PORT",
    "RESTAPI_URL",

}

var Props = GetProps()

func GetProps()(props map[string]string){

    props=make(map[string]string,0)

    // serve
    // flag.Lookup

    ip_version_ptr       := flag.String("ipversion","v4",      "Server ip version"  )
    ip_port_ptr          := flag.String("port",     "80",      "Server port number" )
    ip_addr_ptr          := flag.String("addr",     "0.0.0.0", "Server ip address"  )

    flag.Parse()
    //out_file_ptr       := flag.String("outfile","out.txt","Out file")

    ip_version       := *ip_version_ptr
    ip_port          := *ip_port_ptr
    ip_addr          := *ip_addr_ptr

    props["server_ip_version"]  = ip_version
    props["server_port"]        = ip_port
    props["server_addr"]        = ip_addr

    props["server_addr"] = props["ip_addr"]+":"+props["ip_port"]

    return

}

func GetSettings()(props map[string]string){
    return
}
