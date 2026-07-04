import QtQuick
import QtQuick.Controls
import QtQuick.Layouts

ApplicationWindow {
	width: 400
	height: 400
	minimumWidth: 400
	minimumHeight: 400
	maximumWidth: 400
	maximumHeight: 400
	visible: true
	title: "MemeLord Picker"

	ColumnLayout {
		anchors.fill: parent

		TextField {
			Layout.fillWidth: true
			placeholderText: "Search..."
			onTextChanged: bridge.searchText = text
		}

		GridView {
			Layout.fillWidth: true
			Layout.fillHeight: true
			cellWidth: 200
			cellHeight: 200
			model: memeModel

			delegate: Item {
				width: 200
				height: 200
				
				AnimatedImage {
					id: img
					anchors.fill: parent
					anchors.margins: 4
					source: model.display.filePreviewUrl
					fillMode: Image.PreserveAspectCrop
					asynchronous: true
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