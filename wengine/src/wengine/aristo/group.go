package aristo


import "wengine/core/common"
import "wengine/core/types/aristo_types"

type Group struct {


}


func CreateNewGroup()(g *Group,err error) {
   //group_prop := s.GetProp
   new_group:=make(map[string]interface{},0)
   new_group["id"],_=common.GenId()
   new_query:=Query{Table:GROUPS_T,Type:aristo_types.CREATE_NEW,Body:new_group}
   new_query.Run()
   return g,err
}


func GetGroup(id string)(g *Group,err error){
    return g,err
}

func GetGroupProp(id string)(props map[string]interface{},err error) {
    return props,err
}

