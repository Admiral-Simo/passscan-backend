# Download Go from [here](https://go.dev/).

#### Install Tesseract (a model for detecting text from images):
```sh
sudo apt install tesseract-ocr
```

#### make sure that you have tesseract installed using this command
```sh
tesseract -v
```

#### Compile the program
```sh
go build cmd/main.go
```

###### Usage: ./program --passport example_image.{jpeg, png, jpg ...}
