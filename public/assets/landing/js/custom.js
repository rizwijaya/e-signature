// My Signatures
var signBox = new SignaturePad(document.getElementById('signbox'), {
  backgroundColor: 'rgba(255, 255, 255, 0)',
  penColor: 'rgb(0, 0, 0)'
});

function addSign() {
  var data = signBox.toDataURL('image/png');
  var img = data.replace(/^data:image\/(png|jpg);base64,/, "");

  var settings = {
    "url": "/add-signatures",
    "method": "POST",
    "timeout": 0,
    "headers": {
      "Content-Type": "text/plain"
    },
    "data": "{\r\n    \"unique\": " + user + ",\r\n    \"signature\": \"" + img + "\"\r\n}",
  };

  $.ajax(settings).done(function (response) {
    //console.log(response);
    location.reload();
  });
}

 // var wrapper = document.getElementById("signature-pad");
  // var canvas = wrapper.querySelector("canvas");
  // canvas.width = "400px";
  // canvas.height= "200px";
//   var ratio = Math.max(window.devicePixelRatio || 1, 1);
//  // canvas.getContext("2d").scale(1, 1);
//   function resizeCanvas() {
//    // var ratio = Math.max(window.devicePixelRatio || 1, 1);
//     // if (canvas.offsetWidth == 0 || canvas.offsetHeight == 0) {
//     //   canvas.width = window.innerWidth * ratio;
//     //   canvas.height = window.innerHeight * ratio;
//     //   canvas.getContext("2d").scale(ratio, ratio);
//     // } else {
//       canvas.width = canvas.offsetWidth * ratio;
//       canvas.height = canvas.offsetHeight * ratio;
//       canvas.getContext("2d").scale(ratio, ratio);
//     //}
//     console.log(canvas.offsetHeight);
//     console.log(canvas.offsetWidth);

//     console.log(ratio);
//     console.log(canvas.width);
//     console.log(canvas.height);
//   }
//   window.default
//   window.onresize = resizeCanvas;
//   resizeCanvas();