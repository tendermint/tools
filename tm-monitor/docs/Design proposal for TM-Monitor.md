# Design proposal for Tendermint Monitor

Tendermint monitor is designed to be the tool for all related parties to have a close look about the network. 

## Design Goal for monitoring the network

The design goal of Tendermint Monitor is to track **Liveness** & **Safety** of the whole network and certain validator.

 The liveness of the network is in jeopardy if:

* many validators are offline
* the system barely gather +2/3 votes in recent blocks
* validators have to wait for many rounds before commiting new block


The safety of the network is in jeopardy if:

* There exists a monopoly of Atom token, +1/3

* Not enough Atoms are bonded

* Certain validator attracts too many delegation


## Design Goal for monitoring a validator

The validator must monitor its **availability** and **safety** closely.  There are some early-stage indicators to be noticed:

*  Are most of previous precommits included in blocks?
* How many proposals it missed before?


The safety of a validator is under pressure if :

* it's not connected with enough persistent peers 
* there is not enought peers in its address book

## Tendermint Monitor Atchitecture

The architecture of new version of monitor will be like this:

![WechatIMG114](https://github.com/bianjieai/cosmos-sdk/raw/master/docs/tm-monitor/pic/WechatIMG114.jpeg)

The data will be stored in Mongo to facilitate better querying.

## Metrics 

![metrics](https://github.com/bianjieai/cosmos-sdk/raw/master/docs/tm-monitor/pic/metrics.png)



link: https://docs.google.com/spreadsheets/d/13TbMrG0PlahN-vr_Z8hbQJIQHxqM0g4C4AfF-kDYHPI/edit?usp=sharing 

