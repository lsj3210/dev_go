package db

import (
	"fmt"
)

type PageQuery struct {
	SelectSQL  string
	FromSQL    string
	WhereSQL   string
	OrderBySQL string
}

type PageQueryArgs struct {
	Page       int64          `json:"page"`
	PageSize   int64          `json:"page_size"`
	QueryItems PageQueryItems `json:"query_items"`
}

type PageQueryItems struct {
	IsCriteria string                 `json:"iscriteria"`
	KeyWord    string                 `json:"keyword"`
	Condition  map[string]interface{} `json:"condition"`
}

type PageCount struct {
	Count int64
}

func (obj *PageQuery) GetPageData(args *PageQueryArgs, out interface{}) (total int64, err error) {
	var qSQL string
	var cSQL string
	obj.genSQL(args, &qSQL, &cSQL)

	var count []PageCount
	if err = db.Debug().Raw(cSQL).Find(&count).Error; err != nil {
		return
	}
	if len(count) == 1 {
		total = count[0].Count
	}
	if err = db.Debug().Raw(qSQL).Find(out).Error; err != nil {
		return
	}
	return
}

func (obj *PageQuery) genSQL(args *PageQueryArgs, querySQL *string, countSQL *string) {
	if obj.SelectSQL == "" {
		obj.SelectSQL = "*"
	}
	if obj.OrderBySQL == "" {
		obj.OrderBySQL = "id desc"
	}

	if obj.WhereSQL == "" {
		obj.WhereSQL = "(1=1)"
	}

	isCriteria := args.QueryItems.IsCriteria
	condition := args.QueryItems.Condition
	keyword := args.QueryItems.KeyWord

	if isCriteria == "0" && keyword != "" {
		var tmp string
		var keys []string
		for item := range condition {
			keys = append(keys, item)
		}
		for i := 0; i < len(keys); i++ {
			if i == 0 {
				tmp = fmt.Sprintf("(%s like '%s%s%s'", keys[i], "%", keyword, "%")
			} else if i == len(keys)-1 {
				tmp = fmt.Sprintf("%s or %s like '%s%s%s')", tmp, keys[i], "%", keyword, "%")
			} else {
				tmp = fmt.Sprintf("%s or %s like '%s%s%s'", tmp, keys[i], "%", keyword, "%")
			}
		}
		obj.WhereSQL = fmt.Sprintf("%s and %s", tmp, obj.WhereSQL)
	} else if isCriteria == "1" {
		var tmp string
		var keys []string
		for item, value := range condition {
			if value != "" {
				keys = append(keys, item)
			}
		}
		for i := 0; i < len(keys); i++ {
			if i == 0 {
				tmp = fmt.Sprintf("(%s = '%s'", keys[i], condition[keys[i]])
			} else if i == len(keys)-1 {
				tmp = fmt.Sprintf("%s and %s = '%s')", tmp, keys[i], condition[keys[i]])
			} else {
				tmp = fmt.Sprintf("%s and %s = '%s'", tmp, keys[i], condition[keys[i]])
			}
		}
		obj.WhereSQL = fmt.Sprintf("%s and %s", tmp, obj.WhereSQL)
	}

	offset := (args.Page - 1) * args.PageSize
	*querySQL = fmt.Sprintf("select %s from %s where %s order by %s limit %d, %d;",
		obj.SelectSQL, obj.FromSQL, obj.WhereSQL, obj.OrderBySQL, offset, args.PageSize)

	*countSQL = fmt.Sprintf("select count(*) as count from %s where %s;", obj.FromSQL, obj.WhereSQL)
}
