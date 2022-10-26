const paginationNumbers = document.getElementById("pagination-numbers");
const paginatedList = document.getElementById("paginated-list");
const listItems = paginatedList.querySelectorAll(".step");
const nextButton = document.getElementById("next-button");
const prevButton = document.getElementById("prev-button");
const progressbar = document.getElementById("progress-page");
const progressbutton = document.getElementById("progress-button-1");
const progressbutton2 = document.getElementById("progress-button-2");
const progressbutton3 = document.getElementById("progress-button-3");
const progressbutton4 = document.getElementById("progress-button-4");
var signPage = document.getElementById("signPage");
var signX = document.getElementById("signX");
var signY = document.getElementById("signY");
var signH = document.getElementById("signH");
var signW = document.getElementById("signW");
const paginationLimit = 1;
const pageCount = Math.ceil(listItems.length / paginationLimit);
let currentPage = 1;

const disableButton = (button) => {
  button.classList.add("disabled");
  button.setAttribute("disabled", true);
};

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

  if (pageCount === currentPage) {
    disableButton(nextButton);
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
  if (pageNum == 3) {
    signPage.value = sign_page;
    signX.value = sign_x;
    signY.value = sign_y;
    signH.value = sign_h;
    signW.value = sign_w;
    // console.log(sign_page);
    // console.log("X: " + sign_x + " Y: " + sign_y + " H: " + sign_h + "px W: " + sign_w + "px");
  }
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
    progressbutton4.style.backgroundColor = "#E9ECEF";
    progressbutton4.style.color = "#444444";
  } else if (page == 2) {
    progressbar.style.width = "33.3%";
    progressbutton.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton.style.color = "white";
    progressbutton2.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton2.style.color = "white";
    progressbutton3.style.backgroundColor = "#E9ECEF";
    progressbutton3.style.color = "#444444";
    progressbutton4.style.backgroundColor = "#E9ECEF";
    progressbutton4.style.color = "#444444";
  } else if (page == 3) {
   // $(".step-3").removeClass("hide_page");
    progressbar.style.width = "66.6%";
    progressbutton.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton.style.color = "white";
    progressbutton2.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton2.style.color = "white";
    progressbutton3.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton3.style.color = "white";
    progressbutton4.style.backgroundColor = "#E9ECEF";
    progressbutton4.style.color = "#444444";
  } else if (page == 4) {
    // $(".step-3").addClass("hide_page");
    // $("#SignImg").removeClass("hide_page");
    progressbar.style.width = "100%";
    progressbutton.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton.style.color = "white";
    progressbutton2.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton2.style.color = "white";
    progressbutton3.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton3.style.color = "white";
    progressbutton4.style.backgroundColor = "rgba(65, 84, 241, 1)";
    progressbutton4.style.color = "white";
  }
};

window.addEventListener("load", () => {
  getPaginationNumbers();
  setCurrentPage(1);

  prevButton.addEventListener("click", () => {
    setCurrentPage(currentPage - 1);
  });

  nextButton.addEventListener("click", () => {
    setCurrentPage(currentPage + 1);
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
