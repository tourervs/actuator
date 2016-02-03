package userspace
import "fmt"
import "github.com/gin-gonic/gin"
import "wapour/api/webclient"
import "wapour/auth"


func Index( data  gin.H, wrappers *[]*webclient.WengineWrapper,  params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    navigaton_menu := GetNavigationMenu()
    data["navigation_items"] = navigaton_menu
    return  func(c *gin.Context ){
        //if (auth.IsAuthorized(c,wrappers) && (token_id,user_id,err:=auth.GetTokenFromCookies(c); err==nil) )  {
        // thanks for postman from golang@cjr
        if token_id,user_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c,wrappers) && err==nil {
            dashboards:=webclient.GetUserDashboards(token_id,user_id,wrappers)
            fmt.Printf("\nUserDashboards: %v\n",dashboards)
            c.HTML(200, template_name,  data )
        } else {
            c.Redirect(302,"/auth/login")
        }
    }
}
