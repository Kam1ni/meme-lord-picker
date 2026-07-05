package main

import (
	"fmt"
	"meme-lord-picker/cache"
	"meme-lord-picker/config"
	"meme-lord-picker/memelord"
	"meme-lord-picker/windowrules"
	"net/url"
	"os"
	"os/exec"

	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/qml"
)

const (
	RoleFileThumbnailUrl = int(qt6.UserRole) + iota
	RoleFileUrl
	RoleName
)

func main() {
	config := config.GetConfig()
	if config.MemeLordApiUrl == "" {
		panic("MEME_PICKER_API_URL is not set")
	}
	client := memelord.CreateClient(config.MemeLordApiUrl, config.MemeLordApiToken)
	cachingServer := cache.CreateCachingServer()
	go cachingServer.Run()
	defer cachingServer.Close()

	windowrules.AttemptSetWindowPositionRule()

	qt6.NewQApplication(os.Args)

	result := memelord.MemesResponse{}
	model := qt6.NewQAbstractListModel()
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
			"fileUrl": *qt6.NewQVariant14(url),
			"name":    *qt6.NewQVariant14(m.Title),
		})
	})

	bridge := qml.NewQQmlPropertyMap()
	bridge.OnValueChanged(func(key string, value *qt6.QVariant) {
		switch key {
		case "searchText":
			fmt.Println("Fetching with title", value.ToString())
			if value.ToString() == "" {
				result = memelord.MemesResponse{}
			} else {
				var err error
				result, err = client.FetchMemes(memelord.Query{Title: value.ToString()})
				if err != nil {
					fmt.Println("Failed to fetch memes\n", err.Error())
				} else {
					fmt.Println(result)
				}
			}
			fmt.Println(result.Count)
			model.BeginResetModel()
			model.EndResetModel()
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
			qt6.QCoreApplication_Exit()
		}
	})

	bridge.SetProperty("windowSize", qt6.NewQVariant22(qt6.NewQSize2(config.Window.Width, config.Window.Height)))
	bridge.SetProperty("imageSize", qt6.NewQVariant22(qt6.NewQSize2(config.Window.ImageWidth, config.Window.ImageHeight)))

	engine := qml.NewQQmlApplicationEngine()
	engine.RootContext().SetContextProperty("bridge", bridge.QObject)
	engine.RootContext().SetContextProperty("memeModel", model.QObject)
	engine.Load(qt6.NewQUrl3("qrc:/qml/main.qml"))

	qt6.QApplication_Exec()
}
