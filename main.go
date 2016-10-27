package main

import (
	"fmt"
	"os"
	"raytracer/base"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
	"raytracer/textures"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	args := os.Args[1:]

	// Pixel counts
	nx := 500
	ny := 500
	ns := 1 // 8x Antialiasing
	film := base.NewFilm(nx, ny)

	var camera *base.Camera
	var mat materials.Material
	ambientLight := materials.NewAmbientLight(textures.White)
	lights := make([]materials.Light, 0, 10)
	cmdObjects := objects.NewEmptyObjectList(10)

	for i := 0; i < len(args); i++ {
		if args[i] == "cam" {
			ex, _ := strconv.ParseFloat(args[i+1], 64)
			ey, _ := strconv.ParseFloat(args[i+2], 64)
			ez, _ := strconv.ParseFloat(args[i+3], 64)

			llx, _ := strconv.ParseFloat(args[i+4], 64)
			lly, _ := strconv.ParseFloat(args[i+5], 64)
			llz, _ := strconv.ParseFloat(args[i+6], 64)

			lrx, _ := strconv.ParseFloat(args[i+7], 64)
			lry, _ := strconv.ParseFloat(args[i+8], 64)
			lrz, _ := strconv.ParseFloat(args[i+9], 64)

			ulx, _ := strconv.ParseFloat(args[i+10], 64)
			uly, _ := strconv.ParseFloat(args[i+11], 64)
			ulz, _ := strconv.ParseFloat(args[i+12], 64)

			urx, _ := strconv.ParseFloat(args[i+13], 64)
			ury, _ := strconv.ParseFloat(args[i+14], 64)
			urz, _ := strconv.ParseFloat(args[i+15], 64)

			eye := primitives.NewVec3(ex, ey, ez)
			LL := primitives.NewVec3(llx, lly, llz)
			LR := primitives.NewVec3(lrx, lry, lrz)
			UL := primitives.NewVec3(ulx, uly, ulz)
			UR := primitives.NewVec3(urx, ury, urz)

			camera = base.NewCameraFromCoordinates(LL, LR, UL, UR, eye, float64(nx), float64(ny))
			i += 15
			continue
		} else if args[i] == "sph" {
			cx, _ := strconv.ParseFloat(args[i+1], 64)
			cy, _ := strconv.ParseFloat(args[i+2], 64)
			cz, _ := strconv.ParseFloat(args[i+3], 64)
			r, _ := strconv.ParseFloat(args[i+4], 64)
			cmdObjects.Add(objects.NewSphere(primitives.NewVec3(cx, cy, cz), r, mat))
			i += 4
			continue
		} else if args[i] == "tri" {
			ax, _ := strconv.ParseFloat(args[i+1], 64)
			ay, _ := strconv.ParseFloat(args[i+2], 64)
			az, _ := strconv.ParseFloat(args[i+3], 64)

			bx, _ := strconv.ParseFloat(args[i+4], 64)
			by, _ := strconv.ParseFloat(args[i+5], 64)
			bz, _ := strconv.ParseFloat(args[i+6], 64)

			cx, _ := strconv.ParseFloat(args[i+7], 64)
			cy, _ := strconv.ParseFloat(args[i+8], 64)
			cz, _ := strconv.ParseFloat(args[i+9], 64)

			v1 := primitives.NewVec3(ax, ay, az)
			v2 := primitives.NewVec3(bx, by, bz)
			v3 := primitives.NewVec3(cx, cy, cz)

			cmdObjects.Add(objects.NewTriangle(v1, v2, v3, mat))
			i += 9
			continue
		} else if args[i] == "obj" {
			continue
		} else if args[i] == "ltp" {
			px, _ := strconv.ParseFloat(args[i+1], 64)
			py, _ := strconv.ParseFloat(args[i+2], 64)
			pz, _ := strconv.ParseFloat(args[i+3], 64)

			r, _ := strconv.ParseFloat(args[i+4], 64)
			g, _ := strconv.ParseFloat(args[i+5], 64)
			b, _ := strconv.ParseFloat(args[i+6], 64)
			falloff, err := strconv.ParseInt(args[i+7], 10, 32)
			if err != nil {
				falloff = 0
			}

			location := primitives.NewVec3(px, py, pz)
			color := textures.NewColor(r, g, b)
			lights = append(lights, materials.NewPointLight(location, color, int(falloff)))
			if err != nil {
				i += 6
			} else {
				i += 7
			}
			continue
		} else if args[i] == "ltd" {
			dx, _ := strconv.ParseFloat(args[i+1], 64)
			dy, _ := strconv.ParseFloat(args[i+2], 64)
			dz, _ := strconv.ParseFloat(args[i+3], 64)

			r, _ := strconv.ParseFloat(args[i+4], 64)
			g, _ := strconv.ParseFloat(args[i+5], 64)
			b, _ := strconv.ParseFloat(args[i+6], 64)

			location := primitives.NewVec3(dx, dy, dz)
			color := textures.NewColor(r, g, b)
			lights = append(lights, materials.NewDirectionalLight(location, color))
			i += 6
			continue
		} else if args[i] == "lta" {
			r, _ := strconv.ParseFloat(args[i+1], 64)
			g, _ := strconv.ParseFloat(args[i+2], 64)
			b, _ := strconv.ParseFloat(args[i+3], 64)
			color := textures.NewColor(r, g, b)
			ambientLight = materials.NewAmbientLight(color)
			i += 3
			continue
		} else if args[i] == "mat" {
			kar, _ := strconv.ParseFloat(args[i+1], 64)
			kag, _ := strconv.ParseFloat(args[i+2], 64)
			kab, _ := strconv.ParseFloat(args[i+3], 64)

			kdr, _ := strconv.ParseFloat(args[i+4], 64)
			kdg, _ := strconv.ParseFloat(args[i+5], 64)
			kdb, _ := strconv.ParseFloat(args[i+6], 64)

			ksr, _ := strconv.ParseFloat(args[i+7], 64)
			ksg, _ := strconv.ParseFloat(args[i+8], 64)
			ksb, _ := strconv.ParseFloat(args[i+9], 64)
			phong, _ := strconv.ParseFloat(args[i+10], 64)

			krr, _ := strconv.ParseFloat(args[i+11], 64)
			krg, _ := strconv.ParseFloat(args[i+12], 64)
			krb, _ := strconv.ParseFloat(args[i+13], 64)

			ambient := textures.NewColor(kar, kag, kab)
			diffuse := textures.NewColor(kdr, kdg, kdb)
			specular := textures.NewColor(ksr, ksg, ksb)
			reflective := textures.NewColor(krr, krg, krb)

			mat = materials.NewBlinnphong(ambient, diffuse, specular, reflective, phong, ambientLight)
			i += 13
			continue
		} else if args[i] == "xft" {
			continue
		} else if args[i] == "xfr" {
			continue
		} else if args[i] == "xfs" {
			continue
		} else if args[i] == "random" {
			continue
		} else {
			fmt.Println("Unexpected argument: ", i, args[i])
		}
	}

	// World space
	// origin := primitives.NewVec3(0, 0, 10)
	// // origin := primitives.NewVec3(4, 4, 4)
	// lookat := primitives.NewVec3(0.0, 0.0, 0.0)
	// vertical := primitives.NewVec3(0.0, 1.0, 0.0)
	// distToFocus := 1.0
	// aperature := 0.0
	// camera := base.NewCameraFOV(origin, lookat, vertical, 20,
	// 	float64(nx)/float64(ny), aperature, distToFocus, 0, 1)
	// camera.ToggleBlur()
	//
	// eye := primitives.NewVec3(0, 0, 1)
	// LL := primitives.NewVec3(-2, -1, 0)
	// LR := primitives.NewVec3(2, -1, 0)
	// UL := primitives.NewVec3(-2, 1, 0)
	// UR := primitives.NewVec3(2, 1, 0)
	// camera = base.NewCameraFromCoordinates(LL, LR, UL, UR, eye, float64(nx), float64(ny))

	// Objects
	// world := randomScene()
	// ambient := textures.NewColor(0, 0, 0)
	// diffuse := textures.NewColor(0.5, 0.5, 0.5)
	// specular := textures.NewColor(0.5, 0.5, 0.5)
	// reflective := textures.NewColor(0.5, 0.5, 0.5)
	// phong := 1.0
	// ambientLight = materials.NewAmbientLight(textures.NewColor(0.5, 0.5, 0.5))
	// s1 := objects.NewSphere(primitives.NewVec3(0, 0, -1), 1, materials.NewBlinnphong(ambient, diffuse, specular, reflective, phong, ambientLight))
	// s2 := objects.NewSphere(primitives.NewVec3(-.5, .5, 0), 0.25, materials.NewBlinnphong(ambient, diffuse, specular, reflective, phong, ambientLight))
	// lights = []materials.Light{materials.NewDirectionalLight(primitives.NewVec3(-1, 1, 1), textures.NewColor(.35, .7, 1))}
	//
	// cmdObjects = objects.NewObjectList(2, s1, s2)
	// world := s1
	scene := base.NewScene(camera, film, cmdObjects, lights, ns)
	scene.Render("output")

}
