Following Packt tutorial "Hyperledger for Blockchain Applications".

## Session 1.2: Setup terminal, zsh, and editor (optional)

```
$ sudo apt install zsh
```

Enter zsh:
```
$ zsh
```
Type 0.

```
$ curl -L git.io/antigen > antigen.zsh
$ vim ~/.zshrc
# Delete everything, enter:
source ~/private/viridian/backend/hyperledger/composer/antigen.zsh

antigen use oh-my-zsh

antigen bundle git
antigen bundle zsh-users/zsh-syntax-highlighting

antigen theme https://github.com/denysdovhan/spaceship-prompt spaceship

antigen apply
<ESC>:wq<ENTER>
---
$ source ~/.zshrc
```

Terminal-multiplexer byobu:
```
$ sudo apt install byobu
$ byobu-enable # (byobu-disable to disable it)
```

Install Atom text editor from https://atom.io/.

Run Atom and install the 'composer-atom-plugin' via Edit -> Preferences ->
Install -> enter 'composer-atom-plugin' into text field.



## Session 1.3: Install prerequisites for Hyperledger Composer

```
$ curl -O https://hyperledger.github.io/composer/latest/prereqs-ubuntu.sh
$ source ./prereqs-ubuntu.sh
# Logout and in again
$ exit
```

Install Hyperledger Composer etc.:

```
$ npm i -g composer-cli
$ npm i -g composer-rest-server
$ npm i -g generator-hyperledger-composer
$ npm i -g yo
$ npm i -g composer-playground
```

Maybe update npm packages according to these warnings:

```
npm WARN deprecated core-js@2.3.0: core-js@<2.6.8 is no longer maintained. Please, upgrade to core-js@3 or at least to actual version of core-js@2.
npm WARN deprecated hoek@4.2.1: This version has been deprecated in accordance with the hapi support policy (hapi.im/support). Please upgrade to the latest version to get the best features, bug fixes, and security patches. If you are unable to upgrade at this time, paid support is available for older versions (hapi.im/commercial).
npm WARN deprecated json3@3.3.2: Please use the native JSON object instead of JSON 3
npm WARN deprecated samsam@1.3.0: This package has been deprecated in favour of @sinonjs/samsam
npm WARN deprecated mixin-deep@1.3.1: Critical bug fixed in v2.0.1, please upgrade to the latest version.
npm WARN deprecated set-value@2.0.0: Critical bug fixed in v3.0.1, please upgrade to the latest version.
npm WARN deprecated set-value@0.4.3: Critical bug fixed in v3.0.1, please upgrade to the latest version.
npm WARN deprecated cross-spawn-async@2.2.5: cross-spawn no longer requires a build toolchain, use it instead
npm WARN deprecated mixin-deep@1.3.1: Critical bug fixed in v2.0.1, please upgrade to the latest version.
npm WARN deprecated hawk@3.1.3: This module moved to @hapi/hawk. Please make sure to switch over as this distribution is no longer supported and may contain bugs and critical security issues.
npm WARN deprecated sntp@1.0.9: This module moved to @hapi/sntp. Please make sure to switch over as this distribution is no longer supported and may contain bugs and critical security issues.
npm WARN deprecated hoek@2.16.3: This version has been deprecated in accordance with the hapi support policy (hapi.im/support). Please upgrade to the latest version to get the best features, bug fixes, and security patches. If you are unable to upgrade at this time, paid support is available for older versions (hapi.im/commercial).
npm WARN deprecated cryptiles@2.0.5: This version has been deprecated in accordance with the hapi support policy (hapi.im/support). Please upgrade to the latest version to get the best features, bug fixes, and security patches. If you are unable to upgrade at this time, paid support is available for older versions (hapi.im/commercial).
npm WARN deprecated boom@2.10.1: This version has been deprecated in accordance with the hapi support policy (hapi.im/support). Please upgrade to the latest version to get the best features, bug fixes, and security patches. If you are unable to upgrade at this time, paid support is available for older versions (hapi.im/commercial).

```

Install Fabric:

```
$ mkdir fabric-dev-servers
$ cd fabric-dev-servers
$ curl -O https://raw.githubusercontent.com/hyperledger/composer-tools/master/packages/fabric-dev-servers/fabric-dev-servers.tar.gz
$ tar -xvf fabric-dev-servers.tar.gz
# Keep this output for comparison later:
$ docker images
$ ./downloadFabric.sh
# See the new downloaded fabric images:
$ docker images
```

## Session 3.1: Start a Hyperledger Fabric network

Maybe optional:
```
# Look up HLF version in fabric-dev-servers/_loader.sh
# Set it in .zshenv (if using zsh for Hyperledger development)
vi ~/.zshenv
---
export HL_FABRIC_VERSION=hlfv12
---
```

Optional: Look at steps in the scripts `fabric-dev-servers/fabric-scripts/hlfv12/{downloadFabric.sh,startFabric.sh}`.

Basically, `downloadFabric.sh` is just a series of docker pull commands.

`startFabric.sh` consists basically of setting a few variables and then running
`docker-compose -f $DOCKER_FILE down` and `docker-compose -f $DOCKER_FILE up -d`
and some commands on the peer0 (the only peer) to create the channel and join it:
```
docker exec peer0.org1.example.com peer channel create ...
```
and:
```
docker exec -e "..." peer0.org1.example.com peer channel join ...
```

If you need to modify the Hyperledger Fabric settings, edit the docker file
under `fabric-dev-servers/fabric-scripts/hlfv12/composer/docker-compose.yml`
and/or the downloadFabric/startFabric scripts. You may want to create custom
cryptographic material that goes along with a custom configuration (like MSP
data, CA certificate and key and public and private keys of peers and orderers).

When finished, run:
```
$ ./startFabric.sh
```
and look at running containers with
```
docker ps
```
You can also look at logs of individual containers with
```
docker logs peer0.org1.example.com
docker logs orderer.example.com
...
```
Stop the containers with ./stopFabric.sh, doing `docker-compose -f $DOCKER_FILE stop`.
With `docker-compose -f $DOCKER_FILE start`, you can start them again. With
`docker-compose -f $DOCKER_FILE down`, the containers are deleted
(startFabric.sh first runs docker-compose down, then docker-compose up).

## Session 3.2: Create a peer admin card

Optional: create your own cryptographic material and then run the commands in
createPeerAdminCard.sh. Or simply run:

```
$ ./createPeerAdminCard.sh
```

Look at your imported card with
```
$ composer card list
```
Use
```
$ composer network ping --card PeerAdmin@hlfv1
```
to test the network associated with the network card names 'PeerAdmin@hlfv1'.

Use `composer card delete -c PeerAdmin@hlfv1` to delete the card you created
with name 'PeerAdmin@hlfv1'.

## Session 3.3: Network configuration and creating crpytographic material

Optional for first development, but needed for running in production. However,
the discussed tool `cryptogen` is not suitable for production. Better use the
Hyperledger Fabric CA.

```
$ mkdir fabric-tools
$ cd fabric-tools
$ curl -O https://raw.githubusercontent.com/hyperledger/fabric/master/scripts/bootstrap.sh
# Download only binaries (skip samples and docker images):
$ bash bootstrap.sh -s -d
$ vi ~/.zshenv
# Add:
---
export PATH=$PATH:/path/to/fabric-tools/bin
---
```

This downloaded some Fabric tools, most importantly cryptogen and configtxgen
for generating cryptographic material and a network/channel configuration
transaction.

Edit the file `fabric-dev-servers/fabric-scripts/hlfv12/composer/crypto-config.yaml`
to suit your needs. You define the organizational structure of your network
there.

```
$ cd fabric-dev-servers/fabric-scripts/hlfv12/composer/
$ cryptogen generate --config=crypto-config.yaml --output=new-crypto-config
```

Your new cryptographic material now lives in the directory `new-crypto-config`,
side-by-side with the old one in `crypto-config`.

Edit `fabric-dev-servers/fabric-scripts/hlfv12/composer/configtx.yaml` to suit
your needs. Change MSPDir paths to `new-crypto-config`.

```
$ cd fabric-dev-servers/fabric-scripts/hlfv12/composer/
$ FABRIC_CFG_PATH=. configtxgen -profile ComposerOrdererGenesis -outputBlock new-orderer-genesis.block
$ FABRIC_CFG_PATH=. configtxgen -profile ComposerChannel -channelID composerchannel -outputCreateChannelTx new-composer-channel.tx
```

Edit `fabric-dev-servers/fabric-scripts/hlfv12/composer/docker-compose.yml` to
suit your needs.

## Session 3.4: Play with Composer Playground

Optional: Start the playground and play around with an example, e.g. the marbles
network:

```
$ cd fabric-dev-servers
$ ./startFabric.sh
$ ./createPeerAdminCard.sh
$ composer-playground
```

## Session 4.1: Building a Network

Generate a business network skeleton directory with all the necessary files via
the template generator Yeoman (yo):

```
$ yo hyperledger-composer:businessnetwork
---
Welcome to the business network generator
? Business network name: viridian
? Description: The Viridian Project, an idea for a sustainable and fair economy, https://www.viridian-project.org
? Author name:  Markus Voge
? Author email: info@viridian-project.org
? License: GNU GPLv3
? Namespace: org.viridian
? Do you want to generate an empty template network? Yes: generate an empty template network
---
```

### Session 5.1/4.2: Create business network archive

To test your network definition, create a business network archive that you can load into the Composer Playground.

```
$ cd viridian # change into the directory created by yo
$ composer archive create --sourceType dir --sourceName .
```

Then you can start the Composer playground

```
$ composer-playground
```

and create a new Web Browser connection and upload the file created by
'composer archive create', e.g. `viridian@0.0.1.bna`.
