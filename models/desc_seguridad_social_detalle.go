package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type DescSeguridadSocialDetalle struct {
	Id                        int                      `orm:"column(id);pk;auto"`
	IdDetalleLiquidacion      int                      `orm:"column(id_detalle_liquidacion)"`
	Valor                     int64                    `orm:"column(valor)"`
	IdDescSeguridadSocial     *DescSeguridadSocial     `orm:"column(id_desc_seguridad_social);rel(fk)"`
	IdTipoPagoSeguridadSocial *TipoPagoSeguridadSocial `orm:"column(id_tipo_pago_seguridad_social);rel(fk)"`
}

func (t *DescSeguridadSocialDetalle) TableName() string {
	return "desc_seguridad_social_detalle"
}

func init() {
	orm.RegisterModel(new(DescSeguridadSocialDetalle))
}

// AddDescSeguridadSocialDetalle insert a new DescSeguridadSocialDetalle into database and returns
// last inserted Id on success.
func AddDescSeguridadSocialDetalle(m *DescSeguridadSocialDetalle) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDescSeguridadSocialDetalleById retrieves DescSeguridadSocialDetalle by Id. Returns error if
// Id doesn't exist
func GetDescSeguridadSocialDetalleById(id int) (v *DescSeguridadSocialDetalle, err error) {
	o := orm.NewOrm()
	v = &DescSeguridadSocialDetalle{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDescSeguridadSocialDetalle retrieves all DescSeguridadSocialDetalle matches certain condition. Returns empty list if
// no records exist
func GetAllDescSeguridadSocialDetalle(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DescSeguridadSocialDetalle))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []DescSeguridadSocialDetalle
	qs = qs.OrderBy(sortFields...).RelatedSel("IdDescSeguridadSocial", "IdTipoPagoSeguridadSocial")
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDescSeguridadSocialDetalle updates DescSeguridadSocialDetalle by Id and returns error if
// the record to be updated doesn't exist
func UpdateDescSeguridadSocialDetalleById(m *DescSeguridadSocialDetalle) (err error) {
	o := orm.NewOrm()
	v := DescSeguridadSocialDetalle{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDescSeguridadSocialDetalle deletes DescSeguridadSocialDetalle by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDescSeguridadSocialDetalle(id int) (err error) {
	o := orm.NewOrm()
	v := DescSeguridadSocialDetalle{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DescSeguridadSocialDetalle{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
