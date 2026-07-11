# Meme Lord Picker
A wayland companion app for LRVT's [MemeLord](https://github.com/l4rm4nd/MemeLord). It allows you to quickly search and copy the url of your memes.

## Installation
Arch users can install it from the AUR using their prefered AUR helper
```bash
# For yay
yay -S meme-lord-picker

# For paru
paru -S meme-lord-picker
```

## Build from source
Make sure you have gcc, golang and qt6 installed on your system.
```bash
# Debian / Ubuntu
apt install build-essential golang-go qt6-base-dev

# Fedora
dnf install qt6-qtbase-devel qscintilla-qt6-devel qt6-qtcharts-devel qt6-qtmultimedia-devel qt6-qtpdf-devel qt6-qtpositioning-devel qt6-qtsvg-devel qt6-qttools-devel qt6-qtwebchannel-devel qt6-qtwebengine-devel qt6-qtdeclarative-devel golang

# Arch Linux
pacman -S pkg-config qt6-base gcc go
```
Then run build.sh to compile the application. The first build can take up to 10 minutes.
```bash
./build.sh
```

## Setup
First run the built binary.
```bash
./meme-lord-picker
```
First time you run it it will ask you to complete the config file at $XDG_CONFIG_HOME/meme-lord-picker/config.json.
```
Created default config file at /home/kamil/.config/meme-lord-picker/config.json
Please configure memeLordApiUrl and memeLordApiToken
```
Fill in those two values with your meme lord instance url and your api token. You can also play with window and image sizing. For example:
```json
{
	"memeLordApiToken": "asfiojasd123....",
	"memeLordApiUrl": "https://memelord.my-domain.com",
	"window": {
		"Width": 400,
		"Height": 400,
		"ImageWidth": 200,
		"ImageHeight": 200
	}
}
```

Next time you run meme-lord-picker it will open a window with your memes.

![Screenshot](./assets/picker.png)


## Theming
You can theme MemeLord Picker to match your environment. Create a theme file next to your config file `~/.config/meme-lord-picker/theme`. It uses the same format as an env file.
For exmaple a ugly conifg:
```env
WINDOW=#FF0000
TEXT_FIELD=#000000
TEXT_FIELD_TEXT=#FFFFFF
TEXT_FIELD_PLACEHOLDER=#00FF00
TEXT_FIELD_BORDER=#0000FF
TEXT_FIELD_BORDER_ACTIVE=#FFFFFF
```

## Tips
* Assign a hotkey in your desktop environment to open meme-lord-picker with a press of a button. Use `pkill -9 meme-lord-picke || meme-lord-picker` as the command so it toggles the picker on repeated presses.
* If you're not on hyprland, create a window rule with target class meme-lord-picker so that the window creates at your cursors position.
