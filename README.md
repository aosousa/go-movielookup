# go-movielookup

**Golang application to look up movie or TV show information using the OMDB API.**

### Installation
Download the executable in the [Releases tab](https://github.com/aosousa/go-movielookup/releases).

This tool requires an OMDB API key. You can set it in 2 ways:

#### Configuration file

A `config.json` file present in the same directory as the executable, with the following structure:
```json
{
    "apiKey": "<your-OMDB-API-key>"
}
```

#### Environment variable

Configure your OMDB API key as an environment variable.
```sh
echo 'export OMDB_KEY="<apiKey>"' >> <terminalFile>
```

Regarding the `<terminalFile>`, it depends on your terminal. If you're using
bash, change it for `~/.bashrc`, as for zsh, you can change for `~/.zshenv`.

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

### License

MIT © [André Sousa](https://github.com/aosousa)
