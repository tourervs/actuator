package wengine

var TABLE_TYPE_COMPONENT     int = 1000
var CHART_TYPE_COMPONENT     int = 1001

var TABLE_DATA_CREATE_ACTION int = 2000
var TABLE_DATA_EDIT_ACTION   int = 2001
var TABLE_DATA_DELETE_ACTION int = 2002
var TABLE_DATA_LINK_ACTION   int = 2003


type Dashboard struct {

    Id    string
    Name  string
    Title string
}



type Table struct {

    Name         string
    TableActions []TableAction
    RowActions   []TableAction



}

type TableAction struct {


}

func (a *Api) DashboardList()(err error, dashboards []Dashboard) {

    dashboards = []Dashboard {Dashboard{Name:"mountpoints",Title:"Mountpoints",Id:"ux4bxa2nscr3bsmm"}, Dashboard{Name:"network_settings",Title:"Network Settings",Id:"pdjku29gr9x2naq8"}}

    return nil, dashboards


}

func (d *Dashboard) GetData () {



}

