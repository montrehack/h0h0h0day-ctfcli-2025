package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	allowedThemeHash = []string{
		"72949359d58c4e8993f8ddeb609201d5ee0b41d09aaff1459b5907e6c32ea355",
		"facb0a9deffb33297a7e28c8671f6ee873e51c4c6ea8c41373df587136f3ab74",
		"dd649601078a37cc453d66cc8e3a221eb777e30631317b6ba8b7110f1aa46912",
	}
)

type AppState struct {
	Flag   string
	Author string
}

func CalculateZipHash(zipReader *zip.Reader) (string, error) {
	values := make([]string, 0, 2*len(zipReader.File))
	for _, f := range zipReader.File {
		values = append(values, f.Name)
		open, err := f.Open()
		if err != nil {
			return "", err
		}

		all, err := io.ReadAll(open)
		if err != nil {
			return "", err
		}
		values = append(values, string(all))
	}

	// The order of files in zip doesn't matter
	sort.Strings(values)

	id := strings.Join(values, "-")
	h := sha256.Sum256([]byte(id))
	return hex.EncodeToString(h[:]), nil
}

func main() {
	flag := os.Getenv("FLAG")
	if len(flag) == 0 {
		log.Fatal("flag not set")
	}
	r := gin.Default()

	baseDirectory := "static"

	err := os.MkdirAll(baseDirectory, 0755)
	if err != nil {
		return
	}
	data := AppState{
		Flag:   flag,
		Author: "talgarr",
	}

	r.GET("/", HandleRoot(data, baseDirectory))
	r.GET("/flag", HandleFlag(data, baseDirectory))
	r.GET("/static/*filepath", func(c *gin.Context) {
		fp := c.Param("filepath")
		if strings.HasPrefix(fp, "/") {
			fp = fp[1:]
		}
		fullPath := filepath.Join(baseDirectory, fp)

		absBase, err := filepath.Abs(baseDirectory)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		absFull, err := filepath.Abs(fullPath)
		if err != nil || !strings.HasPrefix(absFull, absBase) {
			c.Status(http.StatusNotFound)
			return
		}

		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		c.File(absFull)
	})
	r.POST("/update_website", UploadTheme(baseDirectory))

	err = r.Run(":8080")
	if err != nil {
		return
	}
}

func HandleFlag(data AppState, baseDirectory string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tmpl, err := template.ParseFiles(filepath.Join(baseDirectory, "flag.html"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading template: %v", err)
			return
		}

		data.Flag = "REDACTED"
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
	}
}

func HandleRoot(data AppState, baseDirectory string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tmpl, err := template.ParseFiles(filepath.Join(baseDirectory, "index.html"))
		if err != nil {
			c.String(http.StatusInternalServerError, "Error loading template: %v", err)
			return
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			c.String(http.StatusInternalServerError, "Error rendering template: %v", err)
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", buf.Bytes())
	}
}

func UploadTheme(baseDirectory string) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("theme")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		open, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot open file"})
			return
		}

		zipReader, err := zip.NewReader(open, file.Size)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid zip file"})
			return
		}

		hashStr, err := CalculateZipHash(zipReader)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot calculate zip hash"})
			return
		}

		if !slices.Contains(allowedThemeHash, hashStr) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Theme not whitelisted", "hash": hashStr})
			return
		}

		for _, f := range zipReader.File {
			destPath := filepath.Join(baseDirectory, f.Name)

			destFile, err := os.Create(destPath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create file"})
				return
			}

			srcFile, err := f.Open()
			if err != nil {
				destFile.Close()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read zip entry"})
				return
			}

			_, err = io.Copy(destFile, srcFile)
			_ = srcFile.Close()
			_ = destFile.Close()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot write file"})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Theme updated successfully"})
	}
}
