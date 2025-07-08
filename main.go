package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 1 unit = 1km
const EarthRadius float32 = 6371 //in KM
const MoonRadius float32 = 1737  //in KM
const shipLength float32 = 15    //in M
const KMtoUnits float32 = 0.001

func main() {
	var shipCam = ShipCamera{ //Setting up custom camera
		Position:      rl.NewVector3(0, 0, 16400),
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
	fmt.Println(EarthRadius * KMtoUnits)
	fmt.Println(shipLength * KMtoUnits)

	var Earth = Planet{
		&Sphere,
		&SphereTex,
		rl.NewVector3(0, 0, 0),
		EarthRadius,
	}

	var Moon = Planet{
		&MoonM,
		&MoonTex,
		rl.NewVector3(0, 0, 384400),
		MoonRadius,
	}

	Sphere.Materials.Maps.Texture = *Earth.Texture
	MoonM.Materials.Maps.Texture = *Moon.Texture

	//Main game loop
	for !rl.WindowShouldClose() {
		//Updating our main camera based on the parameters of our custom one
		rlCamera.Position = rl.NewVector3(0, 0, 0)
		rlCamera.Target = rl.Vector3Normalize(shipCam.Forward)
		rlCamera.Up = shipCam.Up
		//shipCam.Position = rl.NewVector3(0, 0, 0)
		//Draw loop begins
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		//Start 3d mode
		rl.BeginMode3D(rlCamera)
		shipCam.HandleInput()    //Handle input from custom camera
		shipCam.HandleRotation() //Handle rotation

		//getting scaled down positions of the earth, moon, and test sidewinder
		relShipPos := rl.Vector3Subtract(rl.NewVector3(0, 0, 6731), shipCam.Position)
		relEarthPos := rl.Vector3Subtract(Earth.Position, shipCam.Position)
		relMoonPos := rl.Vector3Subtract(Moon.Position, shipCam.Position)

		//drawing our models at the correct scale and positions (scaled down again so that they will be within the view frustum)
		rl.DrawModel(*Earth.Model, rl.Vector3Scale(relEarthPos, KMtoUnits), (Earth.Radius)*KMtoUnits, rl.White)
		rl.DrawModel(*Moon.Model, rl.Vector3Scale(relMoonPos, KMtoUnits), Moon.Radius*KMtoUnits, rl.White)

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
