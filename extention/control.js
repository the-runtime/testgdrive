let downloadId;
let downloadUrl;

console.log("process started");

browser.downloads.onCreated.addListener(handleDownload);

console.log("event strated");

function onError(error) {
    console.log('Error: ${error}');
}


function popOpt(){

}

function directDownDrive(){
    console.log("download redirected");
    let downUrl = "http://127.0.0.1:8000/gdrive/down/?url="+downloadUrl;
    //downUrl = "https://theruntime.software/gdrive/down?url="+downloadUrl;
    console.log(downUrl);
    window.location.replace(downUrl);
    //var wind = window.open(downUrl,"_self");
}

function pauseDownload(downloadId){
    console.log("Download pause started");
    let pausing = browser.downloads.pause(downloadId);
    //pausing.then(()=> {console.log("paused")
    //    }, onError);
    //pausing.then(popOpt,onError);
    pausing.then(directDownDrive,onError);
}


function handleDownload(down){
    console.log("handle download started");
    downloadId = down.id;
    //down[0].id = 12345;
    //downloadId = 12345;
    console.log("Downloadid   "+ downloadId);
    console.log("hello world");
    downloadUrl = down.url;
    console.log(downloadUrl);
    pauseDownload(downloadId);
    //console.log(down.id);
}



