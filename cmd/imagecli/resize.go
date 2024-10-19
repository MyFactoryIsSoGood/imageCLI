// cmd/imagecli/resize.go
package imagecli

import (
	"fmt"
	"github.com/spf13/cobra"
	"imageCLI/pkg/imaging"
	"imageCLI/pkg/loader"
	"imageCLI/pkg/service"
	"log"
	"time"
)

var sizeFlag string

var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Изменение размера изображения",
	Run: func(cmd *cobra.Command, args []string) {
		width, height, err := parseSize(sizeFlag)
		if err != nil {
			log.Fatalf("Неверный формат размера: %v", err)
		}

		ops := imaging.Operations{
			Resize:      &imaging.ResizeParams{Width: width, Height: height},
			UseParallel: !noParallel,
		}

		t := time.Now()
		images, err := loader.LoadImages(inputPath, ops.UseParallel)
		if err != nil {
			log.Fatalf("Ошибка загрузки изображений: %v", err)
		}

		processedImages, err := service.ProcessImages(images, ops)
		if err != nil {
			log.Fatalf("Ошибка обработки: %v", err)
		}

		err = loader.SaveImagesToDir(outputPath, processedImages)
		if err != nil {
			log.Fatalf("Ошибка сохранения изображений: %v", err)
		}
		fmt.Println(time.Since(t))

	},
}

func init() {
	resizeCmd.Flags().StringVarP(&sizeFlag, "size", "s", "", "Новый размер в формате WIDTHxHEIGHT")
	resizeCmd.MarkFlagRequired("size")
}
