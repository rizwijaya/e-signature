var doc;
var fileReader = new FileReader();
var sign_x = 10;
var sign_y = 10;
var sign_h = 0;
var sign_w = 0;
var sign_page = 1;
var sign_status = false;
var curPage;
const pageInput = document.querySelector("#paginationinput");

function sign(klik) {
  if (klik == 1) {
    //IF Button Click to Signing
    //Changing button to cancel
    $(".signing").html(
      `
<div style="position:absolute; top: 0; right:0">
  <i id="signBtnSave" class="btn btn-sm btn-primary"><a onclick="sign(3)">save</a></i>
  <i id="signBtnCancel" class="btn btn-sm btn-danger"><a onclick="sign(2)">cancel</a></i>
</div>
      `
    );

    $("#SignImg").addClass("Sign-Img");
    $("#SignImg").removeClass("hide_page");

    const canvEl = document.querySelector("#PDFSign");
    const canvRect = canvEl.getBoundingClientRect();

    $("#SignImg").data("x", 10);
    $("#SignImg").data("y", 10);
    $("#SignImg").css("width", canvRect.width / 2);
    $("#SignImg").css("transform", "translate(10px, 10px)");

    sign_status = false;
  } else if (klik == 2) {
    //IF Cancel Signing
    //Changing button to sign
    $(".signing").html(
      `
<div style="position:absolute; top: 0; right:0">
  <i id="signBtn" class="btn btn-sm btn-primary"><a onclick="sign(1)">sign</a></i>
</div>
`
    );
    $("#SignImg").addClass("hide_page");
    $("#SignImg").removeClass("Sign-Img");
    sign_status = false;
  } else if (klik == 3) {
    //If Button Save
    //Changing button to sign
    $(".signing").html(
      `
<div style="position:absolute; top: 0; right:0">
  <i id="signBtnEdit" class="btn btn-sm btn-primary"><a onclick="sign(4)">Edit</a></i>
</div>
`
    );
    $("#SignImg").removeClass("Sign-Img-drag");
    $("#SignImg").addClass("Sign-bdr");
    sign_status = true;
    sign_page = curPage;
    window.applySignPosition();
  } else if (klik == 4) {
    //IF Button Edit
    //Changing button to sign
    $(".signing").html(
      `
<div style="position:absolute; top: 0; right:0">
  <i id="signBtnSave" class="btn btn-sm btn-primary"><a onclick="sign(3)">save</a></i>
  <i id="signBtnDelete" class="btn btn-sm btn-danger"><a onclick="sign(2)">delete</a></i>
</div>
`
    );
    $("#SignImg").addClass("Sign-Img-drag");
    $("#SignImg").removeClass("hide_page");
    $("#SignImg").removeClass("Sign-bdr");
    sign_status = false;
  }
  // else {
  //   $(".signing").html( //Changing button to sign
  //   '<i id="signBtn" class="btn btn-sm btn-primary" style="position:absolute; top:2%; left:87%"><a onclick="sign(1)">sign</a></i>'
  //   );
  //   $("#SignImg").addClass("hide_page");
  //   $("#SignImg").removeClass("Sign-Img");
  // }
}

// function setPagination() {
//   $(".pagination.bottom").html(
//     '<li style="margin-right: 5px;" class="btn btn-primary disabled"><a onclick="prevPage()"><i class="material-icons">left</i></a></li><li style="margin-right: 5px;" class="btn btn-primary active"><a onclick="renderingPage(doc, 1)">1</a></li><li style="margin-right: 5px;" class="btn btn-primary disabled"><a onclick="nextPage()"><i class="material-icons">right</i></a></li>'
//   );
//   if (doc.numPages > 1) {
//     var last = $(".pagination.bottom li").last();
//     for (i = 2; i <= doc.numPages; i++) {
//       last.before(
//         '<li style="margin-right: 5px;" class="btn btn-primary"><a onclick="renderingPage(doc, ' +
//           i +
//           ')">' +
//           i +
//           "</a></li>"
//       );
//     }
//     last.removeClass("disabled");
//   }
// }
$("form").bind("keypress", function (e) {
  if (e.keyCode == 13) {
      $("#finish").attr('value');
      //add more buttons here
      return false;
  }
});

function prevPage() {
  pageInput.value = curPage - 1;
  renderingPage(doc, curPage - 1);
}
function nextPage() {
  pageInput.value = curPage + 1;
  renderingPage(doc, curPage + 1);
}

//pageInput.addEventListener('paginationinput', function(e) {
  //Check empty value
  // if (e.detail.value != "") {
  //   console.log(e.target.value);
  //   renderingPage(e.target.value);
  // }
//})
function changePage() {
  if (pageInput.value != "") {
    //console.log(parseInt(pageInput.value));
    renderingPage(doc, parseInt(pageInput.value));
  }
}


function renderingPage(pdf, pageNumber) {
  //If page not available
  if (pageNumber < 1 || pageNumber > pdf.numPages) {
    alert("Enter between 1 and " + pdf.numPages);
    return;
  }
  //Pagination Button
  if (curPage != pageNumber) {
    $(".pagination.bottom .active").removeClass("active");
    $(".pagination.bottom li:nth-child(" + (pageNumber + 1) + ")").addClass(
      "active"
    );
    $(".pagination.bottom li").first().removeClass("disabled");
    $(".pagination.bottom li").last().removeClass("disabled");
  }
  //First Page
  if (pageNumber == 1) {
    $(".pagination.bottom li").first().addClass("disabled");
    if (doc.numPages > 1) {
      $(".pagination.bottom li").last().removeClass("disabled");
    }
  }
  //Last Page
  if (pageNumber == doc.numPages) {
    $(".pagination.bottom li").last().addClass("disabled");
    if (doc.numPages > 1) {
      $(".pagination.bottom li").first().removeClass("disabled");
    }
  }
  //Check Signing or Not Sign in Pages
  if (sign_status) {
    if (pageNumber == sign_page) {
      $("#SignImg").removeClass("hide_page");
    } else {
      $("#SignImg").addClass("hide_page");
    }
  }
  // console.log("Page PDF Reader" + pageNumber)
  // console.log("Current Pages: " + curPage)
  curPage = pageNumber;
  //sign_page = pageNumber;
  //$("#pageinfo").text(pageNumber + "/" + pdf.numPages);
  pdf.getPage(pageNumber).then(function (page) {
    //console.log("Page loaded");

    var scale = 1.5;
    var viewport = page.getViewport({ scale: scale });

    // Prepare canvas using PDF page dimensions
    var canvas = document.getElementById("PDFSign");
    var context = canvas.getContext("2d");
    canvas.height = viewport.height;
    canvas.width = viewport.width;

    // Render PDF page into canvas context
    var renderContext = {
      canvasContext: context,
      viewport: viewport,
    };
    var renderTask = page.render(renderContext);
    renderTask.promise.then(function () {
      console.log("Page rendered");
    });
  });
}
function renderPage(file, pageNumber) {
  fileReader.readAsArrayBuffer(file);

  fileReader.onload = function () {
    //var data = fileReader.result
    var data = new Uint8Array(this.result);
    var loadingTask = pdfjsLib.getDocument({ data: data });
    loadingTask.promise.then(
      function (pdf) {
        doc = pdf;
        //setPagination();
        pageInput.value = 1;
        //$(".navigation").show();
        //console.log("PDF loaded");
        renderingPage(doc, 1);
      },
      function (reason) {
        // PDF loading error
        console.error(reason);
      }
    );
  };
}
