package pdd

import (
	"encoding/json"
)

type ThemeListResponse struct {
	ThemeList []*Theme `json:"theme_list"`
	Total     int      `json:"total"`
}
type ThemeListSearchResponse struct {
	Goodlist []*Goods `json:"goods_list"`
	Total     int      `json:"total"`
}

type Theme struct {
	Id       int    `json:"id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	GoodsNum int    `json:"goods_num"` // 商品数量
}

func (d *DDK) ThemeListGet(page, pageSize int) (res *ThemeListResponse, err error) {
	params := NewParamsWithType(DDK_ThemeListGet)
	params.Set("page", page)
	params.Set("page_size", pageSize)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "theme_list_get_response")
	res = new(ThemeListResponse)
	err = json.Unmarshal(bytes, res)
	return
}
func (d *DDK) ThemeGoodsSearch(themeId int) (res *ThemeListSearchResponse, err error) {
	params := NewParamsWithType(DDK_ThemeGoodsSearch)
	params.Set("theme_id", themeId)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseBytes(r, "theme_list_get_response")
	res = new(ThemeListSearchResponse)
	err = json.Unmarshal(bytes, res)
	return
}
