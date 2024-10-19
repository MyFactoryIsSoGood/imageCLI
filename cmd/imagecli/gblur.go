// cmd/imagecli/gblur.go
package imagecli

import (
	"fmt"
	"imageCLI/pkg/imaging"
	"imageCLI/pkg/loader"
	"imageCLI/pkg/service"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var radius float64

var gblurCmd = &cobra.Command{
	Use:   "gblur",
	Short: "Применение гауссова размытия к изображению",
	Run: func(cmd *cobra.Command, args []string) {
		ops := imaging.Operations{
			Blur:        &imaging.BlurParams{Radius: radius},
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
		if err != nil {
			log.Fatalf("Ошибка обработки: %v", err)
		}
	},
}

func init() {
	gblurCmd.Flags().Float64VarP(&radius, "radius", "r", 0, "Радиус размытия")
	gblurCmd.MarkFlagRequired("radius")
}
