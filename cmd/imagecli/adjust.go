package imagecli

import (
	"fmt"
	"imageCLI/pkg/loader"
	"imageCLI/pkg/service"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var (
	saturation float64
	contrast   float64
	hue        float64
	invert     bool
	redShift   float64
	greenShift float64
	blueShift  float64
)

func NewAdjustCmd(imageService *service.ImageService) *cobra.Command {
	cmd := cobra.Command{
		Use:   "adjust",
		Short: "Настройка цветовых параметров изображения",
		Run: func(cmd *cobra.Command, args []string) {
			adjustParams := service.AdjustParams{
				Saturation: saturation,
				Contrast:   contrast,
				Hue:        hue,
				Invert:     invert,
				RedShift:   redShift,
				GreenShift: greenShift,
				BlueShift:  blueShift,
			}

			ops := service.Operations{
				Adjust:      &adjustParams,
				UseParallel: !noParallel,
			}

			t := time.Now()
			images, err := loader.LoadImages(inputPath, ops.UseParallel)
			if err != nil {
				log.Fatalf("Ошибка загрузки изображений: %v", err)
			}

			processedImages, err := imageService.ProcessImages(images, ops)
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

	cmd.Flags().Float64Var(&saturation, "saturation", 0, "Изменение насыщенности (-100 to 100)")
	cmd.Flags().Float64Var(&contrast, "contrast", 0, "Изменение контрастности (-100 to 100)")
	cmd.Flags().Float64Var(&hue, "hue", 0, "Изменение оттенка (-180 to 180)")
	cmd.Flags().BoolVar(&invert, "invert", false, "Инвертировать цвета")
	cmd.Flags().Float64VarP(&redShift, "red", "r", 0, "Сдвиг красного канала (-255 to 255)")
	cmd.Flags().Float64VarP(&greenShift, "green", "g", 0, "Сдвиг зеленого канала (-255 to 255)")
	cmd.Flags().Float64VarP(&blueShift, "blue", "b", 0, "Сдвиг синего канала (-255 to 255)")

	return &cmd
}
