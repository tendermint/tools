# Cosmos Network Monitor

> Graphical interface to view testnets on the Cosmos Network.

## Build Setup

``` bash
# install dependencies
yarn

# serve with hot reload at localhost:8080
yarn dev

# build for production with minification
yarn build
```

## Change Watched Node

In `./src/store/modules/blockchain.js` and `./src/store/modules/validators.js`, change the `let url = ` line to the Tendermint RPC of your choice.
