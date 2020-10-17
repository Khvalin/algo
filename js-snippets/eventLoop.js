(function() {
    console.log("start");
    Promise.resolve().then(() => {console.log("promise")})
    setTimeout(() => {console.log("setTimeout")}, 0); 
    setInterval(() => {console.log("setInterval")}, 0); 
    console.log("finish");
})();
