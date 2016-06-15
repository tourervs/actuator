package cross

import "fmt"
import "errors"
import "github.com/boltdb/bolt"

var dynima_edit_error =  errors.New("Unable to edit dynima")
var dynima_get_error  =  errors.New("Unable to get dynima")

type Dynima struct {
    //parsers
    Id              string
    ids             []int // id column number
    SourcePath      string
    SourceType      string
    header          []string
    data            [][]string
    filters         []string //FilterList
    template        string
    data_indexes    [][]int
    delim_indexes   [][]int
}


func (d *Dynima) BindFilter (filter_name string)(error) {
    return nil
}

func (d *Dynima) UnbindFilter (filter_name string)(error) {
    return nil
}

//func (d *Dynima) RunFilter (filter_name string)(error) {
//    return nil
//}

func (d *Dynima) Save ()(error) {
    return nil
}

func (d *Dynima) GetData ()(error) {
    return nil
}

func (d *Dynima) SetTemplate()(error) {
    return nil
}

func (d *Dynima) SetSource (sourceType string, sourcePath string)(error) {
    return nil
}


func CreateDynima(id string)(error) {

    fmt.Printf("Storage error %v\nNew dynima id %s\n",STORAGE_INSTANCE.Error,id)
    return nil

}

func EditDynima(d Dynima)(err error){
    if STORAGE_INSTANCE.Error == false {

        db:=STORAGE_INSTANCE.Db
        err=db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            if b==nil{ return nil }
            dynima,err:=b.CreateBucket([]byte(d.Id))
            if err==nil || err==bolt.ErrBucketExists { // If the key exist then its previous value will be overwritten
                err=dynima.Put([]byte("source_path"),[]byte(d.SourcePath))
                if err!=nil{ return err }
                err=dynima.Put([]byte("source_type"),[]byte(d.SourceType))
                if err!=nil{ return err }
                err=dynima.Put([]byte("template"),[]byte(d.template))
                if err!=nil{ return err }
            } else { return err }
            return nil
        });

    }
    return dynima_edit_error
}

func GetDynima(id string)(dynima *Dynima,err error){
    if STORAGE_INSTANCE.Error == false {

        db:=STORAGE_INSTANCE.Db
        err = db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            fmt.Printf("\nDynimas collection doesnt exist\n")
            if b==nil{ return dynima_get_error }
            d:=b.Bucket([]byte(id))
            fmt.Printf("\nDynima doesnt exist\n")
            if d==nil{ return dynima_get_error }
            dynima = &Dynima{}
            source_path := d.Get([]byte("source_path"))
            if source_path == nil { source_path=[]byte("")  }
            source_type      := d.Get([]byte("source_type"))
            if source_type == nil { source_type=[]byte("")  }
            template     := d.Get([]byte("template"))
            if template == nil { template=[]byte("")  }
            dynima.SourcePath   = id
            dynima.SourcePath   = string(source_path)
            dynima.SourceType   = string(source_type)
            dynima.template     = string(template)
            return nil
        });
        if err == nil { return dynima, nil } else { return nil, err }


    }
    return dynima, err
}

