// Resize stuff
var topDiv = document.getElementById('overlay-top');
var topDivResizer = document.createElement('div');
topDivResizer.className = 'resizer';
topDivResizer.style.width = '100%';
topDivResizer.style.height = '2%';
topDivResizer.style.background = '#4e4e4e';
topDivResizer.style.position = 'absolute';
topDivResizer.style.right = 0;
topDivResizer.style.bottom = 0;
topDivResizer.style.cursor = 'se-resize';
topDiv.appendChild(topDivResizer);
topDivResizer.addEventListener('mousedown', initTopResize, false);

function initTopResize(e) {
    window.addEventListener('mousemove', topResize, false);
    window.addEventListener('mouseup', stopTopResize, false);
}
function topResize(e) {
    topDiv.style.height = (e.clientY - topDiv.offsetTop) + 'px';
}
function stopTopResize(e) {
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
bottomDivResizer.style.right = 0;
bottomDivResizer.style.top = 0;
bottomDivResizer.style.cursor = 'se-resize';
bottomDiv.appendChild(bottomDivResizer);
bottomDivResizer.addEventListener('mousedown', initBottomResize, false);

function initBottomResize(e) {
    window.addEventListener('mousemove', bottomResize, false);
    window.addEventListener('mouseup', stopBottomResize, false);
}
function bottomResize(e) {
    bottomDiv.style.height = (window.innerHeight - e.clientY) + 'px';
}
function stopBottomResize(e) {
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
rightDivResizer.style.left = 0;
rightDivResizer.style.bottom = 0;
rightDivResizer.style.cursor = 'se-resize';
rightDiv.appendChild(rightDivResizer);
rightDivResizer.addEventListener('mousedown', initRightResize, false);

function initRightResize(e) {
    window.addEventListener('mousemove', rightResize, false);
    window.addEventListener('mouseup', stopRightResize, false);
}
function rightResize(e) {
    rightDiv.style.width = (window.innerWidth - e.clientX) + 'px';
}
function stopRightResize(e) {
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
leftDivResizer.style.right = 0;
leftDivResizer.style.bottom = 0;
leftDivResizer.style.cursor = 'se-resize';
leftDiv.appendChild(leftDivResizer);
leftDivResizer.addEventListener('mousedown', initLeftResize, false);

function initLeftResize(e) {
    window.addEventListener('mousemove', leftResize, false);
    window.addEventListener('mouseup', stopLeftResize, false);
}
function leftResize(e) {
    leftDiv.style.width = (e.clientX - leftDiv.offsetLeft) + 'px';
}
function stopLeftResize(e) {
    window.removeEventListener('mousemove', leftResize, false);
    window.removeEventListener('mouseup', stopLeftResize, false);
}

// Three-js stuff
var scene = new THREE.Scene();
var camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);

var renderer = new THREE.WebGLRenderer();

var raycaster, mouse = { x: 0, y: 0 };

function raycast(e) {
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
        if (object.position.y == 1.7320508075688772) {
            console.log(e)
        }

        var x_arrow = scene.getObjectByName("x_arrow")
        if (x_arrow != undefined) {
            console.log(x_arrow.position)
            var y_arrow = scene.getObjectByName("y_arrow")
            var z_arrow = scene.getObjectByName("z_arrow")
            scene.remove(x_arrow)
            scene.remove(y_arrow)
            scene.remove(z_arrow)
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
        z_arrowHelper.position.x = object.position.x
        z_arrowHelper.position.y = object.position.y
        z_arrowHelper.position.z = object.position.z
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

var controls = new THREE.OrbitControls(camera, renderer.domElement);

var table_x_and_y = prompt("Specify the length of the table in X,Y format.");
var table_pos = table_x_and_y.split(",");
var table_x = table_pos[0];
var table_y = table_pos[1];

var robot_x_and_y = prompt("Where would you like to initialize the robot and where should it face? Please use the format X,Y,F where f is NORTH, EAST, SOUTH, or WEST.");
var robot_pos = robot_x_and_y.split(",");
var robot_x = robot_pos[0];
var robot_y = robot_pos[1];
var robot_f = robot_pos[2];

var robots = [];
var robot;

var streamSocket

// Simulation stuff
function pauseSimulation() {
    streamSocket.send(JSON.stringify({ type: "pause" }));
}

function resetSimulation() {
    streamSocket.send(JSON.stringify({ type: "reset" }));
}

function playSimulation() {
    if (streamSocket == undefined) {
        streamSocket = new WebSocket("ws://localhost/api/stream-simulation");

        streamSocket.onmessage = function (event) {
            bot = JSON.parse(event.data)[0];
            console.log(bot)
            robot.position.x = bot.X;
            robot.position.y = bot.Y;
            robot.position.z = bot.Z;
        }
    } else {
        streamSocket.send(JSON.stringify({ type: "play" }));
    }
}

function addRobot() {
    var robot_geometry = new THREE.BoxGeometry(1, 1, 1);
    var robot_material = new THREE.MeshBasicMaterial({
        color: "orange",
    });
    var robot = new THREE.Mesh(robot_geometry, robot_material);
    robot.position.z = 0.5;
    robots.push(robot);
    scene.add(robot);
}

async function initRobot(table_x, table_y, robot_x, robot_y, robot_f) {
    const data = { x: parseInt(table_x), y: parseInt(table_y) }

    var response = await fetch("http://localhost/api/set-position", { method: 'POST', body: JSON.stringify(data) });
    if (response.status = 200) {
        response = await fetch("http://localhost/api/place-robot", { method: 'POST', body: JSON.stringify({ x: parseInt(robot_x), y: parseInt(robot_y), f: robot_f }) })
        return response;
    }
}
initRobot(table_x, table_y, robot_x, robot_y, robot_f).then(response => response.json()).then(data => {
    if (!data.f) {
        return;
    }
    var geometry = new THREE.BoxGeometry(table_x, table_y, 0.01);
    var material = new THREE.MeshBasicMaterial({
        color: "grey",
    });
    var floor = new THREE.Mesh(geometry, material);
    scene.add(floor);

    var robot_geometry = new THREE.BoxGeometry(1, 1, 1);
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
    robot.position.z = 1.00;
    robot.position.y = data.y;
    robot.position.x = data.x;

    var animate = function () {
        requestAnimationFrame(animate);
        renderer.render(scene, camera);
    };

    animate();
})