package config

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/qml"
)

func getThemeFilePath() string {
	return path.Join(getConfigDirPath(), "theme")
}

type ColorTheme struct {
	Window                string
	TextField             string
	TextFieldText         string
	TextFieldPlaceholder  string
	TextFieldBorder       string
	TextFieldBorderActive string
}

func GetThemeQPallete(defaultPalette *qt6.QPalette) ColorTheme {
	_, err := os.Stat(getThemeFilePath())
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Theme file %s does not exist\n", getThemeFilePath())
			return ColorTheme{}
		}
		panic(fmt.Sprintf("Failed to check if theme file (%s) exists\n%s", getThemeFilePath(), err.Error()))
	}
	fmt.Printf("Loading theme file %s\n", getThemeFilePath())
	paletteMap, err := godotenv.Read(getThemeFilePath())
	if err != nil {
		panic(fmt.Sprintf("Failed to read theme file (%s)\n%s", getThemeFilePath(), err.Error()))
	}

	getV := func(key string, defaultValue string) string {
		v, ok := paletteMap[key]
		if ok {
			return v
		}
		return defaultValue
	}

	return ColorTheme{
		Window:                getV("WINDOW", defaultPalette.Window().Color().Name()),
		TextField:             getV("TEXT_FIELD", defaultPalette.Base().Color().Name()),
		TextFieldText:         getV("TEXT_FIELD_TEXT", defaultPalette.ButtonText().Color().Name()),
		TextFieldPlaceholder:  getV("TEXT_FIELD_PLACEHOLDER", defaultPalette.PlaceholderText().Color().Name()),
		TextFieldBorder:       getV("TEXT_FIELD_BORDER", defaultPalette.WindowText().Color().Name()),
		TextFieldBorderActive: getV("TEXT_FIELD_BORDER_ACTIVE", defaultPalette.Highlight().Color().Name()),
	}
}

func (c ColorTheme) SetQmlProperties(m *qml.QQmlPropertyMap) {
	m.SetProperty("window", qt6.NewQVariant11(c.Window))
	m.SetProperty("textField", qt6.NewQVariant11(c.TextField))
	m.SetProperty("textFieldText", qt6.NewQVariant11(c.TextFieldText))
	m.SetProperty("textFieldPlaceholder", qt6.NewQVariant11(c.TextFieldPlaceholder))
	m.SetProperty("textFieldBorder", qt6.NewQVariant11(c.TextFieldBorder))
	m.SetProperty("textFieldBorderActive", qt6.NewQVariant11(c.TextFieldBorderActive))
}
