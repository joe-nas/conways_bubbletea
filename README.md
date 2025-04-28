# Conway`s game of Life

This is my Conway`s game of Life implementation written in Go, using the (UI) BubbleTea library.

You can use the arrow keys to navigate map and toggle state(dead|alive) using (e).
Pressing (n) activates the next generation. Pressing (a) toggles autorun.
Furthermore you can switch between normal and stats view by pressing (s).
Pressing (q) will exit the program.

## Normal view:

![Normal view]({AE0EB82A-8065-4943-BBD3-C01FF18E551E}.png)

## Stats view:

- The numbers correspond to the number of living neighbours.
- Red cells will die next generation.
- Blue cells survive till the next generation.
- Green cells will be born the next generation.
  ![Stats view]({6EC9DA57-7FCC-4870-8DB0-EBA667DD4357}.png)
