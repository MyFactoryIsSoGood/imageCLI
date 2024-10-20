// cmd/imagecli/main.go
package main

import (
	"imageCLI/cmd/imagecli"
	"imageCLI/pkg/service"
	"runtime"
)

func main() {
	imageService := service.NewImageService(runtime.NumCPU())

	imagecli.Execute(imageService)
}

// Тест операций: чтобы не вызывать команды
//package main
//
//import (
//	"fmt"
//	"imageCLI/pkg/imaging"
//	"imageCLI/pkg/loader"
//	"os"
//)
//
//func main() {
//	inputPath := "C:/Users/Снежана/GolandProjects/imageCLI/cmd/sources/lenna.png"
//
//	outputBlur := "C:/Users/Снежана/GolandProjects/imageCLI/cmd/outputs/lenna_blur.png"
//	outputResize := "C:/Users/Снежана/GolandProjects/imageCLI/cmd/outputs/lenna_resize.png"
//	outputAdjust := "C:/Users/Снежана/GolandProjects/imageCLI/cmd/outputs/lenna_adjust.png"
//
//	// Ensure the output directory exists
//	err := os.MkdirAll("outputs", os.ModePerm)
//	if err != nil {
//		fmt.Printf("Error creating output directory: %v\n", err)
//		os.Exit(1)
//	}
//
//	// Load the image
//	img, err := loader.LoadImages(inputPath)
//	if err != nil {
//		fmt.Printf("Error loading image '%s': %v\n", inputPath, err)
//		os.Exit(1)
//	}
//
//	// Apply blur
//	fmt.Println("Applying blur...")
//	blurOps := imaging.Operations{
//		blur:        &imaging.BlurParams{Radius: 5},
//		UseParallel: true,
//	}
//	imgBlurred, err := imaging.process(img[0].Img, blurOps)
//	if err != nil {
//		fmt.Printf("Error applying blur: %v\n", err)
//		os.Exit(1)
//	}
//	err = loader.SaveImagesToDir(outputBlur, []loader.ImageFile{{Img: imgBlurred}})
//	if err != nil {
//		fmt.Printf("Error saving blurred image: %v\n", err)
//		os.Exit(1)
//	}
//	fmt.Printf("Blurred image saved to '%s'\n", outputBlur)
//
//	// Apply resize
//	fmt.Println("Applying resize...")
//	resizeOps := imaging.Operations{
//		resize:      &imaging.ResizeParams{Width: 600, Height: 600},
//		UseParallel: true,
//	}
//	imgResized, err := imaging.process(img[0].Img, resizeOps)
//	if err != nil {
//		fmt.Printf("Error applying resize: %v\n", err)
//		os.Exit(1)
//	}
//	err = loader.SaveImagesToDir(outputResize, []loader.ImageFile{{Img: imgResized}})
//	if err != nil {
//		fmt.Printf("Error saving resized image: %v\n", err)
//		os.Exit(1)
//	}
//	fmt.Printf("Resized image saved to '%s'\n", outputResize)
//
//	// Apply adjust
//	fmt.Println("Applying adjust...")
//	adjustOps := imaging.Operations{
//		adjust: &imaging.AdjustParams{
//			Saturation: 0, // Increase saturation by 20%
//			Contrast:   0, // Increase contrast by 10%
//			Hue:        0, // Shift hue by 15 degrees
//			Invert:     false,
//			RedShift:   255,
//			GreenShift: 255,
//			BlueShift:  0,
//		},
//		UseParallel: true,
//	}
//	imgAdjusted, err := imaging.process(img[0].Img, adjustOps)
//	if err != nil {
//		fmt.Printf("Error applying adjust: %v\n", err)
//		os.Exit(1)
//	}
//	err = loader.SaveImagesToDir(outputAdjust, []loader.ImageFile{{Img: imgAdjusted}})
//	if err != nil {
//		fmt.Printf("Error saving adjusted image: %v\n", err)
//		os.Exit(1)
//	}
//	fmt.Printf("Adjusted image saved to '%s'\n", outputAdjust)
//
//	fmt.Println("Image processing completed successfully.")
//}
