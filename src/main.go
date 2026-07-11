package main

import (
	"fmt"
	"meme-lord-picker/cache"
	"meme-lord-picker/config"
	"meme-lord-picker/fetcher"
	"meme-lord-picker/memelord"
	"meme-lord-picker/notify"
	"meme-lord-picker/windowrules"
	"net/url"
	"os"
	"os/exec"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/mainthread"
	"github.com/mappu/miqt/qt6/qml"
)

const (
	RoleFileThumbnailUrl = int(qt6.UserRole) + iota
	RoleFileUrl
	RoleName
)

func main() {
	conf := config.GetConfig()
	if conf.MemeLordApiUrl == "" {
		panic("MEME_PICKER_API_URL is not set")
	}
	client := memelord.CreateClient(conf.MemeLordApiUrl, conf.MemeLordApiToken)
	cachingServer := cache.CreateCachingServer()
	go cachingServer.Run()
	defer cachingServer.Close()

	windowrules.AttemptSetWindowPositionRule()

	qt6.NewQApplication(os.Args)

	result := memelord.MemesResponse{}
	model := qt6.NewQAbstractListModel()
	fetcher := fetcher.CreateFetcher(client, func(mr memelord.MemesResponse) {
		mainthread.Start(func() {
			result = mr
			model.BeginResetModel()
			model.EndResetModel()
		})
	})
	fetcher.QueueFetch("")

	model.OnRowCount(func(parent *qt6.QModelIndex) int {
		return result.Count
	})
	model.OnData(func(index *qt6.QModelIndex, role int) *qt6.QVariant {
		if !index.IsValid() || index.Row() >= result.Count {
			return qt6.NewQVariant()
		}
		m := result.Results[index.Row()]
		url := fmt.Sprintf("%s?url=%s", cachingServer.GetUrl(), url.QueryEscape(m.ImageUrl))
		return qt6.NewQVariant20(map[string]qt6.QVariant{
			"localUrl": *qt6.NewQVariant14(url),
			"fileUrl":  *qt6.NewQVariant14(m.ImageUrl),
			"name":     *qt6.NewQVariant14(m.Title),
		})
	})

	bridge := qml.NewQQmlPropertyMap()
	bridge.OnValueChanged(func(key string, value *qt6.QVariant) {
		switch key {
		case "searchText":
			fetcher.QueueFetch(value.ToString())
		case "selectedPath":
			cmd := exec.Command("wl-copy", value.ToString())
			cmd.Stdin = nil
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				fmt.Println("Failed to run wl-copy command", err.Error())
				return
			}
			fmt.Println("Coppied to clipboard")
			notify.Notify("Meme coppied to clipboard", value.ToString())
			qt6.QCoreApplication_Exit()
		}
	})

	bridge.SetProperty("windowSize", qt6.NewQVariant22(qt6.NewQSize2(conf.Window.Width, conf.Window.Height)))
	bridge.SetProperty("imageSize", qt6.NewQVariant22(qt6.NewQSize2(conf.Window.ImageWidth, conf.Window.ImageHeight)))

	theme := config.GetThemeQPallete(qt6.QApplication_Palette(nil))
	themeMap := qml.NewQQmlPropertyMap()
	theme.SetQmlProperties(themeMap)

	engine := qml.NewQQmlApplicationEngine()
	engine.RootContext().SetContextProperty("bridge", bridge.QObject)
	engine.RootContext().SetContextProperty("memeModel", model.QObject)
	engine.RootContext().SetContextProperty("theme", themeMap.QObject)
	engine.Load(qt6.NewQUrl3("qrc:/qml/main.qml"))

	qt6.QApplication_Exec()
}
