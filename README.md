## Synopsis

A docker compose scaffold for [hyperledger fabric v0.6](https://github.com/hyperledger/fabric). This scaffold contains a list of [custom ENROLLID and ENROLLSECRET](https://github.com/punneng/fabric/blob/7edcc81910b30eef2bf9b20752e8782780e400fd/membersrvc/membersrvc.yaml)

It consists of 3 particular nodes. The first one is Membersrvc where you can set the credentials for all peers in your network.
The next one is the root validation node for the first peer in your network and the last one is the children validation nodes of the root validation node.

[More information](https://hyperledger-fabric.readthedocs.io/en/latest/)

## Prerequisites

docker and docker-compose and vagrant version 1.8.6

## Installation

Fabric baseimage ,
```
docker pull hyperledger/fabric-baseimage:x86_64-0.2.2
docker tag <fabric-baseimage image id> hyperledger/fabric-baseimage:latest
```

Membersrvc node,
```
docker pull punneng/hyperledger-fabric-membersrvc
cd your/workplace/hyperledger-fabric-docker/membersrvc
docker-compose up
```

the second one, Root node with ENROLLID `test_vp0` and ENROLLID `hackathon000`
```
docker pull hyperledger/fabric-peer:latest
cd your/workplace/hyperledger-fabric-docker/osx_or_linux/root
docker-compose up
```

and the last one is the additional validation node with ENROLLID `test_vp1` and ENROLLID `hackathon001`
```
cd your/workplace/hyperledger-fabric-docker/osx_or_linux/peer
docker-compose up
```
