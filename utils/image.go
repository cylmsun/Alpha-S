package utils

import (
	"fmt"
	"github.com/fogleman/gg"
	uuid "github.com/satori/go.uuid"
)

// todo 实现方式需要优化，dc.Clip()不知道为什么不生效
func Text2Image(s string) {
	uuidStr := uuid.NewV4().String()
	path := fmt.Sprintf("./data/temp/temp_%s.png", uuidStr)

	temp_dc := gg.NewContext(1000, 1000)
	err := temp_dc.LoadFontFace("./data/font/LXGWWenKaiLite-Regular.ttf", 26)
	if err != nil {
		fmt.Printf("Text2Image LoadFontFace error:%s", err.Error())
		return
	}
	sWidth, sHeight := temp_dc.MeasureMultilineString(s, 1.8)

	dc := gg.NewContext(int(sWidth+10), int(sHeight+10))
	dc.LoadFontFace("./data/font/LXGWWenKaiLite-Regular.ttf", 26)
	dc.DrawRectangle(0, 0, sWidth+10, sHeight+10)
	//dc.Clip()
	dc.SetHexColor("#4D4D4D")
	dc.Fill()

	dc.SetRGB255(255, 255, 255)
	dc.DrawStringWrapped(s, 0, 6, 0, 0, sWidth+10, 1.5, gg.AlignCenter)
	dc.SavePNG(path)
}
