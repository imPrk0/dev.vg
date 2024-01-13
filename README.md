<center><h1>🛠️&nbsp;dev.vg&nbsp;👨‍💻</h1></center>

## 📖 The Story

One day, I saw a domain name: `dev.vg`, but I didn't think about its role.
I only know that it can provide services for developers.
Suddenly one day, I saw a project [`sdfsdf.dev`](https://sdfsdf.dev),
combined with my own ideas, so there is this site.

## 🤔 How to use

### 1. Shorten URL

The route part is a `6-16` bit string composed of `a-z A-Z 0-9`,
which is a short link.

```
https://dev.vg/[a-zA-Z0-9]{6,16}
```

### 2. Image placeholder

❤️ Thanks to [sdfsdf.dev](https://sdfsdf.dev/),
this will redirect to the corresponding address of [sdfsdf.dev](https://sdfsdf.dev/).

```
https://dev.vg/[width]x[height].[png / jpg]
```

Custom background color:

```
https://dev.vg/[width]x[height].[png / jpg],[background color],[text color]
```

Example:
- 600x400 JPG image  
[https://dev.vg/600x400.jpg](https://dev.vg/600x400.jpg)
- 800x200 PNG image  
[https://dev.vg/800x200.png](https://dev.vg/800x200.png)
- 300×200 PNG image, with blue background color and yellow text color  
[https://dev.vg/300x200.png,blue,yellow](https://dev.vg/300x200.png,blue,yellow)
- 600×400 JPG image, with black background color and white text color.  
[https://dev.vg/600x400.jpg,000000,ffffff](https://dev.vg/600x400.jpg,000000,ffffff)

### 3. Random words, sentences, paragraphs and more

#### Words

```
https://dev.vg/[0-9]+w[,]?
```

Example:
- 5 random words  
[https://dev.vg/5w](https://dev.vg/5w)
- 5 random words array  
[https://dev.vg/5w,](https://dev.vg/5w,)

#### Sentences

```
https://dev.vg/[0-9]+s[,]?
```

Example:
- 5 random sentences  
[https://dev.vg/5s](https://dev.vg/5s)
- 5 random sentences array  
[https://dev.vg/5s,](https://dev.vg/5s,)

#### Paragraphs

```
https://dev.vg/[0-9]+p[,]?
```

Example:
- 3 random paragraphs  
[https://dev.vg/3p](https://dev.vg/3p)
- 3 random paragraphs array  
[https://dev.vg/3p,](https://dev.vg/3p,)

#### Text (Fixed Length of Words)

```
https://dev.vg/[0-9]+t
```

Example:
- Random words with fixed length of 20  
[https://dev.vg/20t](https://dev.vg/20t)

### 4. Random pick

```
https://dev.vg/random,[values with comma separated]
```

Example:
- Randomly select a value from `Apple`, `Banana`, `Orange`, `Pear`, `Peach`  
[https://dev.vg/random,Apple,Banana,Orange,Pear,Peach](https://dev.vg/random,Apple,Banana,Orange,Pear,Peach)

## 📆 Other

I have no more ideas at the moment, and I have run out of inspiration.
If you have better ideas, please put them forward.

## 📮 Contact

- Email: [m@dev.vg](mailto:m@dev.vg)
- X (formerly Twitter): [@imPrk_](https://x.com/imPrk_)
- GitHub: [@imPrk0](https://github.com/imPrk0) | [Project](https://github.com/imPrk0/dev.vg) | [Issues](https://github.com/imPrk0/dev.vg/issues)

## ⚗️ Stacks

- [Go](https://go.dev/) + [Gin](https://github.com/gin-gonic/gin)

## ❤️ Thanks

- CodingStartup: [sdfsdf.dev](https://sdfsdf.dev/) / [YouTube](https://www.youtube.com/@CodingStartup)
- GoReleaser: [Website](https://goreleaser.com/)

## 🚀 Project and Copyright

This project is open source and hosted on GitHub: [imPrk0/dev.vg](https://github.com/imPrk0/dev.vg).

Released under the [MIT License](https://opensource.org/licenses/MIT).

Copyright&nbsp;&copy;&nbsp;2024-present&nbsp;[Prk](https://github.com/imPrk0)&nbsp;All&nbsp;Rights&nbsp;Reserved.
