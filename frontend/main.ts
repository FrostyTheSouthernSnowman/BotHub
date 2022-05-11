// Resize stuff
var topDiv = document.getElementById('overlay-top');
var topDivResizer = document.createElement('div');
topDivResizer.className = 'resizer';
topDivResizer.style.width = '100%';
topDivResizer.style.height = '2%';
topDivResizer.style.background = '#4e4e4e';
topDivResizer.style.position = 'absolute';
topDivResizer.style.right = "0";
topDivResizer.style.bottom = "0";
topDivResizer.style.cursor = 'se-resize';
if (topDiv != null) {
    topDiv.appendChild(topDivResizer);
    topDivResizer.addEventListener('mousedown', initTopResize, false);
}

function initTopResize(e: any) {
    window.addEventListener('mousemove', topResize, false);
    window.addEventListener('mouseup', stopTopResize, false);
}
function topResize(e: any) {
    if (topDiv != null) {
        topDiv.style.height = (e.clientY - topDiv.offsetTop) + 'px';
    }
}
function stopTopResize(e: any) {
    window.removeEventListener('mousemove', topResize, false);
    window.removeEventListener('mouseup', stopTopResize, false);
}

var bottomDiv = document.getElementById('overlay-bottom');
var bottomDivResizer = document.createElement('div');
bottomDivResizer.className = 'resizer';
bottomDivResizer.style.width = '100%';
bottomDivResizer.style.height = '2%';
bottomDivResizer.style.background = '#0000ff';
bottomDivResizer.style.position = 'absolute';
bottomDivResizer.style.right = "0";
bottomDivResizer.style.top = "0";
bottomDivResizer.style.cursor = 'se-resize';
if (bottomDiv != null) {
    bottomDiv.appendChild(bottomDivResizer);
    bottomDivResizer.addEventListener('mousedown', initBottomResize, false);
}

function initBottomResize(e: any) {
    window.addEventListener('mousemove', bottomResize, false);
    window.addEventListener('mouseup', stopBottomResize, false);
}
function bottomResize(e: any) {
    if (bottomDiv != null) {
        bottomDiv.style.height = (window.innerHeight - e.clientY) + 'px';
    }
}
function stopBottomResize(e: any) {
    window.removeEventListener('mousemove', bottomResize, false);
    window.removeEventListener('mouseup', stopBottomResize, false);
}

var rightDiv = document.getElementById('overlay-right');
var rightDivResizer = document.createElement('div');
rightDivResizer.className = 'resizer';
rightDivResizer.style.width = '2%';
rightDivResizer.style.height = '100%';
rightDivResizer.style.background = '#00ff00';
rightDivResizer.style.position = 'absolute';
rightDivResizer.style.left = "0";
rightDivResizer.style.bottom = "0";
rightDivResizer.style.cursor = 'se-resize';
if (rightDiv != null) {
    rightDiv.appendChild(rightDivResizer);
    rightDivResizer.addEventListener('mousedown', initRightResize, false);
}

function initRightResize(e: any) {
    window.addEventListener('mousemove', rightResize, false);
    window.addEventListener('mouseup', stopRightResize, false);
}
function rightResize(e: any) {
    if (rightDiv != null) {
        rightDiv.style.width = (window.innerWidth - e.clientX) + 'px';
    }
}
function stopRightResize(e: any) {
    window.removeEventListener('mousemove', rightResize, false);
    window.removeEventListener('mouseup', stopRightResize, false);
}

var leftDiv = document.getElementById('overlay-left');
var leftDivResizer = document.createElement('div');
leftDivResizer.className = 'resizer';
leftDivResizer.style.width = '2%';
leftDivResizer.style.height = '100%';
leftDivResizer.style.background = '#ff0000';
leftDivResizer.style.position = 'absolute';
leftDivResizer.style.right = "0";
leftDivResizer.style.bottom = "0";
leftDivResizer.style.cursor = 'se-resize';
if (leftDiv != null) {
    leftDiv.appendChild(leftDivResizer);
    leftDivResizer.addEventListener('mousedown', initLeftResize, false);
}

function initLeftResize(e: any) {
    window.addEventListener('mousemove', leftResize, false);
    window.addEventListener('mouseup', stopLeftResize, false);
}
function leftResize(e: any) {
    if (leftDiv != null) {
        leftDiv.style.width = (e.clientX - leftDiv.offsetLeft) + 'px';
    }
}
function stopLeftResize(e: any) {
    window.removeEventListener('mousemove', leftResize, false);
    window.removeEventListener('mouseup', stopLeftResize, false);
}

// Three-js stuff
import * as THREE from 'three';
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls"
import { if_arrow_give_name_or_return_false } from './utils';
import { RigidBodySphereBoundingBox } from './types';

var scene = new THREE.Scene();

var camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);

var orbit = new OrbitControls(camera)

var renderer = new THREE.WebGLRenderer();


// Drag robot stuff
var directionOfDrag: boolean | "x" | "y" | "z" = false;

var stuffToDrag: any = [];

var arrowHeight = 0;

var intersection: any = null;

function initDragRobotAlongAxis() {
    window.addEventListener('mousemove', dragRobotAlongAxis, false);
    window.addEventListener('mouseup', stopDragRobotAlongAxis, false);
}
function dragRobotAlongAxis(e: any) {
    if (directionOfDrag == "z") {
        for (let i in stuffToDrag) {
            stuffToDrag[i].position. z = ((e.clientY / window.innerHeight) * 2 - 1) - arrowHeight;
        }
    } else if (directionOfDrag == "y") {
        for (let i in stuffToDrag) {
            stuffToDrag[i].position.y = ((e.clientX / window.innerWidth) * 2 - 1) - arrowHeight;
        }
    } else if (directionOfDrag == "x") {
        for (let i in stuffToDrag) {
            stuffToDrag[i].position.x = ((e.clientX / window.innerWidth) * 2 - 1) - arrowHeight;
        }
    }
}
function stopDragRobotAlongAxis(e: any) {
    window.removeEventListener('mousemove', dragRobotAlongAxis, false);
    window.removeEventListener('mouseup', stopDragRobotAlongAxis, false);
}

var raycaster: any, mouse = { x: 0, y: 0 };


function raycast(e: any) {
    // Step 1: Detect light helper
    //1. sets the mouse position with a coordinate system where the center
    //   of the screen is the origin
    mouse.x = (e.clientX / window.innerWidth) * 2 - 1;
    mouse.y = - (e.clientY / window.innerHeight) * 2 + 1;

    //2. set the picking ray from the camera position and mouse coordinates
    raycaster.setFromCamera(mouse, camera);

    //3. compute intersections (note the 2nd parameter)
    var intersects = raycaster.intersectObjects(scene.children, true);

    for (var i = 0; i < intersects.length; i++) {
        var intersect = intersects[i]
        var object = intersect.object
        if (object.position.x == 0 && object.position.y == 0 && object.position.z == 0) { continue; }

        //TODO: if_arrow_give_name_or_return_false if a temporary proof-of-concept placeholder function.
        //Todo: Refactor method to return so sort of direction or vector instead of string. 
        //TODO: vector would make it simpler to process the drag-and-drop controls
        directionOfDrag = if_arrow_give_name_or_return_false({ object })
        if (directionOfDrag !== false) {
            stuffToDrag = []
            // Get all object at the same position
            for (let i in objects) {
                if (objects[i].position.x === object.parent.position.x && objects[i].position.y === object.parent.position.y && objects[i].position.z === object.parent.position.z) {
                    stuffToDrag.push(objects[i])
                }
            }
            // Move all objects to mouse's location minus arrow length
            arrowHeight = object.parent.line.scale.y + object.parent.cone.position.y

            intersection = intersect

            initDragRobotAlongAxis()

            continue
        }

        var x_arrow = scene.getObjectByName("x_arrow")
        if (x_arrow != undefined) {
            var y_arrow = scene.getObjectByName("y_arrow")
            var z_arrow = scene.getObjectByName("z_arrow")
            if (y_arrow != undefined && z_arrow != undefined) {
                scene.remove(x_arrow)
                scene.remove(y_arrow)
                scene.remove(z_arrow)
            }
        }

        const x_dir = new THREE.Vector3(1, 0, 0);
        const y_dir = new THREE.Vector3(0, 1, 0);
        const z_dir = new THREE.Vector3(0, 0, 1);

        //normalize the direction vector (convert to vector of length 1)
        x_dir.normalize();
        y_dir.normalize();
        z_dir.normalize();

        const length = 2 * object.geometry.boundingSphere.radius
        const x_hex = 0xff0000;
        const y_hex = 0x00ff00;
        const z_hex = 0x0000ff;

        const x_arrowHelper = new THREE.ArrowHelper(x_dir, object.position, length, x_hex);
        const y_arrowHelper = new THREE.ArrowHelper(y_dir, object.position, length, y_hex);
        const z_arrowHelper = new THREE.ArrowHelper(z_dir, object.position, length, z_hex);
        x_arrowHelper.name = "x_arrow"
        y_arrowHelper.name = "y_arrow"
        z_arrowHelper.name = "z_arrow"
        x_arrowHelper.position.x = object.position.x
        y_arrowHelper.position.y = object.position.y
        z_arrowHelper.position.z = object.position.z
        objects.push(x_arrowHelper)
        objects.push(y_arrowHelper)
        objects.push(z_arrowHelper)
        scene.add(x_arrowHelper);
        scene.add(y_arrowHelper);
        scene.add(z_arrowHelper);
    }
    // Step 2: Detect normal objects
    //1. sets the mouse position with a coordinate system where the center
    //   of the screen is the origin
    mouse.x = (e.clientX / window.innerWidth) * 2 - 1;
    mouse.y = - (e.clientY / window.innerHeight) * 2 + 1;

    //2. set the picking ray from the camera position and mouse coordinates
    raycaster.setFromCamera(mouse, camera);

    //3. compute intersections (no 2nd parameter true anymore)
    var intersects = raycaster.intersectObjects(scene.children);

    for (var i = 0; i < intersects.length; i++) {
        /*
            An intersection has the following properties :
                - object : intersected object (THREE.Mesh)
                - distance : distance from camera to intersection (number)
                - face : intersected face (THREE.Face3)
                - faceIndex : intersected face index (number)
                - point : intersection point (THREE.Vector3)
                - uv : intersection point in the object's UV coordinates (THREE.Vector2)
        */
    }

}

raycaster = new THREE.Raycaster();
renderer.domElement.addEventListener('mousedown', raycast, false);

renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

var table_x: number = 50;
var table_y: number = 50;

var objects: any = []
var robots: any[] = [];
var robot: any;

var streamSocket: WebSocket

// Simulation stuff
function pauseSimulation() {
    streamSocket.send(JSON.stringify({ type: "pause" }));
}

function resetSimulation() {
    streamSocket.send(JSON.stringify({ type: "reset" }));
}

function playSimulation() {
    streamSocket.send(JSON.stringify({ type: "play" }));
}

async function addRobot() {
    streamSocket.send(JSON.stringify({type: "addObject"}))
    console.log("new object added")
    var robot_geometry = new THREE.SphereGeometry(1);
    var robot_material = new THREE.MeshBasicMaterial({
        color: "orange",
    });
    var robot = new THREE.Mesh(robot_geometry, robot_material);
    robot.position.z = 0.5;
    robots.push(robot);
    objects.push(robot)
    scene.add(robot);
}

document.querySelector('#add-button')!.addEventListener('click', (e:Event) => addRobot());
document.querySelector('#play-button')!.addEventListener('click', (e:Event) => playSimulation());
document.querySelector('#pause-button')!.addEventListener('click', (e:Event) => pauseSimulation());
document.querySelector('#reset-button')!.addEventListener('click', (e:Event) => resetSimulation());

var scene_has_initialized: boolean = false

function initializeScene(data: any) {
    var geometry = new THREE.BoxGeometry(table_x, table_y, 0.01);
    var material = new THREE.MeshBasicMaterial({
        color: "grey",
    });
    var floor = new THREE.Mesh(geometry, material);
    scene.add(floor)

    var robot_geometry = new THREE.SphereGeometry(1);
    var robot_material = new THREE.MeshBasicMaterial({
        color: "orange",
    });
    robot = new THREE.Mesh(robot_geometry, robot_material);
    robots.push(robot)
    scene.add(robot);

    camera.position.y = -5;
    camera.position.z = 2;
    camera.position.x = 0;
    camera.rotation.x = 1;
    robot.position.z = data[0].position.z;
    robot.position.y = data[0].position.y;
    robot.position.x = data[0].position.x;
    scene_has_initialized = true

    var animate = function () {
        requestAnimationFrame(animate);
        renderer.render(scene, camera);
    };

    animate();
}

function initRobot(): void {
    streamSocket = new WebSocket("ws://localhost/api/stream-simulation");

    type EventType = {
        data: string
    }

    streamSocket.onmessage = function (event: EventType) {
        let data: RigidBodySphereBoundingBox[] = JSON.parse(event.data.toLowerCase())
        if (!scene_has_initialized) {
            initializeScene(data)

        } else {
            for (let i = 0; i < data.length; i++) {
                robots[i].position.x = data[i].position.x
                robots[i].position.y = data[i].position.y
                robots[i].position.z = data[i].position.z
            }
        }
    }
}

initRobot()