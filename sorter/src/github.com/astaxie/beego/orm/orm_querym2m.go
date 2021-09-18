// Beego (http://beego.me/)

// @description beego is an open-source, high-performance web framework for the Go programming language.

// @link        http://github.com/astaxie/beego for the canonical source repository

// @license     http://github.com/astaxie/beego/blob/master/LICENSE

// @authors     astaxie, slene

package orm

import (
	"reflect"
)

// model to model struct
type queryM2M struct {
	md  interface{}
	mi  *modelInfo
	fi  *fieldInfo
	qs  *querySet
	ind reflect.Value
}

// add models to origin models when creating queryM2M.
// example:
// 	m2m := orm.QueryM2M(post,"Tag")
// 	m2m.Add(&Tag1{},&Tag2{})
//  for _,tag := range post.Tags{}
//
// make sure the relation is defined in post model struct tag.
func (o *queryM2M) Add(mds ...interface{}) (int64, error) {
	fi := o.fi
	mi := fi.relThroughModelInfo
	mfi := fi.reverseFieldInfo
	rfi := fi.reverseFieldInfoTwo

	orm := o.qs.orm
	dbase := orm.alias.DbBaser

	var models []interface{}

	for _, md := range mds {
		val := reflect.ValueOf(md)
		if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
			for i := 0; i < val.Len(); i++ {
				v := val.Index(i)
				if v.CanInterface() {
					models = append(models, v.Interface())
				}
			}
		} else {
			models = append(models, md)
		}
	}

	_, v1, exist := getExistPk(o.mi, o.ind)
	if exist == false {
		panic(ErrMissPK)
	}

	names := []string{mfi.column, rfi.column}

	values := make([]interface{}, 0, len(models)*2)

	for _, md := range models {

		ind := reflect.Indirect(reflect.ValueOf(md))

		var v2 interface{}
		if ind.Kind() != reflect.Struct {
			v2 = ind.Interface()
		} else {
			_, v2, exist = getExistPk(fi.relModelInfo, ind)
			if exist == false {
				panic(ErrMissPK)
			}
		}

		values = append(values, v1, v2)

	}

	return dbase.InsertValue(orm.db, mi, true, names, values)
}

// remove models following the origin model relationship
func (o *queryM2M) Remove(mds ...interface{}) (int64, error) {
	fi := o.fi
	qs := o.qs.Filter(fi.reverseFieldInfo.name, o.md)

	nums, err := qs.Filter(fi.reverseFieldInfoTwo.name+ExprSep+"in", mds).Delete()
	if err != nil {
		return nums, err
	}
	return nums, nil
}

// check model is existed in relationship of origin model
func (o *queryM2M) Exist(md interface{}) bool {
	fi := o.fi
	return o.qs.Filter(fi.reverseFieldInfo.name, o.md).
		Filter(fi.reverseFieldInfoTwo.name, md).Exist()
}

// clean all models in related of origin model
func (o *queryM2M) Clear() (int64, error) {
	fi := o.fi
	return o.qs.Filter(fi.reverseFieldInfo.name, o.md).Delete()
}

// count all related models of origin model
func (o *queryM2M) Count() (int64, error) {
	fi := o.fi
	return o.qs.Filter(fi.reverseFieldInfo.name, o.md).Count()
}

var _ QueryM2Mer = new(queryM2M)

// create new M2M queryer.
func newQueryM2M(md interface{}, o *orm, mi *modelInfo, fi *fieldInfo, ind reflect.Value) QueryM2Mer {
	qm2m := new(queryM2M)
	qm2m.md = md
	qm2m.mi = mi
	qm2m.fi = fi
	qm2m.ind = ind
	qm2m.qs = newQuerySet(o, fi.relThroughModelInfo).(*querySet)
	return qm2m
}