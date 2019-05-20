package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:AportanteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:BeneficiariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:DetalleNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "TrRegistrarDetalle",
            Router: `/tr_registrar_detalle`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:EstadoSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "GetParametroEntidad",
            Router: `/GetParametroEntidad/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "PagosPorPeriodoPago",
            Router: `/PagosPorPeriodoPago`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PagoController"],
        beego.ControllerComments{
            Method: "GetPagos",
            Router: `GetPagos/:idPeriodoPago`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:PeriodoPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:RangoEdadUpcController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoNovedadSeguridadSocialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoPagoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TipoUpcController"],
        beego.ControllerComments{
            Method: "RegistrarTipo",
            Router: `/registrar_valores`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TrPeriodoPagoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:TrPeriodoPagoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:UpcAdicionalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"] = append(beego.GlobalControllerRouter["github.com/udistrital/seguridad_social_crud/controllers:ZonaUpcController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
