#### Build And Deploy
```sh
make dockerize
```

#### To work with this project client side
POST: ip_address:port/get-document-data

#### Request: FORM File

#### Response: Document {
	documentNumber: string;
	documentType:   string;
	countryCode:    string;
	firstName:      string;
	lastName:       string;
	sex:            string;
	birthDate:      string;
	expireDate:     string;
#### }
