function sendCmd(cmd, onFin, onErr) {
	const req = new XMLHttpRequest();
	req.open('GET', 'cmd/' + cmd, true);

	req.onload = function() {
		if (req.status >= 200 && request.status < 400) {
			onFin(req.responseText);
		} else {
			onErr(req.status)
		}
	}
	
	req.onerror = onErr;

	req.send();
}

function registerButton(bId) {
	const button = document.getElementById(bId);
	button.addEventListener('click', function () {
		sendCmd(bId, null, null)
	});
}

registerButton("zoom-in");
registerButton("zoom-out");
registerButton("left");
registerButton("right");
registerButton("up");
registerButton("down");
registerButton("rot-left");
registerButton("rot-right");
registerButton("rot-up");
registerButton("rot-down");
