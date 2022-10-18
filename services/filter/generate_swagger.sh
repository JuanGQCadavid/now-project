#!/bin/bash
echo "Checking installation"
if ! command -v swagger &> /dev/null
then
    echo "Installing go-swagger"
    download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
    curl -o /usr/local/bin/swagger -L'#' "$download_url"
    chmod +x /usr/local/bin/swagger
fi

echo "Creating swagger directory"
mkdir swagger

echo "Generating go-swagger config"
swagger generate spec -o ./swagger/swagger.json --scan-models

echo "Copying dist"
cp -r ../swagger-dist/* ./swagger

echo "Updating json from dist"
#sed -i bak -e 's|url: "https://petstore.swagger.io/v2/swagger.json"|url: "/swagger/swagger.json"|g' swagger/swagger-initializer.js
sed -i 's|url: "https://petstore.swagger.io/v2/swagger.json"|url: "/filter/swagger/swagger.json"|g' swagger/swagger-initializer.js