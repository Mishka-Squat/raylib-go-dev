package rl

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"

	"github.com/Mishka-Squat/gamemath/rect2"
	"github.com/Mishka-Squat/gamemath/vector2"
	"github.com/Mishka-Squat/gamemath/vector3"
	"github.com/Mishka-Squat/goex/image/colorex"
	"golang.org/x/exp/slices"
)

// DrawLine3D - Draw a line in 3D world space
func DrawLine3D(startPos vector3.Float32, endPos vector3.Float32, col colorex.RGBA) {
	cstartPos := cvec3ptr(&startPos)
	cendPos := cvec3ptr(&endPos)
	ccolor := ccolorptr(&col)
	C.DrawLine3D(*cstartPos, *cendPos, *ccolor)
}

// DrawPoint3D - Draw a point in 3D space, actually a small line
func DrawPoint3D(position vector3.Float32, col colorex.RGBA) {
	cposition := cvec3ptr(&position)
	ccolor := ccolorptr(&col)
	C.DrawPoint3D(*cposition, *ccolor)
}

// DrawCircle3D - Draw a circle in 3D world space
func DrawCircle3D(center vector3.Float32, radius float32, rotationAxis vector3.Float32, rotationAngle float32, col colorex.RGBA) {
	ccenter := cvec3ptr(&center)
	cradius := (C.float)(radius)
	crotationAxis := cvec3ptr(&rotationAxis)
	crotationAngle := (C.float)(rotationAngle)
	ccolor := ccolorptr(&col)
	C.DrawCircle3D(*ccenter, cradius, *crotationAxis, crotationAngle, *ccolor)
}

// DrawTriangle3D - Draw a color-filled triangle (vertex in counter-clockwise order!)
func DrawTriangle3D(v1 vector3.Float32, v2 vector3.Float32, v3 vector3.Float32, col colorex.RGBA) {
	cv1 := cvec3ptr(&v1)
	cv2 := cvec3ptr(&v2)
	cv3 := cvec3ptr(&v3)
	ccolor := ccolorptr(&col)
	C.DrawTriangle3D(*cv1, *cv2, *cv3, *ccolor)
}

// DrawCube - Draw cube
func DrawCube(position vector3.Float32, width float32, height float32, length float32, col colorex.RGBA) {
	cposition := cvec3ptr(&position)
	cwidth := (C.float)(width)
	cheight := (C.float)(height)
	clength := (C.float)(length)
	ccolor := ccolorptr(&col)
	C.DrawCube(*cposition, cwidth, cheight, clength, *ccolor)
}

// DrawCubeV - Draw cube (Vector version)
func DrawCubeV(position vector3.Float32, size vector3.Float32, col colorex.RGBA) {
	cposition := cvec3ptr(&position)
	csize := cvec3ptr(&size)
	ccolor := ccolorptr(&col)
	C.DrawCubeV(*cposition, *csize, *ccolor)
}

// DrawCubeWires - Draw cube wires
func DrawCubeWires(position vector3.Float32, width float32, height float32, length float32, col colorex.RGBA) {
	cposition := cvec3ptr(&position)
	cwidth := (C.float)(width)
	cheight := (C.float)(height)
	clength := (C.float)(length)
	ccolor := ccolorptr(&col)
	C.DrawCubeWires(*cposition, cwidth, cheight, clength, *ccolor)
}

// DrawCubeWiresV - Draw cube wires (Vector version)
func DrawCubeWiresV(position vector3.Float32, size vector3.Float32, col colorex.RGBA) {
	cposition := cvec3ptr(&position)
	csize := cvec3ptr(&size)
	ccolor := ccolorptr(&col)
	C.DrawCubeWiresV(*cposition, *csize, *ccolor)
}

// DrawSphere - Draw sphere
func DrawSphere(centerPos vector3.Float32, radius float32, col colorex.RGBA) {
	ccenterPos := cvec3ptr(&centerPos)
	cradius := (C.float)(radius)
	ccolor := ccolorptr(&col)
	C.DrawSphere(*ccenterPos, cradius, *ccolor)
}

// DrawSphereEx - Draw sphere with extended parameters
func DrawSphereEx(centerPos vector3.Float32, radius float32, rings int32, slices int32, col colorex.RGBA) {
	ccenterPos := cvec3ptr(&centerPos)
	cradius := (C.float)(radius)
	crings := (C.int)(rings)
	cslices := (C.int)(slices)
	ccolor := ccolorptr(&col)
	C.DrawSphereEx(*ccenterPos, cradius, crings, cslices, *ccolor)
}

// DrawSphereWires - Draw sphere wires
func DrawSphereWires(centerPos vector3.Float32, radius float32, rings int32, slices int32, col colorex.RGBA) {
	ccenterPos := cvec3ptr(&centerPos)
	cradius := (C.float)(radius)
	crings := (C.int)(rings)
	cslices := (C.int)(slices)
	ccolor := ccolorptr(&col)
	C.DrawSphereWires(*ccenterPos, cradius, crings, cslices, *ccolor)
}

// DrawCylinder - Draw a cylinder/cone
func DrawCylinder(position vector3.Float32, radiusTop float32, radiusBottom float32, height float32, slices int32, col colorex.RGBA) {
	cposition := cvec3ptr(&position)
	cradiusTop := (C.float)(radiusTop)
	cradiusBottom := (C.float)(radiusBottom)
	cheight := (C.float)(height)
	cslices := (C.int)(slices)
	ccolor := ccolorptr(&col)
	C.DrawCylinder(*cposition, cradiusTop, cradiusBottom, cheight, cslices, *ccolor)
}

// DrawCylinderEx - Draw a cylinder with base at startPos and top at endPos
func DrawCylinderEx(startPos vector3.Float32, endPos vector3.Float32, startRadius float32, endRadius float32, sides int32, col colorex.RGBA) {
	cstartPos := cvec3ptr(&startPos)
	cendPos := cvec3ptr(&endPos)
	cstartRadius := (C.float)(startRadius)
	cendRadius := (C.float)(endRadius)
	csides := (C.int)(sides)
	ccolor := ccolorptr(&col)
	C.DrawCylinderEx(*cstartPos, *cendPos, cstartRadius, cendRadius, csides, *ccolor)
}

// DrawCylinderWires - Draw a cylinder/cone wires
func DrawCylinderWires(position vector3.Float32, radiusTop float32, radiusBottom float32, height float32, slices int32, col colorex.RGBA) {
	cposition := cvec3ptr(&position)
	cradiusTop := (C.float)(radiusTop)
	cradiusBottom := (C.float)(radiusBottom)
	cheight := (C.float)(height)
	cslices := (C.int)(slices)
	ccolor := ccolorptr(&col)
	C.DrawCylinderWires(*cposition, cradiusTop, cradiusBottom, cheight, cslices, *ccolor)
}

// DrawCylinderWiresEx - Draw a cylinder wires with base at startPos and top at endPos
func DrawCylinderWiresEx(startPos vector3.Float32, endPos vector3.Float32, startRadius float32, endRadius float32, sides int32, col colorex.RGBA) {
	cstartPos := cvec3ptr(&startPos)
	cendPos := cvec3ptr(&endPos)
	cstartRadius := (C.float)(startRadius)
	cendRadius := (C.float)(endRadius)
	csides := (C.int)(sides)
	ccolor := ccolorptr(&col)
	C.DrawCylinderWiresEx(*cstartPos, *cendPos, cstartRadius, cendRadius, csides, *ccolor)
}

// DrawCapsule - Draw a capsule with the center of its sphere caps at startPos and endPos
func DrawCapsule(startPos, endPos vector3.Float32, radius float32, slices, rings int32, col colorex.RGBA) {
	cstartPos := cvec3ptr(&startPos)
	cendPos := cvec3ptr(&endPos)
	cradius := (C.float)(radius)
	cslices := (C.int)(slices)
	crings := (C.int)(rings)
	ccolor := ccolorptr(&col)
	C.DrawCapsule(*cstartPos, *cendPos, cradius, cslices, crings, *ccolor)
}

// DrawCapsuleWires - Draw capsule wireframe with the center of its sphere caps at startPos and endPos
func DrawCapsuleWires(startPos, endPos vector3.Float32, radius float32, slices, rings int32, col colorex.RGBA) {
	cstartPos := cvec3ptr(&startPos)
	cendPos := cvec3ptr(&endPos)
	cradius := (C.float)(radius)
	cslices := (C.int)(slices)
	crings := (C.int)(rings)
	ccolor := ccolorptr(&col)
	C.DrawCapsuleWires(*cstartPos, *cendPos, cradius, cslices, crings, *ccolor)
}

// DrawPlane - Draw a plane XZ
func DrawPlane(centerPos vector3.Float32, size vector2.Float32, col colorex.RGBA) {
	ccenterPos := cvec3ptr(&centerPos)
	csize := cvec2ptr(&size)
	ccolor := ccolorptr(&col)
	C.DrawPlane(*ccenterPos, *csize, *ccolor)
}

// DrawRay - Draw a ray line
func DrawRay(ray Ray, col colorex.RGBA) {
	cray := ray.cptr()
	ccolor := ccolorptr(&col)
	C.DrawRay(*cray, *ccolor)
}

// DrawGrid - Draw a grid (centered at (0, 0, 0))
func DrawGrid(slices int32, spacing float32) {
	cslices := (C.int)(slices)
	cspacing := (C.float)(spacing)
	C.DrawGrid(cslices, cspacing)
}

// LoadModel - Load model from file
func LoadModel(fileName string) Model {
	cfileName := textAlloc(fileName)
	ret := C.LoadModel(cfileName)
	return *newModelFromPointer(&ret)
}

// LoadModelFromMesh - Load model from mesh data
func LoadModelFromMesh(data Mesh) Model {
	cdata := data.cptr()
	ret := C.LoadModelFromMesh(*cdata)
	return *newModelFromPointer(&ret)
}

// IsModelValid - Check if a model is valid (loaded in GPU, VAO/VBOs)
func IsModelValid(model Model) bool {
	cmodel := model.cptr()
	ret := C.IsModelValid(*cmodel)
	v := bool(ret)
	return v
}

// UnloadModel - Unload model from memory (RAM and/or VRAM)
func UnloadModel(model *Model) {
	cmodel := model.cptr()
	C.UnloadModel(cmodel)
}

// GetModelBoundingBox - Compute model bounding box limits (considers all meshes
func GetModelBoundingBox(model Model) BoundingBox {
	cmodel := model.cptr()
	ret := C.GetModelBoundingBox(*cmodel)
	return *newBoundingBoxFromPointer(&ret)
}

// DrawModel - Draw a model (with texture if set)
func DrawModel(model Model, position vector3.Float32, scale float32, tint colorex.RGBA) {
	cmodel := model.cptr()
	cposition := cvec3ptr(&position)
	cscale := (C.float)(scale)
	ctint := ccolorptr(&tint)
	C.DrawModel(*cmodel, *cposition, cscale, *ctint)
}

// DrawModelEx - Draw a model with extended parameters
func DrawModelEx(model Model, position vector3.Float32, rotationAxis vector3.Float32, rotationAngle float32, scale vector3.Float32, tint colorex.RGBA) {
	cmodel := model.cptr()
	cposition := cvec3ptr(&position)
	crotationAxis := cvec3ptr(&rotationAxis)
	crotationAngle := (C.float)(rotationAngle)
	cscale := cvec3ptr(&scale)
	ctint := ccolorptr(&tint)
	C.DrawModelEx(*cmodel, *cposition, *crotationAxis, crotationAngle, *cscale, *ctint)
}

// DrawModelWires - Draw a model wires (with texture if set)
func DrawModelWires(model Model, position vector3.Float32, scale float32, tint colorex.RGBA) {
	cmodel := model.cptr()
	cposition := cvec3ptr(&position)
	cscale := (C.float)(scale)
	ctint := ccolorptr(&tint)
	C.DrawModelWires(*cmodel, *cposition, cscale, *ctint)
}

// DrawModelWiresEx - Draw a model wires (with texture if set) with extended parameters
func DrawModelWiresEx(model Model, position vector3.Float32, rotationAxis vector3.Float32, rotationAngle float32, scale vector3.Float32, tint colorex.RGBA) {
	cmodel := model.cptr()
	cposition := cvec3ptr(&position)
	crotationAxis := cvec3ptr(&rotationAxis)
	crotationAngle := (C.float)(rotationAngle)
	cscale := cvec3ptr(&scale)
	ctint := ccolorptr(&tint)
	C.DrawModelWiresEx(*cmodel, *cposition, *crotationAxis, crotationAngle, *cscale, *ctint)
}

// DrawModelPoints - Draw a model as points
func DrawModelPoints(model Model, position vector3.Float32, scale float32, tint colorex.RGBA) {
	cmodel := model.cptr()
	cposition := cvec3ptr(&position)
	cscale := (C.float)(scale)
	ctint := ccolorptr(&tint)
	C.DrawModelPoints(*cmodel, *cposition, cscale, *ctint)
}

// DrawModelPointsEx - Draw a model as points with extended parameters
func DrawModelPointsEx(model Model, position vector3.Float32, rotationAxis vector3.Float32, rotationAngle float32, scale vector3.Float32, tint colorex.RGBA) {
	cmodel := model.cptr()
	cposition := cvec3ptr(&position)
	crotationAxis := cvec3ptr(&rotationAxis)
	crotationAngle := (C.float)(rotationAngle)
	cscale := cvec3ptr(&scale)
	ctint := ccolorptr(&tint)
	C.DrawModelPointsEx(*cmodel, *cposition, *crotationAxis, crotationAngle, *cscale, *ctint)
}

// DrawBoundingBox - Draw bounding box (wires)
func DrawBoundingBox(box BoundingBox, col colorex.RGBA) {
	cbox := box.cptr()
	ccolor := ccolorptr(&col)
	C.DrawBoundingBox(*cbox, *ccolor)
}

// DrawBillboard - Draw a billboard texture
func DrawBillboard(camera Camera, texture Texture2D, center vector3.Float32, scale float32, tint colorex.RGBA) {
	ccamera := camera.cptr()
	ctexture := texture.cptr()
	ccenter := cvec3ptr(&center)
	cscale := (C.float)(scale)
	ctint := ccolorptr(&tint)
	C.DrawBillboard(*ccamera, *ctexture, *ccenter, cscale, *ctint)
}

// DrawBillboardRec - Draw a billboard texture defined by sourceRec
func DrawBillboardRec(camera Camera, texture Texture2D, sourceRec rect2.Float32, center vector3.Float32, size vector2.Float32, tint colorex.RGBA) {
	ccamera := camera.cptr()
	ctexture := texture.cptr()
	csourceRec := crect2ptr(&sourceRec)
	ccenter := cvec3ptr(&center)
	csize := cvec2ptr(&size)
	ctint := ccolorptr(&tint)
	C.DrawBillboardRec(*ccamera, *ctexture, *csourceRec, *ccenter, *csize, *ctint)
}

// DrawBillboardPro - Draw a billboard texture with pro parameters
func DrawBillboardPro(camera Camera, texture Texture2D, sourceRec rect2.Float32, position vector3.Float32, up vector3.Float32, size vector2.Float32, origin vector2.Float32, rotation float32, tint colorex.RGBA) {
	ccamera := camera.cptr()
	ctexture := texture.cptr()
	csourceRec := crect2ptr(&sourceRec)
	cposition := cvec3ptr(&position)
	cup := cvec3ptr(&up)
	csize := cvec2ptr(&size)
	corigin := cvec2ptr(&origin)
	crotation := (C.float)(rotation)
	ctint := ccolorptr(&tint)
	C.DrawBillboardPro(*ccamera, *ctexture, *csourceRec, *cposition, *cup, *csize, *corigin, crotation, *ctint)
}

// UpdateMeshBuffer - Update mesh vertex data in GPU for a specific buffer index
func UpdateMeshBuffer(mesh Mesh, index int, data []byte, offset int) {
	cindex := (C.int)(index)
	coffset := (C.int)(offset)
	cdataSize := (C.int)(len(data))
	C.UpdateMeshBuffer(*mesh.cptr(), cindex, unsafe.Pointer(&data[0]), cdataSize, coffset)
}

// UnloadMesh - Unload mesh from memory (RAM and/or VRAM)
func UnloadMesh(mesh *Mesh) {
	// Check list of go-managed mesh IDs
	if slices.Contains(goManagedMeshIDs, mesh.VaoID) {
		// C.UnloadMesh() only needs to read the VaoID & VboID
		// passing a temporary struct with all other fields nil makes it safe for the C code to call free()
		tempMesh := Mesh{
			VaoID: mesh.VaoID,
			VboID: mesh.VboID,
		}
		cmesh := tempMesh.cptr()
		C.UnloadMesh(cmesh)

		// remove mesh VaoID from list
		goManagedMeshIDs = slices.DeleteFunc(goManagedMeshIDs, func(id uint32) bool { return id == mesh.VaoID })
	} else {
		cmesh := mesh.cptr()
		C.UnloadMesh(cmesh)
	}
}

// DrawMesh - Draw a single mesh
func DrawMesh(mesh Mesh, material Material, transform Matrix) {
	C.DrawMesh(*mesh.cptr(), *material.cptr(), *transform.cptr())
}

// DrawMeshInstanced - Draw mesh with instanced rendering
func DrawMeshInstanced(mesh Mesh, material Material, transforms []Matrix, instances int) {
	C.DrawMeshInstanced(*mesh.cptr(), *material.cptr(), transforms[0].cptr(), C.int(instances))
}

// ExportMesh - Export mesh as an OBJ file
func ExportMesh(mesh Mesh, fileName string) {
	cfileName := textAlloc(fileName)
	cmesh := mesh.cptr()
	C.ExportMesh(*cmesh, cfileName)
}

// GetMeshBoundingBox - Compute mesh bounding box limits
func GetMeshBoundingBox(mesh Mesh) BoundingBox {
	cmesh := mesh.cptr()
	ret := C.GetMeshBoundingBox(*cmesh)
	return *newBoundingBoxFromPointer(&ret)
}

// GenMeshPoly - Generate polygonal mesh
func GenMeshPoly(sides int, radius float32) Mesh {
	csides := (C.int)(sides)
	cradius := (C.float)(radius)

	ret := C.GenMeshPoly(csides, cradius)
	return *newMeshFromPointer(&ret)
}

// GenMeshPlane - Generate plane mesh (with subdivisions)
func GenMeshPlane(width, length float32, resX, resZ int) Mesh {
	cwidth := (C.float)(width)
	clength := (C.float)(length)
	cresX := (C.int)(resX)
	cresZ := (C.int)(resZ)

	ret := C.GenMeshPlane(cwidth, clength, cresX, cresZ)
	return *newMeshFromPointer(&ret)
}

// GenMeshCube - Generate cuboid mesh
func GenMeshCube(width, height, length float32) Mesh {
	cwidth := (C.float)(width)
	cheight := (C.float)(height)
	clength := (C.float)(length)

	ret := C.GenMeshCube(cwidth, cheight, clength)
	return *newMeshFromPointer(&ret)
}

// GenMeshSphere - Generate sphere mesh (standard sphere)
func GenMeshSphere(radius float32, rings, slices int) Mesh {
	cradius := (C.float)(radius)
	crings := (C.int)(rings)
	cslices := (C.int)(slices)

	ret := C.GenMeshSphere(cradius, crings, cslices)
	return *newMeshFromPointer(&ret)
}

// GenMeshHemiSphere - Generate half-sphere mesh (no bottom cap)
func GenMeshHemiSphere(radius float32, rings, slices int) Mesh {
	cradius := (C.float)(radius)
	crings := (C.int)(rings)
	cslices := (C.int)(slices)

	ret := C.GenMeshHemiSphere(cradius, crings, cslices)
	return *newMeshFromPointer(&ret)
}

// GenMeshCylinder - Generate cylinder mesh
func GenMeshCylinder(radius, height float32, slices int) Mesh {
	cradius := (C.float)(radius)
	cheight := (C.float)(height)
	cslices := (C.int)(slices)

	ret := C.GenMeshCylinder(cradius, cheight, cslices)
	return *newMeshFromPointer(&ret)
}

// GenMeshCone - Generate cone/pyramid mesh
func GenMeshCone(radius, height float32, slices int) Mesh {
	cradius := (C.float)(radius)
	cheight := (C.float)(height)
	cslices := (C.int)(slices)

	ret := C.GenMeshCone(cradius, cheight, cslices)
	return *newMeshFromPointer(&ret)
}

// GenMeshTorus - Generate torus mesh
func GenMeshTorus(radius, size float32, radSeg, sides int) Mesh {
	cradius := (C.float)(radius)
	csize := (C.float)(size)
	cradSeg := (C.int)(radSeg)
	csides := (C.int)(sides)

	ret := C.GenMeshTorus(cradius, csize, cradSeg, csides)
	return *newMeshFromPointer(&ret)
}

// GenMeshKnot - Generate trefoil knot mesh
func GenMeshKnot(radius, size float32, radSeg, sides int) Mesh {
	cradius := (C.float)(radius)
	csize := (C.float)(size)
	cradSeg := (C.int)(radSeg)
	csides := (C.int)(sides)

	ret := C.GenMeshKnot(cradius, csize, cradSeg, csides)
	return *newMeshFromPointer(&ret)
}

// GenMeshHeightmap - Generate heightmap mesh from image data
func GenMeshHeightmap(heightmap Image, size vector3.Float32) Mesh {
	cheightmap := heightmap.cptr()
	csize := cvec3ptr(&size)

	ret := C.GenMeshHeightmap(*cheightmap, *csize)
	return *newMeshFromPointer(&ret)
}

// GenMeshCubicmap - Generate cubes-based map mesh from image data
func GenMeshCubicmap(cubicmap Image, size vector3.Float32) Mesh {
	ccubicmap := cubicmap.cptr()
	csize := cvec3ptr(&size)

	ret := C.GenMeshCubicmap(*ccubicmap, *csize)
	return *newMeshFromPointer(&ret)
}

// LoadMaterials - Load material data (.MTL)
func LoadMaterials(fileName string) []Material {
	cfileName := textAlloc(fileName)
	ccount := C.int(0)
	ret := C.LoadMaterials(cfileName, &ccount)
	v := (*[1 << 24]Material)(unsafe.Pointer(ret))[:int(ccount)]
	return v
}

// LoadMaterialDefault - Load default material (Supports: DIFFUSE, SPECULAR, NORMAL maps)
func LoadMaterialDefault() Material {
	ret := C.LoadMaterialDefault()
	return *newMaterialFromPointer(&ret)
}

// IsMaterialValid - Check if a material is valid (shader assigned, map textures loaded in GPU)
func IsMaterialValid(material Material) bool {
	cmaterial := material.cptr()
	ret := C.IsMaterialValid(*cmaterial)
	v := bool(ret)
	return v
}

// UnloadMaterial - Unload material textures from VRAM
func UnloadMaterial(material *Material) {
	cmaterial := material.cptr()
	C.UnloadMaterial(cmaterial)
}

// SetMaterialTexture - Set texture for a material map type (MATERIAL_MAP_DIFFUSE, MATERIAL_MAP_SPECULAR...)
func SetMaterialTexture(material *Material, mapType int32, texture Texture2D) {
	cmaterial := material.cptr()
	cmapType := (C.int)(mapType)
	ctexture := texture.cptr()
	C.SetMaterialTexture(cmaterial, cmapType, *ctexture)
}

// SetModelMeshMaterial - Set material for a mesh
func SetModelMeshMaterial(model *Model, meshId int32, materialId int32) {
	cmodel := model.cptr()
	cmeshId := (C.int)(meshId)
	cmaterialId := (C.int)(materialId)
	C.SetModelMeshMaterial(cmodel, cmeshId, cmaterialId)
}

// LoadModelAnimations - Load model animations from file
func LoadModelAnimations(fileName string) []ModelAnimation {
	cfileName := textAlloc(fileName)
	ccount := C.int(0)
	ret := C.LoadModelAnimations(cfileName, &ccount)
	v := (*[1 << 24]ModelAnimation)(unsafe.Pointer(ret))[:int(ccount)]
	return v
}

// UpdateModelAnimation - Update model animation pose (CPU)
func UpdateModelAnimation(model Model, anim ModelAnimation, frame int32) {
	cmodel := model.cptr()
	canim := anim.cptr()
	cframe := (C.int)(frame)
	C.UpdateModelAnimation(*cmodel, *canim, cframe)
}

// UpdateModelAnimationBones - Update model animation mesh bone matrices (GPU skinning)
func UpdateModelAnimationBones(model Model, anim ModelAnimation, frame int32) {
	cmodel := model.cptr()
	canim := anim.cptr()
	cframe := (C.int)(frame)
	C.UpdateModelAnimationBones(*cmodel, *canim, cframe)
}

// UnloadModelAnimation - Unload animation data
func UnloadModelAnimation(anim *ModelAnimation) {
	canim := anim.cptr()
	C.UnloadModelAnimation(canim)
}

// UnloadModelAnimations - Unload animation array data
func UnloadModelAnimations(animations []ModelAnimation) {
	C.UnloadModelAnimations((*C.ModelAnimation)(unsafe.Pointer(&animations[0])), (C.int)(len(animations)))
}

// IsModelAnimationValid - Check model animation skeleton match
func IsModelAnimationValid(model Model, anim ModelAnimation) bool {
	cmodel := model.cptr()
	canim := anim.cptr()
	ret := C.IsModelAnimationValid(*cmodel, *canim)
	v := bool(ret)
	return v
}

// CheckCollisionSpheres - Detect collision between two spheres
func CheckCollisionSpheres(centerA vector3.Float32, radiusA float32, centerB vector3.Float32, radiusB float32) bool {
	ccenterA := cvec3ptr(&centerA)
	cradiusA := (C.float)(radiusA)
	ccenterB := cvec3ptr(&centerB)
	cradiusB := (C.float)(radiusB)
	ret := C.CheckCollisionSpheres(*ccenterA, cradiusA, *ccenterB, cradiusB)
	v := bool(ret)
	return v
}

// CheckCollisionBoxes - Detect collision between two bounding boxes
func CheckCollisionBoxes(box1 BoundingBox, box2 BoundingBox) bool {
	cbox1 := box1.cptr()
	cbox2 := box2.cptr()
	ret := C.CheckCollisionBoxes(*cbox1, *cbox2)
	v := bool(ret)
	return v
}

// CheckCollisionBoxSphere - Detect collision between box and sphere
func CheckCollisionBoxSphere(box BoundingBox, centerSphere vector3.Float32, radiusSphere float32) bool {
	cbox := box.cptr()
	ccenterSphere := cvec3ptr(&centerSphere)
	cradiusSphere := (C.float)(radiusSphere)
	ret := C.CheckCollisionBoxSphere(*cbox, *ccenterSphere, cradiusSphere)
	v := bool(ret)
	return v
}

// GetRayCollisionSphere - Get collision info between ray and sphere
func GetRayCollisionSphere(ray Ray, center vector3.Float32, radius float32) RayCollision {
	cray := ray.cptr()
	ccenter := cvec3ptr(&center)
	cradius := (C.float)(radius)
	ret := C.GetRayCollisionSphere(*cray, *ccenter, cradius)
	return *newRayCollisionFromPointer(&ret)
}

// GetRayCollisionBox - Get collision info between ray and box
func GetRayCollisionBox(ray Ray, box BoundingBox) RayCollision {
	cray := ray.cptr()
	cbox := box.cptr()
	ret := C.GetRayCollisionBox(*cray, *cbox)
	return *newRayCollisionFromPointer(&ret)
}

// GetRayCollisionMesh - Get collision info between ray and mesh
func GetRayCollisionMesh(ray Ray, mesh Mesh, transform Matrix) RayCollision {
	cray := ray.cptr()
	cmesh := mesh.cptr()
	ctransform := transform.cptr()
	ret := C.GetRayCollisionMesh(*cray, *cmesh, *ctransform)
	return *newRayCollisionFromPointer(&ret)
}

// GetRayCollisionTriangle - Get collision info between ray and triangle
func GetRayCollisionTriangle(ray Ray, p1, p2, p3 vector3.Float32) RayCollision {
	cray := ray.cptr()
	cp1 := cvec3ptr(&p1)
	cp2 := cvec3ptr(&p2)
	cp3 := cvec3ptr(&p3)
	ret := C.GetRayCollisionTriangle(*cray, *cp1, *cp2, *cp3)
	return *newRayCollisionFromPointer(&ret)
}

// GetRayCollisionQuad - Get collision info between ray and quad
func GetRayCollisionQuad(ray Ray, p1, p2, p3, p4 vector3.Float32) RayCollision {
	cray := ray.cptr()
	cp1 := cvec3ptr(&p1)
	cp2 := cvec3ptr(&p2)
	cp3 := cvec3ptr(&p3)
	cp4 := cvec3ptr(&p4)
	ret := C.GetRayCollisionQuad(*cray, *cp1, *cp2, *cp3, *cp4)
	return *newRayCollisionFromPointer(&ret)
}
