package static

import (
    "fmt"
    "log/slog"
    "mime"
    "net/http"
    "os"
    "path/filepath"
    "strconv"
    "strings"
    
    "github.com/gin-gonic/gin"
    "humpback/config"
    "humpback/pkg/utils"
)

type staticResourceInfo struct {
    FileName     string
    Content      []byte
    LastModified string
    ContentType  string
    Size         int64
}

var staticCache = map[string]*staticResourceInfo{}

func InitStaticsResource() (err error) {
    if config.HtmlArgs().Load {
        staticConfig := config.HtmlArgs()
        slog.Info("[Api] Init static resource to cache...")
        if staticCache, err = readFileToCache(staticConfig.Dir); err != nil {
            return fmt.Errorf("init static resource failed: %s", err)
        }
        slog.Info("[Api] Init static resource to cache complted.")
    }
    return nil
}

func readFileToCache(htmlDir string) (map[string]*staticResourceInfo, error) {
    cache := map[string]*staticResourceInfo{}
    err := filepath.Walk(htmlDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            content, err := os.ReadFile(path)
            if err != nil {
                return fmt.Errorf("read file %s failed: %s", path, err)
            }
            relPath, err := filepath.Rel(htmlDir, path)
            if err != nil {
                return fmt.Errorf("parse path %s failed: %s", path, err)
            }
            key := fmt.Sprintf("/%s", strings.ReplaceAll(filepath.Clean(relPath), `\`, "/"))
            contentType := mime.TypeByExtension(filepath.Ext(path))
            if contentType == "" {
                contentType = http.DetectContentType(content)
            }
            cache[key] = &staticResourceInfo{
                FileName:     info.Name(),
                Content:      content,
                LastModified: info.ModTime().Format(http.TimeFormat),
                ContentType:  contentType,
                Size:         info.Size(),
            }
        }
        return nil
    })
    if err != nil {
        return nil, err
    }
    return cache, nil
}

// web  每次从文件读取静态资源
func web(c *gin.Context) {
    htmlConfig := config.HtmlArgs()
    if c.Request.URL.String() != "/" && utils.FileExist(fmt.Sprintf("%s/%s", htmlConfig.Dir, c.Request.URL.Path)) {
        c.File(fmt.Sprintf("%s/%s", htmlConfig.Dir, c.Request.URL.Path))
    } else {
        c.File(fmt.Sprintf("%s/index.html", htmlConfig.Dir))
    }
}

// Web  从缓存中读取静态资源
func Web(c *gin.Context) {
    resourceInfo := staticCache["/index.html"]
    isIndexHtml := true
    if c.Request.URL.String() != "/" && staticCache[c.Request.URL.Path] != nil {
        resourceInfo = staticCache[c.Request.URL.Path]
        isIndexHtml = false
    }
    c.Header("Accept-Ranges", "bytes")
    c.Header("Content-Type", resourceInfo.ContentType)
    c.Header("Content-Length", strconv.FormatInt(resourceInfo.Size, 10))
    c.Header("Last-Modified", resourceInfo.LastModified)
    if isIndexHtml {
        c.Header("Cache-Control", "public, max-age=86400")
    } else {
        c.Header("Cache-Control", "public, max-age=31536000")
    }
    c.Data(http.StatusOK, resourceInfo.ContentType, resourceInfo.Content)
}
