# Simple Shamir's Secret Sharing (s4)

With **S**imple **S**hamir's **S**ecret **S**haring (s4) I want to provide you an easy to use interface for this beautiful little piece of math.

Please note that s4 is provided as it is and I do not take responsibility for any bugs. s4 is a tiny layer around [hashicorp vault shamir](https://github.com/hashicorp/vault) and golang's [AES encryption](https://github.com/gtank/cryptopasta/blob/master/encrypt.go).

## 📚 Usage as library

You can use s4 as normal go library in your go projects by importing it via `import "github.com/simonfrey/s4"` and en-/decrypt
bytes to byte shares and vice versa

## 🏠 Building the WASM for the frontend

I assume you have a [go](https://golang.org/) build environment setup in your machine. 

In order to build & pack the web assembly file for the frontend please use the following command in the top level directory:
```
./build.sh
```

This will build you the required file. You now can copy to `build` folder to your web server (or use it locally) and it
should run s4 as intended.

## 💸 Report Bugs & Tip

Please use [Github Issues](https://github.com/simonfrey/s4/issues) in order to report bugs

💸 If you want to tip me for my work on this project [feel free to do so](https://simon-frey.com/tip) 💸

## 🗣 Discussions

s4 was broadly discussed on [HackerNews](https://news.ycombinator.com/item?id=23541949), and was mentioned in [golang weekly](https://golangweekly.com/issues/317) in June 2020.

## 📃 License
[MIT License](https://github.com/simonfrey/s4/blob/master/LICENSE)
