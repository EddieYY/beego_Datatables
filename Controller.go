/*
* 显示datatables列表页数据
 */
func (this *UserController) List() {
 
	aColumns := []string{
		"Id",
		"Username",
		"Realname",
		"Status",
		"Gid",
	}
 
	var where = make(map[string]interface{})
 
	err := this.CheckRule("isUserAll")
	if err != nil {
		where["gid"] = this.GetUserGid()
	}
 
	fmt.Println(where)
 
	maps, count, counts := d.Datatables(aColumns, user, this.Ctx.Input, where)
 
	var output = make([][]interface{}, len(maps))
	for i, m := range maps {
		for _, v := range aColumns {
			if v == "Lasttime" {
				output[i] = append(output[i], m[v].(time.Time).Format("2006-01-02 15:04:05"))
			} else {
				output[i] = append(output[i], m[v])
			}
		}
	}
 
	data := make(map[string]interface{}, count)
	data["sEcho"] = this.GetString("sEcho")
	data["iTotalRecords"] = counts
	data["iTotalDisplayRecords"] = count
	data["aaData"] = output
 
	this.Data["json"] = data
	this.ServeJson()
 
}
