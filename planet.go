package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Planet struct {
	Model         *rl.Model
	Texture       *rl.Texture2D
	Position      rl.Vector3
	Radius        float32
	RotationSpeed float32
	AxialTilt     float64
	Axis          rl.Vector3
	Rotation      float32
}
