# jp: A command line tool to prettyfy JSON

## Install

    $ git clone https://github.com/Moosemorals/jp.git
    $ cd jp
    $ make install

The command `jp` will be installed into `$GOROOT/bin` (assuming the 
tests pass)

## Run

    $ jp < infile.json > outfile.json

## Options

  *  -compact : Strip out any non-meaningful whitespace
  *  -indent n : Set the indent width (default 2)

## Licence

This package is covered by the [ISC licence](LICENCE)
