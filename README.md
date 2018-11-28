---
description: NKN mining web application.
---

# nkn-mining

## How to build 🍀 

 **Note: This build script only tested on Mac OS X**

1. intall npm & golang first
2. clone or download this repository
3. cd to the directory
4. run ./build.sh script
5. when script done, you will find the output file under ./dist directory

## How to use 🍀 

### What do you need first!

1. A computer to run the NKNMining application
2. A public ipv4 address
3. Permission of port open between 30000 ~ 30003

### step 0 

Upload the package to your server and run NKNMining program.

NKNMining will output a 40 character length ‘serial number’ .

```text
$ ./NKNMining
$ serial number(sn): NKN-fef5d6b5-ba56-11e8-9c0b-260049909001
```

### step 1

Open http://ip:8181/web/  to create your control account.e.g:

```text
http://127.0.0.1:8181/web/
```

![](.gitbook/assets/step1.png)

**Note: ‘System initialization serial number’ here is the ‘serial number’ string from step 0.**

### step 2

Generate a wallet for node running.

![](.gitbook/assets/step2.png)

**Note: password in this page is for wallet.**

### step 3

In this step you will see your wallet information.Also you can click ‘Download wallet’ button to save your wallet file.

![](.gitbook/assets/step3.png)

### step 4

The NKN node is setup already! Now you can turn on/off your node to mining NKN or transfer some NKN you mined to someone else.

![](.gitbook/assets/step4.png)

### Q&A

Discord: [https://discord.gg/wUBNKFm](https://discord.gg/wUBNKFm)

 **Note: This repository is in the early development stage and may not have all functions working properly. It should be used only for testing now.**

