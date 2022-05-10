interface Vector3 {
    x: number
    y: number
    z: number
}

interface RigidBodySphereBoundingBox {
    position: Vector3
    velocity: Vector3
    mass: number
    radius: number
}

export {Vector3, RigidBodySphereBoundingBox}