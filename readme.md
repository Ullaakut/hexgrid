# HexGrid

This is a Go library used to handle hexagons. It's based on the
[following algorithms](http://www.redblobgames.com/grids/hexagons/implementation.html).

## Usage

### Importing

```go
import "github.com/Ullaakut/hexgrid"
```

### Examples

#### Creating Hexagons

```go
hexagonA := hexgrid.NewHex(1, 2) // at axial coordinates Q=1 R=2
hexagonB := hexgrid.NewHex(2, 3) // at axial coordinates Q=2 R=3
```

#### Measuring the Distance Between Two Hexagons

```go
distance := hexgrid.Distance(hexagonA, hexagonB)
```

#### Getting the Array of Hexagons on the Path Between Two Hexagons

```go
origin := hexgrid.NewHex(10, 20)
destination := hexgrid.NewHex(30, 40)
path := hexgrid.Line(origin, destination) 
```

#### Creating a Layout

```go
origin := hexgrid.Point{X: 0, Y: 0}     // The coordinate that corresponds to the center of hexagon 0,0
size := hexgrid.Point{X: 100, Y: 100}  // The length of an hexagon side => 100
layout: = hexgrid.Layout{Origin: origin, Size: size, Orientation: OrientationFlatTop}
```

#### Obtaining the Pixel that Corresponds to a Given Hexagon

```go
hex := hexgrid.NewHex(1, 0)
pixel := hexgrid.HexToPixel(layout, hex) // Pixel that corresponds to the center of hex 1,0 (in the given layout)
```

#### Obtaining the Hexagon that Contains the Given Pixel

```go
point := hexgrid.Point{X: 10, Y: 20}
hex := hexgrid.PixelToHex(layout, point).Round()
```

## Credits

* [ishmulyan](https://github.com/ishmulyan)
* [Pedro Sousa](https://github.com/pmcxs)
* [Red Blob Games](http://www.redblobgames.com/grids/hexagons/implementation.html)