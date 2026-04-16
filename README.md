# GoView 
[![License: GPL v2](https://img.shields.io/badge/License-GPL_v2-blue.svg)](https://github.com/Scryv/GoView/blob/main/LICENSE)  
**GoView** is a open source image viewer made in golang together with cross platform ui toolkit called fyne it for now only supports png, jpg etc but will be able to work for a larger scope of file types in future  
> Still being worked on and for now just does the basics
<img src="https://github.com/Scryv/GoView/blob/main/gopherlogo/gophergoview.png" width="600px">

---

## How to install
> can also use the installer.sh script that got added
1. `git clone https://github.com/Scryv/GoView.git`
2. `cd GoView`
3. `go build -o goview main.go`
4. `sudo mv goview /usr/local/bin/`
5. `GoView` in a preferred directory<br>
**And enjoy :)**

### Add as primary ImgViewer(OPTIONAL)
> Image Cycling does not work here yet just single view
1. `nano ~/.local/share/applications/goview.desktop`
<pre>
[Desktop Entry]
Name=GoView Image Viewer
Exec=/usr/local/bin/goview %f
Type=Application
MimeType=image/png;image/jpg;
NoDisplay=true
</pre>
3. `xdg-mime default goview.desktop image/jpg`
4. `xdg-mime default goview.desktop image/png`
  
## Features  
### Cycling trough images 
when opening GoView from terminal it takes path of where u opened it and open the first picture u can use **Ctrl+Tab** to cycle forward and **Ctrl+Alt+tab** to go backwards
