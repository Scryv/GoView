#!/bin/bash
set -euo pipefail

require_root() {
  if [[ $EUID -ne 0 ]]; then
    echo "Error: must run as root"
    exit 1
  fi
}


require_root

go build -o goview main.go

sudo mv goview /usr/local/bin/

rm -f ~/.local/share/applications/goview.desktop
tee ~/.local/share/applications/goview.desktop > /dev/null <<EOF
[Desktop Entry]
Name=GoView Image Viewer
Exec=/usr/local/bin/goview %f
Type=Application
MimeType=image/png;image/jpg;
NoDisplay=true
EOF

xdg-mime default goview.desktop image/jpg
xdg-mime default goview.desktop image/png

echo "Installer done u can now use goview as command or just click images!"
