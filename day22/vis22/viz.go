package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/g3n/engine/animation"

	"github.com/askreet/aoc2023/day22"
)

import (
	"github.com/g3n/engine/app"
	"github.com/g3n/engine/camera"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/engine/window"
)

func main() {

	// Create application and scene
	a := app.App()
	scene := core.NewNode()

	// Set the scene to be managed by the gui manager
	gui.Manager().Set(scene)

	// Create perspective camera
	cam := camera.New(1)
	cam.SetPosition(0, 0, 3)
	scene.Add(cam)

	// Set up orbit control for the camera
	camera.NewOrbitControl(cam)

	// Set up callback to update viewport and camera aspect ratio when the window is resized
	onResize := func(evname string, ev interface{}) {
		// Get framebuffer size and update viewport accordingly
		width, height := a.GetSize()
		a.Gls().Viewport(0, 0, int32(width), int32(height))
		// Update the camera's aspect ratio
		cam.SetAspect(float32(width) / float32(height))
	}
	a.Subscribe(window.OnWindowSize, onResize)
	onResize("", nil)

	var tower *day22.Tower
	var shapeNodes []*core.Node
	removeAllShapes := func() {
		for _, n := range shapeNodes {
			scene.Remove(n)
		}
	}
	addAllShapes := func() {
		var colors = []math32.Color{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1},
		}

		if len(shapeNodes) != tower.Len() {
			shapeNodes = make([]*core.Node, tower.Len())
		}
		for i, shape := range tower.Shapes() {
			color := colors[i%len(colors)]
			node := NodeFor(shape, color)
			shapeNodes[i] = node
			scene.Add(node)
		}
	}

	var anims []*animation.Animation

	settleChan := make(chan *day22.Tower, 1)

	a.Subscribe(window.OnKeyDown, func(s string, ev any) {
		keyEvent := ev.(*window.KeyEvent)

		// Load (E)xample
		if keyEvent.Key == window.KeyE {
			removeAllShapes()
			tower = day22.Parse(bytes.NewBufferString(day22.ExampleInput))
			addAllShapes()
		}

		// Load (I)nput
		if keyEvent.Key == window.KeyI {
			data, err := os.ReadFile("in/day22.txt")
			if err != nil {
				fmt.Println("failed to read file")
				return
			}

			removeAllShapes()
			tower = day22.Parse(bytes.NewBuffer(data))
			addAllShapes()
		}

		// (S)ettle Tower
		if keyEvent.Key == window.KeyS {
			// Clear any active animations.
			anims = anims[0:0]

			if tower.Len() == 0 {
				return
			}

			if keyEvent.Mods&window.ModShift > 0 {
				go func() {
					settleChan <- tower.Settle()
				}()

			} else {
				removeAllShapes()
				tower.SettleStep()
				addAllShapes()
			}

		}
	})

	a.Subscribe("SettleComputationComplete", func(s string, i interface{}) {
		newShapes := i.(*day22.Tower).Shapes()
		shapes := tower.Shapes()

		for i := range shapes {
			if shapes[i] != newShapes[i] {
				fmt.Println("shape", i, "was updated")
			}
			// Use the target mesh node order to find target positions. This could be smarter.
			targetMesh := NodeFor(newShapes[i], math32.Color{0, 0, 0})
			for chIdx, child := range shapeNodes[i].Children() {
				fallDistance := shapes[i].MinZ() - newShapes[i].MinZ()
				if fallDistance == 0 {
					// No animation required...
					continue
				}

				anim := animation.NewAnimation()
				anim.SetLoop(false)

				keyframes := math32.NewArrayF32(0, 2)
				keyframes.Append(0, 0.1*float32(fallDistance))

				posValues := math32.NewArrayF32(0, 6)
				currentPos := child.Position()
				targetPos := targetMesh.ChildAt(chIdx).Position()
				posValues.AppendVector3(
					&currentPos,
					&targetPos)

				posChan := animation.NewPositionChannel(child)
				posChan.SetBuffers(keyframes, posValues)

				anim.AddChannel(posChan)
				anims = append(anims, anim)
			}
		}
	})
	// Add tower to the scene.

	// Create and add a button to the scene
	//btn := gui.NewButton("Make Red")
	//btn.SetPosition(100, 40)
	//btn.SetSize(40, 40)
	//btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
	//	mat.SetColor(math32.NewColor("DarkRed"))
	//})
	//scene.Add(btn)

	// Create and add lights to the scene
	scene.Add(light.NewAmbient(&math32.Color{1.0, 1.0, 1.0}, 0.8))
	pointLight := light.NewPoint(&math32.Color{1, 1, 1}, 100.0)
	pointLight.SetPosition(5, 5, 2)
	pointLight.SetLinearDecay(0.25)
	scene.Add(pointLight)

	// Create and add an axis helper to the scene
	scene.Add(helper.NewAxes(0.5))
	scene.Add(helper.NewGrid(100, 1, &math32.Color{0, 0, 0}))

	// Set background color to gray
	a.Gls().ClearColor(0.5, 0.5, 0.5, 1.0)

	// Run the application
	a.Run(func(renderer *renderer.Renderer, deltaTime time.Duration) {
		a.Gls().Clear(gls.DEPTH_BUFFER_BIT | gls.STENCIL_BUFFER_BIT | gls.COLOR_BUFFER_BIT)
		// run all active animations
		for _, anim := range anims {
			anim.Update(float32(deltaTime.Seconds()))
		}

		if len(settleChan) == 1 {
			newShapes := <-settleChan
			a.Dispatch("SettleComputationComplete", newShapes)
		}

		renderer.Render(scene, cam)
	})
}

func NodeFor(shape day22.Shape, color math32.Color) *core.Node {
	var node = core.NewNode()

	for x := shape.X1; x <= shape.X2; x++ {
		for y := shape.Y1; y <= shape.Y2; y++ {
			for z := shape.Z1; z <= shape.Z2; z++ {
				geom := geometry.NewCube(0.95)
				mat := material.NewStandard(&color)
				mesh := graphic.NewMesh(geom, mat)
				//
				x, y, z := ShapePosToGenPos(x, y, z)
				mesh.SetPosition(x, y, z)

				node.Add(mesh)
			}
		}
	}

	return node
}

func ShapePosToGenPos(x, y, z int) (float32, float32, float32) {
	return float32(y), float32(z), float32(x)
}
