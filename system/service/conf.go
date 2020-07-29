package service

import (
	"path/filepath"

	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/utils"
)

// GetFileConfPath get file configuration path
func GetFileConfPath(baseDir, fileName string, h *handler.Handler) (string, error) {
	filename := fileName
	filePath := baseDir

	// first location
	file := filepath.Join(filePath, filename)
	exist, _ := utils.FileIsExist(file)

	if exist == false {
		// check in default location [./] -> binary dir
		filePath = "./"
		file = filepath.Join(filePath, filename)
		exist, _ := utils.FileIsExist(file)

		if exist == false {
			// check in configuration directory [./conf]
			filePath = "./conf"
			file = filepath.Join(filePath, filename)
			exist, _ := utils.FileIsExist(file)

			if exist == false && h != nil {
				// check in config.yaml configuration path [dirLocations.conf]
				cfg, err := h.GetConfig()
				if err != nil {
					return "", err
				}

				filePath := cfg.DirLocations.Conf
				file = filepath.Join(filePath, filename)
				_, err = utils.FileIsExist(file)

				// give up, return error
				if err != nil {
					return "", err
				}

				// return file, nil
			}			
			// return file, nil
		}
		// return file, nil
	}

	return file, nil
}