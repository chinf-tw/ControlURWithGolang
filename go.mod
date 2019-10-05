module UR3Demo

go 1.13

require (
	DualArmControl v0.0.0
	UR3Handler v0.0.0
)

replace (
	DualArmControl => ../DualArmControl
	UR3Handler => ../UR3Handler
)
