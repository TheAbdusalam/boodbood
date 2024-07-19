# Boodbood
waa baddalka dabiiciga(xun) ee autojump, maadama iigu shaqayn waayay systemkayga.

waxa uu ka caawinaya inaad dhax gooshtid folder-yadaada sida `cd` lakin ka baranaya sida aad u gooshtid, side?
> mar kasta oo aad booqatid directory waxa ugu darsamaya miisaan(sida google indexing uu u shaqeeyo)

# Installing(dhisid?)
```bash
$ make
$ bd
```

## Isticmaalka
```bash
$ bd list        -   waxa uu taxi 10 directory ee ugu danbeeyay miisaan ahaan
$ bd j      -   waxa uu tagaya directoriga ugu horeya ee uu ku jiro xarafka `j` (sidoo kale miisaan ahaan)
$ bd j a    -   marlabaad waxa uu gali directoriga ugu horeya ee uu ku jiro `j` uuna kusi dhaxjiro `a`
``` 

# Boodbood
is an organic(shittier) alternative to autojump, because i couldn't get it to work on my machine

works like `cd` but learn from the way you navigate your filesystem, how?
> the more you visit a directory the more weight it adds to when you navigate(like how google indexing works)

## usage
```bash
$ bd list        -   list of the last 10 navigated to directories by weight
$ bd j      -   goes to the first dir that contains j (also by weight)
$ bd j a    -   again the first instance of the tree after the j dir

### isticmaal-ka(usage)
for now zsh kali ah ayuu u shaqeeynaya, will be adding support for others soon.
```bash
$ make

# Add this to the last line of your .zshrc
[ -f ~/.config/boodbood/boodbood.sh ] && source ~/.config/boodbood/boodbood.sh

$ bd
```

# Task list
- [x] cd to child dirs
- [x] list last 10 dirs with respect to weight
- [] group dirs into trees and add folders as branches
