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

func (cam *ShipCamera) HandleInput() {
	var rotationSpeed float32 = 0.01
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
}

func (cam *ShipCamera) HandleRotation() {
	rotationMatrix := rl.QuaternionToMatrix(cam.Rotation)

	cam.Forward = rl.Vector3Transform(rl.NewVector3(0, 0, -1), rotationMatrix)
	cam.Right = rl.Vector3Transform(rl.NewVector3(1, 0, 0), rotationMatrix)
	cam.Up = rl.Vector3Transform(rl.NewVector3(0, 1, 0), rotationMatrix)
}
