package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShipCamera struct {
	Position rl.Vector3
	Forward  rl.Vector3
	Up       rl.Vector3
	Right    rl.Vector3
	Rotation rl.Quaternion
}

func (cam *ShipCamera) HandleInput() { //Takes players input and updates camera's rotation
	var rotationSpeed float32 = 0.05
	var movementSpeed float32 = 0.05

	rotMatrix := rl.QuaternionToMatrix(cam.Rotation)
	forward := rl.Vector3Transform(rl.NewVector3(0, 0, -1), rotMatrix)
	right := rl.Vector3Transform(rl.NewVector3(1, 0, 0), rotMatrix)

	//Rotation functions
	if rl.IsKeyDown(rl.KeyRight) {
		// Yaw right
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(0, 1, 0), -rotationSpeed)
		cam.Rotation = rl.QuaternionMultiply(q, cam.Rotation)
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		// Yaw left
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(0, 1, 0), rotationSpeed)
		cam.Rotation = rl.QuaternionMultiply(q, cam.Rotation)
	}

	if rl.IsKeyDown(rl.KeyUp) {
		// Pitch up
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(1, 0, 0), -rotationSpeed)
		cam.Rotation = rl.QuaternionMultiply(q, cam.Rotation)
	}

	if rl.IsKeyDown(rl.KeyDown) {
		// Pitch down
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(1, 0, 0), rotationSpeed)
		cam.Rotation = rl.QuaternionMultiply(q, cam.Rotation)
	}

	if rl.IsKeyDown(rl.KeyQ) {
		// Roll counter-clockwise
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(0, 0, 1), rotationSpeed)
		cam.Rotation = rl.QuaternionMultiply(q, cam.Rotation)
	}

	if rl.IsKeyDown(rl.KeyE) {
		// Roll clockwise
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(0, 0, 1), -rotationSpeed)
		cam.Rotation = rl.QuaternionMultiply(q, cam.Rotation)
	}

	//Movement functions
	if rl.IsKeyDown(rl.KeyW) {
		cam.Position = rl.Vector3Add(cam.Position, rl.Vector3Scale(forward, movementSpeed))
	}
	if rl.IsKeyDown(rl.KeyS) {
		cam.Position = rl.Vector3Subtract(cam.Position, rl.Vector3Scale(forward, movementSpeed))
	}
	if rl.IsKeyDown(rl.KeyA) {
		cam.Position = rl.Vector3Subtract(cam.Position, rl.Vector3Scale(right, movementSpeed))
	}
	if rl.IsKeyDown(rl.KeyD) {
		cam.Position = rl.Vector3Add(cam.Position, rl.Vector3Scale(right, movementSpeed))
	}
}

func (cam *ShipCamera) HandleRotation() { //Actually changes camera's rotation
	rotationMatrix := rl.QuaternionToMatrix(cam.Rotation)

	cam.Forward = rl.Vector3Transform(rl.NewVector3(0, 0, -1), rotationMatrix)
	cam.Right = rl.Vector3Transform(rl.NewVector3(1, 0, 0), rotationMatrix)
	cam.Up = rl.Vector3Transform(rl.NewVector3(0, 1, 0), rotationMatrix)
}
