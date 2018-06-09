# go-movielookup

**Golang CMD Utility Tool to look up movie or TV show information using the OMDB API.**

No installation required. Either download the executable separately, or clone the entire repository through the following command:

`$ git clone https://github.com/aosousa/go-movielookup.git`

### Usage

```
go-movielookup.exe [-m | --movie | -s | --show | -h | --help | -v | --version]
```

### Options

```
-h, --help                    Prints the list of available commands.
-v, --version                 Prints the version of the application
-m, --movie TITLE [(YEAR)]    Search for a movie
You can add the year in front of the movie title to search for a movie from a specific year. Look at the Examples section for more information.
-s, --show TITLE [S1 | S1 E1] Search for a TV show. 
You can also search for a TV show season or TV show episode. 
In case you want the TV show from a specific year, you can add the year in front of the show title. Look at the Examples section for more information.
```

### Examples

#### Find movie

`$ go-movielookup.exe -m Avengers: Infinity War`

![ScreenShot](/img/findmovie.png)

#### Find movie from a specific year

`$ go-movielookup.exe -m Ghostbusters (1984)`

![ScreenShot](/img/findmovie_withyear.png)

#### Find TV show

`$ go-movielookup.exe -s Game of Thrones`

![ScreenShot](/img/findshow.png)

#### Find TV show from a specific year

`$ go-movielookup.exe -s House of Cards (2013)`

![ScreenShot](/img/findshow_withyear.png)

#### Find TV show season

`$ go-movielookup.exe -s Game of Thrones S1`

![ScreenShot](/img/findshowseason.png)

#### Find TV show season from a specific year

`$ go-movielookup.exe -s House of Cards (2013) S1`

![ScreenShot](/img/findshowseason_withyear.png)

#### Find TV show episode

`$ go-movielookup.exe -s Game of Thrones S1 E5`

![ScreenShot](/img/findshowepisode.png)

#### Find TV show episode from a specific year

`$ go-movielookup.exe -s House of Cards (2013) S1 E1`

![ScreenShot](/img/findshowepisode_withyear.png)

### Contribute

Found a bug? Have a feature you'd like to see added or something you'd like to see improved? You can do so by [opening a new issue](https://github.com/aosousa/go-movielookup/issues)!
