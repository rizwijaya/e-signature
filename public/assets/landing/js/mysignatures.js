// My Signatures
var signBox = new SignaturePad(document.getElementById("signbox"), {
    backgroundColor: "rgba(255, 255, 255, 0)",
    penColor: "rgb(0, 0, 0)",
});

function addSign() {
    var data = signBox.toDataURL("image/png");
    var img = data.replace(/^data:image\/(png|jpg);base64,/, "");

    var settings = {
        url: "/add-signatures",
        method: "POST",
        timeout: 0,
        headers: {
            "Content-Type": "text/plain",
        },
        data: '{\r\n    "unique": ' + user + ',\r\n    "signature": "' + img + '"\r\n}',
    };

    $.ajax(settings).done(function (response) {
        //console.log(response);
        location.reload();
    });
}
