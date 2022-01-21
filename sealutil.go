// Package libseal
//
// @author: xwc1125
package libseal

import (
	"errors"
	"image/color"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

// Polygon TODO
func Polygon(n int, x, y, r float64) []Point {
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
	}
	return result
}

// BuildPersonalSeal 生成私人印章图片，并保存到指定路径
// @param lineSize 边线宽度
// @param font 字体对象
// @param addString 追加字符
func BuildPersonalSeal(userName string, fontFamily string) (*gg.Context, error) {
	drawFont := &DrawFont{
		FontText:   userName,
		IsBold:     false,
		FontFamily: fontFamily,
		FontSize:   80,
		FontColor: color.RGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		},
	}
	const S = 200
	const SPACE = 10
	dc := gg.NewContext(S, S)

	fullName := drawFont.FontText
	temp := []rune(drawFont.FontText)
	fontTextLen := len(temp)
	if fontTextLen < 2 || fontTextLen > 4 {
		return nil, errors.New("姓名长度不对")
	}

	switch fontTextLen {
	case 2:
		fullName = fullName + "之印"
		break
	case 3:
		fullName = fullName + "印"
		break
	}
	temp = []rune(fullName)
	// 2.字体大小，默认根据字体长度动态设定
	fontSize := drawFont.FontSize

	if drawFont.FontColor == nil {
		dc.SetRGB(1, 0, 0)
	} else {
		dc.SetColor(drawFont.FontColor)
	}

	dc.SetLineWidth(SPACE)
	dc.DrawRoundedRectangle(SPACE, SPACE, S-2*SPACE, S-2*SPACE, 10)
	dc.Stroke()
	dc.Push()

	defaultFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	// 3.字体样式
	if drawFont.FontFamily == "" {
		face := truetype.NewFace(defaultFont, &truetype.Options{
			Size: fontSize,
		})
		dc.SetFontFace(face)
	} else {
		if err := dc.LoadFontFace(drawFont.FontFamily, fontSize); err != nil {
			return nil, err
		}
	}

	_, h := dc.MeasureString(drawFont.FontText)

	// 左上
	dc.DrawStringAnchored(string(temp[2]), h, h, 0.5, 0.5)
	// 右上
	dc.DrawStringAnchored(string(temp[0]), S-h, h, 0.5, 0.5)
	// 左下
	dc.DrawStringAnchored(string(temp[3]), h, S-h, 0.5, 0.5)
	// 右下
	dc.DrawStringAnchored(string(temp[1]), S-h, S-h, 0.5, 0.5)
	dc.Push()
	return dc, nil
}

// BuildCompanySeal 生成印章图片，并保存到指定路径
// @param conf 配置文件
func BuildCompanySeal(companyName string, code string, fontFamily string) (*gg.Context, error) {
	const S = 400
	const radius = S / 2
	dc := gg.NewContext(S, S)
	point := Point{
		X: radius,
		Y: radius,
	}

	// 五角星
	star := &Star{
		Point:      point,
		Radius:     70,
		LineSize:   2,
		PointCount: 5,
		LineColor:  color.RGBA{255, 0, 0, 255},
		FillColor:  color.RGBA{255, 0, 0, 255},
	}
	DrawStar(dc, star)

	// 圆
	circle := &Circle{
		Point:     point,
		LineSize:  10,
		Radius:    radius - 10,
		LineColor: color.RGBA{255, 0, 0, 255},
		FillColor: color.RGBA{0, 0, 0, 0},
	}
	DrawCircle(dc, circle)
	if fontFamily == "" {
		dir, _ := os.Getwd()
		fontFamily = dir + "/Songti.ttc"
	}
	drawFont := &DrawFont{
		FontText:   companyName,
		IsBold:     false,
		FontFamily: fontFamily,
		FontSize:   42,
		FontColor:  color.RGBA{255, 0, 0, 255},
	}
	err := DrawFont4Arc(dc, drawFont, point, radius-20, 260, true)
	if err != nil {
		return nil, err
	}

	if code != "" {
		drawFont.FontText = code
		drawFont.FontSize = 20
		DrawFont4Arc(dc, drawFont, point, radius-20, 70, false)
	}

	return dc, nil
}

// GetCurrentDirectory TODO
func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) // 返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	return strings.Replace(dir, "\\", "/", -1)       // 将\替换成/
}

// DrawCircle 画圆
// 如果lineSize=0时，那么使用默认线宽=直径的1/35
func DrawCircle(dc *gg.Context, circle *Circle) {
	dc.DrawCircle(circle.Point.X, circle.Point.Y, circle.Radius)
	if circle.FillColor != nil {
		dc.SetColor(circle.FillColor)
	}
	dc.FillPreserve()
	if circle.LineColor == nil {
		dc.SetRGBA(1, 0, 0, 0)
	} else {
		dc.SetColor(circle.LineColor)
	}
	if circle.LineSize == 0 {
		dc.SetLineWidth(circle.Radius / 35)
	} else {
		dc.SetLineWidth(circle.LineSize)
	}
	dc.Stroke()
	dc.Push()
}

// DrawStar 画五角星
func DrawStar(dc *gg.Context, star *Star) {
	// 五角星
	points := Polygon(star.PointCount, star.Point.X, star.Point.Y, star.Radius)
	// dc.SetHexColor("fff")
	for i := 0; i < star.PointCount+1; i++ {
		index := (i * 2) % star.PointCount
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	if star.FillColor == nil {
		dc.SetRGBA(1, 0, 0, 1)
	} else {
		dc.SetColor(star.FillColor)
	}
	dc.SetFillRule(gg.FillRuleWinding)
	dc.FillPreserve()
	// 线条颜色
	if star.LineColor != nil {
		dc.SetColor(star.LineColor)
	}
	dc.SetLineWidth(star.LineSize)
	dc.Stroke()
	dc.Push()
}

// DrawFont4Arc 在圆弧上写字
// @param dc
// @param drawFont 文本内容
// @param point 圆心
// @param circleRadius 圆半径
// @param allArc 总分配角度
// @param isTop 是否上面写入
func DrawFont4Arc(dc *gg.Context, drawFont *DrawFont, point Point, circleRadius float64, allArc float64, isTop bool) error {
	// 1.字体长度,中文字体需要用utf8.RuneCountInString统计
	temp := []rune(drawFont.FontText)
	fontTextLen := len(temp)
	// 2.字体大小，默认根据字体长度动态设定
	fontSize := drawFont.FontSize
	if fontSize <= 0 {
		fontSize = float64(55 - fontTextLen*2)
	}
	if fontSize <= 0 {
		fontSize = 35
	}
	// 3.字体样式
	if drawFont.FontFamily == "" {
		font, err := truetype.Parse(goregular.TTF)
		if err != nil {
			panic(err)
		}
		face := truetype.NewFace(font, &truetype.Options{
			Size: fontSize,
		})
		dc.SetFontFace(face)
	} else {
		if err := dc.LoadFontFace(drawFont.FontFamily, fontSize); err != nil {
			panic(err)
		}
	}

	// Draw the text.
	// 计算字体长度和总宽度
	_, h := dc.MeasureString(drawFont.FontText)

	if drawFont.FontColor == nil {
		dc.SetRGB(1, 0, 0)
	} else {
		dc.SetColor(drawFont.FontColor)
	}

	// 幅度间隔
	// fmt.Println("字体个数", fontTextLen)
	radianPerInterval := allArc / float64(fontTextLen-1)

	var thetaX, thetaY float64
	ry := circleRadius - h

	if isTop {
		// 上面部分
		for i := 0; i < fontTextLen; i++ {
			// 2. 计算x,y相对长度
			f := allArc/2 - radianPerInterval*float64(i)
			ras1 := gg.Radians(f)

			// fmt.Println("度数：", f, "ras1：", ras1)
			thetaX = ry * math.Sin(ras1)
			thetaY = ry * math.Cos(ras1)
			// fmt.Println("thetaX1：", thetaX, "thetaY1：", thetaY)
			// 左侧
			dc.Push()
			dc.RotateAbout(-ras1, point.X-thetaX, point.Y-thetaY)
			dc.DrawStringAnchored(string(temp[i]), point.X-thetaX, point.Y-thetaY, 0.5, 0.5)
			dc.Pop()
		}
	} else {
		// 下面部分
		for i := 0; i < fontTextLen; i++ {
			// 2. 计算x,y相对长度
			f := allArc/2 - radianPerInterval*float64(i)
			ras1 := gg.Radians(f)

			// fmt.Println("度数：", f, "ras1：", ras1)
			thetaX = ry * math.Sin(ras1)
			thetaY = ry * math.Cos(ras1)
			// fmt.Println("thetaX2：", thetaX, "thetaY2：", thetaY)
			// 左侧
			dc.Push()
			dc.RotateAbout(ras1, point.X-thetaX, point.Y+thetaY)
			dc.DrawStringAnchored(string(temp[i]), point.X-thetaX, point.Y+thetaY, 0.5, 0.5)
			dc.Pop()
		}
	}
	return nil
}
