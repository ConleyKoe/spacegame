package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 1 unit = 1km
const KMtoUnits float32 = 0.001
const EarthRadius float32 = 6371 //in KM
const EarthRotationSpeed float32 = float32((2 * math.Pi) / 86164.0)
const MoonRadius float32 = 1737 //in KM
const shipLength float32 = 1    //in KM

func main() {
	var shipCam = ShipCamera{ //Setting up custom camera
		Position:      rl.NewVector3(0, 0, 67311),
		Forward:       rl.NewVector3(0, 0, 1),
		Up:            rl.NewVector3(0, 1, 0),
		Right:         rl.NewVector3(1, 0, 0),
		Rotation:      rl.QuaternionIdentity(),
		MovementSpeed: 1000,
		RotationSpeed: 0.01,
	}
	var rlCamera = rl.Camera3D{ //Setting up raylib camera
		Position:   rl.NewVector3(0, 0, 0),
		Target:     rl.Vector3Add(shipCam.Position, shipCam.Forward),
		Up:         shipCam.Up,
		Fovy:       50,
		Projection: rl.CameraPerspective,
	}

	rl.InitWindow(1000, 500, "space game")
	rl.SetTargetFPS(60)

	var Sphere = rl.LoadModel("3dmodels/testplanet.obj")
	var MoonM = rl.LoadModel("3dmodels/testplanet.obj")
	var MoonTex = rl.LoadTexture("lalunetexture.jpg")
	var SphereTex = rl.LoadTexture("2_no_clouds_8k.jpg")

	sidewinder := rl.LoadModel("sidewinder.obj")
	sideTex := rl.LoadTexture("sidewindertextest.png")
	//sidewinder.Materials.Maps.Texture = sideTex

	var Earth = Planet{
		Model:         &Sphere,
		Texture:       &SphereTex,
		Position:      rl.NewVector3(0, 0, 0),
		Radius:        EarthRadius,
		RotationSpeed: EarthRotationSpeed,
		AxialTilt:     23.5,
		Axis:          rl.NewVector3(0, 1, 0),
		Rotation:      0,
	}

	var Moon = Planet{
		Model:         &MoonM,
		Texture:       &MoonTex,
		Position:      rl.NewVector3(0, 0, 384400),
		Radius:        MoonRadius,
		RotationSpeed: 0.05,
		AxialTilt:     6.68,
		Axis:          rl.NewVector3(0, 1, 0),
		Rotation:      0,
	}

	Sphere.Materials.Maps.Texture = *Earth.Texture
	MoonM.Materials.Maps.Texture = *Moon.Texture

	Earth.Axis = rl.NewVector3(
		float32(math.Sin(float64(Earth.AxialTilt*rl.Deg2rad))),
		float32(math.Cos(float64(Earth.AxialTilt*rl.Deg2rad))),
		0,
	)
	fmt.Println(Earth.Axis)

	//Main game loop
	for !rl.WindowShouldClose() {

		//Updating our main camera based on the parameters of our custom one
		rlCamera.Position = rl.NewVector3(0, 0, 0)
		rlCamera.Target = rl.Vector3Normalize(shipCam.Forward)
		rlCamera.Up = shipCam.Up
		//shipCam.Position = rl.NewVector3(0, 0, 0)

		deltaTime := rl.GetFrameTime()

		Earth.Rotation = Earth.Rotation + Earth.RotationSpeed*deltaTime
		//Draw loop begins
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		//Start 3d mode
		rl.BeginMode3D(rlCamera)
		shipCam.HandleInput()    //Handle input from custom camera
		shipCam.HandleRotation() //Handle rotation

		//getting scaled down positions of the earth, moon, and test sidewinder
		relShipPos := rl.Vector3Subtract(rl.NewVector3(0, 0, 67311), shipCam.Position)
		relEarthPos := rl.Vector3Subtract(Earth.Position, shipCam.Position)
		relMoonPos := rl.Vector3Subtract(Moon.Position, shipCam.Position)

		//drawing our models at the correct scale and positions (scaled down again so that they will be within the view frustum)
		rl.DrawModelEx(
			*Earth.Model,
			rl.Vector3Scale(relEarthPos, KMtoUnits),
			Earth.Axis,
			Earth.Rotation*rl.Rad2deg,
			rl.NewVector3((Earth.Radius)*KMtoUnits, (Earth.Radius)*KMtoUnits, (Earth.Radius)*KMtoUnits),
			rl.White)
		rl.DrawModel(*Moon.Model, rl.Vector3Scale(relMoonPos, KMtoUnits), Moon.Radius*KMtoUnits, rl.White)

		//Earth Axis
		lineEnd := rl.Vector3Scale(rl.Vector3Add(relEarthPos, rl.Vector3Scale(Earth.Axis, 10000.0)), KMtoUnits)
		rl.DrawLine3D(rl.Vector3Scale(relEarthPos, KMtoUnits), lineEnd, rl.Red)

		//draws our test sidewinder and its wireframe
		rl.DrawModel(sidewinder, rl.Vector3Scale(relShipPos, KMtoUnits), shipLength*KMtoUnits, rl.Red)
		rl.DrawModelWires(sidewinder, rl.Vector3Scale(relShipPos, KMtoUnits), shipLength*KMtoUnits, rl.White)
		//rl.DrawCubeV(rl.NewVector3(0, 0, 6.377), rl.NewVector3(0.01, 0.01, 0.01), rl.Green)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.DrawText(fmt.Sprint(shipCam.Position), 20, 20, 10, rl.White)
		rl.DrawText(fmt.Sprint(shipCam.MovementSpeed), 300, 20, 10, rl.White)

		rl.EndDrawing()
	}

	rl.UnloadModel(Sphere)
	rl.UnloadModel(MoonM)
	rl.UnloadModel(sidewinder)
	rl.UnloadTexture(SphereTex)
	rl.UnloadTexture(sideTex)
	rl.UnloadTexture(MoonTex)

	rl.CloseWindow()

}
