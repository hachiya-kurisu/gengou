# gengou (元号)

find japanese era names from dates

## usage

```
now := time.Time.now()
era := gengou.Find(now).Name
```

can properly calculate the era year as well:

```
gengou.EraYear(time.Parse("2006.01.02 MST", "2019.04.30 JST")) // "平成31年"

gengou.EraYear(time.Parse("2006.01.02 MST", "2019.05.01 JST")) // "令和元年"

gengou.EraYear(time.Parse("2006.01.02 MST", "2020.01.01 JST")) // "令和2年"
```

## cli

```
% date
2024年 12月21日 土曜日 14時00分20秒 JST
% gengou
令和6年
% gengou -w
令和６年
% gengou -w 2019.04.30 2019.05.01
平成３１年
令和元年
% gengou -f 2006 1991 # go date layouts
平成3年
```
