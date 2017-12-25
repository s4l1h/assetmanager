rm generatedAsset.go
rm -rf binarybuild
mkdir binarybuild
go run main.go --build yes
#GOOS=linux go build -ldflags="-s -w"
go build -ldflags="-s -w"
rm generatedAsset.go
mv example_binary binarybuild/example_binary