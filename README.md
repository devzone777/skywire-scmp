# skywire-scmp
Network utilities for Skywire Mainnet

This package provides the **SCMP** (Skywire Control Management Protocol) network level utilities for network engineers and developers to rapidly troubleshoot Skywire network issues and integrate Skywire into network monitoring applications. 

## Commands
 - visorverify >>>> Verifies if a visor pubkey is active on the Skywire network or not
 - visorcount >>>> Displays the number of visors currently active on the Skywire network
 - activevisors >>> Returns a list of all visors currently active on the Skywire network
 - genkeypair >>>> Generates a new Skywire Pubkey and SecKey

## Usage
Run:

```visorverify -k <pubkey>```

```visorcount```

```activevisors```

```genkeypair```

It is recommended that you use "grep" and or something like "jq" to enhance the output of activevisors

New features and utilities are actively being developed and will be added

### Installation

```make build```

Copy or move all binaries into /usr/local/bin


***Install Go***
Skywire supports go1.10+

***Go 1.10+ Installation and Setup***

[Golang 1.10+ Installation/Setup](https://github.com/devzone777/skycoin/blob/develop/INSTALLATION.md)
