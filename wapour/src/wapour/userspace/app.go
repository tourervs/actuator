package userspace
import "github.com/gin-gonic/gin"
import "wapour/api/webclient"
import "wapour/auth"
import "wapour/settings"
import "wapour/core/common"


func Index()(func (c *gin.Context)) {
    template_name  := "index.html"
    self_link      := "/userspace"
    return  func(c *gin.Context ){
        //if (auth.IsAuthorized(c,wrappers) && (user_id,token_id,err:=auth.GetTokenFromCookies(c); err==nil) )  {
        // thanks for postman from golang@cjr
        if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            session_id,_   :=common.GenId()
            dashboards     :=webclient.GetUserDashboards(user_id, token_id)
            navigaton_menu := CreateNavigationMenu(dashboards)
            data           := settings.APP_SETTINGS
            data["navigation_items"] = navigaton_menu
            data["session_id"]       = session_id
            data["websocket"]        = "true"
            data["app_data_url"]     = settings.USERSPACE_DATA_URL
            c.HTML(200, template_name,  data )
        } else {
            c.Redirect(302,"/auth/login?redirect_to="+self_link)
        }
    }
}

func AppData()(func (c *gin.Context)) {
    //template_name  := "index.html"
    return  func(c *gin.Context ){
        if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            session_id,_   :=common.GenId()
            dashboards     :=webclient.GetUserDashboards(user_id, token_id)
            //navigaton_menu := CreateNavigationMenu(dashboards)
            data           := settings.APP_SETTINGS
            data["navigation_items"] = dashboards
            data["session_id"]       = session_id
            //c.HTML(200, template_name,  data )
             c.JSON(200, data)
        } else {
            c.Redirect(302,"/auth/login?redirect_to=/index")
        }
    }
}


func GetDashboardData()(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            dashboardId        := c.Param("dashboardId")
            _                  = c.Param("dashboardGroupId")
            dashboard_data,_ := webclient.GetDashboardData(user_id, token_id, dashboardId)
            c.JSON(200, gin.H{"status": "ok","data":dashboard_data})
        } else {
            c.JSON(200, gin.H{"status": "error"})
        }
    }
}
