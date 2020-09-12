const stop = document.getElementById('stop');
const record = document.getElementById('record');
stop.disabled = true;

  
const handleSuccess = function(stream) {
	let chunks = [];
    const options = {mimeType: 'audio/webm'};
	const mediaRecorder = new MediaRecorder(stream, options);
	console.log("start recording")
		record.onclick = function() {
			mediaRecorder.start();
			console.log(mediaRecorder.state);
			console.log("recorder started");
			record.style.background = "red";
	
			stop.disabled = false;
			record.disabled = true;
        }
        stop.onclick = function() {
			mediaRecorder.stop();
			console.log(mediaRecorder.state);
			console.log("recorder stopped");
			console.log(mediaRecorder.state);
			record.style.background = "";
			record.style.color = "";
			stop.disabled = true;
			record.disabled = false;
        }
	console.log("listen...")
	mediaRecorder.onstop = function() {
		console.log("start downloading")
		let formdata = new FormData() ;
		let soundBlob=new Blob(chunks);
		formdata.append('soundBlob', soundBlob,  'test.wav'); 
		var serverUrl = '/upload/';
		req = new XMLHttpRequest()
		req.open('POST', serverUrl, true);
		req.send(formdata);
	};

	mediaRecorder.ondataavailable= function(e) {
		console.log("data available",e.data,e.data.size)
      	if (e.data && e.data.size > 0) {
			console.log("data push")  
			chunks.push(e.data);
      	}
	};
};

	navigator.mediaDevices.getUserMedia({ audio: true, video: false })
		.then(handleSuccess);