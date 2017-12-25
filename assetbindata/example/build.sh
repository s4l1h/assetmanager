rm -rf build
mkdir build
go run main.go --build yes  # make temp asset File
#GOOS=linux go build -ldflags="-s -w";upx example
go build -ldflags="-s -w"
rm generatedAsset.go # remove temp asset File
mv example build/example #copy binary file to build directory