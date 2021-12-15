package qrcode

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)


func Poster()(image.Image,error) {
	var (
		bgFile    *os.File
		bgImg     image.Image
		qrCodeImg image.Image
		offset    image.Point
	)

	// 01: 打开背景图片
	bgFile, err = os.Open("./3hamburger-01.png")
	if err != nil {
		fmt.Println("打开背景图片失败", err)
		return nil, err
	}

	defer bgFile.Close()

	// 02: 编码为图片格式
	bgImg, err = png.Decode(bgFile)
	if err != nil {
		fmt.Println("背景图片编码失败:", err)
		return nil, err
	}

	// 03: 生成二维码
	qrCodeImg, err = CreateAvatar()
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return nil, err
	}

	offset = image.Pt(200, 200) //用于调整二维码在背景图片上的位置

	b := bgImg.Bounds()

	m := image.NewRGBA(b)

	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)

	draw.Draw(m, qrCodeImg.Bounds().Add(offset), qrCodeImg, image.Point{X: 0, Y: 0}, draw.Over)

	return m,err

}


