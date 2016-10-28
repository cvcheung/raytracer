# raytracer
## Introduction
raytracer is a pure-go implementation of a raytracer that follows Peter Shirley's Raytracing Minibooks series with minor modifications to adapt to the specifications of [CS184](http://inst.eecs.berkeley.edu/~cs184/fa16/assignments/as2/assignment-02.pdf). The reason to use `go` is to take advantage of the built-in concurrency.

## Dependencies
raytracer utilizes the [gonum](github.com/gonum/matrix/) package to support homogenous transformations.

`go get github.com/gonum/matrix/mat64`

## How to use
The basic features supported from command line are as follows.

```Shell
-aa uint
    Sets the antialising amount. (default 8)
-apt float
    Sets the aperature of the camera, requires fovcam.
-blur
    Turns on camera blur, effects change based on camera.
-dist float
    Sets the distance to focus. (default 1)
-f string
    File to load.
-fovcam
    Use a camera with a specified field of view.
-o string
    The filename. (default "output")
-r	Generate a random scene.
-vfov float
    Sets the camera fov, requires fovcam. (default 60)
-x uint
    Specifies the width of the image. (default 500)
-y uint
    Specifies the height of the image. (default 500)
```

Files are supported in the following format:
* The camera is specified by the coordinates of the eye and 4 corners.
  * `cam ex ey ez llx lly llz lrx lry lrz ulx uly ulz urx ury urz`
* Spheres and triangles can be specified with the following.
  * `sph cx cy cz r`
  * `tri ax ay az bx by bz cx cy cz`
  * `obj "file name”`
* The current supported lights through files are point, directional, and ambient.
  * `ltp px py pz r g b [falloff]`, falloff is 0, 1, or 2
  * `ltd dx dy dz r g b`
  * `lta r g b`
* You can specify the Blinnphong shading model here.
  * `mat kar kag kab kdr kdg kdb ksr ksg ksb ksp krr krg krb`
* Supported transformations include translation, rotation, scaling. `xfz` resets the transformation.
  * `xft tx ty tz`
  * `xfr rx ry rz`
  * `xfs sx sy sz`
  * `xfz`

## Futurework
* triangles
* perlin noise
* matrix transformations

## Sample Images
![shiny]
![blur]
![arealight]
![wide]
![checkered]
![stress]

[shiny]: sample/shiny.png
[blur]: sample/blur.png
[arealight]: sample/area_light.png
[wide]: sample/wide_view.png
[checkered]: sample/checkered.png
[stress]: sample/stress.png
