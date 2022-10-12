// Loaded via <script> tag, create shortcut to access PDF.js exports.
var pdfjsLib = window["pdfjs-dist/build/pdf"];
// The workerSrc property shall be specified.
pdfjsLib.GlobalWorkerOptions.workerSrc =
  "https://mozilla.github.io/pdf.js/build/pdf.worker.js";

document.querySelectorAll(".drop-zone-pdf__input").forEach((inputElement) => {
  const dropZoneElement = inputElement.closest(".drop-zone-pdf");

  dropZoneElement.addEventListener("click", (e) => {
    inputElement.click();
  });

  inputElement.addEventListener("change", (e) => {
    if (inputElement.files.length) {
      if (inputElement.files[0].size > 5242880) {
        console.log(
          "Sorry, the file size is too big. Try something smaller than 5MB."
        );
        return;
      } else {
        if (inputElement.files[0].type != "application/pdf") {
          console.log(
            "Sorry, the file type is not supported. Try something else."
          );
          return;
        } else {
          updateThumbnail(dropZoneElement, inputElement.files[0]);
        }
      }
    }
  });

  dropZoneElement.addEventListener("dragover", (e) => {
    e.preventDefault();
    dropZoneElement.classList.add("drop-zone-pdf--over");
  });

  ["dragleave", "dragend"].forEach((type) => {
    dropZoneElement.addEventListener(type, (e) => {
      dropZoneElement.classList.remove("drop-zone-pdf--over");
    });
  });

  dropZoneElement.addEventListener("drop", (e) => {
    e.preventDefault();

    if (e.dataTransfer.files.length) {
      inputElement.files = e.dataTransfer.files;
      updateThumbnail(dropZoneElement, e.dataTransfer.files[0]);
    }

    dropZoneElement.classList.remove("drop-zone-pdf--over");
  });
});

/**
 * Updates the thumbnail on a drop zone element.
 *
 * @param {HTMLElement} dropZoneElement
 * @param {File} file
 */
function updateThumbnail(dropZoneElement, file) {
  let thumbnailElement = dropZoneElement.querySelector(".drop-zone-pdf__thumb");

  // First time - remove the prompt
  if (dropZoneElement.querySelector(".drop-zone-pdf__prompt")) {
    dropZoneElement.querySelector(".drop-zone-pdf__prompt").remove();
  }

  // First time - there is no thumbnail element, so lets create it
  if (!thumbnailElement) {
    thumbnailElement = document.createElement("canvas");
    thumbnailElement.classList.add("page");
    thumbnailElement.id = "pdfViewer";
    thumbnailElement.style = "display: none;";
    dropZoneElement.appendChild(thumbnailElement);
  }

  thumbnailElement.dataset.label = file.name;
  //Show thumbnail for pdf files
  var fileReader = new FileReader();
  fileReader.onload = function () {
    var pdfData = new Uint8Array(this.result);
    // Using DocumentInitParameters object to load binary data.
    var loadingTask = pdfjsLib.getDocument({ data: pdfData });
    loadingTask.promise.then(
      function (pdf) {
        //console.log('PDF loaded');
        // Fetch the first page
        var pageNumber = 1;
        pdf.getPage(pageNumber).then(function (page) {
          //console.log('Page loaded');
          var desiredWidth = 250;
          var viewport = page.getViewport({ scale: 1 });
          var scale = desiredWidth / viewport.width;
          var scaledViewport = page.getViewport({ scale: scale });
          // Prepare canvas using PDF page dimensions
          var canvas = $("#pdfViewer")[0];
          var context = canvas.getContext("2d");
          canvas.height = scaledViewport.height;
          canvas.width = scaledViewport.width;
          // Render PDF page into canvas context
          var renderContext = {
            canvasContext: context,
            viewport: scaledViewport,
          };
          var renderTask = page.render(renderContext);
          renderTask.promise.then(function () {
            //console.log('Page rendered');
            var png_data = canvas.toDataURL("image/png");
            document.getElementById("pdf_preview").src = png_data;
            document.getElementById("titlenya").style.display = "none";
          });
        });
      },
      function (reason) {
        // PDF loading error
        //console.error(reason);
      }
    );
  };
  fileReader.readAsArrayBuffer(file);
}
