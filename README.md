# Download Go from [here](https://go.dev/).

#### Install Tesseract (a model for detecting text from images):
```sh
sudo apt install tesseract-ocr
```

#### make sure that you have tesseract installed using this command
```sh
tesseract -v
```

#### build the server
```sh
make build
```

###### Usage: ./main
###### it will listen on port 8080
###### you can in the main.go

###### you should upload the file using http to this route /get-passport-data
###### You'll get as a response the data attached to the Passport Card or ID Card
