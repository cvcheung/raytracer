package parsers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"raytracer/base"
	"raytracer/materials"
	"raytracer/objects"
	"raytracer/primitives"
	"raytracer/textures"
	"raytracer/transformations"
	"strconv"
	"strings"

	"github.com/gonum/matrix/mat64"
)

// ParseFile ...
func ParseFile(filename string, opt *Options) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseLine(strings.Fields(scanner.Text()), opt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parseLine(line []string, opt *Options) {
	for i := 0; i < len(line); i++ {
		if line[i] == "cam" {
			ex, _ := strconv.ParseFloat(line[i+1], 64)
			ey, _ := strconv.ParseFloat(line[i+2], 64)
			ez, _ := strconv.ParseFloat(line[i+3], 64)

			llx, _ := strconv.ParseFloat(line[i+4], 64)
			lly, _ := strconv.ParseFloat(line[i+5], 64)
			llz, _ := strconv.ParseFloat(line[i+6], 64)

			lrx, _ := strconv.ParseFloat(line[i+7], 64)
			lry, _ := strconv.ParseFloat(line[i+8], 64)
			lrz, _ := strconv.ParseFloat(line[i+9], 64)

			ulx, _ := strconv.ParseFloat(line[i+10], 64)
			uly, _ := strconv.ParseFloat(line[i+11], 64)
			ulz, _ := strconv.ParseFloat(line[i+12], 64)

			urx, _ := strconv.ParseFloat(line[i+13], 64)
			ury, _ := strconv.ParseFloat(line[i+14], 64)
			urz, _ := strconv.ParseFloat(line[i+15], 64)

			eye := primitives.NewVec3(ex, ey, ez)
			LL := primitives.NewVec3(llx, lly, llz)
			LR := primitives.NewVec3(lrx, lry, lrz)
			UL := primitives.NewVec3(ulx, uly, ulz)
			UR := primitives.NewVec3(urx, ury, urz)

			nx := float64(opt.nx)
			ny := float64(opt.ny)
			camera := base.NewCameraFromCoordinates(LL, LR, UL, UR, eye, nx, ny,
				opt.vfov, opt.aperture, opt.distFocus, opt.fovcam)
			opt.SetCamera(camera)
			i += 15
			continue
		} else if line[i] == "sph" {
			cx, _ := strconv.ParseFloat(line[i+1], 64)
			cy, _ := strconv.ParseFloat(line[i+2], 64)
			cz, _ := strconv.ParseFloat(line[i+3], 64)
			r, _ := strconv.ParseFloat(line[i+4], 64)
			if len(opt.transforms) > 0 {
				transform := transformations.Coalesce(opt.transforms)
				opt.AddObjects(objects.NewSphereWithTransform(
					primitives.NewVec3(cx, cy, cz), r, opt.mat, transform))
			} else {
				opt.AddObjects(objects.NewSphere(primitives.NewVec3(cx, cy, cz), r,
					opt.mat))
			}
			i += 4
			continue
		} else if line[i] == "tri" {
			ax, _ := strconv.ParseFloat(line[i+1], 64)
			ay, _ := strconv.ParseFloat(line[i+2], 64)
			az, _ := strconv.ParseFloat(line[i+3], 64)

			bx, _ := strconv.ParseFloat(line[i+4], 64)
			by, _ := strconv.ParseFloat(line[i+5], 64)
			bz, _ := strconv.ParseFloat(line[i+6], 64)

			cx, _ := strconv.ParseFloat(line[i+7], 64)
			cy, _ := strconv.ParseFloat(line[i+8], 64)
			cz, _ := strconv.ParseFloat(line[i+9], 64)

			v1 := primitives.NewVec3(ax, ay, az)
			v2 := primitives.NewVec3(bx, by, bz)
			v3 := primitives.NewVec3(cx, cy, cz)

			opt.AddObjects(objects.NewTriangle(v1, v2, v3, opt.mat))
			i += 9
			continue
		} else if line[i] == "obj" {
			i++
			// TODO
			// vertices, normals := ParseObj(line[i])
			continue
		} else if line[i] == "ltp" {
			px, _ := strconv.ParseFloat(line[i+1], 64)
			py, _ := strconv.ParseFloat(line[i+2], 64)
			pz, _ := strconv.ParseFloat(line[i+3], 64)

			r, _ := strconv.ParseFloat(line[i+4], 64)
			g, _ := strconv.ParseFloat(line[i+5], 64)
			b, _ := strconv.ParseFloat(line[i+6], 64)
			falloff, err := strconv.ParseInt(line[i+7], 10, 32)
			if err != nil {
				falloff = 0
			}

			location := primitives.NewVec3(px, py, pz)
			color := textures.NewColor(r, g, b)
			opt.AddLights(materials.NewPointLight(location, color, int(falloff)))
			if err != nil {
				i += 6
			} else {
				i += 7
			}
			continue
		} else if line[i] == "ltd" {
			dx, _ := strconv.ParseFloat(line[i+1], 64)
			dy, _ := strconv.ParseFloat(line[i+2], 64)
			dz, _ := strconv.ParseFloat(line[i+3], 64)

			r, _ := strconv.ParseFloat(line[i+4], 64)
			g, _ := strconv.ParseFloat(line[i+5], 64)
			b, _ := strconv.ParseFloat(line[i+6], 64)

			location := primitives.NewVec3(dx, dy, dz)
			color := textures.NewColor(r, g, b)
			opt.AddLights(materials.NewDirectionalLight(location, color))
			i += 6
			continue
		} else if line[i] == "lta" {
			r, _ := strconv.ParseFloat(line[i+1], 64)
			g, _ := strconv.ParseFloat(line[i+2], 64)
			b, _ := strconv.ParseFloat(line[i+3], 64)
			color := textures.NewColor(r, g, b)
			opt.SetAmbientLight(materials.NewAmbientLight(color))
			i += 3
			continue
		} else if line[i] == "mat" {
			kar, _ := strconv.ParseFloat(line[i+1], 64)
			kag, _ := strconv.ParseFloat(line[i+2], 64)
			kab, _ := strconv.ParseFloat(line[i+3], 64)

			kdr, _ := strconv.ParseFloat(line[i+4], 64)
			kdg, _ := strconv.ParseFloat(line[i+5], 64)
			kdb, _ := strconv.ParseFloat(line[i+6], 64)

			ksr, _ := strconv.ParseFloat(line[i+7], 64)
			ksg, _ := strconv.ParseFloat(line[i+8], 64)
			ksb, _ := strconv.ParseFloat(line[i+9], 64)
			phong, _ := strconv.ParseFloat(line[i+10], 64)

			krr, _ := strconv.ParseFloat(line[i+11], 64)
			krg, _ := strconv.ParseFloat(line[i+12], 64)
			krb, _ := strconv.ParseFloat(line[i+13], 64)

			ambient := textures.NewColor(kar, kag, kab)
			diffuse := textures.NewColor(kdr, kdg, kdb)
			specular := textures.NewColor(ksr, ksg, ksb)
			reflective := textures.NewColor(krr, krg, krb)

			opt.SetMat(materials.NewBlinnphong(ambient, diffuse, specular,
				reflective, phong, opt.ambientLight))
			i += 13
			continue
		} else if line[i] == "xft" {
			tx, _ := strconv.ParseFloat(line[i+1], 64)
			ty, _ := strconv.ParseFloat(line[i+2], 64)
			tz, _ := strconv.ParseFloat(line[i+3], 64)
			opt.transforms = append(opt.transforms,
				transformations.NewTranslationMatrix(tx, ty, tz))
			i += 3
			continue
		} else if line[i] == "xfr" {
			rx, _ := strconv.ParseFloat(line[i+1], 64)
			ry, _ := strconv.ParseFloat(line[i+2], 64)
			rz, _ := strconv.ParseFloat(line[i+3], 64)
			opt.transforms = append(opt.transforms,
				transformations.NewRotationMatrix(rx, ry, rz))
			i += 3
			continue
		} else if line[i] == "xfs" {
			sx, _ := strconv.ParseFloat(line[i+1], 64)
			sy, _ := strconv.ParseFloat(line[i+2], 64)
			sz, _ := strconv.ParseFloat(line[i+3], 64)
			opt.transforms = append(opt.transforms,
				transformations.NewScalingMatrix(sx, sy, sz))
			i += 3
			continue
		} else if line[i] == "xfz" {
			opt.transforms = make([]*mat64.Dense, 0, 3)
			continue
		} else {
			fmt.Println("Unexpected argument: ", i, line[i])
		}
	}
}
