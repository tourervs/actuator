package webclient
import "net/http"
import "encoding/json"
import "wapour/api/wengine"
import "wapour/settings"


var USER_DASHBOARD_LIST_URL string = "/rest/user/get-my-dashboards"

func GetUserDashboards(user_id string,token_id string, wrappers *[]*WengineWrapper)(dashboards wengine.DashboardList){

    url  := settings.RESTAPI_URL + USER_DASHBOARD_LIST_URL
    client := &http.Client{}
    req,_ := http.NewRequest("GET", url, nil)
    SetAuthCookie(req,user_id,token_id)
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    decoder := json.NewDecoder(resp.Body)
    _ = decoder.Decode(&dashboards)
    return

}
