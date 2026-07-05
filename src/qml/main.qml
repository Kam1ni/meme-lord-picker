import QtQuick
import QtQuick.Controls
import QtQuick.Layouts

ApplicationWindow {
	width: bridge.windowSize.width
	height: bridge.windowSize.height
	minimumWidth: bridge.windowSize.width
	minimumHeight: bridge.windowSize.height
	maximumWidth: bridge.windowSize.width
	maximumHeight: bridge.windowSize.height
	visible: true
	title: "MemeLord Picker"

	ColumnLayout {
		anchors.fill: parent

		TextField {
			Layout.fillWidth: true
			placeholderText: "Search (Use # to search by hashtags (comma seperated))"
			onTextChanged: bridge.searchText = text
			focus: true
		}

		GridView {
			Layout.fillWidth: true
			Layout.fillHeight: true
			cellWidth: bridge.imageSize.width
			cellHeight: bridge.imageSize.height
			model: memeModel
			clip: true

			delegate: Item {
				width: 200
				height: 200
				
				AnimatedImage {
					id: img
					anchors.fill: parent
					anchors.margins: 4
					source: model.display.localUrl
					fillMode: Image.PreserveAspectCrop
					playing: true
					cache: true
				}

				HoverHandler {id: hover}

				TapHandler {
					onTapped: bridge.selectedPath = model.display.fileUrl
				}
			}
		}
	}
}