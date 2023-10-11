# Welcome to RabbitX Bot Repository

This repository contains the code for creating bots for the RabbitX API. If you're new to this, don't worry! We've got you covered. Here's a step-by-step guide to get you started.

## Getting Started

1. **Environment Setup:** Create a .env file in the root directory of the project. This file should contain the following variables:
```bash
API_URL = "https://api.testnet.rabbitx.io"
WS_URL = "wss://api.testnet.rabbitx.io/ws"

WALLET = ""
PRIVATE_KEY = ""
API_SECRET = ""
API_KEY = ""
REFRESH_TOKEN = ""
PUBLIC_JWT = ""
PRIVATE_JWT = ""
API_KEY_EXPIRED = ""
```
These variables are used to authenticate your bot with the RabbitX API. Make sure to replace the empty strings with your actual credentials.

2. **Running the Bot:** Once you've set up your environment, you can launch the bot on the testnet using the following command:
```bash
make run
```
This command will start your bot and connect it to the RabbitX API.

3. **Documentation:** To understand more about the code and how it works, you can generate and view the documentation. Use the following command to generate the documentation:
```bash
go get golang.org/x/tools/cmd/godoc

sudo apt install golang-golang-x-tools

make doc

# then open http://127.0.0.1:6060/pkg/rabbitx-client/
```
This command will generate the documentation and provide a URL. Open this URL in your web browser to view the documentation.

And that's it! You're now ready to start creating your own bots for the RabbitX API. Happy coding!
