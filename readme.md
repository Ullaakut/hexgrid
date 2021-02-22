# Hexgrid

This is a GO (Golang) library used to handle regular hexagons.
It's based on the algorithms described at http://www.redblobgames.com/grids/hexagons/implementation.html

## Installation
```bash
    go get github.com/ishmulyan/hexgrid
```

## Usage
#### Importing

```go
import "github.com/ishmulyan/hexgrid"
```

### Examples

#### Creating hexagons

```go
hexagonA := hexgrid.NewHex(1, 2) //at axial coordinates Q=1 R=2
hexagonB := hexgrid.NewHex(2, 3) //at axial coordinates Q=2 R=3
```

#### Measuring the distance (in hexagons) between two hexagons

```go
distance := hexgrid.Distance(hexagonA, hexagonB)
```

#### Getting the array of hexagons on the path between two hexagons

```go
origin := hexgrid.NewHex(10, 20)
destination := hexgrid.NewHex(30, 40)
path := hexgrid.Line(origin, destination) 
```

#### Creating a layout

```go
origin := hexgrid.Point{X: 0, Y: 0}     // The coordinate that corresponds to the center of hexagon 0,0
size := hexgrid.Point{X: 100, Y: 100}  // The length of an hexagon side => 100
layout: = hexgrid.Layout{Origin: origin, Size: size, Orientation: OrientationFlatTop}
```

#### Obtaining the pixel that corresponds to a given hexagon

```go
hex := hexgrid.NewHex(1, 0)             
pixel := hexgrid.HexToPixel(layout, hex)  // Pixel that corresponds to the center of hex 1,0 (in the given layout)
```


#### Obtaining the hexagon that contains the given pixel (and rounding it)

```go
point := hexgrid.Point{X: 10, Y: 20}
hex := hexgrid.PixelToHex(layout, point).Round()
```

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

## History

0.1. First version

## Credits

* Pedro Sousa
* Red Blob Games (http://www.redblobgames.com/grids/hexagons/implementation.html)

## License

MIT