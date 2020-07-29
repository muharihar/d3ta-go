package feature

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
	"github.com/muharihar/d3ta-go/interface/restapi/response"
	"github.com/muharihar/d3ta-go/system/config"
)

// GenerateCaptchaID generate Captcha
func GenerateCaptchaID(cfg *config.Config, c echo.Context) map[string]interface{} {
	captchaID := captcha.NewLen(cfg.Securities.Captcha.KeyLong)

	// encode for security
	captchaIDEncoded := EncodeCaptcha(captchaID, c.RealIP())

	resp := map[string]interface{}{
		"CaptchaID": captchaIDEncoded,
		"PicPath":   fmt.Sprintf(cfg.Securities.Captcha.ImgURL, c.Request().Host, captchaIDEncoded),
	}

	return resp
}

// CaptchaServeHTTP serve captcha image/sound
func CaptchaServeHTTP(cfg *config.Config, c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer

	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	idEncoded := file[:len(file)-len(ext)]
	if ext == "" || idEncoded == "" {
		return response.FailWithMessageWithCode(http.StatusNotFound, "Not Found", c)
	}

	id, err := DecodeCaptcha(idEncoded, c.RealIP())
	if err != nil {
		return response.FailWithMessageWithCode(http.StatusNotFound, err.Error(), c)
	}

	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}

	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if serve(w, r, id, ext, lang, download, cfg.Securities.Captcha.ImgWidth, cfg.Securities.Captcha.ImgHeight) == captcha.ErrNotFound {
		return response.FailWithMessageWithCode(http.StatusNotFound, "Notfound", c)
	}

	return nil
}

// EncodeCaptcha encode captcha
func EncodeCaptcha(captcha string, salt string) string {
	captchaIDSec := []byte(fmt.Sprintf(`%s@%s`, captcha, salt))
	return base64.StdEncoding.EncodeToString(captchaIDSec)
}

// DecodeCaptcha decode Captcha
func DecodeCaptcha(captchaEncoded string, salt string) (string, error) {
	// simple security check
	// decode from base64
	cDecoded, err := base64.StdEncoding.DecodeString(captchaEncoded)
	if err != nil {
		return "", err
	}

	// extract information
	// format: captchaID@IPAddress
	cDecodeds := strings.Split(string(cDecoded), "@")
	if len(cDecodeds) != 2 {
		return "", fmt.Errorf("Invalid Captcha")
	}
	if cDecodeds[1] != salt {
		return "", fmt.Errorf("Malfarmed Captcha")
	}
	c := cDecodeds[0]

	return c, nil
}

// VerifyString verify captcha
func VerifyString(captchaID, digits string) bool {
	return captcha.VerifyString(captchaID, digits)
}

func serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
