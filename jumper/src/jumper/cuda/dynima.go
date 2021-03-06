package cuda

import "sync"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"
import "jumper/cuda/templating"
import "jumper/cuda/result"
import "jumper/cuda/handling"

type Dynima struct {
    //  : :
    //  : :
    //  : :  dynima stores  
    //  : :  each file may got several dynimas binded to itself
    //  : :
    //  : :
    sync.RWMutex                               // mutex will be used to freze operations over dynima while changing filters or modifying targets
    filters             filtering.FilterList   // 
    targets             targets.TargetListPtrs // ? seems it is not necessary to store file and directory content inside dynima
    templates           templating.Template    // 
    configured          bool                   //
    offset              int64                  // for log files, just when dynima instance binded to single file
    id                  string                 //
    name                string                 //
    primaryKeysIndexes  []int                  // list of data indexes whose helps us to determine each uniq line inside file
    //  dataSet  []Data                        // data will collected while targets processing
    //  : :
    //  : :
    //  : :
    //  : :
}

/*
type Target struct {  // interface {
    //
    //  Get       ()(lineAsArray [][]string, err error)
    //  GetType ()(typ int)
    //  Gather    ()(error)
    //  PushPart  ([][]string)(error)
    //
    typ         int
    path        string
    lineAsArray [][]int
}
*/


func(d *Dynima)RunFilters()(resultSet result.ResultSet) {
    //
    // apply filters targets data
    //
    d.Lock()
    defer d.Unlock()
    //
    // test block
    //
    var readableTargets targets.TargetList
    //
    //
    for i := range d.targets {
        target := d.targets[i]
        target.Gather()
        if target.GatherIsFailed() == false && target.IsConfigured() == true {
            readableTargets.Append(target)
        }
    }
    //
    //
    //
    // var resultSet result.ResultSet
    // _ = resultSet
    //
    //
    //  handler := handling.NewHandler(nil)
    //  handler.AddFilters(d.filters)
    //
    //
    for i := range readableTargets {
        //
        //
        target   :=  readableTargets[i]
        handler  :=  handling.NewHandler(nil)
        //
        //
        handler.AddFilters( d.filters )
        //
        //
        handler.AddTargetPtr( &target )
        //
        //
        result,err :=  handler.Handle()
        //
        //
        if err == nil {
            resultSet.Append(result)
        }
        //
        //
        //blankResult := GetResult(target)
        //_           =  blankResult
        //
        //
    }
    //
    //
    //
    return //&resultSet
    //
    //
    //
}

func(d *Dynima)AppendFilter( f filtering.Filter )( error ){
    //
    //
    return d.filters.Append(f)
    //
    //
}

func(d *Dynima)SetSource( t *targets.Target )( error ){
    // unnecessary 
    //
    return nil
    //
    //
}

func(d *Dynima)AppendTarget(t *targets.Target)(error){
    // 
    //
    return d.targets.Append(t)
    //
    //
}
//
//
func(d *Dynima)RemoveTarget(t *targets.Target)( error ){
    // 
    //
    return nil
    //
    //
}
//
//
func(d *Dynima)getTarget(tgt_id int)( t *targets.Target, err error ){
    //
    //
    return
    //
    //
}
//
//
func(d *Dynima)getChildTargets(parent_target_id int)(child_targets *[]targets.Target, err error){
    //
    //
    return
    //
    //
}
//
//
func NewDynima()( *Dynima ){
    //
    //
    //
    var d Dynima
    d.filters     = make( filtering.FilterList, 0      )
    d.targets     = make( targets.TargetListPtrs,   0  )
    return &d
    //
    //
    //
}
