package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	var shipCam = ShipCamera{ //Setting up custom camera
		Position: rl.NewVector3(1, 1, -20),
		Forward:  rl.NewVector3(0, 0, 1),
		Up:       rl.NewVector3(0, 1, 0),
		Right:    rl.NewVector3(1, 0, 0),
		Rotation: rl.QuaternionIdentity(),
	}
	var rlCamera = rl.Camera3D{ //Setting up raylib camera
		Position:   rl.NewVector3(1, 1, 0),
		Target:     rl.Vector3Add(shipCam.Position, shipCam.Forward),
		Up:         shipCam.Up,
		Fovy:       60,
		Projection: rl.CameraPerspective,
	}

	rl.InitWindow(640, 480, "space game")
	rl.SetTargetFPS(60)
	sphere := rl.LoadModel("3dmodels/testplanet.obj")
	sphereTex := rl.LoadTexture("3dmodels/testsphere.png")
	sphere.Materials.Maps.Texture = sphereTex

	sidewinder := rl.LoadModel("sidewinder.obj")
	sideTex := rl.LoadTexture("sidewindertextest.png")
	//sidewinder.Materials.Maps.Texture = sideTex

	//Main game loop
	for !rl.WindowShouldClose() {
		//Updating our main camera based on the parameters of our custom one
		rlCamera.Position = shipCam.Position
		rlCamera.Target = rl.Vector3Add(shipCam.Position, shipCam.Forward)
		rlCamera.Up = shipCam.Up
		//Draw loop begins
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		//Start 3d mode
		rl.BeginMode3D(rlCamera)
		shipCam.HandleInput()    //Handle input from custom camera
		shipCam.HandleRotation() //Handle rotation
		//rl.DrawModel(sphere, rl.NewVector3(1, 1, -8), 1000, rl.White)
		rl.DrawModel(sidewinder, rl.NewVector3(0, 0, 0), 1, rl.Red)
		rl.DrawModelWires(sidewinder, rl.NewVector3(0, 0, 0), 1, rl.White)
		rl.EndMode3D()
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}

	rl.UnloadModel(sphere)
	rl.UnloadModel(sidewinder)
	rl.UnloadTexture(sphereTex)
	rl.UnloadTexture(sideTex)

	rl.CloseWindow()

}
