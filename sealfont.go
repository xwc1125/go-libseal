// Package libseal
//
// @author: xwc1125
package libseal

import "image/color"

type Point struct {
	X, Y float64
}

type Circle struct {
	Point     Point       // 圆心
	Radius    float64     // 半径
	LineSize  float64     // 线宽度
	LineColor color.Color // 线条颜色
	FillColor color.Color // 填充颜色
}

type Star struct {
	Point      Point       // 圆心
	Radius     float64     // 半径
	LineSize   float64     // 线宽度
	PointCount int         // 星角个数
	LineColor  color.Color // 线条颜色
	FillColor  color.Color // 填充颜色
}

type DrawFont struct {
	FontText   string      // 字体内容
	IsBold     bool        // 是否加粗
	FontFamily string      "宋体" // 字形名，默认为宋体
	FontSize   float64     // 字体大小
	FontColor  color.Color // 线颜色
}
