package index

import "net/http"
import "fmt"
import "github.com/gin-gonic/gin"
import "wapour/settings"
import "wapour/api/webclient"
import "wapour/auth"
import "wapour/salvo"

var userstorage = salvo.UserStorageInstance


func Index()(func (c *gin.Context)) {

    template_name  := "index.html"
    navigaton_menu := GetNavigationMenu()
    data:=gin.H{"navigation_items":navigaton_menu,"static_url":settings.STATIC_URL }
    return  func(c *gin.Context ){
        fmt.Printf("\n>>>>Referer:%v\n", c.Request.Header.Get("Referer"))
        if auth.IsAuthorized(c) == true { c.HTML(200, template_name,  data ) } else { c.Redirect(302,"/auth/login") }
    }
}

func Login( ) (func (c *gin.Context)) {
    template_name := "login.html"
    server_addr   := settings.SERVER_ADDR
    server_proto  := settings.SERVER_PROTO
    server_port   := settings.SERVER_PORT
    post_url      := server_proto+"://"+server_addr+":"+server_port+"/auth/login"
    data          :=gin.H{"post_url":post_url, "static_url":settings.STATIC_URL }
    return  func(c *gin.Context ){
        fmt.Printf("\n>>>>RefererLogin:%v\n", c.Request.Header.Get("Referer"))
        if auth.IsAuthorized(c) == true {
            c.Redirect(302,"/index")
        } else {
            c.HTML(200, template_name,  data )
        }
    }
}

func Logout() (func (c *gin.Context)) {
    return  func(c *gin.Context ){
        user_id,token_id,err:=auth.GetTokenFromCookies(c)
        wrapper:=userstorage.GetWrapper(user_id,token_id)
        if wrapper != nil {
            webclient.Disconnect(wrapper)
            userstorage.RemoveWrapper(user_id,token_id)
            fmt.Printf("Logout:: FindWrapper:: %v", userstorage.FindWrapper(user_id,token_id))
        }
        if err == nil {
           c.Redirect(302,settings.SERVER_URL+"/auth/login")
        } else {
            c.Redirect(302,settings.SERVER_URL+"/index")
        }
        //c.HTML(200, template_name,  data )
    }
}


func LoginPost () (func (c *gin.Context)) {
    return func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        w, err :=  webclient.Init(username, password)
        if err != nil { c.Redirect(302,settings.SERVER_URL+"/auth/login") } else {
            user := userstorage.FindWrapper(w.UserId, w.TokenId)
            if user == nil {
                userstorage.AddWrapper(w)
            }
            cookie_userid := &http.Cookie{Name:settings.USERID_COOKIE_FIELD_NAME, Value:w.UserId, Path:"/", Domain:settings.SERVER_ADDR }
            cookie_token  := &http.Cookie{Name:settings.TOKEN_COOKIE_FIELD_NAME,  Value:w.TokenId, Path:"/", Domain:settings.SERVER_ADDR }
            http.SetCookie(c.Writer, cookie_userid)
            http.SetCookie(c.Writer, cookie_token)

            c.Redirect(302,settings.SERVER_URL+"/index")
        }
    }
}
