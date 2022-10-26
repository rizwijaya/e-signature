var doc;
var fileReader = new FileReader();

function sign(klik) { 
  if (klik == 1) { //IF Button Click to Signing
    $(".signing").html( //Changing button to cancel
      '<i id="signBtn" class= "btn btn-sm btn-danger" style="position:absolute; top:2%; left:87%"><a onclick="sign(2)">cancel</a></i>'
    );
    $("#SignImg").addClass("Sign-Img");
    $("#SignImg").removeClass("hide_page");
   // $("#PDFSign").attr('id', 'signImg') //Adding id images to canvas
  } else if(klik == 2) { //IF Cancel Signing
    $(".signing").html( //Changing button to sign
      '<i id="signBtn" class="btn btn-sm btn-primary" style="position:absolute; top:2%; left:87%"><a onclick="sign(1)">sign</a></i>'
    );
    $("#SignImg").addClass("hide_page");
    $("#SignImg").removeClass("Sign-Img");
  } else {
    $(".signing").html( //Changing button to sign
    '<i id="signBtn" class="btn btn-sm btn-primary" style="position:absolute; top:2%; left:87%"><a onclick="sign(1)">sign</a></i>'
    );
    $("#SignImg").addClass("hide_page");
    $("#SignImg").removeClass("Sign-Img");
  }
}

function setPagination() {
  $(".pagination.bottom").html(
    '<li style="margin-right: 5px;" class="btn btn-primary disabled"><a onclick="prevPage()"><i class="material-icons">left</i></a></li><li style="margin-right: 5px;" class="btn btn-primary active"><a onclick="renderingPage(doc, 1)">1</a></li><li style="margin-right: 5px;" class="btn btn-primary disabled"><a onclick="nextPage()"><i class="material-icons">right</i></a></li>'
  );
  if (doc.numPages > 1) {
    var last = $(".pagination.bottom li").last();
    for (i = 2; i <= doc.numPages; i++) {
      last.before(
        '<li style="margin-right: 5px;" class="btn btn-primary"><a onclick="renderingPage(doc, ' +
          i +
          ')">' +
          i +
          "</a></li>"
      );
    }
    last.removeClass("disabled");
  }
}

function prevPage() {
  renderingPage(doc, curPage - 1);
}
function nextPage() {
  renderingPage(doc, curPage + 1);
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

//   console.log("Page PDF Reader" + pageNumber)
//   console.log("Current Pages: " + curPage)
  curPage = pageNumber;
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
    loadingTask.promise.then(function (pdf) {
        doc = pdf;
        setPagination();
        //$(".navigation").show();
        //console.log("PDF loaded");
        renderingPage(doc, 1);
      },
      function (reason) { // PDF loading error
        console.error(reason);
      }
    );
  };
}