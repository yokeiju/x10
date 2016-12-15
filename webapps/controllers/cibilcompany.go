package controllers

import (
	. "eaciit/x10/webapps/helper"
	. "eaciit/x10/webapps/models"
	"github.com/eaciit/knot/knot.v1"
	tk "github.com/eaciit/toolkit"
)

type CibilCompanyController struct {
	*BaseController
}

func (c *CibilCompanyController) Default(k *knot.WebContext) interface{} {
	access := c.LoadBase(k)
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputTemplate
	DataAccess := Previlege{}

	for _, o := range access {
		DataAccess.Create = o["Create"].(bool)
		DataAccess.View = o["View"].(bool)
		DataAccess.Delete = o["Delete"].(bool)
		DataAccess.Process = o["Process"].(bool)
		DataAccess.Delete = o["Delete"].(bool)
		DataAccess.Edit = o["Edit"].(bool)
		DataAccess.Menuid = o["Menuid"].(string)
		DataAccess.Menuname = o["Menuname"].(string)
		DataAccess.Approve = o["Approve"].(bool)
		DataAccess.Username = o["Username"].(string)
		DataAccess.Fullname = o["Fullname"].(string)
	}

	k.Config.OutputType = knot.OutputTemplate
	k.Config.IncludeFiles = []string{"shared/filter.html", "shared/loading.html"}

	return DataAccess
}

func (c *CibilCompanyController) GetData(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson

	param := tk.M{}

	err := k.GetPayload(&param)
	if err != nil {
		return CreateResult(false, nil, err.Error())
	}

	data, err := new(CibilReportModel).GetData(param.GetInt("CustomerId"), param.GetString("DealNo"))

	if err != nil {
		return CreateResult(false, nil, err.Error())
	}

	return CreateResult(true, data, "")
}

func (c *CibilCompanyController) Update(k *knot.WebContext) interface{} {
	k.Config.OutputType = knot.OutputJson

	param := CibilReportModel{}

	err := k.GetPayload(&param)
	tk.Println(err)
	if err != nil {
		return CreateResult(false, nil, err.Error())
	}

	err = new(CibilReportModel).Update(param)

	if err != nil {
		return CreateResult(false, nil, err.Error())
	}

	return CreateResult(true, nil, "")
}
