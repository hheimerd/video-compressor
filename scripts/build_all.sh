wails build -platform windows/amd64
wails build -platform darwin/universal
cd ./build/bin
create-dmg --icon VideoCompressor.app 100 150 --app-drop-link 300 150 VideoCompressor.dmg VideoCompressor.app
cd ../..
