## Dii
Portable DNS and IP resolver tool

## Installation
### Source
```sh
# Clone the repo
git clone https://github.com/kiwimoe/dii.git

# Install golang and make and run
make

# To install to path run
sudo make install
```
### Prebuild
Head over to [releases](https://github.com/kiwimoe/dii/releases) and grab the latest one. (Only for x86_64 CPU arch)

## Usage
```sh
# Retrives info about the IP address
## Note: IP "37.19.205.148" is a VPN IP which is located at Japan
dii -ip 37.19.205.148

# Retrives info about the DNS
dii -dns linux.org
```

## Images
![image](https://user-images.githubusercontent.com/71683721/184000240-0b7a15c5-cf28-4a8b-b323-9eb24679d38a.png)

![image](https://user-images.githubusercontent.com/71683721/184000479-542d3e13-ae36-4ba5-a330-e86dd67a2def.png)
