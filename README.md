# gengou （元号）

find japanese era names from dates

![gengou in action](gengou.gif)

## usage

```
now := time.Now()
era, err := gengou.Find(now).Name
```

can properly calculate the era year as well:

```
// returns "平成31年"
gengou.EraYear(time.Parse("2006.01.02 MST", "2019.04.30 JST"))

// returns "令和元年"
gengou.EraYear(time.Parse("2006.01.02 MST", "2019.05.01 JST"))

// returns "令和2年"
gengou.EraYear(time.Parse("2006.01.02 MST", "2020.01.01 JST"))
```

## cli

```
$ date
2024年 12月21日 土曜日 14時00分20秒 JST
$ gengou
令和6年
$ gengou -w
令和６年
$ gengou -w 2019.04.30 2019.05.01
平成３１年
令和元年
$ gengou -f 2006 1991 # go date layouts
平成3年
$ gengou -d # show the full date
令和6年12月27日
```

## author

[蜂谷栗栖](https://blekksprut.net/)
## installation

### go

```
go install blekksprut.net/gengou/cmd/gengou@latest
```

### arch linux

[gengou](https://aur.archlinux.org/packages/gengou)
is available as a package in the AUR

it can be installed with an AUR helper (e.g. yay):
```
$ yay -S gengou
```

