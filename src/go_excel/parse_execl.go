package go_excel

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/xuri/excelize/v2"
	_ "go_simple/src/go_seelog"
	"strings"
)

var filePath = "C:/Users/Administrator/Desktop/学历编号数据导入_学历数据导出.xlsx"

func ParseExcel() {
	defer log.Flush()
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Error("Fatal error ", err.Error())
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Error("文件关闭失败")
		}
	}()
	list := file.GetSheetList()
	if err != nil {
		log.Info("Fatal error ", err.Error())
		return
	}
	file.RemoveRow(list[0], 1)
	rows, err := file.Rows(list[0])
	sql := "select * from hr_user_edu_ex where user_id in ( select user_id from mdm_prod.mdm_user where user_code in ("
	var build strings.Builder
	for rows.Next() {
		row, err := rows.Columns()
		if row[1] == "" {
			break
		}
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			return
		}
		//build.WriteString(sql)
		build.WriteString("'")
		build.WriteString(row[0])
		build.WriteString("'")
		build.WriteString(",")
	}
	build.WriteString("))")
	sql = sql + build.String()
	index := strings.LastIndex(sql, ",)")
	if index > 0 {
		sql = strings.Replace(sql, ",)", ")", index)
	}
	log.Debug("-- debug log --")
	log.Info("-- 解析完成, 输出sql --")
	log.Trace("-- trace log --")
	log.Warn("-- warn log --")
	log.Error("-- error log --")
	log.Info(sql)
}
