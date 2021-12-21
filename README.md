# Substrate External Adapter

The Substrate Adapter is used to send the Oracle answer back to our parachain.

## Installation
```go install```

install [subkey](https://docs.substrate.io/v3/tools/subkey/)

Install [polkadot wallet](https://polkadot.js.org/extension/) to create account

## Configuration

### Inspecting keys
In order to get the **SA_PRIVATE_KEY**
- 1.Create a account in polkadot wallet
- 2.Inspecting keys using command below
```
subkey inspect "caution juice atom organ advance problem want pledge someone senior holiday very"
```
output
```
Secret phrase `caution juice atom organ advance problem want pledge someone senior holiday very` is account:
  Secret seed:       0xc8fa03532fb22ee1f7f6908b9c02b4e72483f0dbd66e4cd456b8f34c6230b849
  Public key (hex):  0xd6a3105d6768e956e9e5d41050ac29843f98561410d3a47f9dd5b3b227ab8746
  Public key (SS58): 5Gv8YYFu8H1btvmrJy9FjjAWfb99wrhV3uhPFoNEr918utyR
  Account ID:        0xd6a3105d6768e956e9e5d41050ac29843f98561410d3a47f9dd5b3b227ab8746
  SS58 Address:      5Gv8YYFu8H1btvmrJy9FjjAWfb99wrhV3uhPFoNEr918utyR
```
The Secret seed is our **SA_PRIVATE_KEY**

Then, open **adapter.sh** to change your information:
We need to configure the connection settings between the adapter and the substrate chain:
- **SA_ENDPOINT** : our parachainchain web socket endpoint, _ws://172.17.0.1:9944_
- **SA_PRIVATE_KEY** : your parachain oracle account's Secret seed

Then, run ```sh adapter.sh```it will automatically generate substrate_external_initiator.env in current directory

## Run
And we can run the substrate-adapter container:

we will use [Docker](https://hub.docker.com/r/smartcontract/substrate-adapter) to run this component, and find the necessary documentation on its [github repository](https://github.com/smartcontractkit/substrate-adapter).

 ```shell
docker run -d -p 8081:8080 --name substrate-adapter --env-file substrate_adapter.env smartcontract/substrate-adapter
```