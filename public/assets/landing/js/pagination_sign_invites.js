const paginationNumbers = document.getElementById("pagination-numbers");
const paginatedList = document.getElementById("paginated-list");
const listItems = paginatedList.querySelectorAll(".step");
const nextButton = document.getElementById("next-button");
const prevButton = document.getElementById("prev-button");
const progressbar = document.getElementById("progress-page");
const progressbutton = document.getElementById("progress-button-1");
const progressbutton2 = document.getElementById("progress-button-2");
const progressbutton3 = document.getElementById("progress-button-3");
const prevconf = document.getElementById("prevconf");
const emptydata = document.getElementById("emptydata");
const submitform = document.getElementById("submitform");
const paginationLimit = 1;
const pageCount = Math.ceil(listItems.length / paginationLimit);
let currentPage = 1;

$(document).ready(function() {
  $(".add-email").click(function() {
    var html = $(".copy-fields").html();
    $(".after-add-email").after(html);
  });
  
  $("body").on("click", ".remove", function() {
    $(this)
      .parents(".control-group")
      .remove();
  });
});

const disableButton = (button) => {
  button.classList.add("disabled");
  button.setAttribute("disabled", true);
};

function submitForm() {
  if(document.getElementsByClassName("email[]").value == "") {
    emptydata.style.display = "block";
  } else {
    submitform.style.display = "block";
  }
}

function exitSubmit() {
  submitform.style.display = "none";
}

$('#submit').click(function(){
  submitform.style.display = "none";
  $('#form-documents').submit();
})

const enableButton = (button) => {
  button.classList.remove("disabled");
  button.removeAttribute("disabled");
};

const handlePageButtonsStatus = () => {
  if (currentPage === 1) {
    disableButton(prevButton);
  } else {
    enableButton(prevButton);
  }

  if (currentPage === 3) {
    $("#next-button").addClass("hide_page");
    $("#finish").removeClass("hide_page");
    //disableButton(nextButton);
  } else {
    enableButton(nextButton);
  }
};

const handleActivePageNumber = () => {
  document.querySelectorAll(".pagination-number").forEach((button) => {
    button.classList.remove("active");
    const pageIndex = Number(button.getAttribute("page-index"));
    if (pageIndex == currentPage) {
      button.classList.add("active");
    }
  });
};

const appendPageNumber = (index) => {
  const pageNumber = document.createElement("button");
  pageNumber.className = "pagination-number";
  pageNumber.innerHTML = index;
  pageNumber.setAttribute("page-index", index);
  pageNumber.setAttribute("aria-label", "Page " + index);

  paginationNumbers.appendChild(pageNumber);
};

const getPaginationNumbers = () => {
  for (let i = 1; i <= pageCount; i++) {
    appendPageNumber(i);
  }
};

const setCurrentPage = (pageNum) => {
  currentPage = pageNum;
  progressPage(pageNum);
  handleActivePageNumber();
  handlePageButtonsStatus();

  const prevRange = (pageNum - 1) * paginationLimit;
  const currRange = pageNum * paginationLimit;
  listItems.forEach((item, index) => {
    item.classList.add("hide_page");
    if (index >= prevRange && index < currRange) {
      item.classList.remove("hide_page");
    }
  });
};

const progressPage = (page) => {
  //Math.round(progressbar.style.width = (page / pageCount) * 100 + "%")
  if (page == 1) {
    progressbar.style.width = "0%";
    progressbutton.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton.style.color = "white";
    progressbutton2.style.backgroundColor = "#E9ECEF";
    progressbutton2.style.color = "#444444";
    progressbutton3.style.backgroundColor = "#E9ECEF";
    progressbutton3.style.color = "#444444";
  } else if (page == 2) {
    progressbar.style.width = "50%";
    progressbutton.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton.style.color = "white";
    progressbutton2.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton2.style.color = "white";
    progressbutton3.style.backgroundColor = "#E9ECEF";
    progressbutton3.style.color = "#444444";
  } else if (page == 3) {
    progressbar.style.width = "100%";
    progressbutton.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton.style.color = "white";
    progressbutton2.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton2.style.color = "white";
    progressbutton3.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton3.style.color = "white";
  }
};
function yesPrev() {
  prevconf.style.display = "none";
  if(currentPage == 3) {
    $("#finish").addClass("hide_page");
    $("#next-button").removeClass("hide_page");
    setCurrentPage(currentPage - 1);
  } else {
    setCurrentPage(currentPage - 1);
  }
}

function exitPrev() {
  prevconf.style.display = "none";
}

function exitEmpty() {
  emptydata.style.display = "none";
}

window.addEventListener("load", () => {
  getPaginationNumbers();
  setCurrentPage(1);

  prevButton.addEventListener("click", () => {
    prevconf.style.display = "block";
  });

  nextButton.addEventListener("click", () => {
    if (currentPage + 1 == 2 && document.getElementById("file doc").files.length == 0) {
        emptydata.style.display = "block";
        setCurrentPage(currentPage);
    } else {
      emptydata.style.display = "none";
      setCurrentPage(currentPage + 1);
    }
  });

  document.querySelectorAll(".pagination-number").forEach((button) => {
    const pageIndex = Number(button.getAttribute("page-index"));

    if (pageIndex) {
      button.addEventListener("click", () => {
        setCurrentPage(pageIndex);
      });
    }
  });
});
