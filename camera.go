package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShipCamera struct {
	Position      rl.Vector3
	Forward       rl.Vector3
	Up            rl.Vector3
	Right         rl.Vector3
	Rotation      rl.Quaternion
	RotationSpeed float32
	MovementSpeed float32
}

func (cam *ShipCamera) HandleInput() { //Takes players input and updates camera's rotation
	var rotationSpeed float32 = cam.RotationSpeed
	var movementSpeed float32 = cam.MovementSpeed

	//Creates a rotation matrix from our camera's current orientation to align the global forward and right vectors to our current
	//view direction
	rotMatrix := rl.QuaternionToMatrix(cam.Rotation)
	//Aligns our direction vectors by multiplying the global forward and right directions
	forward := rl.Vector3Transform(rl.NewVector3(0, 0, -1), rotMatrix)
	right := rl.Vector3Transform(rl.NewVector3(1, 0, 0), rotMatrix)

	//Debug camera functions
	if rl.IsKeyDown(rl.KeyT) {
		cam.MovementSpeed = 0.001
	}
	if rl.IsKeyDown(rl.KeyG) {
		cam.MovementSpeed = 0.1
	}
	if rl.IsKeyDown(rl.KeyB) {
		cam.MovementSpeed = 100
	}
	if rl.IsKeyDown(rl.KeyH) {
		cam.MovementSpeed = 1000
	}
	if rl.IsKeyDown(rl.KeyN) {
		cam.MovementSpeed = 1
	}
	if rl.IsKeyDown(rl.KeyP) {
		cam.Position = rl.NewVector3(0, 0, 384300)
	}
	if rl.IsKeyDown(rl.KeyL) {
		cam.Position = rl.NewVector3(0, 0, 100)
	}

	//Rotation functions
	if rl.IsKeyDown(rl.KeyRight) {
		// Yaw right
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(0, 1, 0), rotationSpeed)
		cam.Rotation = rl.QuaternionMultiply(q, cam.Rotation)
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		// Yaw left
		q := rl.QuaternionFromAxisAngle(rl.NewVector3(0, 1, 0), -rotationSpeed)
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
		//Move forwards
		cam.Position = rl.Vector3Add(cam.Position, rl.Vector3Scale(forward, movementSpeed))
	}
	if rl.IsKeyDown(rl.KeyS) {
		//Move backwards
		cam.Position = rl.Vector3Subtract(cam.Position, rl.Vector3Scale(forward, movementSpeed))
	}
	if rl.IsKeyDown(rl.KeyA) {
		//Move left
		cam.Position = rl.Vector3Subtract(cam.Position, rl.Vector3Scale(right, movementSpeed))
	}
	if rl.IsKeyDown(rl.KeyD) {
		//Move right
		cam.Position = rl.Vector3Add(cam.Position, rl.Vector3Scale(right, movementSpeed))
	}
}

func (cam *ShipCamera) HandleRotation() { //Actually changes camera's rotation
	rotationMatrix := rl.QuaternionToMatrix(cam.Rotation)

	cam.Forward = rl.Vector3Transform(rl.NewVector3(0, 0, -1), rotationMatrix)
	cam.Right = rl.Vector3Transform(rl.NewVector3(1, 0, 0), rotationMatrix)
	cam.Up = rl.Vector3Transform(rl.NewVector3(0, 1, 0), rotationMatrix)
}
