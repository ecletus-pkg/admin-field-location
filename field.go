package admin_field_location

import (
	"github.com/moisespsena-go/tzdb"
	"github.com/ecletus/core"
	"github.com/moisespsena-go/i18n-modular/i18nmod"
	path_helpers "github.com/moisespsena-go/path-helpers"

	"github.com/ecletus/admin"
)

var group = i18nmod.PkgToGroup(path_helpers.GetCalledDir())

func Setup(res *admin.Resource, fieldName string) {
	res.Meta(&admin.Meta{
		Name: fieldName,
		Config: &admin.SelectOneConfig{
			Collection: tzdb.Db.Pair(),
			AllowBlank: true,
		},
		Label: group + ".Location",
	}).NewValuer(func(meta *admin.Meta, old admin.MetaValuer, recorde interface{}, ctx *core.Context) interface{} {
		if value := old(recorde, ctx); value == nil || value == "" {
			return tzdb.LocationCity("")
		} else {
			switch vt := value.(type) {
			case tzdb.Location:
				return vt
			case *tzdb.Location:
				return *vt
			case string:
				return tzdb.LocationCity(vt)
			case *string:
				return tzdb.LocationCity(*vt)
			}
			return tzdb.LocationCity("")
		}
	})
}
